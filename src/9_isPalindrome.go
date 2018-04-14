package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	var tmp int = x
	var digits int = 1
	if x < 0 {
		return false
	}
	for tmp > 9 {
		digits *= 10
		tmp /= 10
	}

	var high, low int
	for x > 0 {
		high = x / digits
		low = x % 10
		// fmt.Println(x, high, low)
		if high != low {
			return false
		}

		x = (x - (digits * high) - low) / 10
		digits /= 100
	}

	return true
}

func main() {
	fmt.Println(isPalindrome(25652))
	fmt.Println(isPalindrome(-121))
}
