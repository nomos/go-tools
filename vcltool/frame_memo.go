// 由res2go自动生成，不要编辑。
package vcltool

import (
    "github.com/ying32/govcl/vcl"
    "github.com/ying32/govcl/vcl/types"
    "github.com/ying32/govcl/vcl/types/colors"
    "time"
)

type TMemoFrame struct {
    *vcl.TFrame
    Memo      *vcl.TMemo

    msgChan chan string
    ticker  *time.Ticker
}


func NewMemoFrame(owner vcl.IComponent) (root *TMemoFrame)  {
    vcl.CreateResFrame(owner, &root)
    return
}

func (this *TMemoFrame) setup(){
    this.SetAlign(types.AlClient)
    this.Memo = vcl.NewMemo(this)
    this.Memo.SetAlign(types.AlClient)
    this.Memo.SetParent(this)
}

func (this *TMemoFrame) OnCreate() {
    this.setup()
    this.msgChan = make(chan string, 200)
    this.ticker = time.NewTicker(100*time.Millisecond)
    this.update()
}

type onChangeFunc func(sender vcl.IObject)

func (this *TMemoFrame) SetOnChange(bindFunc onChangeFunc) {
    this.Memo.SetOnChange(func(sender vcl.IObject) {
        bindFunc(sender)
    })
}

func (this *TMemoFrame) SetColorRed() {
    this.Memo.Font().SetColor(colors.ClRed)
}
func (this *TMemoFrame) SetColorDefault() {
    this.Memo.Font().SetColor(colors.ClSysDefault)
}

func (this *TMemoFrame) Text() string {
    return this.Memo.Lines().Text()
}

func (this *TMemoFrame) Add(s string) {
    this.Memo.Lines().Add(s)
}

func (this *TMemoFrame) SetText(s string) {
    this.Memo.Lines().SetText(s)
}

func (this *TMemoFrame) Clear() {
    this.Memo.Clear()
}

func (this *TMemoFrame) Write(p []byte) (int, error) {
    this.msgChan <- string(p)
    return 0, nil
}

func (this *TMemoFrame) WriteString(s string) {
    this.msgChan <- s
}

func (this *TMemoFrame) write(s string) {
    this.Memo.Lines().Add(s)
}

func (this *TMemoFrame) update() {
    go func() {
        for {
            select {
            case <-this.ticker.C:
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
            }
        }
    }()
}
