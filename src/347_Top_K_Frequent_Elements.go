package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	var result []int = make([]int, k)
	m := make(map[int]int)
	length := len(nums)

	for i := 0; i < length; i++ {
		m[nums[i]]++
	}

	for i := 0; i < k; i++ {
		maxKey, maxValue := 0, 0
		for k1, v1 := range m {
			if v1 > maxValue {
				maxKey = k1
				maxValue = v1
			}
		}
		result[i] = maxKey
		delete(m, maxKey)
	}

	return result
}

func main() {
	ints := []int{7, 2, 2, 4}
	fmt.Println(topKFrequent(ints, 2))
}
