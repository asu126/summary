package main

import (
	"fmt"
)

func countSmaller(nums []int) []int {
	var length int = len(nums)
	var ret []int = make([]int, length)
	var count int

	for i := 0; i < length; i++ {
		count = 0
		for j := i + 1; j < length; j++ {
			if nums[j] < nums[i] {
				count++
			}
		}
		ret[i] = count
	}

	return ret
}

func main() {
	ints := []int{5, 2, 6, 1}
	fmt.Println(countSmaller(ints))
}
