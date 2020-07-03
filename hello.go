package main

import "fmt"

func main() {
	//fmt.Println(test(100))
	//testA(10)
	testB(10, 20)
}

//测试递归
func test(n int) int {
	if n == 1 {
		return 1
	}
	return n + test(n-1)
}

//顺序打印递归
func testA(n int) {
	if n > 0 {
		testA(n - 1)
	}
	fmt.Println(n)
}

//指定区间递归
func testB(begin int, end int) {
	if begin > end {
		return
	}
	fmt.Println(begin)
	testB(begin+1, end)
}
