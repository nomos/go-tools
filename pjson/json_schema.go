package pjson

import (
	"errors"
	"github.com/nomos/go-lokas/util"
	"math"
	"reflect"
	"strconv"
)

type Schema struct {
	Type      Type
	Name      string
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
		Name:      "",
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
		d1 := d.(map[string]interface{})
		for k, v := range d1 {
			schema:=this.AppendChild()
			schema.Unmarshal(k,-1,v)
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
	return s
}

func (this *Schema) AppendChild()*Schema {
	s:=NewSchema()
	s.parent = this
	this.Container = append(this.Container, s)
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
		ret = append(ret, root.innerIndex)
		root = root.Parent()
	}
	ret1:=make([]int,0)
	for i:=len(ret)-1;i>=0;i-- {
		ret1 = append(ret1, ret[i])
	}
	return ret
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

func (this *Schema) Detach(s *Schema)*Schema {
	index:=-1
	newContainer:=make([]*Schema,0)
	for i,v:=range this.Container {
		if v== s {
			index = i
			continue
		}
		newContainer = append(newContainer, v)
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