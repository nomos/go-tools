package pjson

import (
	"encoding/json"
	"github.com/iancoleman/orderedmap"
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/util"
	"github.com/nomos/go-tools/ui"
	"github.com/nomos/go-tools/ui/icons"
	"github.com/nomos/go-tools/ui/treeview"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var _ ui.IFrame = (*JsonEditor)(nil)

type JsonEditor struct {
	*vcl.TFrame
	ui.ConfigAble

	lastTreeItem *vcl.TTreeNode
	textEdit     *ui.MemoFrame
	valueEditor  *ui.ValueEditorFrame
	tree         *treeview.TreeView

	clipSchema *Schema
	schema     *Schema
	editingSchema *Schema

	keyEditTag bool
	remodeling bool
	editing    bool
	assigning bool
	mutex      sync.Mutex
}

func NewEditor(owner vcl.IComponent, option ...ui.FrameOption) (root *JsonEditor) {
	vcl.CreateResFrame(owner, &root)
	for _, o := range option {
		o(root)
	}
	return
}

func (this *JsonEditor) setup() {
	this.SetAlign(types.AlClient)
	ui.CreateSplitter(this, types.AlLeft, 6)
	leftPanel := ui.CreatePanel(types.AlLeft, this)
	leftPanel.Constraints().SetMinWidth(300)
	leftPanel.SetWidth(350)
	this.textEdit = ui.NewMemoFrame(leftPanel)
	this.textEdit.OnCreate()
	this.textEdit.SetAlign(types.AlClient)
	this.textEdit.SetParent(leftPanel)
	this.textEdit.BorderSpacing().SetLeft(6)
	this.textEdit.BorderSpacing().SetRight(6)
	rightPanel := ui.CreatePanel(types.AlClient, this)
	line2 := ui.CreateLine(types.AlTop, 0, 0, 32, rightPanel)
	line2.BorderSpacing().SetBottom(3)
	line3 := ui.CreateLine(types.AlTop, 6, 0, 32, leftPanel)
	line3.BorderSpacing().SetBottom(3)
	ui.CreateSeg(3, line2)
	ui.CreateSpeedBtn("sort_up", icons.GetImageList(32, 32), line2)
	ui.CreateSeg(3, line2)
	ui.CreateSpeedBtn("sort_down", icons.GetImageList(32, 32), line2)
	ui.CreateSeg(3, line2)
	ui.CreateSpeedBtn("null_icon", icons.GetImageList(32, 32), line2).SetOnClick(func(sender vcl.IObject) {
		this.AddNewSchemaAtSelected(Null)
	})
	ui.CreateSeg(3, line2)
	ui.CreateSpeedBtn("object_box", icons.GetImageList(32, 32), line2).SetOnClick(func(sender vcl.IObject) {
		this.AddNewSchemaAtSelected(Object)
	})
	ui.CreateSeg(3, line2)
	ui.CreateSpeedBtn("array_box", icons.GetImageList(32, 32), line2).SetOnClick(func(sender vcl.IObject) {
		this.AddNewSchemaAtSelected(Array)
	})
	ui.CreateSeg(3, line2)
	ui.CreateSpeedBtn("boolean_icon", icons.GetImageList(32, 32), line2).SetOnClick(func(sender vcl.IObject) {
		this.AddNewSchemaAtSelected(Boolean)
	})
	ui.CreateSeg(3, line2)
	ui.CreateSpeedBtn("number_icon", icons.GetImageList(32, 32), line2).SetOnClick(func(sender vcl.IObject) {
		this.AddNewSchemaAtSelected(Number)
	})
	ui.CreateSeg(3, line2)
	ui.CreateSpeedBtn("string_icon", icons.GetImageList(32, 32), line2).SetOnClick(func(sender vcl.IObject) {
		this.AddNewSchemaAtSelected(String)
	})

	ui.CreateSpeedBtn("cancel", icons.GetImageList(32, 32), line3)
	ui.CreateSeg(3, line3)
	ui.CreateSpeedBtn("save", icons.GetImageList(32, 32), line3)
	ui.CreateSeg(3, line3)
	ui.CreateSpeedBtn("folder", icons.GetImageList(32, 32), line3)
	this.tree = treeview.New(rightPanel)
	this.tree.OnCreate()
	this.tree.SetAlign(types.AlClient)
	this.tree.SetParent(rightPanel)
	this.tree.BorderSpacing().SetBottom(6)
	this.tree.BorderSpacing().SetRight(6)
	rightPanel.Constraints().SetMinWidth(400)
	rightPanel.SetWidth(400)
	this.valueEditor = ui.NewValueEditorFrame(leftPanel,ui.WithConfig(this.Config()))
	this.valueEditor.SetHeight(200)
	this.valueEditor.OnCreate()
	this.valueEditor.SetParent(leftPanel)
	this.valueEditor.SetAlign(types.AlBottom)
	ui.CreateSplitter(leftPanel, types.AlBottom, 6)
}

func (this *JsonEditor) OnCreate() {
	this.setup()

	this.tree.BuildMenuFunc = this.buildMenu
	this.tree.OnEditedFunc = this.onEdited
	this.tree.OnEditingFunc = this.onEditing
	this.tree.OnChangingFunc = this.onChanging
	this.tree.OnUpdateFunc = this.onUpdate
	this.valueEditor.OnSetSchema = this.onSetSchema
	this.textEdit.SetOnChange(func(sender vcl.IObject) {
		if !this.remodeling {
			log.Warn("textEdit SetOnChange parseString")
			this.tree.Clear()
			this.parseString()
		}
	})

}

func (this *JsonEditor) onSetSchema(s ui.ITreeSchema){
	this.editingSchema = nil
	this.assigning = true
	if s== nil {
		this.valueEditor.Clear()
		this.assigning = false
		return
	}
	log.Warnf("Type",s.(*Schema).Type.String())
	this.valueEditor.SetType(s.(*Schema).Type.String())
	if s.Idx() !=-1 {
		this.valueEditor.SetKey("[item"+strconv.Itoa(s.Idx())+"]")
	} else {
		this.valueEditor.SetKey(s.Key())
	}
	this.valueEditor.SetValue(s.(*Schema).ToString(this.valueEditor.Pretty()))
	this.editingSchema = s.(*Schema)
	this.assigning = false
}

func (this *JsonEditor) onChanging() {
	if this.editing == true {
		this.editing = false
		this.parseModel()
	}
}

func (this *JsonEditor) onEditing(schema ui.ITreeSchema) {
	this.editing = true
	if util.IsNil(schema.Parent()) {
		return
	}
	node:=schema.Node()
	if this.keyEditTag||schema.(*Schema).Type==Array||schema.(*Schema).Type==Object {
		if schema.Parent() != nil {
			node.SetText(schema.Key())
		}
	} else {
		node.SetText(schema.Value())
	}
}

func (this *JsonEditor) onEdited(schema ui.ITreeSchema, s *string) {
	if util.IsNil(schema.Parent()) {
		this.keyEditTag = false
		log.Warnf("onEdited",schema.Image(),schema.String())
		schema.UpdateNode()
		return
	}
	log.Warnf("start edit")
	if schema != nil {
		update := false
		if this.keyEditTag||schema.(*Schema).Type==Array||schema.(*Schema).Type==Object {
			update = schema.(*Schema).TrySetKey(*s)
		} else {
			update = schema.(*Schema).TrySetValue(*s)
		}
		log.Warnf(update,"edited",schema.Image())
		this.parseModel()
		this.tree.Tree.SetSelected(nil)
		schema.UpdateNode()
	}
	this.keyEditTag = false

}

func (this *JsonEditor) onUpdate(view *treeview.TreeView, parent *vcl.TTreeNode, schema ui.ITreeSchema) {
	switch schema.(*Schema).Type {
	case Object, Array:
		node := view.AddNode(parent, schema)
		for _, v := range schema.Children() {
			this.onUpdate(view, node, v)
		}
	case String:
		view.AddNode(parent, schema)
	case Number:
		view.AddNode(parent, schema)
	case Boolean:
		view.AddNode(parent, schema)
	case Null:
		view.AddNode(parent, schema)
	default:

	}
}

func (this *JsonEditor) OnDestroy() {
	this.tree.OnDestroy()
	this.textEdit.OnDestroy()
	this.valueEditor.OnDestroy()
}

func (this *JsonEditor) parseString() {
	text := this.textEdit.Text()
	i := orderedmap.New()
	err := json.Unmarshal([]byte(text), i)
	if err != nil {
		this.textEdit.Font().SetColor(colors.ClRed)
		return
	}
	this.textEdit.Font().SetColor(colors.ClSysDefault)
	this.schema = NewSchema()
	this.schema.Unmarshal("", -1, *i)
	this.tree.UpdateTree(this.schema)
}

func (this *JsonEditor) Clear() {
	this.tree.ClearTreeNodes()
	this.lastTreeItem = nil
}

func (this *JsonEditor) newDropMenuItem(img string, caption string, shortcut string, menu *vcl.TPopupMenu, f func(schema ui.ITreeSchema)) *vcl.TMenuItem {
	ret := vcl.NewMenuItem(menu)
	if img != "" {
		ret.SetImageIndex(icons.GetImageList(16, 16).GetImageIndex(img))
	}
	ret.SetCaption(caption)
	ret.SetShortCutFromString(shortcut)
	if f != nil {
		ret.SetOnClick(func(sender vcl.IObject) {
			f(this.tree.GetSelectSchema())
		})
	}

	return ret
}
func (this *JsonEditor) createNewFunc(t Type) func(schema ui.ITreeSchema) {
	return func(schema ui.ITreeSchema) {
		new := schema.Insert(t.CreateDefaultSchema())
		this.parseModel()
		this.tree.SetSelectSchema(new)
	}
}
func (this *JsonEditor) buildMenu(s ui.ITreeSchema) *vcl.TPopupMenu {
	types := []Type{Object, Array, String, Number, Boolean, Null}
	iconTypes := []string{"object_box", "array_box", "string_icon", "number_icon", "boolean_icon", "null_icon"}
	funcTypes := []func(schema ui.ITreeSchema){}
	for _, t := range types {
		funcTypes = append(funcTypes, this.createNewFunc(t))
	}

	ret := vcl.NewPopupMenu(this.tree)
	ret.SetImages(icons.GetImageList(16, 16).ImageList())

	newMenu := this.newDropMenuItem("", "新建[N]", "N", ret, nil)
	for idx, t := range types {
		newMenu.Add(this.newDropMenuItem(iconTypes[idx], t.String(), "", ret, funcTypes[idx]))
	}
	ret.Items().Add(newMenu)
	if !util.IsNil(s.Parent()) {
		ret.Items().Add(this.newDropMenuItem("", "修改键:"+s.Key(), "", ret, func(schema ui.ITreeSchema) {
			this.keyEditTag = true
			schema.Node().EditText()
		}))
	}

	if !util.IsNil(s.Parent())&&s.(*Schema).Type!=Object&&s.(*Schema).Type!=Array {
		ret.Items().Add(this.newDropMenuItem("", "修改值:"+s.Value(), "", ret, func(schema ui.ITreeSchema) {
			schema.Node().EditText()
		}))
	}
	ret.Items().Add(this.newDropMenuItem("", "复制[C]", "C", ret, func(schema ui.ITreeSchema) {
		this.ActionCopy()
	}))
	if this.clipSchema != nil {
		parseMenu := vcl.NewMenuItem(this)
		parseMenu.SetCaption("粘贴[V]")
		parseMenu.SetShortCutFromString("V")
		parseMenu.SetOnClick(func(sender vcl.IObject) {
			this.ActionParse()
		})
		ret.Items().Add(parseMenu)
	}
	ret.Items().Add(this.newDropMenuItem("", "剪切[X]", "X", ret, func(schema ui.ITreeSchema) {
		this.ActionCopy()
	}))
	ret.Items().Add(this.newDropMenuItem("", "删除[D]", "D", ret, func(schema ui.ITreeSchema) {
		this.ActionDel()
	}))
	ret.Items().Add(this.newDropMenuItem("", "收起", "", ret, func(schema ui.ITreeSchema) {
		this.ActionCollapse()
	}))
	return ret
}

func (this *JsonEditor) parseModel() {
	if this.tree.Root != nil {
		defer func() {
			if r := recover(); r != nil {
				if err,ok:=r.(error);ok {
					log.Error(err.Error())
					buf := make([]byte, 1<<16)
					runtime.Stack(buf, true)
					log.Error(string(buf))
				}
			}
		}()
		this.remodeling = true
		text := this.schema.ToString(true)
		go func() {
			time.Sleep(time.Millisecond*100)
			vcl.ThreadSync(func() {
				defer func() {
					if r := recover(); r != nil {
						if err,ok:=r.(error);ok {
							log.Error(err.Error())
							buf := make([]byte, 1<<16)
							runtime.Stack(buf, true)
							log.Error(string(buf))
						}
					}
				}()
				this.textEdit.Clear()
				this.textEdit.SetText(text)
				this.remodeling = false
			})
		}()
		this.tree.UpdateTree(nil)
		this.textEdit.Font().SetColor(colors.ClSysDefault)
	}
}

func (this *JsonEditor) ActionCopy() {
	if this.tree.GetSelectSchema() != nil {
		this.clipSchema = this.tree.GetSelectSchema().Clone().(*Schema)
	}
}

func (this *JsonEditor) ActionClip() {
	if this.tree.GetSelectSchema() != nil {
		this.clipSchema = this.tree.GetSelectSchema().Clone().(*Schema)
		if this.tree.GetSelectSchema().IsRoot() {
			this.Clear()
			this.schema = nil
			this.textEdit.Clear()
			this.tree.SetSelectSchema(nil)
			//TODO编辑器
			//this.jsonValueEditFrame.SetSchema(nil)
			return
		}
		this.tree.GetSelectSchema().(*Schema).DetachFromParent()
		this.tree.SetSelectSchema(nil)
		//TODO编辑器
		//this.jsonValueEditFrame.SetSchema(nil)
		this.parseModel()
	}
}

func (this *JsonEditor) ActionDel() {
	if this.tree.GetSelectSchema() != nil {
		if this.tree.GetSelectSchema().IsRoot() {
			this.Clear()
			this.schema = nil
			this.textEdit.Clear()
			this.tree.SetSelectSchema(nil)
			//TODO编辑器
			//this.jsonValueEditFrame.SetSchema(nil)
			return
		}
		this.tree.GetSelectSchema().(*Schema).DetachFromParent()
		this.tree.SetSelectSchema(nil)
		//TODO编辑器
		//this.jsonValueEditFrame.SetSchema(nil)
		this.parseModel()
	}
}

func (this *JsonEditor) ActionParse() {
	if this.tree.GetSelectSchema() != nil {
		if this.clipSchema != nil {
			s1 := this.tree.GetSelectSchema().Insert(this.clipSchema)
			this.parseModel()
			this.tree.SetSelectSchema(s1)
		}
	}
}

func (this *JsonEditor) ActionCollapse() {
	if this.tree.Tree.Selected() != nil {
		vcl.ThreadSync(func() {
			this.tree.Tree.Selected().Collapse(true)
		})
	} else {
		vcl.ThreadSync(func() {
			this.tree.Tree.FullCollapse()
		})
	}
}

func (this *JsonEditor) AddNewSchemaAtSelected(t Type) {
	if this.tree.GetSelectSchema() != nil {
		s := this.tree.GetSelectSchema().Insert(t.CreateDefaultSchema())
		this.parseModel()
		this.tree.SetSelectSchema(s)
	}
}
