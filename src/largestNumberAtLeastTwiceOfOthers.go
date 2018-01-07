package main

import (
	"fmt"
)

func dominantIndex(nums []int) int {
	var max, index, next_max int = -1, -1, 0
	var length int = len(nums)
	if length == 1 {
		return 0
	}

	// find max
	for i := 0; i < length; i++ {
		if nums[i] > max {
			max = nums[i]
			index = i
		}
	}

	// find next max
	for i := 0; i < length; i++ {
		if nums[i] > next_max && nums[i] != max {
			next_max = nums[i]
		}
	}

	if max >= (2 * next_max) {
		return index
	}
	return -1
}

func main() {
	// array1 := [4]int{3, 6, 1, 0}

	// 使用切片创建数组
	// array := []int{3, 6, 1, 0}
	// array := []int{0, 0, 3, 2}
	// array := []int{1, 6, 6, 3}
	array := []int{0}
	fmt.Println(array)
	fmt.Println(dominantIndex(array))
}
