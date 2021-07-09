// 由res2go自动生成，不要编辑。
package ui

import (
    "github.com/ying32/govcl/vcl"
    "github.com/ying32/govcl/vcl/types"
    "github.com/ying32/govcl/vcl/types/colors"
    "time"
)

type MemoFrame struct {
    *vcl.TFrame
    ConfigAble
    Memo      *vcl.TMemo

    msgChan chan string
    ticker  *time.Ticker
}


func NewMemoFrame(owner vcl.IComponent,option... FrameOption) (root *MemoFrame)  {
    vcl.CreateResFrame(owner, &root)
    for _,o:=range option {
        o(root)
    }
    return
}

func (this *MemoFrame) setup(){
    this.SetAlign(types.AlClient)
    this.Memo = vcl.NewMemo(this)
    this.Memo.SetAlign(types.AlClient)
    this.Memo.SetParent(this)
}

func (this *MemoFrame) OnCreate() {
    this.setup()
    this.msgChan = make(chan string, 200)
    this.ticker = time.NewTicker(100*time.Millisecond)
    this.update()
}

func (this *MemoFrame) OnDestroy(){

}

type onChangeFunc func(sender vcl.IObject)

func (this *MemoFrame) SetOnChange(bindFunc onChangeFunc) {
    this.Memo.SetOnChange(func(sender vcl.IObject) {
        bindFunc(sender)
    })
}

func (this *MemoFrame) SetColorRed() {
    this.Memo.Font().SetColor(colors.ClRed)
}
func (this *MemoFrame) SetColorDefault() {
    this.Memo.Font().SetColor(colors.ClSysDefault)
}

func (this *MemoFrame) Text() string {
    return this.Memo.Lines().Text()
}

func (this *MemoFrame) Add(s string) {
    this.Memo.Lines().Add(s)
}

func (this *MemoFrame) SetText(s string) {
    this.Memo.Lines().SetText(s)
}

func (this *MemoFrame) Clear() {
    this.Memo.Clear()
}

func (this *MemoFrame) Write(p []byte) (int, error) {
    this.msgChan <- string(p)
    return 0, nil
}

func (this *MemoFrame) WriteString(s string) {
    this.msgChan <- s
}

func (this *MemoFrame) write(s string) {
    this.Memo.Lines().Add(s)
}

func (this *MemoFrame) update() {
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
