package main

import (
	"fmt"
)

func lengthOfLastWord(s string) int {
	var length int = len(s)
	if length == 0 {
		return 0
	}
	startNotBalck, endNotBlank := 0, 0
	for i := 0; i < length; i++ {
		if s[i] != ' ' {
			startNotBalck = i
			break
		}
	}
	for i := length - 1; i >= 0; i-- {
		if s[i] != ' ' {
			endNotBlank = i
			break
		}
	}
	var lastBlank int = startNotBalck
	for i := startNotBalck; i <= endNotBlank; i++ {
		if s[i] == ' ' {
			lastBlank = i
			// fmt.Println(s[i])
		}
	}
	// fmt.Println(startNotBalck, endNotBlank, lastBlank)
	if lastBlank == startNotBalck && s[startNotBalck] != ' ' {
		// fmt.Println(s[lastBlank : endNotBlank+1])
		return endNotBlank + 1 - startNotBalck
	}

	// fmt.Println(s[lastBlank+1 : endNotBlank+1])
	return endNotBlank - lastBlank
}

func main() {
	// s := "hello dd"
	// s := "    day"
	// s := " a"
	s := " "
	// s := ""
	// s := " a "
	fmt.Println(lengthOfLastWord(s))
}
