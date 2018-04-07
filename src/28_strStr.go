package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	lengthA := len(haystack)
	lengthB := len(needle)
	if lengthB == 0 {
		return 0
	}
	for i := 0; i < lengthA; i++ {
		var tag bool = true
		var j int
		for j = 0; j < lengthB; j++ {
			if i+j < lengthA {
				if haystack[i+j] != needle[j] {
					tag = false
					break
				}
			} else {
				tag = false
				break
			}
		}

		if tag && j == lengthB {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("aaa", "aaaa"))
}
