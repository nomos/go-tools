package pjson

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/iancoleman/orderedmap"
	"github.com/nomos/go-lokas/util"
	"github.com/nomos/go-lokas/util/log"
	"math"
	"reflect"
	"strconv"
)

type Schema struct {
	Type      Type
	Index     int
	innerIndex int
	Key       string
	Value     string
	parent    *Schema
	Container []*Schema
}

func NewSchema() *Schema {
	return &Schema{
		Type:      0,
		Index:     -1,
		innerIndex: -1,
		Key:       "",
		Value:     "",
		parent:    nil,
		Container: make([]*Schema, 0),
	}
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
		this.Key = key
	}
	if index>=0 {
		this.Index = index
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
		this.Value = d.(string)
		break
	case Number:
		if math.IsNaN(d.(float64)) {
			this.Value = "NaN"
			break
		}
		this.Value = strconv.FormatFloat(d.(float64),'f', -1, 64)
	case Boolean:
		if d.(bool) {
			this.Value = "true"
		} else {
			this.Value = "false"
		}
		break
	case Null:
		this.Value = "null"
		break
	}
	return this,nil
}

func (this *Schema) AddChild(s *Schema)*Schema {
	s.parent = this
	this.Container = append(this.Container, s)
	s.innerIndex = len(this.Container)-1
	return s
}

func (this *Schema) AppendChild()*Schema {
	s:=NewSchema()
	s.parent = this
	this.Container = append(this.Container, s)
	s.innerIndex = len(this.Container)-1
	return s
}

func (this *Schema) Parent()*Schema {
	return this.parent
}

func (this *Schema) Root()*Schema {
	root:=this
	for {
		if root.Parent()==nil {
			return root
		}
		root = root.Parent()
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
		root = root.Parent()
	}
	ret1:=make([]int,0)
	for i:=len(ret)-1;i>=0;i-- {
		ret1 = append(ret1, ret[i])
	}
	return ret1
}

func (this *Schema) ToLineString()string {
	ret:=""
	if this.Index!= -1 {
		ret+="[Item"+strconv.Itoa(this.Index)+"]:   "
	}
	if this.Key!="" {
		ret+=`"`+this.Key+`":`
	}
	ret = util.AddStringGap(ret,10,6)

	switch this.Type {
	case Object:
		return ret+"object"
	case Array:
		return ret+"array"
	case String:
		return ret+`"`+this.Value+`"`
	default:
		return ret+this.Value
	}
}

func (this *Schema) ToObj()interface{} {
	switch this.Type {
	case Object:
		ret:=orderedmap.New()
		for _,container:=range this.Container {
			ret.Set(container.Key,container.ToObj())
		}
		return ret
	case Array:
		ret:=make([]interface{},0)
		for _,container:=range this.Container {
			ret = append(ret, container.ToObj())
		}
		return ret
	case String:
		return this.Value
	case Number:
		ret,err:= strconv.ParseFloat(this.Value,64)
		if err != nil {
			log.Error(err.Error())
			return nil
		}
		return ret
	case Boolean:
		if this.Value == "true" {
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
	if this.Index!= -1 {
		return "["+strconv.Itoa(this.Index)+"]"
	}
	return this.Key
}

func (this *Schema) ToValueString()string {
	switch this.Type {
	case Object:
		return "[object]"
	case Array:
		return "[array]"
	default:
		return this.Value
	}
}

func (this *Schema) Clone()*Schema {
	return this.clone(true)
}

func (this *Schema) clone(root bool)*Schema {
	schema:=NewSchema()
	schema.Type = this.Type
	schema.Value = this.Value
	schema.Key = this.Key
	schema.parent = nil
	if !root {
		schema.Index = schema.Index
		schema.innerIndex = schema.innerIndex
	} else {
		schema.Index = -1
		schema.innerIndex = -1
	}
	for _,s:=range this.Container{
		s1:=s.clone(false)
		schema.AddChild(s1)
	}
	return schema
}

func (this *Schema) IsRoot()bool {
	return this.parent== nil
}

func (this *Schema) IsObjectElem()bool {
	return !this.IsRoot()&&this.Index == -1
}

func (this *Schema) IsArrayElem()bool {
	return !this.IsRoot()&&this.Index != -1
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
		return this.Value
	}
}

func (this *Schema) Detach(s *Schema)*Schema {
	index:=-1
	newContainer:=make([]*Schema,0)
	for i,v:=range this.Container {
		if v== s {
			index = i
			continue
		}
		newContainer = append(newContainer, v)
		v.innerIndex = len(newContainer)-1
		if this.Type == Array {
			v.Index = v.innerIndex
		}
	}
	this.Container = newContainer
	if index!= -1 {
		this.parent = nil
		if this.Index!= -1 {
			this.Index = -1
		}
		this.innerIndex = -1
	}
	return this
}

func (this *Schema) Insert(s *Schema)*Schema {
	parent:=this
	index:=0
	if this.Type.IsValue() {
		parent = this.parent
		index = this.innerIndex+1
	}
	if parent.Type == Array {
		s.Key = ""
	}
	parent.insert(index,s)
	return s
}

func (this *Schema) insert(pos int,s *Schema) {
	newC := make([]*Schema,0)
	for _,v:=range this.Container[0:pos] {
		newC = append(newC, v)
		v.innerIndex = len(newC)-1
		if this.Type == Array {
			v.Index = v.innerIndex
		}
	}
	newC = append(newC,s)
	s.innerIndex = len(newC)-1
	if this.Type == Array {
		s.Index = s.innerIndex
	}
	for _,v:=range this.Container[pos:] {
		newC = append(newC, v)
		v.innerIndex = len(newC)-1
		if this.Type == Array {
			v.Index = v.innerIndex
		}
	}
	this.Container = newC
	s.parent = this
}

func (this *Schema) ChangeType(t Type)bool {
	if this.Type == t {
		return false
	}
	if this.Type == Object || this.Type == Array {
		if len(this.Container)>0 {
			return false
		}
	}
	this.Type = t
	this.Value = t.Default()
	return true
}

func (this *Schema) TrySetKey(s string) bool {
	if this.IsRoot()||(this.parent!=nil&&this.parent.Type==Array) {
		return false
	}
	this.Key = s
	return true
}

func (this *Schema) TrySetValue(s string) bool {
	switch this.Type {
	case Object,Array:
		return false
	default:
		s,ok:=this.Type.CheckValue(s)
		if ok {
			this.Value = s
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
	newContainer:=make([]*Schema,0)
	for _,v:=range this.Container[0:innerIndex-1] {
		newContainer = append(newContainer, v)
		v.innerIndex = len(newContainer) -1
		if this.Type == Array {
			v.Index = v.innerIndex
		}
	}
	newContainer = append(newContainer,s)
	s.innerIndex = len(newContainer) -1
	if this.Type == Array {
		s.Index = s.innerIndex
	}
	for _,v:=range this.Container[innerIndex-1:] {
		if v == s {
			continue
		}
		newContainer = append(newContainer, v)
		v.innerIndex = len(newContainer) -1
		if this.Type == Array {
			v.Index = v.innerIndex
		}
	}
	this.Container = newContainer
}

func (this *Schema) MoveDown()bool {
	if this.parent==nil {
		return false
	}
	if this.innerIndex == len(this.parent.Container)-1 {
		return false
	}
	this.parent.moveDown(this)
	return true
}

func (this *Schema) moveDown(s *Schema){
	innerIndex := s.innerIndex
	newContainer:=make([]*Schema,0)
	for _,v:=range this.Container[0:innerIndex+2] {
		if v == s {
			continue
		}
		newContainer = append(newContainer, v)
		v.innerIndex = len(newContainer) -1
		if this.Type == Array {
			v.Index = v.innerIndex
		}
	}
	newContainer = append(newContainer,s)
	s.innerIndex = len(newContainer) -1
	if this.Type == Array {
		s.Index = s.innerIndex
	}
	for _,v:=range this.Container[innerIndex+2:] {
		newContainer = append(newContainer, v)
		v.innerIndex = len(newContainer) -1
		if this.Type == Array {
			v.Index = v.innerIndex
		}
	}
	this.Container = newContainer
}