// 由res2go自动生成。
// 在这里写你的事件。

package vcltool

import (
    "github.com/nomos/go-events"
    "github.com/nomos/go-log/log"
    "github.com/nomos/go-lokas/util"
    "github.com/nomos/go-tools/pjson"
    "github.com/ying32/govcl/vcl"
    "github.com/ying32/govcl/vcl/types"
    "github.com/ying32/govcl/vcl/types/keys"
    "strconv"
    "time"
)

//::private::
type TJsonValueEditFrameFields struct {
    ConfigAble
    events.EventEmmiter
    schema *pjson.Schema
    onSchemaChange func()
    assigning bool
    conf *util.AppConfig
}

func (this *TJsonValueEditFrame) OnCreate(){
    this.EventEmmiter = events.New()
    this.FormatCheck.SetChecked(this.conf.GetBool("value_edit_format"))
    this.FormatCheck.SetOnChange(func(sender vcl.IObject) {
        this.conf.Set("value_edit_format",this.FormatCheck.Checked())
        if this.schema!=nil {
            this.SetSchema(this.schema)
        }
    })
    this.ValueEdit.SetOnExit(func(sender vcl.IObject) {
        log.Warnf("ValueEdit SetOnExit",this.ValueEdit.Text())
        if this.schema == nil {
            log.Warnf("nildasdasd")
            this.ValueEdit.SetText("")
            return
        }
        switch this.schema.Type {
        case pjson.Object,pjson.Array:
            this.ValueEdit.SetText(this.schema.ToString(this.FormatCheck.Checked()))
            return
        default:
            s,ok:=this.schema.Type.CheckValue(this.ValueEdit.Text())
            log.Warnf("s ok",ok,s)
            if ok {
                this.ValueEdit.SetText(s)
                this.schema.Value = s
                this.Emit("schema_change")
            } else {
                go func() {
                    time.Sleep(time.Millisecond)
                    vcl.ThreadSync(func() {
                        this.ValueEdit.SetText(this.schema.Value)
                    })
                }()
            }
        }
    })
    this.KeyEdit.SetOnChange(func(sender vcl.IObject) {
        if this.schema == nil {
            if !this.assigning {
                this.KeyEdit.SetText("")
            }
            return
        }
        if this.schema.IsRoot() {
            this.KeyEdit.SetText("")
            return
        }
        if this.schema.IsArrayElem() {
            this.KeyEdit.SetText(this.schema.ToKeyString())
            return
        }
        if this.schema.IsObjectElem() {
            this.schema.Key = this.KeyEdit.Text()
            this.Emit("schema_change")
            return
        }
    })
    this.ValueEdit.SetOnKeyDown(func(sender vcl.IObject, key *types.Char, shift types.TShiftState) {
        if *key == keys.VkReturn {
            if this.schema.Type.IsValue() {
                this.SetFocus()
            }
        }
    })
    this.ValueEdit.SetOnChange(func(sender vcl.IObject) {
        if this.schema!=nil&&this.schema.Type == pjson.String {
            this.schema.Value = this.ValueEdit.Text()
            this.Emit("schema_change")
        }
        if this.schema!=nil&&this.schema.Type == pjson.Number {
            if s,ok:=pjson.Number.CheckValue(this.ValueEdit.Text());ok {
                this.schema.Value = s
                this.Emit("schema_change")
            }
        }
        if this.schema!=nil&&this.schema.Type == pjson.Boolean {
            if s,ok:=pjson.Boolean.CheckValue(this.ValueEdit.Text());ok {
                this.schema.Value = s
                this.SetSchema(this.schema)
                this.Emit("schema_change")
            }
        }
    })
    this.TypeList.SetOnSelect(func(sender vcl.IObject) {
        if this.schema == nil {
            return
        }
        t:=pjson.GetTypeByString(this.TypeList.Text())
        if t == -1 {
            this.TypeList.SetText(this.schema.Type.String())
        }
        success:=this.schema.ChangeType(t)
        log.Warnf(t.String(),success)
        if success {
            this.SetSchema(this.schema)
            this.Emit("schema_change")
        }
    })
}

func (this *TJsonValueEditFrame) SetOnSchemaChange(bindFunc func()){
    this.On("schema_change", func(i ...interface{}) {
        log.Warnf("SetOnSchemaChange",this.schema.Value)
        bindFunc()
    })
}

func (this *TJsonValueEditFrame) SetSchema(s *pjson.Schema){
    this.schema = nil
    this.assigning = true
    if s== nil {
        this.KeyEdit.SetText("")
        this.ValueEdit.SetText("")
        this.TypeList.SetText("")
        this.assigning = false
        return
    }
    log.Warnf("Type",s.Type.String())
    this.TypeList.SetText(s.Type.String())
    if s.Index!=-1 {
        this.KeyEdit.SetText("[item"+strconv.Itoa(s.Index)+"]")
    } else {
        this.KeyEdit.SetText(s.Key)
    }
    this.ValueEdit.SetText(s.ToString(this.FormatCheck.Checked()))
    this.schema = s
    this.assigning = false
}

func (this *TJsonValueEditFrame) OnTypeListChange(sender vcl.IObject) {

}

