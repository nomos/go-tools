package erpc

import (
	"errors"
	"github.com/nomos/go-lokas"
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/lox"
	"github.com/nomos/go-lokas/lox/flog"
	"github.com/nomos/go-lokas/protocol"
	"github.com/nomos/go-lokas/util"
	"go.uber.org/zap"
	"time"
)

const (
	TimeOut            = time.Second * 15
	UpdateTime = time.Second*15
)

type SessionOption func(*Session)

var _ lokas.ISession = &Session{}
var _ lokas.IActor = &Session{}

func NewSession(conn lokas.IConn, id util.ID, manager lokas.ISessionManager, opts ...SessionOption) *Session {
	s := &Session{
		Actor:    lox.NewActor(),
		Messages: make(chan []byte, 100),
		Conn:     conn,
		manager:  manager,
		timeout:  TimeOut,
		ticker:   time.NewTicker(UpdateTime),
	}
	for _, o := range opts {
		o(s)
	}
	s.SetType("ErpcSession")
	s.SetId(id)
	return s
}

type Session struct {
	*lox.Actor
	Messages         chan []byte
	Conn             lokas.IConn
	manager          lokas.ISessionManager
	done             chan struct{}
	OnCloseFunc      func(conn lokas.IConn)
	OnOpenFunc       func(conn lokas.IConn)
	ClientMsgHandler func(msg *protocol.BinaryMessage,session *Session)
	timeout          time.Duration
	ticker           *time.Ticker
}

func (this *Session) Load(conf lokas.IConfig) error {
	panic("implement me")
}

func (this *Session) Unload() error {
	panic("implement me")
}

func (this *Session) OnStart() error {
	panic("implement me")
}

func (this *Session) OnStop() error {
	panic("implement me")
}

func (this *Session) OnCreate() error {
	return nil
}

func (this *Session) Start() error {
	return nil
}

func (this *Session) Stop() error {
	return nil
}

func (this *Session) CallAdminCommand(command *lox.AdminCommand)([]byte,error){
	res,err:=this.Call(0,command)
	if err != nil {
		log.Error(err.Error())
		return nil,err
	}
	cmd:=res.(*lox.AdminCommandResult)
	if !cmd.Success {
		return nil,errors.New(string(cmd.Data))
	}
	return cmd.Data,nil
}

func (this *Session) SendMessage(actorId util.ID, transId uint32, msg protocol.ISerializable) error {
	_,err:=msg.GetId()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	data,err:=protocol.MarshalMessage(transId,msg,protocol.BINARY)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	this.Conn.Write(data)
	return nil
}

func (this *Session) OnDestroy() error {
	return nil
}

func (this *Session) GetConn() lokas.IConn {
	return this.Conn
}

func (this *Session) StartMessagePump() {

	this.done = make(chan struct{})
	go func() {
		defer func() {
			r := recover()
			if r != nil {
				if e, ok := r.(error); ok {
					log.Errorf(e.Error())
					this.Conn.Close()
				}
			}
		}()
	Loop:
		for {
			select {
			case <-this.ticker.C:
			case data := <-this.Messages:
				cmdId := protocol.GetCmdId16(data)
				msg, err := protocol.UnmarshalMessage(data, protocol.BINARY)
				if err != nil {
					log.Error("unmarshal client message error",
						zap.Any("cmdId", cmdId),
					)
					msg, _ := protocol.NewError(protocol.ERR_MSG_FORMAT).Marshal()
					_, err := this.Conn.Write(msg)
					if err != nil {
						log.Error(err.Error())
					}
					this.Conn.Close()
					break Loop
				}
				if msg.CmdId == lox.TAG_AdminCmd {
					this.handAdminCommand(msg)
				} else if this.ClientMsgHandler != nil {
					this.ClientMsgHandler(msg,this)
				} else {
					log.Errorf("no msg handler found",msg.CmdId)
				}
			case <-this.done:
				this.closeSession()
				break Loop
			}
		}
		close(this.done)
		close(this.Messages)
	}()
}

func (this *Session) handAdminCommand(msg *protocol.BinaryMessage){
	cmd:=msg.Body.(*lox.AdminCommand)
	if handler, ok := rpcHandlers[cmd.Command]; ok {
		res,err:=handler(cmd,cmd.ParamsValue())
		if err!=nil {
			log.Error(err.Error())
			ret:=lox.NewAdminCommandResult(cmd,false,[]byte(err.Error()))
			data, _ := protocol.MarshalMessage(msg.TransId, ret, protocol.BINARY)
			this.Conn.Write(data)
			return
		}
		ret:=lox.NewAdminCommandResult(cmd,true,res)
		data, _ := protocol.MarshalMessage(msg.TransId, ret, protocol.BINARY)
		this.Conn.Write(data)
	} else {
		log.Errorf("Admin Command not found",cmd.Command)
		ret:=lox.NewAdminCommandResult(cmd,false,[]byte("admin cmd not found"))
		data, _ := protocol.MarshalMessage(msg.TransId, ret, protocol.BINARY)
		this.Conn.Write(data)
	}
}

func (this *Session) closeSession() {
	if this.manager != nil {
		this.manager.RemoveSession(this.GetId())
	}
}

func (this *Session) stop() {
	close(this.done)
}

func (this *Session) OnOpen(conn lokas.IConn) {
	this.StartMessagePump()
	log.Info("ErpcSession:OnOpen", flog.ActorInfo(this)...)
	if this.manager != nil {
		this.manager.AddSession(this.GetId(), this)
	}
	if this.OnOpenFunc != nil {
		this.OnOpenFunc(conn)
	}
}

func (this *Session) OnClose(conn lokas.IConn) {
	if this.manager != nil {
		this.manager.RemoveSession(this.GetId())
	}
	log.Info("ErpcSession:OnClose", flog.ActorInfo(this)...)
	if this.OnCloseFunc != nil {
		this.OnCloseFunc(conn)
	}
	this.stop()
}

func (this *Session) OnRecv(conn lokas.IConn, data []byte) {
	d := make([]byte, len(data), len(data))
	copy(d, data)
	this.Messages <- d
}
