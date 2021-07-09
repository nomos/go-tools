package treeview

import "github.com/nomos/go-tools/ui"

var _ ui.ITreeSchema = (*Node)(nil)

type Node struct {
	parent *Node
	innerIndex int
	children []ui.ITreeSchema
}

func NewNode()*Node{
	ret:=&Node{
		innerIndex: -1,
		children: []ui.ITreeSchema{},
	}
	return ret
}

func (this *Node) InnerIdx()int{
	return this.innerIndex
}

func (this *Node) Key() string {
	panic("implement me")
}

func (this *Node) Value() string {
	panic("implement me")
}

func (this *Node) Parent() ui.ITreeSchema {
	return this.parent
}

func (this *Node) Children() []ui.ITreeSchema {
	return this.children
}

func (this *Node) String()string {
	return ""
}

func (this *Node) AddNode(node *Node){

}

func (this *Node) GetRootTree()[]int {
	ret:=make([]int,0)
	var root ui.ITreeSchema = this
	for {
		if root.Parent()==nil {
			break
		}
		if root.InnerIdx()== -1 {
			return make([]int,0)
		}
		ret = append(ret, root.InnerIdx())
		root = root.Parent()
	}
	ret1:=make([]int,0)
	for i:=len(ret)-1;i>=0;i-- {
		ret1 = append(ret1, ret[i])
	}
	return ret1
}





