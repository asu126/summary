package main

import (
	"fmt"
)

func singleNumber(nums []int) []int {

	m := make(map[int]int)

	var result []int = make([]int, 0, 2)

	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	for j := range m {
		// fmt.Println(j)
		// fmt.Println(m[j])
		if m[j] == 1 {
			result = append(result, j)
		}
	}
	return result
}

func main() {
	var arr []int = []int{2, 2, 3, 2}
	fmt.Println(singleNumber(arr))

	var arr1 []int = []int{1, 2, 1, 3, 2, 5}
	fmt.Println(singleNumber(arr1))
}
