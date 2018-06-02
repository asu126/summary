package main

import (
	"fmt"
)

func findPeakElement(nums []int) int {
	var length int = len(nums)

	if length < 1 {
		return 0
	}
	if length == 1 {
		return 0
	}

	for i := 0; i < length; i++ {
		if i == 0 {
			if i+1 < length && nums[0] > nums[1] {
				return i
			}
		} else if i == length-1 {
			if i-1 >= 0 && nums[i] > nums[i-1] {
				return i
			}
		} else {
			if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
				return i
			}
		}
	}
	return 0
}

func main() {
	a := []int{1, 2, 3, 1}
	fmt.Println(findPeakElement(a)) //2

	a1 := []int{1, 2, 1, 3, 5, 6, 4}
	fmt.Println(findPeakElement(a1)) //1

	a2 := []int{1}
	fmt.Println(findPeakElement(a2)) //0

	a3 := []int{5, 6, 7, 8, 9}
	fmt.Println(findPeakElement(a3)) //4
}
