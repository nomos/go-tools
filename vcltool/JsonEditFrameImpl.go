// 由res2go自动生成。
// 在这里写你的事件。

package vcltool

import (
	"encoding/json"
	"github.com/iancoleman/orderedmap"
	"github.com/nomos/go-log/log"
	"github.com/nomos/go-lokas/util"
	"github.com/nomos/go-tools/pjson"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/keys"
	"io/ioutil"
	"sync"
	"time"
)

type JsonType int

const (
	JT_OBJECT JsonType = iota
	JT_ARRAY
	JT_STRING
	JT_NUMBER
	JT_BOOL
	JT_NULL
)

//::private::
type TJsonEditFrameFields struct {
	ConfigAble
	jsonFormat         interface{}
	jsonValueEditFrame *TJsonValueEditFrame
	lastTreeItem       *vcl.TTreeNode
	schema             *pjson.Schema
	selectedSchema     *pjson.Schema
	textFrame          *TTextFrame
	remodeling bool
	clipSchema   *pjson.Schema
	keyEditTag bool
	mutex sync.Mutex
}

func (this *TJsonEditFrame) OnCreate() {
	this.initTextFrame()
	this.initValueEditor()
	this.initTreeOps()
	this.initOpButtons()
	this.initFileActions()
}

func (this *TJsonEditFrame) OnDestroy() {
}

func (this *TJsonEditFrame) initTextFrame(){
	this.textFrame = NewTextFrame(this)
	this.textFrame.OnCreate()
	this.textFrame.SetParent(this.ParsePanel)
	this.textFrame.SetOnChange(func(sender vcl.IObject) {
		if !this.remodeling {
			this.parseString()
		}
	})
}

func (this *TJsonEditFrame) initValueEditor(){
	this.jsonValueEditFrame = NewJsonValueEditFrame(this)
	this.jsonValueEditFrame.conf =this.conf
	this.jsonValueEditFrame.OnCreate()
	this.jsonValueEditFrame.SetParent(this.EditPanel)
	this.jsonValueEditFrame.SetOnSchemaChange(func() {
		log.Warnf("SetOnSchemaChange",this.selectedSchema)
		this.parseModel()
		node:=this.GetNodeBySchema(this.selectedSchema)
		if node!=nil {
			node.SetSelected(true)
		}
	})
}

func (this *TJsonEditFrame) initTreeOps (){
	this.TreePanel.SetOnChanging(func(sender vcl.IObject, node *vcl.TTreeNode, allowChange *bool) {
		log.Warnf("changing", node.Text())
		if node.SelectedIndex() == 2 {

		}
	})
	this.TreePanel.SetOnEdited(func(sender vcl.IObject, node *vcl.TTreeNode, s *string) {
		if schema:=this.GetSchemaByNode(node);schema!=nil {
			log.Warnf("edited", node.Text())
			update:=false
			if this.keyEditTag {
				update=schema.TrySetKey(*s)
			} else {
				update=schema.TrySetValue(*s)
			}
			if update {
				this.parseModel()
			}
			this.keyEditTag = false
			go func() {
				time.Sleep(time.Millisecond*100)
				vcl.ThreadSync(func() {
					log.Warnf("SetOnEdited")
					this.SetSelectSchema(schema)
					node:=this.GetNodeBySchema(schema)
					if node!=nil {
						node.SetText(schema.ToLineString())
					}
				})
			}()
		}
	})
	this.TreePanel.SetOnEditing(func(sender vcl.IObject, node *vcl.TTreeNode, allowEdit *bool) {
		log.Warnf("editing", node.Text())
		if schema:=this.GetSchemaByNode(node);schema!=nil {
			if this.keyEditTag {
				if schema.Parent()!=nil&&schema.Parent().Type==pjson.Object {
					node.SetText(schema.Key)
				}
			} else {
				node.SetText(schema.Value)
			}
		}
	})
	this.TreePanel.SetOnChange(func(sender vcl.IObject, node *vcl.TTreeNode) {
		if this.schema == nil {
			return
		}
	})
	this.TreePanel.SetOnMouseDown(func(sender vcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
		if this.schema == nil {
			return
		}
		if button == types.MbRight {
			node := this.TreePanel.GetNodeAt(x, y)
			if node != nil {
				this.selectedSchema = this.GetSchemaByNode(node)
				node.SetSelected(true)
				this.jsonValueEditFrame.SetSchema(this.selectedSchema)
				menu :=this.buildPopMenuBySchema(this.selectedSchema)
				p := vcl.Mouse.CursorPos()
				menu.Popup(p.X,p.Y)
			}
		} else if button == types.MbLeft {
			node := this.TreePanel.GetNodeAt(x, y)
			if this.selectedSchema!= nil {
				node1:=this.GetNodeBySchema(this.selectedSchema)
				if node == nil {
					return
				}
				if node1 == node {
					return
				}
			}
			if node != nil {
				this.selectedSchema = this.GetSchemaByNode(node)
				node.SetText(this.selectedSchema.ToLineString())
				this.jsonValueEditFrame.SetSchema(this.selectedSchema)
			}
		}
	})
	this.TreePanel.SetOnKeyDown(func(sender vcl.IObject, key *types.Char, shift types.TShiftState) {
		if (util.IsMac()&&shift==128)||(!util.IsMac()&&shift==4) {
			switch *key {
			case keys.VkC:
				this.ActionCopy()
			case keys.VkV:
				this.ActionParse()
			case keys.VkX:
				this.ActionClip()
			case keys.VkD:
				this.ActionDel()
			}
		}

	})
}

func (this *TJsonEditFrame) initOpButtons() {
	this.Button_Bool.SetOnClick(func(sender vcl.IObject) {
		this.AddNewSchemaAtSelected(pjson.Boolean)
	})
	this.Button_Array.SetOnClick(func(sender vcl.IObject) {
		this.AddNewSchemaAtSelected(pjson.Array)
	})
	this.Button_Null.SetOnClick(func(sender vcl.IObject) {
		this.AddNewSchemaAtSelected(pjson.Null)
	})
	this.Button_Number.SetOnClick(func(sender vcl.IObject) {
		this.AddNewSchemaAtSelected(pjson.Number)
	})
	this.Button_Object.SetOnClick(func(sender vcl.IObject) {
		this.AddNewSchemaAtSelected(pjson.Object)
	})
	this.Button_String.SetOnClick(func(sender vcl.IObject) {
		this.AddNewSchemaAtSelected(pjson.String)
	})

	this.UpButton.SetOnClick(func(sender vcl.IObject) {
		if this.selectedSchema!= nil {
			update:=this.selectedSchema.MoveUp()
			if update {
				this.parseModel()
				this.SetSelectSchema(this.selectedSchema)
			}
		}
	})
	this.DownButton.SetOnClick(func(sender vcl.IObject) {
		if this.selectedSchema!= nil {
			update:=this.selectedSchema.MoveDown()
			if update {
				this.parseModel()
				this.SetSelectSchema(this.selectedSchema)
			}
		}
	})
}

func (this *TJsonEditFrame) initFileActions () {
	openDialog := vcl.NewOpenDialog(this)
	openDialog.SetOptions(openDialog.Options().Include(types.OfShowHelp, types.OfAllowMultiSelect)) //rtl.Include(, types.OfShowHelp))
	openDialog.SetFilter("Json文件(*.json)|*.json|所有文件(*.*)|*.*")
	openDialog.SetTitle("打开Json文件")
	this.OpenDirButton.SetOnClick(func(sender vcl.IObject) {
		if this.conf.GetString("load_file_path") != "" {
			openDialog.SetInitialDir(this.conf.GetString("load_file_path"))
		}
		if openDialog.Execute() {
			path := openDialog.FileName()
			this.conf.Set("load_file_path", path)
			log.Warnf(path)
			this.loadFile(path)
		}
	})
	saveDialog:=vcl.NewSaveDialog(this)
	saveDialog.SetOptions(saveDialog.Options().Include(types.OfShowHelp,types.OfAllowMultiSelect))
	saveDialog.SetDefaultExt(".json")
	saveDialog.SetTitle("保存到json文件")
	this.SaveToButton.SetOnClick(func(sender vcl.IObject) {
		if this.conf.GetString("load_file_path") != "" {
			saveDialog.SetInitialDir(this.conf.GetString("load_file_path"))
		}
		if saveDialog.Execute() {
			path := saveDialog.FileName()
			this.conf.Set("load_file_path", path)
			this.saveFile(path)
		}
	})
	this.SaveButton.SetOnClick(func(sender vcl.IObject) {
		this.saveFile(this.conf.GetString("load_file_path"))
	})
}

func (this *TJsonEditFrame) buildPopMenuBySchema(s *pjson.Schema) *vcl.TPopupMenu {
	ret := vcl.NewPopupMenu(this)
	newMenu:=vcl.NewMenuItem(this)
	newMenu.SetCaption("新建[N]")
	newMenu.SetShortCutFromString("N")
	ret.Items().Add(newMenu)
	ts:=s.AllTypes()
	for _,t:=range ts {
		newTMenu:=vcl.NewMenuItem(this)
		newTMenu.SetCaption(t.String())
		newTMenu.SetOnClick(func(sender vcl.IObject) {
			this.AddNewSchemaAt(s,pjson.GetTypeByString(vcl.AsMenuItem(sender).Caption()))
		})
		newMenu.Add(newTMenu)
	}
	keyMenu:=vcl.NewMenuItem(this)
	keyMenu.SetCaption("修改键:"+s.ToKeyString())
	keyMenu.SetOnClick(func(sender vcl.IObject) {
		this.keyEditTag = true
		this.GetNodeBySchema(s).EditText()
	})
	ret.Items().Add(keyMenu)
	valueMenu:=vcl.NewMenuItem(this)
	valueMenu.SetCaption("修改值:"+s.ToValueString())
	valueMenu.SetOnClick(func(sender vcl.IObject) {
		this.GetNodeBySchema(s).EditText()
	})
	ret.Items().Add(valueMenu)
	copyMenu:=vcl.NewMenuItem(this)
	copyMenu.SetCaption("复制[C]")
	copyMenu.SetShortCutFromString("C")
	copyMenu.SetOnClick(func(sender vcl.IObject) {
		this.ActionCopy()
	})
	ret.Items().Add(copyMenu)
	clipMenu:=vcl.NewMenuItem(this)
	clipMenu.SetCaption("剪切[X]")
	clipMenu.SetShortCutFromString("X")
	clipMenu.SetOnClick(func(sender vcl.IObject) {
		this.ActionClip()
	})
	ret.Items().Add(clipMenu)
	delMenu:=vcl.NewMenuItem(this)
	delMenu.SetCaption("删除[D]")
	delMenu.SetShortCutFromString("D")
	delMenu.SetOnClick(func(sender vcl.IObject) {
		this.ActionDel()
	})
	ret.Items().Add(delMenu)
	if this.clipSchema!=nil {
		parseMenu:=vcl.NewMenuItem(this)
		parseMenu.SetCaption("粘贴[V]")
		parseMenu.SetShortCutFromString("V")
		parseMenu.SetOnClick(func(sender vcl.IObject) {
			this.ActionParse()
		})
		ret.Items().Add(parseMenu)
	}

	changeTypeMenu:=vcl.NewMenuItem(this)
	changeTypeMenu.SetCaption("更改类型:"+s.Type.String())
	ret.Items().Add(changeTypeMenu)
	tsOther:=s.GetOtherTypes()
	for _,t:=range tsOther {
		modTMenu:=vcl.NewMenuItem(this)
		modTMenu.SetCaption(t.String())
		modTMenu.SetOnClick(func(sender vcl.IObject) {
			this.AddNewSchemaAt(s,pjson.GetTypeByString(vcl.AsMenuItem(sender).Caption()))
		})
		changeTypeMenu.Add(modTMenu)

	}
	return ret
}

func (this *TJsonEditFrame) saveFile (path string) {
	err:=ioutil.WriteFile(path,[]byte(this.schema.ToString(true)),0644)
	if err != nil {
		log.Errorf(err.Error())
	}
}

func (this *TJsonEditFrame) loadFile (path string) {
	data,err:=ioutil.ReadFile(path)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	this.textFrame.SetText(string(data))
}

func (this *TJsonEditFrame) parseModel() {
	if this.schema!=nil {
		this.remodeling = true
		text := this.schema.ToString(true)
		this.textFrame.Clear()
		this.textFrame.SetText(text)
		this.mutex.Lock()
		defer this.mutex.Unlock()
		this.TreePanel.Items().BeginUpdate()
		this.TreePanel.Items().Clear()
		this.ParseTree(nil, this.schema)
		this.TreePanel.FullExpand()
		this.TreePanel.Items().EndUpdate()
		this.remodeling = false
		this.textFrame.SetColorDefault()
	}

}

func (this *TJsonEditFrame) parseString() {
	text := this.textFrame.Text()
	i := orderedmap.New()
	err := json.Unmarshal([]byte(text), i)
	if err != nil {
		log.Warnf(err.Error())
		this.textFrame.SetColorRed()
		return
	}
	this.textFrame.SetColorDefault()
	this.schema = pjson.NewSchema()
	this.schema.Unmarshal("", -1, *i)
	this.Clear()
	this.TreePanel.Items().BeginUpdate()
	this.ParseTree(nil, this.schema)
	this.TreePanel.FullExpand()
	this.TreePanel.Items().EndUpdate()
}

func (this *TJsonEditFrame) SetSelectSchema(s *pjson.Schema){
	this.selectedSchema = s
	this.jsonValueEditFrame.SetSchema(s)
	node:=this.GetNodeBySchema(s)
	if node!=nil&&!node.Selected() {
		node.SetSelected(true)
	}
}

func (this *TJsonEditFrame) GetNodeBySchema(s *pjson.Schema)*vcl.TTreeNode {
	defer this.mutex.Unlock()
	this.mutex.Lock()
	if this.schema!=nil {
		rootTree:=s.GetRootTree()
		item:=this.TreePanel.TopItem()
		index:=0
		for {
			if index>len(rootTree)-1 {
				return item
			}
			if int32(rootTree[index])>item.Count()-1 {
				return nil
			}
			item = item.Item(int32(rootTree[index]))
			index++
		}
	}
	return nil
}

func (this *TJsonEditFrame) GetSchemaByNode(node *vcl.TTreeNode) *pjson.Schema {
	rootTree := make([]int, 0)
	for {
		if node.Parent() == nil {
			break
		}
		rootTree = append(rootTree, int(node.Index()))
		node = node.Parent()
	}
	schema := this.schema
	for i := len(rootTree) - 1; i >= 0; i-- {
		schema = schema.Container[rootTree[i]]
	}
	return schema
}

func (this *TJsonEditFrame) AddNewSchemaAtSelected(t pjson.Type){
	if this.selectedSchema!= nil {
		s:=this.selectedSchema.Insert(t.CreateDefaultSchema())
		this.parseModel()
		this.SetSelectSchema(s)
	}
}

func (this *TJsonEditFrame) AddNewSchemaAt(s *pjson.Schema,t pjson.Type){
	s1:=s.Insert(t.CreateDefaultSchema())
	this.parseModel()
	this.SetSelectSchema(s1)
}


func (this *TJsonEditFrame) ActionCopy(){
	if this.selectedSchema!=nil {
		this.clipSchema = this.selectedSchema.Clone()
	}
}

func (this *TJsonEditFrame) ActionClip(){
	if this.selectedSchema!=nil {
		this.clipSchema = this.selectedSchema
		if this.selectedSchema.IsRoot() {
			this.TreePanel.Items().Clear()
			this.schema = nil
			this.textFrame.Clear()
			this.selectedSchema = nil
			this.jsonValueEditFrame.SetSchema(nil)
			return
		}
		this.selectedSchema.DetachFromParent()
		this.selectedSchema = nil
		this.jsonValueEditFrame.SetSchema(nil)
		this.parseModel()
	}
}

func (this *TJsonEditFrame) ActionDel(){
	if this.selectedSchema!=nil {
		if this.selectedSchema.IsRoot() {
			this.TreePanel.Items().Clear()
			this.schema = nil
			this.textFrame.Clear()
			this.selectedSchema = nil
			this.jsonValueEditFrame.SetSchema(nil)
			return
		}
		this.selectedSchema.DetachFromParent()
		this.selectedSchema = nil
		this.jsonValueEditFrame.SetSchema(nil)
		this.parseModel()
	}
}

func (this *TJsonEditFrame) ActionParse(){
	if this.selectedSchema!=nil {
		if this.clipSchema!=nil {
			s1:=this.selectedSchema.Insert(this.clipSchema)
			this.parseModel()
			this.SetSelectSchema(s1)
		}
	}
}

func (this *TJsonEditFrame) ChangeType(s *pjson.Schema,t pjson.Type)  {
	success := this.schema.ChangeType(t)
	if success {
		this.parseModel()
		this.SetSelectSchema(s)
	}

}

func (this *TJsonEditFrame) RootNode(node *vcl.TTreeNode) *vcl.TTreeNode {
	for {
		if node.Parent() == nil {
			return node
		}
		log.Warnf("node", node.Index())
		node = node.Parent()
	}
}

func (this *TJsonEditFrame) ParseTree(parent *vcl.TTreeNode, schema *pjson.Schema) {
	switch schema.Type {
	case pjson.Object, pjson.Array:
		node := this.TreePanel.Items().AddChild(parent, schema.ToLineString())
		this.SetNodeImage(node, schema.Type)
		for _, v := range schema.Container {
			this.ParseTree(node, v)
		}
	case pjson.String:
		node := this.TreePanel.Items().AddChild(parent, schema.ToLineString())
		this.SetNodeImage(node, schema.Type)
	case pjson.Number:
		node := this.TreePanel.Items().AddChild(parent, schema.ToLineString())
		this.SetNodeImage(node, schema.Type)
	case pjson.Boolean:
		node := this.TreePanel.Items().AddChild(parent, schema.ToLineString())
		this.SetNodeImage(node, schema.Type)
	case pjson.Null:
		node := this.TreePanel.Items().AddChild(parent, schema.ToLineString())
		this.SetNodeImage(node, schema.Type)
	default:

	}
}

func (this *TJsonEditFrame) SetNodeImage(node *vcl.TTreeNode, t pjson.Type) {
	node.SetImageIndex(int32(t))
	node.SetStateIndex(int32(t))
	node.SetSelectedIndex(int32(t))
}

func (this *TJsonEditFrame) Clear() {
	this.TreePanel.Items().Clear()
	this.lastTreeItem = nil
}

func (this *TJsonEditFrame) OnEditPanelClick(sender vcl.IObject) {

}
