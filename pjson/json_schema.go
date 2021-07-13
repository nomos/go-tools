package pjson

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/iancoleman/orderedmap"
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/util/stringutil"
	"github.com/nomos/go-tools/ui"
	"github.com/ying32/govcl/vcl"
	"math"
	"reflect"
	"strconv"
)

var _ ui.ITreeSchema = (*Schema)(nil)

type Schema struct {
	Type       Type
	index      int
	innerIndex int
	key        string
	value      string
	parent     *Schema
	children   []ui.ITreeSchema
	collapse bool

	node *vcl.TTreeNode
	tree ui.ITree
}

func (this *Schema) SetNode(node *vcl.TTreeNode){
	this.node = node
}

func (this *Schema) Node()*vcl.TTreeNode{
	return this.node
}

func (this *Schema) Collapse() bool {
	return this.collapse
}

func (this *Schema) SetCollapse(b bool) {
	if this.collapse!= b {
		this.tree.UpdateTree(this.Root())
	}
}

func NewSchema() *Schema {
	return &Schema{
		Type:       0,
		index:      -1,
		innerIndex: -1,
		key:        "",
		value:      "",
		parent:     nil,
		children:   make([]ui.ITreeSchema, 0),
	}
}

func (this *Schema) SetKey(s string) {
	this.key = s
}

func (this *Schema) Key()string {
	return this.key
}

func (this *Schema) Value()string {
	return this.value
}

func (this *Schema) Idx()int{
	return this.index
}

func (this *Schema) SetIdx(id int){
	this.index = id
}

func (this *Schema) InnerIdx()int{
	return this.innerIndex
}

func (this *Schema) SetInnerIdx(id int){
	this.innerIndex = id
}

func (this *Schema) Parent()ui.ITreeSchema {
	return this.parent
}

func (this *Schema) Children()[]ui.ITreeSchema {
	return this.children
}

func (this *Schema) Image()string{
	switch this.Type {
	case Null:
		return "null_icon"
	case Number:
		return "number_icon"
	case Boolean:
		return "boolean_icon"
	case String:
		return "string_icon"
	case Object:
		return "object_box"
	case Array:
		return "array_box"
	}
	return ""
}

func (this *Schema) typeof(d interface{}) Type {
	if d == nil {
		return Null
	}
	switch reflect.TypeOf(d).Kind().String() {
	case "ptr","struct":
		return Object
	case "map":
		return Object
	case "slice":
		return Array
	case "string":
		return String
	case "float64":
		return Number
	case "bool":
		return Boolean
	}
	return Null
}

func (this *Schema) Unmarshal(key string,index int, d interface{}) (*Schema, error) {
	t := this.typeof(d)
	if this.parent == nil && t != Object && t != Array {
		return nil, errors.New("parent cant be " + t.String())
	}
	if key!="" {
		this.key = key
	}
	if index>=0 {
		this.index = index
	}
	this.Type = t
	switch t {
	case Object:
		d1:=d.(orderedmap.OrderedMap)
		for _,k:=range d1.Keys() {
			schema:=this.AppendChild()
			v,ok:=d1.Get(k)
			if ok {
				schema.Unmarshal(k,-1,v)
			}
		}
		break
	case Array:
		d1 := d.([]interface{})
		for k, v := range d1 {
			schema:=this.AppendChild()
			schema.Unmarshal("",k,v)
		}
		break
	case String:
		this.value = d.(string)
		break
	case Number:
		if math.IsNaN(d.(float64)) {
			this.value = "NaN"
			break
		}
		this.value = strconv.FormatFloat(d.(float64),'f', -1, 64)
	case Boolean:
		if d.(bool) {
			this.value = "true"
		} else {
			this.value = "false"
		}
		break
	case Null:
		this.value = "null"
		break
	}
	return this,nil
}

func (this *Schema) AddChild(s ui.ITreeSchema)ui.ITreeSchema {
	child:=s.(*Schema)
	child.parent = this
	this.children = append(this.children, s)
	child.innerIndex = len(this.children)-1
	return s
}

func (this *Schema) AppendChild()*Schema {
	s:=NewSchema()
	s.parent = this
	this.children = append(this.children, s)
	s.innerIndex = len(this.children)-1
	return s
}

func (this *Schema) Root()ui.ITreeSchema {
	root:=this
	for {
		if root.Parent()==nil {
			return root
		}
		root = root.parent
	}
}

func (this *Schema) DetachFromParent()*Schema {
	this.parent.Detach(this)
	return this
}

func (this *Schema) GetRootTree()[]int {
	ret:=make([]int,0)
	root:=this
	for {
		if root.Parent()==nil {
			break
		}
		if root.innerIndex== -1 {
			return make([]int,0)
		}
		ret = append(ret, root.innerIndex)
		root = root.parent
	}
	ret1:=make([]int,0)
	for i:=len(ret)-1;i>=0;i-- {
		ret1 = append(ret1, ret[i])
	}
	return ret1
}

func (this *Schema) String()string {
	ret:=""
	if this.index != -1 {
		ret+="[Item"+strconv.Itoa(this.index)+"]:   "
	}
	if this.key !="" {
		ret+=`"`+this.key +`":`
	}
	ret = stringutil.AddStringGap(ret,10,6)

	switch this.Type {
	case Object:
		return ret+"object"
	case Array:
		return ret+"array"
	case String:
		return ret+`"`+this.value +`"`
	default:
		return ret+this.value
	}
}

func (this *Schema) ToObj()interface{} {
	switch this.Type {
	case Object:
		ret:=orderedmap.New()
		for _,container:=range this.children {
			ret.Set(container.Key(),container.ToObj())
		}
		return ret
	case Array:
		ret:=make([]interface{},0)
		for _,container:=range this.children {
			ret = append(ret, container.ToObj())
		}
		return ret
	case String:
		return this.value
	case Number:
		ret,err:= strconv.ParseFloat(this.value,64)
		if err != nil {
			log.Error(err.Error())
			return nil
		}
		return ret
	case Boolean:
		if this.value == "true" {
			return true
		} else {
			return false
		}
	case Null:
		return nil
	}
	return nil
}

func (this *Schema) ToKeyString()string {
	if this.index != -1 {
		return "["+strconv.Itoa(this.index)+"]"
	}
	return this.key
}

func (this *Schema) ToValueString()string {
	switch this.Type {
	case Object:
		return "[object]"
	case Array:
		return "[array]"
	default:
		return this.value
	}
}

func (this *Schema) Clone()*Schema {
	return this.clone(true)
}

func (this *Schema) clone(root bool)*Schema {
	schema:=NewSchema()
	schema.Type = this.Type
	schema.value = this.value
	schema.key = this.key
	schema.parent = nil
	if !root {
		schema.index = schema.index
		schema.innerIndex = schema.innerIndex
	} else {
		schema.index = -1
		schema.innerIndex = -1
	}
	for _,s:=range this.children {
		s1:=s.(*Schema).clone(false)
		schema.AddChild(s1)
	}
	return schema
}

func (this *Schema) IsRoot()bool {
	return this.parent== nil
}

func (this *Schema) IsObjectElem()bool {
	return !this.IsRoot()&&this.index == -1
}

func (this *Schema) IsArrayElem()bool {
	return !this.IsRoot()&&this.index != -1
}

func (this *Schema) ToString(format bool)string {
	switch this.Type {
	case Object,Array:
		ret,err:=json.Marshal(this.ToObj())
		if err != nil {
			log.Error(err.Error())
			return ""
		}
		if format {
			var str bytes.Buffer
			_ = json.Indent(&str, ret, "", "    ")
			return str.String()
		} else {
			return string(ret)
		}
	default:
		return this.value
	}
}

func (this *Schema) Detach(s ui.ITreeSchema)ui.ITreeSchema {
	index:=-1
	newContainer:=make([]ui.ITreeSchema,0)
	for i,v:=range this.children {
		if v== s {
			index = i
			continue
		}
		newContainer = append(newContainer, v)
		v.SetInnerIdx(len(newContainer)-1)
		if this.Type == Array {
			v.SetIdx(v.InnerIdx())
		}
	}
	this.children = newContainer
	if index!= -1 {
		this.parent = nil
		if this.index != -1 {
			this.index = -1
		}
		this.innerIndex = -1
	}
	return this
}

func (this *Schema) Insert(s ui.ITreeSchema)ui.ITreeSchema {
	parent:=this
	index:=0
	if this.Type.IsValue() {
		parent = this.parent
		index = this.innerIndex+1
	}
	if parent.Type == Array {
		s.SetKey("")
	}
	parent.insert(index,s.(*Schema))
	return s
}

func (this *Schema) insert(pos int,s *Schema) {
	newC := make([]ui.ITreeSchema,0)
	for _,v:=range this.children[0:pos] {
		newC = append(newC, v)
		v.SetInnerIdx(len(newC)-1)
		if this.Type == Array {
			v.SetIdx(v.InnerIdx())
		}
	}
	newC = append(newC,s)
	s.innerIndex = len(newC)-1
	if this.Type == Array {
		s.index = s.innerIndex
	}
	for _,v:=range this.children[pos:] {
		newC = append(newC, v)
		v.SetInnerIdx(len(newC)-1)
		if this.Type == Array {
			v.SetIdx(v.InnerIdx())
		}
	}
	this.children = newC
	s.parent = this
}

func (this *Schema) ChangeType(t Type)bool {
	if this.Type == t {
		return false
	}
	if this.Type == Object || this.Type == Array {
		if len(this.children)>0 {
			return false
		}
	}
	this.Type = t
	this.value = t.Default()
	return true
}

func (this *Schema) TrySetKey(s string) bool {
	if this.IsRoot()||(this.parent!=nil&&this.parent.Type==Array) {
		return false
	}
	this.key = s
	return true
}

func (this *Schema) TrySetValue(s string) bool {
	switch this.Type {
	case Object,Array:
		return false
	default:
		s,ok:=this.Type.CheckValue(s)
		if ok {
			this.value = s
			return true
		}
	}
	return false
}

func (this *Schema) MoveUp()bool {
	if this.parent==nil {
		return false
	}
	if this.innerIndex == 0 {
		return false
	}
	this.parent.moveUp(this)
	return true
}

func (this *Schema) moveUp(s *Schema){
	innerIndex := s.innerIndex
	newContainer:=make([]ui.ITreeSchema,0)
	for _,v:=range this.children[0:innerIndex-1] {
		newContainer = append(newContainer, v)
		v.SetInnerIdx(len(newContainer) -1)
		if this.Type == Array {
			v.SetIdx(v.InnerIdx())
		}
	}
	newContainer = append(newContainer,s)
	s.innerIndex = len(newContainer) -1
	if this.Type == Array {
		s.index = s.innerIndex
	}
	for _,v:=range this.children[innerIndex-1:] {
		if v == s {
			continue
		}
		newContainer = append(newContainer, v)
		v.SetInnerIdx(len(newContainer) -1)
		if this.Type == Array {
			v.SetIdx(v.InnerIdx())
		}
	}
	this.children = newContainer
}

func (this *Schema) MoveDown()bool {
	if this.parent==nil {
		return false
	}
	if this.innerIndex == len(this.parent.children)-1 {
		return false
	}
	this.parent.moveDown(this)
	return true
}

func (this *Schema) moveDown(s *Schema){
	innerIndex := s.innerIndex
	newContainer:=make([]ui.ITreeSchema,0)
	for _,v:=range this.children[0:innerIndex+2] {
		if v == s {
			continue
		}
		newContainer = append(newContainer, v)
		v.SetInnerIdx(len(newContainer) -1)
		if this.Type == Array {
			v.SetIdx(v.InnerIdx())
		}
	}
	newContainer = append(newContainer,s)
	s.innerIndex = len(newContainer) -1
	if this.Type == Array {
		s.index = s.innerIndex
	}
	for _,v:=range this.children[innerIndex+2:] {
		newContainer = append(newContainer, v)
		v.SetInnerIdx(len(newContainer) -1)
		if this.Type == Array {
			v.SetIdx(v.InnerIdx())
		}
	}
	this.children = newContainer
}