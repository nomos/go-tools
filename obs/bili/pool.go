package bili

import (
	"github.com/nomos/go-lokas/log"
)

// Pool`s fields map CMD value
type Pool struct {
	UserMsg          chan string
	UserGift         chan string
	UserEnter        chan string
	UserGuard        chan string
	MsgUncompressed  chan string
	UserEntry        chan string
	UserInteractWord chan string
}

func NewPool() *Pool {
	return &Pool{
		UserMsg:          make(chan string, 10),
		UserGift:         make(chan string, 10),
		UserEnter:        make(chan string, 10),
		UserGuard:        make(chan string, 10),
		MsgUncompressed:  make(chan string, 10),
		UserEntry:        make(chan string, 10),
		UserInteractWord: make(chan string, 10),
	}
}

func (pool *Pool) Handle() {
	for {
		select {
		case uc := <-pool.MsgUncompressed:
			// 目前只处理未压缩数据的关注数变化信息
			if cmd := json.Get([]byte(uc), "cmd").ToString(); CMD(cmd) == CMD_ROOM_REAL_TIME_MESSAGE_UPDATE {
				fans := json.Get([]byte(uc), "data", "fans").ToInt()
				log.Infof("当前房间关注数变动：", fans)
			}
		case src := <-pool.UserMsg:
			m := NewDanmu()
			m.GetDanmuMsg([]byte(src))
			log.Infof("%d-%s | %d-%s: %s", m.MedalLevel, m.MedalName, m.Ulevel, m.Uname, m.Text)
		case src := <-pool.UserGift:
			g := NewGift()
			g.GetGiftMsg([]byte(src))
			log.Infof("%s %s 价值 %d 的 %s", g.UUname, g.Action, g.Price, g.GiftName)
		case src := <-pool.UserEnter:
			name := json.Get([]byte(src), "data", "uname").ToString()
			log.Infof("欢迎VIP %s 进入直播间", name)
		case src := <-pool.UserGuard:
			name := json.Get([]byte(src), "data", "username").ToString()
			log.Infof("欢迎房管 %s 进入直播间", name)
		case src := <-pool.UserEntry:
			cw := json.Get([]byte(src), "data", "copy_writing").ToString()
			log.Infof("%s", cw)
		case src := <-pool.UserInteractWord:
			cw := json.Get([]byte(src), "data", "uname").ToString()
			log.Infof("玩家 %s 进入直播间", cw)
		}
	}
}
