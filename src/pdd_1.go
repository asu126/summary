package main

import (
	"fmt"
)

func print(m int) {
	var result [][]int = make([][]int, m)
	// var dp []int = make([]int, 3)  //error

	for i := 0; i < m; i++ {
		var nn []int = make([]int, m)
		result[i] = nn
	}

	var left, right, low, high int = 0, m - 1, 0, m - 1
	var number int = 1

	for left <= right && low <= high {
		for i := left; i <= right; i++ {
			// fmt.Println("1", left, right, low, high)
			result[low][i] = number
			number++
		}
		low++

		for i := low; i <= high; i++ {
			// fmt.Println("2", left, right, low, high)
			result[i][right] = number
			number++
		}
		right--

		for i := right; i >= left; i-- {
			// fmt.Println("3", left, right, low, high)
			result[high][i] = number
			number++
		}
		high--

		for i := high; i >= low; i-- {
			// fmt.Println("4", left, right, low, high)
			result[i][left] = number
			number++
		}
		left++

	}
	fmt.Println(result)
}

func main() {
	print(1)
	print(2)
	print(3)
	print(4)
}
