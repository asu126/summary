package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	m := make(map[int]int)

	var result int

	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	for j := range m {
		// fmt.Println(j)
		// fmt.Println(m[j])
		if m[j] == 1 {
			result = j
			break
		}
	}
	return result
}

func main() {
	var arr []int = []int{2, 2, 3, 2}
	fmt.Println(singleNumber(arr))

	var arr1 []int = []int{0, 1, 0, 1, 0, 1, 99}
	fmt.Println(singleNumber(arr1))
}
