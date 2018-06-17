package main

import "fmt"

func majorityElement(nums []int) []int {
	var length int = len(nums)
	var m map[int]int = make(map[int]int)
	var result []int = make([]int, 0)
	var temp = length / 3

	for i := 0; i < length; i++ {
		m[nums[i]]++
	}

	for k, v := range m {
		if v > temp {
			result = append(result, k)
		}
	}
	return result
}

func main() {
	// 	输入: [3,2,3]
	// 输出: [3]
	var arr []int = []int{3, 2, 3}
	fmt.Println(majorityElement(arr))

	// 输入: [1,1,1,3,3,2,2,2]
	// 输出: [1,2]
	var arr1 []int = []int{1, 1, 1, 3, 3, 2, 2, 2}
	fmt.Println(majorityElement(arr1))
}
