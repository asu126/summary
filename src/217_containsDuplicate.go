package main

import (
	"fmt"
)

func containsDuplicate(nums []int) bool {
	var length int = len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}

	return false
}

func main() {
	a := []int{1, 2, 3, 4, 2}
	fmt.Println(containsDuplicate(a))
}
