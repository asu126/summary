package main

import (
	"fmt"
)

func trailingZeroes(n int) int {
	var countZero int = 0
	for n > 0 {
		countZero += n / 5
		n = n / 5
	}
	return countZero
}

func main() {
	fmt.Println(trailingZeroes(1))
	fmt.Println(trailingZeroes(2))
	fmt.Println(trailingZeroes(4))
	fmt.Println(trailingZeroes(8))
	fmt.Println(trailingZeroes(10))
	fmt.Println(trailingZeroes(101))
}
