package bili

import (
	"github.com/nomos/go-lokas/log"
)

// Pool`s fields map CMD value
type Pool struct {
	UserMsg         chan string
	UserGift        chan string
	VipEnter        chan string
	GuardEnter      chan string
	MsgUncompressed chan string
	UserEntry       chan string
	UserEnter       chan string
}

func NewPool() *Pool {
	return &Pool{
		UserMsg:         make(chan string, 10),
		UserGift:        make(chan string, 10),
		VipEnter:        make(chan string, 10),
		GuardEnter:      make(chan string, 10),
		MsgUncompressed: make(chan string, 10),
		UserEntry:       make(chan string, 10),
		UserEnter:       make(chan string, 10),
	}
}

func (this *Client) Handle() {
	for {
		select {
		case uc := <-this.MsgUncompressed:
			// 目前只处理未压缩数据的关注数变化信息
			if cmd := json.Get([]byte(uc), "cmd").ToString(); CMD(cmd) == CMD_ROOM_REAL_TIME_MESSAGE_UPDATE {
				fans := json.Get([]byte(uc), "data", "fans").ToInt()
				log.Infof("当前房间关注数变动：", fans)
			}
		case src := <-this.UserMsg:
			m := NewDanmu()
			m.GetDanmuMsg([]byte(src))
			log.Infof("%d-%s | %d-%s: %s", m.MedalLevel, m.MedalName, m.Ulevel, m.Uname, m.Text)
		case src := <-this.UserGift:
			g := NewGift()
			g.GetGiftMsg([]byte(src))
			log.Infof("%s %s 价值 %d 的 %s", g.UUname, g.Action, g.Price, g.GiftName)
		case src := <-this.VipEnter:
			name := json.Get([]byte(src), "data", "uname").ToString()
			log.Infof("欢迎VIP %s 进入直播间", name)
		case src := <-this.GuardEnter:
			name := json.Get([]byte(src), "data", "username").ToString()
			log.Infof("欢迎房管 %s 进入直播间", name)
		case src := <-this.UserEntry:
			cw := json.Get([]byte(src), "data", "copy_writing").ToString()
			log.Infof("%s", cw)
		case src := <-this.UserEnter:
			cw := json.Get([]byte(src), "data", "uname").ToString()
			log.Infof("玩家 %s 进入直播间", cw)
		}
	}
}
