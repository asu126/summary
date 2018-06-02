package main

import (
	"fmt"
)

import "strings"

func isPalindrome(s string) bool {
	str := strings.ToLower(s)

	var array []byte = make([]byte, 0)

	for i := 0; i < len(str); i++ {
		if (str[i] >= 'a' && str[i] <= 'z') || (str[i] >= '0' && str[i] <= '9') {
			array = append(array, str[i])
		}
	}

	// fmt.Println(array)
	for i, j := 0, len(array)-1; i < j; {
		if array[i] != array[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))

	fmt.Println(isPalindrome("race a car"))

	fmt.Println(isPalindrome(".,"))
}
