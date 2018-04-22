package main

import (
	"fmt"
)

func getMax(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func canJump(nums []int) bool {
	var length int = len(nums)
	var maxIdx int = 0
	for i := 0; i < length; i++ {
		if i > maxIdx || maxIdx >= length-1 {
			break
		}
		maxIdx = getMax(maxIdx, i+nums[i])
	}
	return maxIdx >= length-1
}

func main() {
	var a1 []int = []int{2, 3, 1, 1, 4}
	fmt.Println(canJump(a1))

	var a2 []int = []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(a2))
}
