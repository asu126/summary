package main

import (
	"fmt"
)

func isMonotonic(A []int) bool {
	// var up bool = false
	var length int = len(A)

	if length < 2 {
		return true
	}

	if A[length-1] >= A[0] {
		for i := 0; i < length-1; i++ {
			if A[i+1] < A[i] {
				return false
			}
		}
	} else { // A[length-1] <=A[0]
		for i := 0; i < length-1; i++ {
			if A[i+1] > A[i] {
				return false
			}
		}
	}

	return true
}

func main() {
	array := []int{1, 1, 5, 3, 6, 1}
	fmt.Println(isMonotonic(array))

	// 输入：[1,2,2,3]
	// 输出：true
	fmt.Println(isMonotonic([]int{1, 2, 2, 3}))

	// 输入：[6,5,4,4]
	// 输出：true
	fmt.Println(isMonotonic([]int{6, 5, 4, 4}))

	// 输入：[1,3,2]
	// 输出：false
	fmt.Println(isMonotonic([]int{1, 3, 2}))

	// 输入：[1,2,4,5]
	// 输出：true
	fmt.Println(isMonotonic([]int{1, 2, 4, 5}))

	// 	输入：[1,1,1]
	// 输出：true
	fmt.Println(isMonotonic([]int{1, 1, 1}))
}
