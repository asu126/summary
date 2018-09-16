package main

import "fmt"

func repeatedSubstringPattern(s string) bool {
	var length int = len(s)
	if length <= 1 {
		return false
	}

	var i int = length / 2
	var tag bool
	for ; i > 0; i-- {
		tag = true
		if length%i == 0 {
			var j int
			for j = 0; j < i; j++ {
				for k := j + i; k < length; {
					if s[k-i] != s[k] {
						tag = false
						break
					}
					k += i

				}
			}
			if tag {
				return true
			}
		}

	}

	return false
}

func main() {
	// 输入: "abab"
	// 输出: True
	fmt.Println(repeatedSubstringPattern("abab"))

	// 输入: "aba"
	// 输出: False
	fmt.Println(repeatedSubstringPattern("aba"))

	// 输入: "abcabcabcabc"
	// 输出: True
	fmt.Println(repeatedSubstringPattern("abcabcabcabc"))

	// 输入: "abcabcaabcabc"
	// 输出: False
	fmt.Println(repeatedSubstringPattern("abcabcaabcabc"))

	// 输入: "aaa"
	// 输出: True
	fmt.Println(repeatedSubstringPattern("aaa"))

	// 输入: "babbbbaabb"
	// 输出: False
	fmt.Println(repeatedSubstringPattern("babbbbaabb"))
}
