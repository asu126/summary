package main

import (
	"fmt"
)

const INT_MAX = int32(^uint32(0) >> 1)
const INT_MIN = ^INT_MAX

func isPowerOfFour(num int) bool {
	if num <= 0 || num > int(INT_MAX) {
		return false
	}

	var tmp int = 1
	for tmp <= num {
		if tmp == num {
			return true
		}
		tmp *= 4
	}

	return false
}

func main() {
	fmt.Println(isPowerOfFour(1))
	fmt.Println(isPowerOfFour(-1))
	fmt.Println(isPowerOfFour(4))
	fmt.Println(isPowerOfFour(5))
	fmt.Println(isPowerOfFour(16))

	fmt.Println(isPowerOfFour(2))
	fmt.Println(isPowerOfFour(8))

}
