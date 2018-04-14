package main

import (
	"fmt"
)

const INT_MAX = int32(^uint32(0) >> 1)
const INT_MIN = ^INT_MAX

func reverse(x int) int {
	var arr []int = make([]int, 0, 32)
	var is_negative bool = false
	var result int
	if x < 0 {
		is_negative = true
		x = -x
	}

	for x != 0 {
		arr = append(arr, x%10)
		x /= 10
	}

	var tmp int = 1
	for i := len(arr) - 1; i >= 0; i-- {
		result += arr[i] * tmp
		tmp *= 10
	}

	if is_negative {
		result = -result
	}
	if result >= int(INT_MIN) && result <= int(INT_MAX) {
		return result
	}

	return 0
}

func main() {
	fmt.Println(reverse(1534236469))
	fmt.Println(reverse(-123))
	fmt.Println(int((^uint(0)) >> 1))
}
