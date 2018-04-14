package main

import (
	"fmt"
)

func getMin(a, b int) int {
	if a >= b {
		return b
	} else {
		return a
	}
}

func getMax(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func maxArea(height []int) int {
	var max int = 0

	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			max = getMax(max, getMin(height[i], height[j])*(j-i))
		}
	}

	return max
}

func main() {
	a := []int{1, 1}
	fmt.Println(maxArea(a))
	a1 := []int{1, 2, 1, 2, 3, 0}
	fmt.Println(maxArea(a1))
}
