package treeview

type Node struct {
	Name string
	Nodes []*Node
}

func NewNode()*Node{
	ret:=&Node{
		Name:  "",
		Nodes: []*Node{},
	}
	return ret
}

func (this *Node) String()string {
	return ""
}

func (this *Node) AddNode(node *Node){

}





