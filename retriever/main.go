package main

import (
	"fmt"
	"go/retriever/mock"
	real2 "go/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"this is vteso.com"}
	r = real2.Retriever{}
	fmt.Println(download(r))
}
