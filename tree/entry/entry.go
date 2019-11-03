package main

import (
	"fmt"
	"go/tree"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	root.Traverse()
	fmt.Println()
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	//root.Traverse()
	// // nodes := []treeNode{
	// // 	{value:3},
	// // 	{},
	// // 	{6,nil,&root},
	// // }
	// // fmt.Println(root)

	// // fmt.Println(root.left.right)

	// root.right.left.print()

	// // root.print()
	// root.setValue(100)

	// //此处如果加了&号，代表此处引用root变量的内存地址，不重新创建一个新的
	// pRoot := &root
	// // pRoot.print()
	// pRoot.setValue(200)
	// // root.print()

	// var nRoot *treeNode
	// nRoot.setValue(200)
	// nRoot = &root
	// nRoot.setValue(300)
	// nRoot.print()

}
