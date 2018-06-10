package main

import (
	"fmt"
)

func missingNumber(nums []int) int {
	var sum int = 0
	var length int = len(nums)
	if length < 1 {
		return 0
	}

	var min, max int = nums[0], nums[0]
	for i := 0; i < length; i++ {
		sum += nums[i]

		if nums[i] > max {
			max = nums[i]
		}

		if nums[i] < min {
			min = nums[i]
		}
	}

	shulie := ((length + 1) * length) / 2

	return shulie - sum
}

func main() {
	var arr []int = []int{3, 0, 1}
	fmt.Println(missingNumber(arr)) //2

	var arr1 []int = []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Println(missingNumber(arr1)) //8

	var arr2 []int = []int{1}
	fmt.Println(missingNumber(arr2)) //0

}
