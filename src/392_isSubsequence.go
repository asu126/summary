package main

import "fmt"

func isSubsequence(s string, t string) bool {
	var lengthS, lengthT int = len(s), len(t)
	var i, j int = 0, 0

	for i = 0; i < lengthS; i++ {
		for ; j < lengthT; j++ {
			if s[i] == t[j] {
				// j++
				break
			}
		}
		if j >= lengthT {
			return false
		}

		j++
	}

	return true
}

func main() {
	// s = "abc", t = "ahbgdc"
	// 返回 true.
	fmt.Println(isSubsequence("abc", "ahbgdc"))

	// s = "axc", t = "ahbgdc"
	// 返回 false.
	fmt.Println(isSubsequence("axc", "ahbgdc"))

	fmt.Println(isSubsequence("aabc", "abaceew"))
	fmt.Println(isSubsequence("aabc", "abc"))
}
