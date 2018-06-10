package main

import "fmt"

import "strings"

// 同205. 同构字符串
func wordPattern(pattern string, str string) bool {
	strArray := strings.Split(str, " ")
	var lengthS, lengthT int = len(pattern), len(strArray)

	if lengthS != lengthT {
		return false
	}
	var m map[byte]string = make(map[byte]string)
	var m1 map[string]byte = make(map[string]byte)

	for i := 0; i < lengthS; i++ {
		if value, ok := m[pattern[i]]; ok {
			if value != strArray[i] {
				return false
			}
		} else if value, ok := m1[strArray[i]]; ok {
			if value != pattern[i] {
				return false
			}
		} else {
			m[pattern[i]] = strArray[i]
			m1[strArray[i]] = pattern[i]
		}
	}

	return true
}

func main() {
	// 	输入: pattern = "abba", str = "dog cat cat dog"
	// 输出: true
	fmt.Println(wordPattern("abba", "dog cat cat dog"))

	// 输入:pattern = "abba", str = "dog cat cat fish"
	// 输出: false
	fmt.Println(wordPattern("abba", "dog cat cat fish"))

	// 	输入: pattern = "aaaa", str = "dog cat cat dog"
	// 输出: false
	fmt.Println(wordPattern("aaaa", "dog cat cat dog"))

	// 	输入: pattern = "abba", str = "dog dog dog dog"
	// 输出: false
	fmt.Println(wordPattern("abba", "dog dog dog dog"))
}
