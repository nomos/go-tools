package pjson

import (
	"encoding/json"
	"github.com/iancoleman/orderedmap"
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/util"
	"github.com/nomos/go-tools/ui"
	"github.com/nomos/go-tools/ui/icons"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
	"github.com/ying32/govcl/vcl/types/keys"
	"io/ioutil"
	"regexp"
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
	tree         *ui.TreeView

	clipSchema    *Schema
	schema        *Schema
	editingSchema *Schema

	keyEditTag   bool
	remodeling   bool
	editing      bool
	valueEditing bool
	assigning    bool
	mutex        sync.Mutex
}

func NewEditor(owner vcl.IComponent, option ...ui.FrameOption) (root *JsonEditor) {
	vcl.CreateResFrame(owner, &root)
	for _, o := range option {
		o(root)
	}
	return
}

func (this *JsonEditor) SetJsonString(s string){
	this.textEdit.SetText(s)
	this.parseString()
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
	line2 := ui.CreateLine(types.AlTop, 32, rightPanel)
	line2.BorderSpacing().SetBottom(3)
	line3 := ui.CreateLine(types.AlTop, 32, leftPanel)
	line3.BorderSpacing().SetLeft(6)
	line3.BorderSpacing().SetBottom(3)
	ui.CreateSeg(3, line2)
	ui.CreateSpeedBtn("sort_up", icons.GetImageList(32, 32), line2).SetOnClick(func(sender vcl.IObject) {
		schema:=this.tree.GetSelectSchema()
		if !util.IsNil(schema) {
			s:=schema.(*Schema)
			update:=s.MoveUp()
			if update {
				this.parseModel()
				this.tree.SetSelectSchema(s)
			}
		}
	})
	ui.CreateSeg(3, line2)
	ui.CreateSpeedBtn("sort_down", icons.GetImageList(32, 32), line2).SetOnClick(func(sender vcl.IObject) {
		schema:=this.tree.GetSelectSchema()
		if !util.IsNil(schema) {
			s:=schema.(*Schema)
			update:=s.MoveDown()
			if update {
				this.parseModel()
				this.tree.SetSelectSchema(s)
			}
		}
	})
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



	ui.CreateSpeedBtn("cancel", icons.GetImageList(32, 32), line3).SetOnClick(func(sender vcl.IObject) {
		this.Clear()
	})
	ui.CreateSeg(3, line3)

	saveDialog:=vcl.NewSaveDialog(this)
	saveDialog.SetOptions(saveDialog.Options().Include(types.OfShowHelp,types.OfAllowMultiSelect))
	saveDialog.SetDefaultExt(".json")
	saveDialog.SetTitle("保存到json文件")
	ui.CreateSpeedBtn("save", icons.GetImageList(32, 32), line3).SetOnClick(func(sender vcl.IObject) {
		if this.Config().GetString("load_file_path") != "" {
			saveDialog.SetInitialDir(this.Config().GetString("load_file_path"))
		}
		if saveDialog.Execute() {
			path := saveDialog.FileName()
			this.Config().Set("load_file_path", path)
			err:=ioutil.WriteFile(path,[]byte(this.schema.ToString(true)),0644)
			if err != nil {
				log.Errorf(err.Error())
			}
		}
	})
	ui.CreateSeg(3, line3)

	openDialog := vcl.NewOpenDialog(this)
	openDialog.SetOptions(openDialog.Options().Include(types.OfShowHelp, types.OfAllowMultiSelect)) //rtl.Include(, types.OfShowHelp))
	openDialog.SetFilter("Json文件(*.json)|*.json|所有文件(*.*)|*.*")
	openDialog.SetTitle("打开Json文件")
	ui.CreateSpeedBtn("folder", icons.GetImageList(32, 32), line3).SetOnClick(func(sender vcl.IObject) {
		if this.Config().GetString("load_file_path") != "" {
			openDialog.SetInitialDir(this.Config().GetString("load_file_path"))
		}
		if openDialog.Execute() {
			path := openDialog.FileName()
			this.Config().Set("load_file_path", path)
			log.Warnf(path)
			data,err:=ioutil.ReadFile(path)
			if err != nil {
				log.Errorf(err.Error())
				return
			}
			this.textEdit.SetText(string(data))
		}
	})
	this.tree = ui.NewTreeView(rightPanel)
	this.tree.OnCreate()
	this.tree.SetAlign(types.AlClient)
	this.tree.SetParent(rightPanel)
	this.tree.BorderSpacing().SetBottom(6)
	this.tree.BorderSpacing().SetRight(6)
	rightPanel.Constraints().SetMinWidth(400)
	rightPanel.SetWidth(400)
	this.valueEditor = ui.NewValueEditorFrame(leftPanel, ui.WithConfig(this.Config()))
	this.valueEditor.SetHeight(200)
	this.valueEditor.OnCreate()
	this.valueEditor.SetParent(leftPanel)
	this.valueEditor.SetAlign(types.AlBottom)
	for _, t := range Types {
		this.valueEditor.AddType(t.String())
	}
	ui.CreateSplitter(leftPanel, types.AlBottom, 6)
}

func (this *JsonEditor) OnCreate() {
	this.setup()

	this.tree.BuildMenuFunc = this.buildMenu
	this.tree.OnEditedFunc = this.onEdited
	this.tree.OnEditingFunc = this.onEditing
	this.tree.OnChangingFunc = this.onChanging
	this.tree.OnUpdateFunc = this.onUpdate
	this.tree.OnSelectFunc = this.onSelect
	this.valueEditor.OnSetSchema = this.onEditorSetSchema
	this.valueEditor.OnKeyChange = this.onEditorKeyChange
	this.valueEditor.OnValueChange = this.onEditorValueChange
	this.valueEditor.OnKeyKeyDown = this.onEditorKeyKeyDown
	this.valueEditor.OnValueKeyDown = this.onEditorValueKeyDown
	this.valueEditor.OnSelect = this.onEditorSelect
	this.valueEditor.OnKeyExit = this.onEditorKeyExit
	this.valueEditor.OnValueExit = this.onEditorValueExit

	this.textEdit.SetOnChange(func(sender vcl.IObject) {
		if !this.remodeling && !this.valueEditing {
			this.tree.Clear()
			this.parseString()
		}
	})
}

func (this *JsonEditor) setEditorKey(s string) {
	vcl.ThreadSync(func() {
		this.valueEditing = true
		this.valueEditor.SetKey(s)
		this.valueEditing = false
	})
}

func (this *JsonEditor) setEditorValue(s string) {
	vcl.ThreadSync(func() {
		this.valueEditing = true
		this.valueEditor.SetValue(s)
		this.valueEditing = false
	})

}

func (this *JsonEditor) onEditorKeyChange(schema ui.ITreeSchema) {
	if this.valueEditing {
		return
	}
	if schema == nil {
		if !this.assigning {
			vcl.ThreadSync(func() {
				this.setEditorKey("")
			})
		}
		return
	}
	if schema.IsRoot() {
		vcl.ThreadSync(func() {
			this.setEditorKey("")
		})
		return
	}
	if schema.(*Schema).IsArrayElem() {
		vcl.ThreadSync(func() {
			this.setEditorKey(schema.(*Schema).ToKeyString())
		})
		return
	}
	if schema.(*Schema).IsObjectElem() {
		schema.SetKey(this.valueEditor.Key())
		schema.(*Schema).key = this.valueEditor.Key()
		schema.UpdateNode()
		this.parseModelTextOnly()
		return
	}
	schema.SetKey(this.valueEditor.Key())
	schema.UpdateNode()
	this.parseModelTextOnly()
}

func (this *JsonEditor) onEditorValueChange(schema ui.ITreeSchema) {
	if this.valueEditing {
		return
	}
	if schema == nil {
		return
	}
	switch schema.(*Schema).Type {
	case String:
		schema.SetValue(this.valueEditor.Value())
		schema.UpdateNode()
		this.parseModelTextOnly()
		this.valueEditing = false
	case Number:
		if s, ok := Number.CheckValue(this.valueEditor.Value()); ok {
			schema.SetValue(s)
			schema.UpdateNode()
			this.parseModelTextOnly()
			return
		}
		this.valueEditing = false
	case Boolean:
		if s, ok := Boolean.CheckValue(this.valueEditor.Value()); ok {
			schema.SetValue(s)
			this.valueEditor.SetSchema(schema)
			schema.UpdateNode()
			this.parseModelTextOnly()
			return
		}
		this.valueEditing = false
	default:
		this.valueEditing = false
	}
}

func (this *JsonEditor) onEditorKeyKeyDown(key types.Char, shift types.TShiftState, schema ui.ITreeSchema) {
	if key == keys.VkReturn {

	}
}

func (this *JsonEditor) onEditorValueKeyDown(key types.Char, shift types.TShiftState, schema ui.ITreeSchema) {
	if key == keys.VkReturn {

	}
}

func (this *JsonEditor) onEditorSelect(schema ui.ITreeSchema, s string) {
	if schema == nil {
		return
	}
	t := GetTypeByString(this.valueEditor.GetType())
	if t == -1 {
		this.valueEditor.SetType(schema.(*Schema).Type.String())
	}
	success := schema.(*Schema).ChangeType(t)
	schema.UpdateNode()
	if success {
		this.valueEditor.SetSchema(schema)
	}
}

func (this *JsonEditor) onEditorKeyExit(schema ui.ITreeSchema) {


}

func (this *JsonEditor) onEditorValueExit(schema ui.ITreeSchema) {
	defer func() {
		if r := recover(); r != nil {
			util.Recover(r,false)
		}
	}()
	if schema == nil {
		log.Warnf("nil schema")
		this.setEditorValue("")
		return
	}
	switch schema.(*Schema).Type {
	case Object, Array:
		this.setEditorValue(schema.(*Schema).ToString(this.valueEditor.Pretty()))
		return
	default:
		s, ok := schema.(*Schema).Type.CheckValue(this.valueEditor.Value())
		log.Warnf("s ok", ok, s)
		if ok {
			this.setEditorValue(schema.Value())
			schema.(*Schema).value = s
			schema.UpdateNode()
			this.parseModelTextOnly()
		} else {
			this.setEditorValue(schema.Value())
		}
	}
}

func (this *JsonEditor) onEditorSetSchema(s ui.ITreeSchema) {
	defer func() {
		if r := recover(); r != nil {
			util.Recover(r,false)
		}
	}()
	this.editingSchema = nil
	this.assigning = true
	if util.IsNil(s) {
		this.valueEditor.Clear()
		this.assigning = false
		return
	}
	this.valueEditor.SetType(s.(*Schema).Type.String())
	if s.Idx() != -1 {
		this.setEditorKey("[item" + strconv.Itoa(s.Idx()) + "]")
	} else {
		this.setEditorKey(s.Key())
	}
	this.setEditorValue(s.(*Schema).ToString(this.valueEditor.Pretty()))
	this.editingSchema = s.(*Schema)
	this.assigning = false
}

func (this *JsonEditor) onSelect(schema ui.ITreeSchema) {
	this.valueEditor.SetSchema(schema)
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
	node := schema.Node()
	if node == nil {
		return
	}
	if this.keyEditTag || schema.(*Schema).Type == Array || schema.(*Schema).Type == Object {
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
		schema.UpdateNode()
		return
	}
	if schema != nil {
		update := false
		if this.keyEditTag || schema.(*Schema).Type == Array || schema.(*Schema).Type == Object {
			update = schema.(*Schema).TrySetKey(*s)
		} else {
			update = schema.(*Schema).TrySetValue(*s)
		}
		this.tree.Tree.SetSelected(nil)
		if update {
			this.parseModel()
			schema.UpdateNode()
		}
	}
	this.keyEditTag = false

}

func (this *JsonEditor) onUpdate(view *ui.TreeView, parent *vcl.TTreeNode, schema ui.ITreeSchema) {
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
	if text=="" {
		return
	}
	if regexp.MustCompile(`^\[`).MatchString(text) {
		log.Warnf("parseArray")
		this.parseArrayRoot(text)
	} else {
		log.Warnf("parseObject")
		this.parseObjectRoot(text)
	}
}

func (this *JsonEditor) parseArrayRoot(text string){
	i := []interface{}{}
	err := json.Unmarshal([]byte(text), &i)
	if err != nil {
		this.textEdit.Font().SetColor(colors.ClRed)
		return
	}
	this.textEdit.Font().SetColor(colors.ClSysDefault)
	this.schema = NewSchema()
	this.schema.Unmarshal("", -1, i)
	this.tree.UpdateTree(this.schema)
}

func (this *JsonEditor) parseObjectRoot(text string){
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

func (this *JsonEditor) Clear(){
	this.schema = nil
	defer func() {
		if r := recover(); r != nil {
			util.Recover(r,false)
		}
	}()
	vcl.ThreadSync(func() {
		this.tree.Clear()
		this.textEdit.Memo.Clear()
		this.valueEditing = true
		this.valueEditor.Clear()
		this.valueEditing = false
	})
}

func (this *JsonEditor) ClearTree() {
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
	iconTypes := []string{"object_box", "array_box", "string_icon", "number_icon", "boolean_icon", "null_icon"}
	funcTypes := []func(schema ui.ITreeSchema){}
	for _, t := range Types {
		funcTypes = append(funcTypes, this.createNewFunc(t))
	}

	ret := vcl.NewPopupMenu(this.tree)
	ret.SetImages(icons.GetImageList(16, 16).ImageList())

	newMenu := this.newDropMenuItem("", "新建[N]", "N", ret, nil)
	for idx, t := range Types {
		newMenu.Add(this.newDropMenuItem(iconTypes[idx], t.String(), "", ret, funcTypes[idx]))
	}
	ret.Items().Add(newMenu)
	if !util.IsNil(s.Parent()) {
		ret.Items().Add(this.newDropMenuItem("", "修改键:"+s.Key(), "", ret, func(schema ui.ITreeSchema) {
			this.keyEditTag = true
			schema.Node().EditText()
		}))
	}

	if !util.IsNil(s.Parent()) && s.(*Schema).Type != Object && s.(*Schema).Type != Array {
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

func (this *JsonEditor) parseModelTextOnly() {
	defer func() {
		if r := recover(); r != nil {
			util.Recover(r,false)
		}
	}()
	if this.tree.Root != nil {
		text := this.schema.ToString(true)
		go func() {
			time.Sleep(time.Millisecond*100)
			vcl.ThreadSync(func() {
				defer func() {
					if r := recover(); r != nil {
						util.Recover(r,false)
					}
				}()
				this.remodeling = true
				this.textEdit.SetText(text)
				this.remodeling = false
			})
		}()
		this.textEdit.Font().SetColor(colors.ClSysDefault)
	}
}

func (this *JsonEditor) parseModel() {
	if this.tree.Root != nil {
		defer func() {
			if r := recover(); r != nil {
				util.Recover(r,false)
			}
		}()
		text := this.schema.ToString(true)
		go func() {
			time.Sleep(time.Millisecond * 100)
			vcl.ThreadSync(func() {
				defer func() {
					if r := recover(); r != nil {
						util.Recover(r,false)
					}
				}()
				this.remodeling = true
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
			this.ClearTree()
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
			this.ClearTree()
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

func (this *JsonEditor) OnEnter() {
	panic("implement me")
}

func (this *JsonEditor) OnExit() {
	panic("implement me")
}
