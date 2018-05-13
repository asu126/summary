package main

import (
	"fmt"
	"math"
)

func judgeSquareSum(c int) bool {
	var i, j, tmp int
	j = int(math.Sqrt(float64(c)))

	for i = 0; i <= j; {
		tmp = i*i + j*j
		if tmp == c {
			return true
		} else if tmp > c {
			j--
		} else {
			i++
		}
	}

	return false
}

func main() {
	fmt.Println(judgeSquareSum(5))
	fmt.Println(judgeSquareSum(8))
	fmt.Println(judgeSquareSum(1))
}
