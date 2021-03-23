package vclui

import "github.com/ying32/govcl/vcl"

type Manager struct {
	owner vcl.IComponent
	Resource *TResource
}

func NewManager(owner vcl.IComponent)*Manager {
	ret :=&Manager{owner: owner}
	ret.init()
	return ret
}

func (this *Manager) init(){
	this.Resource = NewResource(this.owner)
	this.Resource.OnCreate()
}

func (this *Manager) NewLayout()*TLayout{
	return NewLayout(this.owner)
}



