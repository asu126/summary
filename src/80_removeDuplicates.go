package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	var length int = len(nums)
	var result int = 0
	var before int
	var beforeCount int

	if length > 0 {
		result++
		beforeCount++
		before = nums[0]
	}

	for i := 1; i < length; i++ {
		// fmt.Println(result, nums)
		if before == nums[result] {
			beforeCount++
		} else {
			beforeCount = 1
			before = nums[result]
		}

		if beforeCount > 2 {
			// fmt.Println("...")
			for j := result; j < length-1; j++ {
				nums[j] = nums[j+1]
			}
		} else {
			result++
		}

	}

	return result
}

func main() {
	// [2,0,2,1,1,0]
	var arr1 []int = []int{1, 1, 1, 2, 2, 3}
	removeDuplicates(arr1)
	fmt.Println(arr1)

	var arr2 []int = []int{1, 1, 2, 2, 3}
	removeDuplicates(arr2)
	fmt.Println(arr2)

	var arr3 []int = []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	removeDuplicates(arr3)
	fmt.Println(arr3)
}
