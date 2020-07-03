package main

import (
	"bufio"
	"fmt"
	"go/functional/fib"
	"os"
)

//何时使用defer 调用
// open / close
// lock / unlock
// printHeader / PrintFooter

func tryDefer() {
	//defer里面是先进后出，不怕被return 或是panic
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//fmt.Println(3)
	//panic("text error")
	//fmt.Println(4)

	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}

}

func main() {
	//writeFile("fib.txt")
	tryDefer()
}
