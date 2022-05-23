package bili

import (
	"bytes"
	"encoding/binary"
	"github.com/gorilla/websocket"
	"github.com/nomos/go-lokas/log"
	"time"
)

type DanMuMsg struct {
	UID        uint32 `json:"uid"`
	Uname      string `json:"uname"`
	Ulevel     uint32 `json:"ulevel"`
	Text       string `json:"text"`
	MedalLevel uint32 `json:"medal_level"`
	MedalName  string `json:"medal_name"`
}

type InteractWord struct {
	UID        uint32 `json:"uid"`
	Uname      string `json:"uname"`
	Ulevel     uint32 `json:"ulevel"`
	MedalLevel uint32 `json:"medal_level"`
	MedalName  string `json:"medal_name"`
}

func NewDanmu() *DanMuMsg {
	return &DanMuMsg{
		UID:        0,
		Uname:      "",
		Ulevel:     0,
		Text:       "",
		MedalLevel: 0,
		MedalName:  "无勋章",
	}
}

type Gift struct {
	UUname   string `json:"u_uname"`
	Action   string `json:"action"`
	Price    uint32 `json:"price"`
	GiftName string `json:"gift_name"`
}

func NewGift() *Gift {
	return &Gift{
		UUname:   "",
		Action:   "",
		Price:    0,
		GiftName: "",
	}
}

type WelCome struct {
}

type Notice struct {
}

type CMD string

var (
	CMD_DANMU_MSG                     CMD = "DANMU_MSG"     // 普通弹幕信息
	CMD_SEND_GIFT                     CMD = "SEND_GIFT"     // 普通的礼物，不包含礼物连击
	CMD_WELCOME                       CMD = "WELCOME"       // 欢迎VIP
	CMD_WELCOME_GUARD                 CMD = "WELCOME_GUARD" // 欢迎房管
	CMD_INTERACT_WORD                 CMD = "INTERACT_WORD"
	CMD_ENTRY_EFFECT                  CMD = "ENTRY_EFFECT"                  // 欢迎舰长等头衔
	CMD_ROOM_REAL_TIME_MESSAGE_UPDATE CMD = "ROOM_REAL_TIME_MESSAGE_UPDATE" // 房间关注数变动
)

func (c *Client) SendPackage(packetlen uint32, magic uint16, ver uint16, typeID uint32, param uint32, data []byte) (err error) {
	packetHead := new(bytes.Buffer)

	if packetlen == 0 {
		packetlen = uint32(len(data) + 16)
	}
	var pdata = []interface{}{
		packetlen,
		magic,
		ver,
		typeID,
		param,
	}

	// 将包的头部信息以大端序方式写入字节数组
	for _, v := range pdata {
		if err = binary.Write(packetHead, binary.BigEndian, v); err != nil {
			log.Infof("binary.Write err: ", err)
			return
		}
	}

	// 将包内数据部分追加到数据包内
	sendData := append(packetHead.Bytes(), data...)

	// log.Info("本次发包消息为：", sendData)

	if err = c.conn.WriteMessage(websocket.BinaryMessage, sendData); err != nil {
		log.Infof("c.conn.Write err: ", err)
		return
	}

	return
}

func (c *Client) ReceiveMsg() {
	pool := NewPool()
	go pool.Handle()
	defer func() {
		r := recover()
		if r != nil {
			if e, ok := r.(error); ok {
				log.Errorf(e.Error())
				log.Error("客户端协议出错")
			}
		}
	}()
	for {
		if !c.Connecting && !c.Connected {
			continue
		}
		err := c.conn.SetReadDeadline(time.Now().Add(time.Millisecond * 15000))
		if err != nil {
			log.Error(err.Error())
		}
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			log.Infof("ReadMsg err :", err.Error())
			c.Reconnect()
			continue
		}

		switch msg[11] {
		case 8:
			log.Info("握手包收发完毕，连接成功")
			c.Connected = true
			c.Connecting = false
		case 3:
			onlineNow := ByteArrToDecimal(msg[16:])
			if uint32(onlineNow) != c.Room.Online {
				c.Room.Online = uint32(onlineNow)
				log.Infof("当前房间人气变动：", uint32(onlineNow))
			}
		case 5:
			if inflated, err := ZlibInflate(msg[16:]); err != nil {
				// 代表是未压缩数据
				pool.MsgUncompressed <- string(msg[16:])
			} else {
				for len(inflated) > 0 {
					l := ByteArrToDecimal(inflated[:4])
					m := json.Get(inflated[16:l], "cmd").ToString()
					var data []byte = make([]byte, l-16, l-16)
					copy(data, inflated[16:l])
					switch CMD(m) {
					case CMD_DANMU_MSG:
						pool.UserMsg <- string(data)
					case CMD_SEND_GIFT:
						pool.UserGift <- string(data)
					case CMD_WELCOME:
						pool.UserEnter <- string(data)
					case CMD_WELCOME_GUARD:
						pool.UserGuard <- string(data)
					case CMD_ENTRY_EFFECT:
						pool.UserEntry <- string(data)
					case CMD_INTERACT_WORD:
						pool.UserInteractWord <- string(data)
					default:
						log.Infof(json.Get(inflated[16:l]).ToString())
					}
					inflated = inflated[l:]
				}
			}
		default:
			log.Infof(msg[11])
		}
		time.Sleep(1)
	}
}

func (c *Client) HeartBeat() {
	for {
		if c.Connected && !c.Connecting {
			obj := []byte("5b6f626a656374204f626a6563745d")
			err := c.SendPackage(0, 16, 1, 2, 1, obj)
			if err != nil {
				log.Errorf("heart beat err: ", err)
			}
		}
		time.Sleep(30 * time.Second)
	}
}

func (c *Client) Reconnect() {
	log.Infof("Reconnect")
	c.rcMutex.Lock()
	defer c.rcMutex.Unlock()
	c.Connected = false
	c.Connecting = true
	if c.conn != nil {

		c.conn = nil
	}
	c.connect()
}
