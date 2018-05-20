package main

import (
	"fmt"
)

func topKFrequent(words []string, k int) []string {
	var result []string = make([]string, k)
	m := make(map[string]int)
	length := len(words)

	for i := 0; i < length; i++ {
		m[words[i]]++
	}

	for i := 0; i < k; i++ {
		var maxKey string = ""
		var maxValue int = 0
		for k1, v1 := range m {
			// 如果不同的单词有相同出现频率，按字母顺序排序
			if v1 > maxValue || (v1 == maxValue && k1 < maxKey) {
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
	var arr []string = []string{"i", "love", "leetcode", "i", "love", "coding"}
	fmt.Println(topKFrequent(arr, 2))

	var arr1 []string = []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}
	fmt.Println(topKFrequent(arr1, 4))
}
