package main

import (
	"fmt"
)

func findMaxConsecutiveOnes(nums []int) int {
	var result, tmp int = 0, 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			tmp++
		} else {
			if tmp > result {
				result = tmp
			}
			tmp = 0
		}
	}

	if tmp > result {
		result = tmp
	}
	return result
}

func main() {
	var arr []int = []int{1, 1, 0, 1, 1, 1}
	fmt.Println(findMaxConsecutiveOnes(arr)) // 3

	var arr1 []int = []int{1, 1}
	fmt.Println(findMaxConsecutiveOnes(arr1)) // 2

	var arr2 []int = []int{0, 1}
	fmt.Println(findMaxConsecutiveOnes(arr2)) // 1
}
