package main

import (
	"fmt"
)

func isPerfectSquare(num int) bool {
	var sub int = 1
	for num > 0 {
		num -= sub
		sub += 2
	}
	return num == 0
}

func main() {
	fmt.Println(isPerfectSquare(1))
	fmt.Println(isPerfectSquare(2))
	fmt.Println(isPerfectSquare(4))
	fmt.Println(isPerfectSquare(10))
	fmt.Println(isPerfectSquare(16))
}
