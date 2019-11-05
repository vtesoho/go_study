package main

func main() {

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)
	root.right.left.setValue(4)

	root.traverse()
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
