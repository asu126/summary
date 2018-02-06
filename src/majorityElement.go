package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	m := make(map[int]int)
	var t int = len(nums) / 2
	var majority int = nums[t]

	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	for j := range m {
		// fmt.Println(j)
		// fmt.Println(m[j])
		if m[j] > t {
			majority = j
			break
		}
	}
	return majority
}

func main() {
	// m := make(map[string]int)
	// m["a"] = 1
	// fmt.Println(m["a"])
	// m["b"]++
	// fmt.Println(m["b"])
	// fmt.Println(len(m))

	// var arr []int = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	// var arr []int = []int{2, 2}
	var arr []int = []int{6, 5, 5}
	fmt.Println(majorityElement(arr))
}
