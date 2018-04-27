package main

import (
	"fmt"
)

func isUgly(num int) bool {
	if num == 0 {
		return false
	}
	for num%2 == 0 || num%3 == 0 || num%5 == 0 {
		for num%2 == 0 {
			num /= 2
		}
		for num%3 == 0 {
			num /= 3
		}
		for num%5 == 0 {
			num /= 5
		}
	}

	if num == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(isUgly(0))
	fmt.Println(isUgly(1))
	fmt.Println(isUgly(2))
	fmt.Println(isUgly(3))
	fmt.Println(isUgly(4))
	fmt.Println(isUgly(5))
	fmt.Println(isUgly(14))
}
