package tools

import (
	"fmt"
	"github.com/nomos/go-tools/ui"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

var _ ui.IFrame = (*GitProxyTool)(nil)

type GitProxyTool struct {
	*vcl.TFrame
	ui.ConfigAble
	GitProxyEdit *ui.EditLabel

}

func NewGitProxyTool(owner vcl.IWinControl,option... ui.FrameOption) (root *GitProxyTool)  {
	vcl.CreateResFrame(owner, &root)
	for _,o:=range option {
		o(root)
	}
	return
}

func (this *GitProxyTool) setup(){
	this.SetAlign(types.AlClient)
	this.SetConfig(this.Config().Sub("git"))
	line:=ui.CreateLine(types.AlTop,44,this)
	toggle:=ui.NewToggleFrame(line,false)
	toggle.SetParent(line)
	toggle.SetAlign(types.AlLeft)
	toggle.OnCreate()
	ui.CreateSeg(10,line)
	this.GitProxyEdit=ui.NewEditLabel(line,"git proxy",140,ui.EDIT_TYPE_STRING)
	this.GitProxyEdit.SetParent(line)
	this.GitProxyEdit.SetChangeOnEdit(true)
	this.GitProxyEdit.SetAlign(types.AlLeft)
	toggle.BorderSpacing().SetTop(12)
	toggle.OnChecked = func(b bool) {
		if b {
			this.Config().Set("git_proxy_on",true)
			this.gitProxyOn()
		} else {
			this.Config().Set("git_proxy_on",false)
			this.gitProxyOff()
		}
	}
	this.GitProxyEdit.OnCreate()
	if this.Config().GetString("git_proxy")!= "" {
		this.GitProxyEdit.SetString(this.Config().GetString("git_proxy"))
	}
	if this.Config().GetBool("git_proxy_on") {
		toggle.Check()
	}
	this.GitProxyEdit.OnValueChange = func(label *ui.EditLabel, editType ui.EDIT_TYPE, value interface{}) {
		this.Config().Set("git_proxy",this.GitProxyEdit.String())
	}
	if this.Config().GetBool("git_proxy_on") {
		toggle.Check()
	} else {
		toggle.Uncheck()
	}
}

func (this *GitProxyTool) gitProxyOn(){
	if this.Config().GetString("git_proxy") == "" {
		this.GitProxyEdit.SetFocus()
		return
	}
	this.GetLogger().(*ui.ConsoleShell).ExecShellCmd(
		fmt.Sprintf("git config --global http.proxy %s;git config --global https.proxy %s",
			this.Config().GetString("git_proxy"),
			this.Config().GetString("git_proxy"),
		),false).Await()
}

func (this *GitProxyTool) gitProxyOff(){
	this.GetLogger().(*ui.ConsoleShell).ExecShellCmd("git config --global --unset http.proxy;git config --global  --unset https.proxy",false).Await()
}

func (this *GitProxyTool) OnCreate(){
	this.setup()
}

func (this *GitProxyTool) OnDestroy(){

}

func (this *GitProxyTool) OnEnter(){

}

func (this *GitProxyTool) OnExit(){

}

func (this *GitProxyTool) Clear(){

}


