package main

import (
	"fmt"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	var length int = len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i] == nums[j] && j-i <= k {
				return true
			}
		}
	}

	return false
}

func main() {
	a := []int{1, 2, 3, 4, 2}
	fmt.Println(containsNearbyDuplicate(a, 2))
	a1 := []int{-1, -1}
	fmt.Println(containsNearbyDuplicate(a1, 1)) // true
}
