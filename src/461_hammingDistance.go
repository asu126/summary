package main

import "fmt"

func hammingDistance(x int, y int) int {
	var result int = 0
	var temp int = 1 << 31

	var i int = 0
	for ; i < 32; i++ {
		if (x&temp)^(y&temp) > 0 {
			result++
		}
		temp = temp >> 1
	}

	return result
}

func main() {
	fmt.Println(hammingDistance(1, 4)) // 2
}
