package main

import (
	"fmt"
)

import "strings"

func findWords(words []string) []string {
	var result []string = make([]string, 0, 128)
	var line1, line2, line3 int
	var str string

	for i := 0; i < len(words); i++ {
		str = strings.ToLower(words[i])
		line1, line2, line3 = 0, 0, 0
		for j := 0; j < len(str); j++ {
			switch str[j] {
			case 'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p':
				line1 = 1
			case 'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l':
				line2 = 1
			case 'z', 'x', 'c', 'v', 'b', 'n', 'm':
				line3 = 1
			}
		}

		if (line1 + line2 + line3) <= 1 {
			result = append(result, words[i])
		}
	}

	return result
}

func main() {
	var strs []string = []string{"Hello", "Alaska", "Dad", "Peace"}
	fmt.Println(findWords(strs))
}
