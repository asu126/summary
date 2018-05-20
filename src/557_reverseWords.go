package main

import (
	"fmt"
)

func reverseWords(s string) string {
	var length int = len(s)
	bs := []byte(s)

	var i, j, k int = 0, 0, 0
	for k = 0; k < length; k++ {
		if bs[k] == ' ' {
			for j = k - 1; i < j; {
				bs[i], bs[j] = bs[j], bs[i] // swap
				i++
				j--
			}
			i = k + 1

		}
	}
	for j = k - 1; i < j; {
		bs[i], bs[j] = bs[j], bs[i] // swap
		i++
		j--
	}
	return string(bs)
}

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))

	fmt.Println(reverseWords("world"))

	fmt.Println(reverseWords("hello"))
}
