package main

import (
	"fmt"
)

func nextGreaterElement(findNums []int, nums []int) []int {
	var lenA, lenB int = len(findNums), len(nums)
	var result []int
	if lenA < 1 {
		return result
	}

	result = make([]int, lenA)

	var j int
	for i := 0; i < lenA; i++ {
		for j = 0; j < lenB; j++ {
			if nums[j] == findNums[i] {
				break
			}
		}
		for ; j < lenB; j++ {
			if nums[j] > findNums[i] {
				result[i] = nums[j]
				break
			}
		}
		if j == lenB {
			result[i] = -1
		}
	}

	return result
}

func main() {
	a1 := []int{4, 1, 2}
	a2 := []int{1, 3, 4, 2}
	fmt.Println(nextGreaterElement(a1, a2))
}
