package main

import (
	"fmt"
)

func isUpChar(b byte) bool {
	if b >= 'A' && b <= 'Z' {
		return true
	}

	return false
}

func detectCapitalUse(word string) bool {
	var length int = len(word)
	if length <= 1 {
		return true
	}

	var firstUpChar bool = isUpChar(word[0])
	var other bool = isUpChar(word[1])

	for i := 2; i < length; i++ {
		if isUpChar(word[i]) != other {
			return false
		}
	}

	if firstUpChar == false && other {
		return false
	}

	return true
}

func main() {
	fmt.Println(detectCapitalUse("USA"))  //true
	fmt.Println(detectCapitalUse("Flag")) //true
	fmt.Println(detectCapitalUse("FlaG")) //false
	fmt.Println(detectCapitalUse("fLAG")) //false
}
