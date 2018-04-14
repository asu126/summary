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

func lengthOfLongestSubstring(s string) int {
	if len(s) < 1 {
		return 0
	}
	var max int = 1
	var i, j, k int
	var flag bool = false
	for i = 0; i < len(s); i++ {
		for j = i + 1; j < len(s); j++ {
			flag = false
			for k = i; k < j; k++ {
				if s[j] == s[k] {
					flag = true
					// fmt.Println(i, j, k)
					break
				}
			}
			if flag {
				break
			}
		}
		max = getMax(max, j-i)
	}
	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abc"))
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("c"))
	fmt.Println(lengthOfLongestSubstring(""))
}
