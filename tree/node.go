package tree

import "fmt"

//go包的封装
/*
// 首字母大写代表public
// 首字母小写代表private
每个目录一个包，main包包含可执行入口
为结构定义的方法必须放在同一包内
可以是不同文件
*/
type treeNode struct {
	value       int
	left, right *treeNode
}

func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

//方法一和方法二都是一样的，只不过在使用的语法上不一样
//方法一是root.print()来调用
//方法二是print(root)来调用

//方法一
func (node treeNode) print() {
	fmt.Print(node.value, " ")
}

//方法二
// func print(node treeNode){
// 	fmt.Print(node.value)
// }

//此处加了*号代表是直接传递一个指针地址，这样就能改变这个treeNode的值，如果不加*号，就代表是一个值传递，会拷贝一个给他
//只有指针才可以改变结构内存，nil指针也可以调用方法
func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("setting value to nil node.")
		return
	}
	node.value = value
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	fmt.Println("node", node)
	node.left.traverse()
	node.print()
	node.right.traverse()
}
