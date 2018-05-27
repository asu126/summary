package main

import (
	"fmt"
)

func isIsomorphic(s string, t string) bool {
	// var result
	var lengthS, lengthT int = len(s), len(t)
	if lengthS != lengthT {
		return false
	}
	var m map[byte]byte = make(map[byte]byte)
	var m1 map[byte]byte = make(map[byte]byte)

	for i := 0; i < lengthS; i++ {
		if value, ok := m[s[i]]; ok {
			if value != t[i] {
				return false
			}
		} else if value, ok := m1[t[i]]; ok {
			if value != s[i] {
				return false
			}
		} else {
			// if s[i] != t[i] {
			// 	return false
			// }
			m[s[i]] = t[i]
			m1[t[i]] = s[i]
		}
	}

	return true
}

func main() {
	//  s = "egg", t = "add"
	// true
	fmt.Println(isIsomorphic("egg", "add"))

	//  s = "foo", t = "bar"
	//  false
	fmt.Println(isIsomorphic("foo", "bar"))

	// 	s = "paper", t = "title"
	// 输出: true
	fmt.Println(isIsomorphic("paper", "paper"))

	fmt.Println(isIsomorphic("paper", "papwr")) //true

	// "ab"
	// "aa"
	fmt.Println(isIsomorphic("ab", "aa")) //false
	fmt.Println(isIsomorphic("aa", "ab")) //false
	fmt.Println(isIsomorphic("ab", "ca")) //true
}
