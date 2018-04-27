package main

import (
	"fmt"
)

func isPowerOfThree(num int) bool {
	if num <= 0 {
		return false
	}

	var tmp int = 1
	for tmp <= num {
		if tmp == num {
			return true
		}
		tmp *= 3
	}

	return false
}

func main() {
	fmt.Println(isPowerOfThree(1))
	fmt.Println(isPowerOfThree(-1))
	fmt.Println(isPowerOfThree(3))
	fmt.Println(isPowerOfThree(5))
	fmt.Println(isPowerOfThree(16))

	fmt.Println(isPowerOfThree(6))
	fmt.Println(isPowerOfThree(8))

}
