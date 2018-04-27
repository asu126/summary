package main

import (
	"fmt"
)

func isPowerOfTwo(num int) bool {
	if num <= 0 {
		return false
	}

	var tmp int = 1
	for tmp <= num {
		if tmp == num {
			return true
		}
		tmp *= 2
	}

	return false
}

func main() {
	fmt.Println(isPowerOfTwo(1))
	fmt.Println(isPowerOfTwo(-1))
	fmt.Println(isPowerOfTwo(4))
	fmt.Println(isPowerOfTwo(5))
	fmt.Println(isPowerOfTwo(16))

	fmt.Println(isPowerOfTwo(2))
	fmt.Println(isPowerOfTwo(8))

}
