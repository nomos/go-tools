// 由res2go自动生成。
// 在这里写你的事件。

package vcltool

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types/colors"
	"time"
)

//::private::
type TTextFrameFields struct {
	msgChan chan string
	ticker  *time.Ticker
}

func (this *TTextFrame) OnCreate() {
	this.msgChan = make(chan string, 200)
	this.ticker = time.NewTicker(time.Millisecond)
	this.update()
}

type onChangeFunc func(sender vcl.IObject)

func (this *TTextFrame) SetOnChange(bindFunc onChangeFunc) {
	this.Memo.SetOnChange(func(sender vcl.IObject) {
		bindFunc(sender)
	})
}

func (this *TTextFrame) SetColorRed(){
	this.Memo.Font().SetColor(colors.ClRed)
}
func (this *TTextFrame) SetColorDefault(){
	this.Memo.Font().SetColor(colors.ClSysDefault)
}

func (this *TTextFrame) Text() string {
	return this.Memo.Lines().Text()
}

func (this *TTextFrame) Add(s string) {
	this.Memo.Lines().Add(s)
}

func (this *TTextFrame) SetText(s string) {
	this.Memo.Lines().SetText(s)
}

func (this *TTextFrame) Clear() {
	this.Memo.Clear()
}

func (this *TTextFrame) Write(p []byte) (int, error) {
	this.msgChan <- string(p)
	return 0, nil
}

func (this *TTextFrame) WriteString(s string) {
	this.msgChan <- s
}

func (this *TTextFrame) write(s string) {
	this.Memo.Lines().Add(s)
}

func (this *TTextFrame) update() {
	go func() {
		for {
			select {
			case <-this.ticker.C:
				vcl.ThreadSync(func() {
					count := 1
				loop:
					for {
						count++
						select {
						case msg := <-this.msgChan:
							this.write(msg)
						default:
							break loop
						}
						if count > 10 {
							break loop
						}
					}
				})
			}
		}
	}()
}
