package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	var i int = 0
	for ; i < len(nums); i++ {
		if nums[i] >= target {
			break
		}
	}

	return i
}

func main() {
	a1 := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(a1, 5)) // 2
	fmt.Println(searchInsert(a1, 2)) // 1
	fmt.Println(searchInsert(a1, 7)) // 4
	fmt.Println(searchInsert(a1, 0)) // 0
}
