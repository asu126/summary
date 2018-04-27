package main

import (
	"fmt"
)

func addDigits(num int) int {
	var result int = num
	for num >= 10 {
		result = 0
		for num != 0 {
			result += num % 10
			num = num / 10
			// fmt.Println(num, result)
		}
		num = result
	}

	return result
}

func main() {
	// fmt.Println(addDigits(11))
	// fmt.Println(addDigits(38))
	fmt.Println(addDigits(1))
}
