package main

import (
	"fmt"
)

//go语言只有值传递
func printArray(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

// 也可以通过指针变成引用传递
func printArrayA(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	// 定义数组
	// 不设置初始值为0
	// 数组是值类型
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	//原始遍历方式
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	//range
	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	printArray(arr1)
	printArray(arr3)
	// printArray(arr2)
	fmt.Println(arr1, arr3)

	fmt.Println("---------------------")

	printArrayA(&arr1)
	printArrayA(&arr3)
	// printArray(arr2)
	fmt.Println(arr1, arr3)

}
