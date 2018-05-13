package main

import (
	"fmt"
)

func isPerfectSquare(num int) bool {
	var i, j, mid, tmp int = 0, num, 0, 0

	for i = 0; i <= j; {
		mid = i + (j-i)/2
		tmp = mid * mid

		if tmp == num {
			return true
		} else if tmp > num {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}

	return false
}

func main() {
	fmt.Println(isPerfectSquare(5))
	fmt.Println(isPerfectSquare(8))
	fmt.Println(isPerfectSquare(1))
	fmt.Println(isPerfectSquare(16))
}
