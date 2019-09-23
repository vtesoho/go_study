package main

import "fmt"

// func main() {
// 	arrName := [...]int{2, 5, 5, 12}
// 	// fmt.Println(arrName[0])
// 	var targeta = 10
// 	twoSum(arrName, targeta)

// }

func twoSum(nums [4]int, target int) [2]int {
	var start int
	var end int
	for i := 0; i < len(nums)-1; i++ {
		fmt.Println("i---", i)
		for j := i + 1; j < len(nums); j++ {
			fmt.Println("j---", j)
			if (nums[i] + nums[j]) == target {
				start = i
				end = j
				break
			}
		}
		if start != 0 && end != 0 {
			break
		}
	}
	fmt.Println(start, end)
	return [2]int{start, end}
}
