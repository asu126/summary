package main

import (
	"fmt"
)

func peakIndexInMountainArray(A []int) int {
	var length int = len(A)
	var indexMountain int = 0
	if length < 3 {
		return 0
	}

	var i int = 0
	for i = 1; i < length; i++ {
		if A[indexMountain] < A[i] {
			indexMountain = i
		}
	}

	return indexMountain
}

func main() {
	array := []int{0, 2, 1, 0}
	fmt.Println(peakIndexInMountainArray(array)) // 1

	// 输入：[1,2,2,3]
	// 输出：true
	fmt.Println(peakIndexInMountainArray([]int{0, 1, 0})) // 1
}
