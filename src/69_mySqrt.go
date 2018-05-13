package main

import (
	"fmt"
)

func mySqrt(x int) int {
	var i, j, mid, tmp int = 0, x, 0, 0

	for i = 0; i < j; {
		mid = i + (j-i)/2
		tmp = mid * mid

		if tmp == x {
			return mid
		} else if tmp > x {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}

	if i*i > x {
		return i - 1
	}
	return i
}

func main() {
	fmt.Println(mySqrt(5))
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(1))
	fmt.Println(mySqrt(16))
}
