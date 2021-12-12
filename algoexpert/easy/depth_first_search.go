package easy

type NodeDepthFirst struct {
	Name     string
	Children []*NodeDepthFirst
}

func (n *NodeDepthFirst) AddChild(name string) *NodeDepthFirst {
	n.Children = append(n.Children, &NodeDepthFirst{
		Name:     name,
		Children: []*NodeDepthFirst{},
	})
	return n
}

func (n *NodeDepthFirst) DepthFirstSearch(array *[]string) *[]string {
	*array = append(*array, n.Name)
	for _, child := range n.Children {
		child.DepthFirstSearch(array)
	}
	return array
}
