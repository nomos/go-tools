package treeview

import (
	"github.com/nomos/go-tools/ui"
	"github.com/nomos/go-tools/ui/icons"
	"github.com/ying32/govcl/vcl"
)

type DropMenuItem struct {
	Menu *vcl.TMenuItem
	Parent *DropMenuItem
	SubMenus []*DropMenuItem
	Func func(schema ui.ITreeSchema)
}

func NewDropMenuItem (img string,caption string,shortcut string,view *TreeView,f func(schema ui.ITreeSchema))*DropMenuItem{
	ret:=&DropMenuItem{
		Menu:     vcl.NewMenuItem(view),
		SubMenus: []*DropMenuItem{},
		Func:     f,
	}
	if img!= "" {
		ret.Menu.SetImageIndex(icons.GetImageList(16,16).GetImageIndex(img))
	}
	ret.Menu.SetCaption(caption)
	ret.Menu.SetShortCutFromString(shortcut)
	return ret
}

func (this *DropMenuItem) SetParent(parent *DropMenuItem)*DropMenuItem{
	this.Parent = parent
	return this
}

func (this *DropMenuItem) AddDropMenu(img string,caption string,shortcut string,view *TreeView,f func(schema ui.ITreeSchema)){
	m:=NewDropMenuItem(img,caption,shortcut,view,f).SetParent(this)
	this.SubMenus = append(this.SubMenus,m)
	this.Menu.Add(m.Menu)
}
