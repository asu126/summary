package main

import (
	"fmt"
)

func countBits(num int) []int {
	var result []int = make([]int, num+1)

	var flag int
	var count int

	for i := 0; i <= num; i++ {
		count = 0
		flag = 1
		for flag != 0 {
			if (flag & i) != 0 {
				count++
			}
			flag = flag << 1
		}
		result[i] = count
	}

	// fmt.Println(count)
	return result
}

func main() {
	// fmt.Println(uint(1) << 64)
	fmt.Println(countBits(0))
	fmt.Println(countBits(1))
	fmt.Println(countBits(2))
	fmt.Println(countBits(3))
	fmt.Println(countBits(4))
	fmt.Println(countBits(5))
}
