package main

import (
	"fmt"
)

// // timeout
// func isAnagram(s string, t string) bool {
// 	var len1, len2 = len(s), len(t)
// 	if len1 != len2 {
// 		return false
// 	}

// 	var flag []bool = make([]bool, len2)
// 	var i, j int
// 	for i = 0; i < len1; i++ {
// 		for j = 0; j < len2; j++ {
// 			if s[i] == t[j] && !flag[j] {
// 				flag[j] = true
// 				break
// 			}
// 		}
// 		if j == len2 {
// 			return false
// 		}
// 	}

// 	return true
// }

func isAnagram(s string, t string) bool {
	var len1, len2 = len(s), len(t)

	var begin int = int('a')
	var flag []int = make([]int, 26)
	for i := 0; i < len1; i++ {
		flag[int(s[i])-begin]++
	}
	for i := 0; i < len2; i++ {
		flag[int(t[i])-begin]--
	}
	for i := 0; i < 26; i++ {
		if flag[i] != 0 {
			return false
		}
	}
	return true
}

func main() {
	// 	s = "anagram"，t = "nagaram"，返回 true
	// s = "rat"，t = "car"，返回 false
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
	fmt.Println(isAnagram("aacc", "ccac"))
}
