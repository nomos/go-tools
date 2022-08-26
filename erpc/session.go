package erpc

import (
	"encoding/json"
	"errors"
	"github.com/nomos/go-lokas"
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/lox"
	"github.com/nomos/go-lokas/lox/flog"
	"github.com/nomos/go-lokas/protocol"
	"github.com/nomos/go-lokas/rpc"
	"github.com/nomos/go-lokas/util"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"regexp"
	"time"
)

const (
	TimeOut    = time.Second * 15
	UpdateTime = time.Second * 15
)

type SessionOption func(*Session)

func WithSessionHandler(handler func(msg *protocol.BinaryMessage, session *Session)) SessionOption {
	return func(session *Session) {
		session.ClientMsgHandler = handler
	}
}

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
	s.ComposeLogger = log.NewComposeLogger(true, log.ConsoleConfig(""), 1)
	s.ComposeLogger.SetConsoleWriter(s)
	s.SetType("ErpcSession")
	s.SetId(id)
	return s
}

type Session struct {
	*lox.Actor
	*log.ComposeLogger
	Messages         chan []byte
	Conn             lokas.IConn
	manager          lokas.ISessionManager
	done             chan struct{}
	OnCloseFunc      func(conn lokas.IConn)
	OnOpenFunc       func(conn lokas.IConn)
	ClientMsgHandler func(msg *protocol.BinaryMessage, session *Session)
	timeout          time.Duration
	ticker           *time.Ticker
	outchan          chan []byte
}

func (this *Session) WriteString(s string) {
	this.Write([]byte(s))
}

func (this *Session) Write(p []byte) (int, error) {
	err := this.SendMessage(0, 0, lox.NewConsoleEvent(string(p)))
	return 0, err
}

func (this *Session) WriteConsole(e zapcore.Entry, p []byte) error {
	return nil
}
func (this *Session) WriteJson(e zapcore.Entry, p []byte) error {
	var j map[string]interface{}
	json.Unmarshal(p, &j)
	level := j["level"].(string)
	time := j["time"].(string)
	caller := j["caller"].(string)
	msg := j["msg"].(string)
	o := make(map[string]interface{})
	for k, v := range j {
		if k != "level" && k != "time" && k != "caller" && k != "msg" {
			o[k] = v
		}
	}
	level = regexp.MustCompile(`[[][A-Z]+[]][A-z]*`).FindString(level)
	jstr, _ := json.Marshal(o)
	str := time + " " + level + "   " + caller + " " + msg + " " + string(jstr)
	this.Write([]byte(str))
	return nil
}
func (this *Session) WriteObject(e zapcore.Entry, o map[string]interface{}) error {
	return nil
}

func (this *Session) Load(conf lokas.IConfig) error {
	return nil
}

func (this *Session) Unload() error {
	return nil
}

func (this *Session) OnStart() error {
	return nil
}

func (this *Session) OnStop() error {
	return nil
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

func (this *Session) CallAdminCommand(command *lox.AdminCommand) ([]byte, error) {
	res, err := this.Call(0, command)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	cmd := res.(*lox.AdminCommandResult)
	if !cmd.Success {
		return nil, errors.New(string(cmd.Data))
	}
	return cmd.Data, nil
}

func (this *Session) SendMessage(actorId util.ID, transId uint32, msg protocol.ISerializable) error {
	_, err := msg.GetId()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	data, err := protocol.MarshalMessage(transId, msg, protocol.BINARY)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	this.outchan <- data
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
	this.outchan = make(chan []byte, 100)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				util.Recover(r, false)
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
				if msg.CmdId == lox.TAG_ADMIN_CMD {
					this.handAdminCommand(msg)
				} else if this.ClientMsgHandler != nil {
					this.ClientMsgHandler(msg, this)
				} else {
					log.Errorf("no msg handler found", msg.CmdId)
				}
			case <-this.done:
				this.closeSession()
				break Loop
			}
		}
		close(this.done)
		close(this.Messages)
	}()
	go func() {
		defer func() {
			if r := recover(); r != nil {
				util.Recover(r, false)
			}
		}()
	WRITE_LOOP:
		for {
			select {
			case msg := <-this.outchan:
				this.Conn.Write(msg)
			case <-this.done:
				break WRITE_LOOP
			}
		}
		close(this.outchan)
	}()
}

func (this *Session) handAdminCommand(msg *protocol.BinaryMessage) {
	cmd := msg.Body.(*lox.AdminCommand)
	log.Info("adminCmd", zap.String("cmd", cmd.Command), zap.Any("values", cmd.Params))
	if handler, ok := rpc.GetRpcHandlers()[cmd.Command]; ok {
		res, err := handler(cmd, cmd.ParamsValue(), this)
		if err != nil {
			log.Error(err.Error())
			ret := lox.NewAdminCommandResult(cmd, false, []byte(err.Error()))
			this.SendMessage(0, msg.TransId, ret)
			return
		}
		if res == nil {
			res = []byte("")
		}
		ret := lox.NewAdminCommandResult(cmd, true, res)
		this.SendMessage(0, msg.TransId, ret)
	} else {
		log.Errorf("Admin Command not found", cmd.Command)
		ret := lox.NewAdminCommandResult(cmd, false, []byte("admin cmd not found"))
		this.SendMessage(0, msg.TransId, ret)
	}
}

func (this *Session) closeSession() {
	if this.manager != nil {
		this.manager.RemoveSession(this.GetId())
	}
}

func (this *Session) stop() {
	this.done <- struct{}{}
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
