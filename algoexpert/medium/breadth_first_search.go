package medium

type GraphNode struct {
	Name string
	Children []*GraphNode
}


func (g *GraphNode) AddChild(name string) *GraphNode {
	g.Children = append(g.Children, &GraphNode{Name: name})
	return g
}


func (g *GraphNode) BreadthFirstSearch() []string {
	var results []string
	queue := []*GraphNode{g}
	for len(queue) > 0 {
		var x *GraphNode
		x, queue = queue[len(queue) - 1], queue[:len(queue) - 1]
		results = append(results, x.Name)
	}
	return results
}