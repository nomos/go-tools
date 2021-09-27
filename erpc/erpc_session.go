package erpc

import (
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

type ErpcSessionOption func(*Session)

var _ lokas.ISession = &Session{}
var _ lokas.IActor = &Session{}

func NewSession(conn lokas.IConn, id util.ID, manager lokas.ISessionManager, opts ...ErpcSessionOption) *Session {
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
	ClientMsgHandler func(msg *protocol.BinaryMessage)
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
		//ClientSideLoop
	CLIENT_LOOP:
		for {
			select {
			case <-this.ticker.C:
				//ClientSide MessageLoop
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
					break CLIENT_LOOP
				}
				if this.ClientMsgHandler != nil {
					this.ClientMsgHandler(msg)
				} else {
					log.Error("no msg handler found")
				}
			case <-this.done:
				this.closeSession()
				break CLIENT_LOOP
			}
		}
		close(this.done)
		close(this.Messages)
	}()
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
	this.GetProcess().RemoveActor(this)
	this.stop()
}

func (this *Session) OnRecv(conn lokas.IConn, data []byte) {
	d := make([]byte, len(data), len(data))
	copy(d, data)
	this.Messages <- d
}
