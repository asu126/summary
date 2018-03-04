package main

import (
	"fmt"
)

func isVowels(b byte) bool {
	var vowels []byte = []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	for i := 0; i < len(vowels); i++ {
		if vowels[i] == b {
			return true
		}
	}
	return false
}

func reverseVowels(s string) string {
	bs := []byte(s)
	var i, j int = 0, len(s) - 1
	for i < j {
		// fmt.Println(i, j)
		if isVowels(s[i]) == false {
			i++
			continue
		}
		if isVowels(s[j]) == false {
			j--
			continue
		}
		bs[j] = s[i]
		bs[i] = s[j]
		i++
		j--
	}

	return string(bs)
}

func main() {
	var s string = "leetcode"
	fmt.Println(reverseVowels(s))
}
