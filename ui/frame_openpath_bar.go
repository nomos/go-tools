package ui

import (
	"github.com/nomos/go-lokas/protocol"
	"github.com/nomos/go-tools/ui/icons"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
)

var _ IFrame = (*OpenPathBar)(nil)

type OPEN_TYPE protocol.Enum

const (
	OPEN_FILE OPEN_TYPE = iota
	OPEN_DIR  OPEN_TYPE = iota
	SAVE_FILE OPEN_TYPE = iota
)

type OpenPathBar struct {
	*vcl.TFrame
	ConfigAble

	label          *vcl.TLabel
	btn            *vcl.TSpeedButton
	name string
	width int32
	edit           *vcl.TEdit
	path           string
	openDirDialog  *vcl.TSelectDirectoryDialog
	openFileDialog *vcl.TOpenDialog
	saveDialog     *vcl.TSaveDialog
	openType       OPEN_TYPE
	color types.TColor
	OnOpen         func(string)
	OnEdit         func(string)
}

func WithOpenDirDialog(name string) FrameOption {
	return func(frame IFrame) {
		f := frame.(*OpenPathBar)
		f.SetOpenDirDialog(name)
	}
}

func WithOpenFileDialog(name string, filter string) FrameOption {
	return func(frame IFrame) {
		f := frame.(*OpenPathBar)
		f.SetOpenFileDialog(name, filter)
	}
}

func WithSaveDialog(name string, ext string) FrameOption {
	return func(frame IFrame) {
		f := frame.(*OpenPathBar)
		f.SetSaveFileDialog(name, ext)
	}
}

func NewOpenPathBar(owner vcl.IWinControl,name string,width int32, option ...FrameOption) (root *OpenPathBar) {
	vcl.CreateResFrame(owner, &root)
	for _, o := range option {
		o(root)
	}
	root.name = name
	root.width =width

	return
}

func (this *OpenPathBar) SetSaveFileDialog(name string, ext string) {
	this.openType = SAVE_FILE
	this.saveDialog = vcl.NewSaveDialog(this)
	this.saveDialog.SetDefaultExt(ext)
	this.saveDialog.SetTitle(name)
	this.saveDialog.SetOptions(this.saveDialog.Options().Include(types.OfShowHelp, types.OfAllowMultiSelect))
}

func (this *OpenPathBar) SetOpenDirDialog(name string) {
	this.openType = OPEN_DIR
	this.openDirDialog = vcl.NewSelectDirectoryDialog(this)
	this.openDirDialog.SetTitle(name)
}

func (this *OpenPathBar) SetOpenFileDialog(name string, filter string) {
	this.openType = OPEN_FILE
	this.openFileDialog = vcl.NewOpenDialog(this)
	this.openFileDialog.SetFilter(filter)
	this.openFileDialog.SetTitle(name)
	this.openFileDialog.SetOptions(this.openFileDialog.Options().Include(types.OfShowHelp, types.OfAllowMultiSelect))
}

func (this *OpenPathBar) onOpen(s string) {
	if this.OnOpen!=nil {
		this.OnOpen(s)
	} else {
		if this.OnEdit!=nil {
			this.OnEdit(s)
		}
	}
}

func (this *OpenPathBar) setup() {
	this.SetAlign(types.AlLeft)
	this.SetHeight(42)
	this.SetWidth(this.width)
	line1:=CreateLine(types.AlTop,13,this)
	line1.BorderSpacing().SetBottom(0)
	this.label = CreateText(this.name,line1)
	this.label.BorderSpacing().SetLeft(2)
	this.label.Font().SetSize(9)
	this.label.Font().SetColor(colors.ClBlack)
	this.label.SetColor(colors.ClWhite)
	this.edit = CreateEdit(types.TConstraintSize(this.width-36), this)
	this.edit.SetAlign(types.AlClient)
	this.edit.SetHeight(28)
	this.edit.BorderSpacing().SetLeft(0)
	this.btn = CreateSpeedBtn("folder", icons.GetImageList(32, 32), this)
	this.btn.BorderSpacing().SetTop(1)
	this.btn.BorderSpacing().SetBottom(1)
	this.edit.SetOnChange(func(sender vcl.IObject) {
		text := this.edit.Text()
		if this.OnEdit!= nil {
			this.OnEdit(text)
		}
	})
	this.btn.SetOnClick(func(sender vcl.IObject) {
		var p string
		switch this.openType {
		case OPEN_FILE:
			if this.openFileDialog == nil {
				return
			}
			if this.openFileDialog.Execute() {
				p = this.openFileDialog.FileName()
				this.path = p
				this.edit.SetText(p)
				this.onOpen(p)
			}
		case OPEN_DIR:
			if this.openDirDialog == nil {
				return
			}
			if this.openDirDialog.Execute() {
				p = this.openDirDialog.FileName()
				this.path = p
				this.edit.SetText(p)
				this.onOpen(p)
			}
		case SAVE_FILE:
			if this.saveDialog == nil {
				return
			}
			if this.saveDialog.Execute() {
				p = this.saveDialog.FileName()
				this.path = p
				this.edit.SetText(p)
				this.onOpen(p)
			}
		}
	})
}

func (this *OpenPathBar) SetColor(c types.TColor){
	this.color = c
	this.label.SetColor(c)
	this.label.Font().SetColor(c.RGB(255-c.R(),255-c.G(),255-c.B()))
}

func (this *OpenPathBar) SetInitialDir(p string) {
	if this.saveDialog != nil {
		this.saveDialog.SetInitialDir(p)
	}
	if this.openFileDialog != nil {
		this.openFileDialog.SetInitialDir(p)
	}
}

func (this *OpenPathBar) SetPath(p string) {
	this.path = p
	if this.edit.Text()!=p {
		this.edit.SetText(p)
	}
}

func (this *OpenPathBar) GetPath() string {
	return this.path
}

func (this *OpenPathBar) OnCreate() {
	this.setup()
}

func (this *OpenPathBar) OnDestroy() {

}

func (this *OpenPathBar) OnEnter() {

}

func (this *OpenPathBar) OnExit() {

}

func (this *OpenPathBar) Clear() {
	this.edit.SetText("")
}
