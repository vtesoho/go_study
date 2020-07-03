package tree

import "fmt"

func (node *Node) Traverse() {
	if node == nil {
		return
	}
	fmt.Println("node", node)
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}
