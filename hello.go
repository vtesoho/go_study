package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v,len=%d, cap= %d\n", s, len(s), cap(s))
}

func main() {
	fmt.Println("create splie")
	var s []int // zero value for slice is nil

	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	//创建一个已知长度的slice
	s2 := make([]int, 16)
	printSlice(s2)
	s3 := make([]int, 10, 32)
	s3 = append(s3, 20)
	printSlice(s3)

	fmt.Println("copying slice")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("deleting elements from slice")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("popping form front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	printSlice(s2)

	fmt.Println("popping form back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]

	fmt.Println(tail)
	printSlice(s2)

}
