package erpc

import (
	"github.com/nomos/go-lokas"
	"github.com/nomos/go-lokas/lox/flog"
	"github.com/nomos/go-lokas/protocol"
	"github.com/nomos/go-lokas/util"
	"go.uber.org/zap"
	"time"
)

type ErpcSessionOption func(*Session)

var _ lokas.ISession = &Session{}
var _ lokas.IActor = &Session{}

func NewErpcSession(conn lokas.IConn, id util.ID, manager lokas.ISessionManager, opts ...ErpcSessionOption) *Session {
	s := &Session{
		Actor:    NewActor(),
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
	*Actor
	Verified         bool
	Messages         chan []byte
	Conn             lokas.IConn
	Protocol         protocol.TYPE
	manager          lokas.ISessionManager
	done             chan struct{}
	OnCloseFunc      func(conn lokas.IConn)
	OnOpenFunc       func(conn lokas.IConn)
	ClientMsgHandler func(msg *protocol.BinaryMessage)
	AuthFunc         func(data []byte) error
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
	log.Info("ErpcSession:StartMessagePump", flog.ActorInfo(this)...)

	this.msgChan = make(chan *protocol.RouteMessage, 100)
	this.done = make(chan struct{})
	go func() {
		defer func() {
			r := recover()
			if r != nil {
				if e, ok := r.(error); ok {
					log.Errorf(e.Error())
					log.Error("客户端协议出错")
					this.Conn.Close()
				}
			}
		}()
		//ClientSideLoop
	CLIENT_LOOP:
		for {
			select {
			case <-this.ticker.C:
				if this.OnUpdateFunc != nil && this.Verified {
					this.OnUpdateFunc()
				}
				//ClientSide MessageLoop
			case data := <-this.Messages:
				cmdId := protocol.GetCmdId16(data)
				if !this.Verified && cmdId != protocol.TAG_HandShake {
					var msg []byte
					msg, err := protocol.MarshalMessage(0, protocol.NewError(protocol.ERR_AUTH_FAILED), this.Protocol)
					if err != nil {
						log.Error(err.Error())
						this.Conn.Wait()
						this.Conn.Close()
						return
					}
					log.Errorf("Auth Failed", cmdId)
					this.Conn.Write(msg)
					this.Conn.Wait()
					this.Conn.Close()
					return
				}
				msg, err := protocol.UnmarshalMessage(data, this.Protocol)
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
				if cmdId == protocol.TAG_HandShake {
					var err error
					if this.AuthFunc != nil {
						err = this.AuthFunc(msg.Body.(*protocol.HandShake).Data)
					}
					if err != nil {
						log.Error(err.Error())
						msg, err := protocol.MarshalMessage(msg.TransId, protocol.NewError(protocol.ERR_AUTH_FAILED), this.Protocol)
						if err != nil {
							log.Error(err.Error())
							this.Conn.Wait()
							this.Conn.Close()
							return
						}
						log.Errorf("Auth Failed", cmdId)
						this.Conn.Write(msg)
						this.Conn.Wait()
						this.Conn.Close()
						break CLIENT_LOOP
					}
					_, err = this.Conn.Write(data)
					if err != nil {
						log.Error(err.Error())
						this.Conn.Close()
						break CLIENT_LOOP
					}
					this.Verified = true
					continue
				}
				if cmdId == protocol.TAG_Ping {
					//ping:=msg.Body.(*Protocol.Ping)
					pong := &protocol.Pong{Time: time.Now()}
					data, err := protocol.MarshalMessage(msg.TransId, pong, this.Protocol)
					if err != nil {
						log.Error(err.Error())
						this.Conn.Wait()
						this.Conn.Close()
						break CLIENT_LOOP
					}
					_, err = this.Conn.Write(data)
					if err != nil {
						log.Error(err.Error())
						this.Conn.Close()
						break CLIENT_LOOP
					}
					continue
				}
				if this.ClientMsgHandler != nil {
					this.ClientMsgHandler(msg)
				} else {
					log.Error("no msg handler found")
				}
			case <-this.done:
				this.closeSession()
				//log.Warnf("done")
				break CLIENT_LOOP
			}
		}
		close(this.msgChan)
		close(this.done)
		close(this.Messages)
	}()

	go func() {
		defer func() {
			r := recover()
			if r != nil {
				if e, ok := r.(error); ok {
					log.Errorf(e.Error())
					log.Error("服务端协议出错")
					this.Conn.Close()
				}
			}
		}()
	SERVER_LOOP:
		for {
			select {
			//ServerSideMsgLoop
			case rMsg := <-this.msgChan:
				this.OnMessage(rMsg)
			case <-this.done:
				break SERVER_LOOP
			}
		}
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
