package main

import (
	"fmt"
)

func findMin(nums []int) int {
	var length int = len(nums)
	var left, right, mid int = 0, length - 1, 0
	var min int

	if length > 0 {
		min = nums[0]
	}

	for left <= right {
		mid = left + (right-left)/2

		// if nums[mid] < min {
		// 	min = nums[mid]
		// }

		if nums[mid] >= nums[left] {
			if nums[left] < min {
				min = nums[left]
			}

			left = mid + 1

		} else {
			if nums[mid] < min {
				min = nums[mid]
			}

			right = mid - 1

		}
	}

	return min
}

func main() {
	// 2,5,6,0,0,1,2
	var arr1 []int = []int{3, 4, 5, 1, 2}
	fmt.Println(findMin(arr1)) //1

	var arr2 []int = []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(findMin(arr2)) //0

	var arr3 []int = []int{1, 2, 3, 4, 0}
	fmt.Println(findMin(arr3)) //0

	var arr4 []int = []int{4, 0, 1, 2, 3}
	fmt.Println(findMin(arr4)) //0

	var arr5 []int = []int{1}
	fmt.Println(findMin(arr5)) //1
}
