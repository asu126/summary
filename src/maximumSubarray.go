package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var maxSum, tmp int = nums[0], 0

	for i := 0; i < len(nums); i++ {
		tmp = nums[i]
		if tmp > maxSum {
			maxSum = tmp
		}
		for j := i + 1; j < len(nums); j++ {
			tmp += nums[j]
			if tmp > maxSum {
				maxSum = tmp
			}
		}
	}

	return maxSum
}

func main() {
	// var arr []int = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	var arr []int = []int{-2, 0, -1}
	fmt.Println(maxSubArray(arr))

}
