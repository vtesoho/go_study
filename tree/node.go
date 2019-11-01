package main

import "fmt"

type treeNode struct{
	value int
	left,right *treeNode
}

func createNode(value int) *treeNode{
	return &treeNode{value:value}
}

//方法一和方法二都是一样的，只不过在使用的语法上不一样
//方法一是root.print()来调用
//方法二是print(root)来调用

//方法一
func (node treeNode) print(){
	fmt.Print(node.value, " ")
}

//方法二
// func print(node treeNode){
// 	fmt.Print(node.value)
// }

//此处加了*号代表是直接传递一个指针地址，这样就能改变这个treeNode的值，如果不加*号，就代表是一个值传递，会拷贝一个给他
//只有指针才可以改变结构内存，nil指针也可以调用方法
func (node *treeNode) setValue(value int) {
	if node == nil{
		fmt.Println("setting value to nil node.")
		return
	}
	node.value = value
}

func (node *treeNode) traverse(){
	if node == nil {
		return
	}
	fmt.Println("node",node)
	node.left.traverse()
	node.print()
	node.right.traverse()
}




func main(){
	var root treeNode

	root = treeNode{value:3}
	root.left = &treeNode{}
	root.right = &treeNode{5,nil,nil}
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