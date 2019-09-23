package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(reverse(11))
	fmt.Println(reverse(1563847412))

}

func reverse(x int) int {
	// fmt.Println(2 << 31)

	num := x
	arr := []int{}
	if x < 0 {
		num = x - x*2
	}
	inde := 1
	for {
		if (num / inde) < 1 {
			inde = inde / 10
			break
		}
		inde = inde * 10
	}
	cu := 0
	indea := inde
	for {
		if inde < 1 {
			break
		}
		cu = int(math.Floor(float64(num / inde)))
		arr = append(arr, cu)
		num = num - (cu * inde)
		inde = inde / 10
	}
	re := 0
	index := len(arr) - 1
	for {
		if index < 0 {
			break
		}
		re = re + arr[index]*indea
		indea = indea / 10
		index--
	}
	if x < 0 {
		re = re - re*2
	}
	if re < -(2<<30) || re > (2<<30)-1 {
		re = 0
	}
	return re
}
