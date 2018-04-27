package main

import (
	"fmt"
)

func titleToNumber(s string) int {
	var result int
	var tmp int = 1

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] <= 'Z' && s[i] >= 'A' {
			result += ((int(s[i]) - 65 + 1) * tmp)
		}
		tmp *= 26
	}

	return result
}

func main() {

	// A -> 1
	//  B -> 2
	//  C -> 3
	//  ...
	//  Z -> 26
	//  AA -> 27
	//  AB -> 28
	fmt.Println(titleToNumber("A"))
	fmt.Println(titleToNumber("Z"))
	fmt.Println(titleToNumber("AA"))
	fmt.Println(titleToNumber("AB"))
	fmt.Println(titleToNumber("AAA"))
	// fmt.Println(int('A'))
	// fmt.Println(int('Z'))
}
