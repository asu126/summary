package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	var length int = len(digits)
	var result []int = make([]int, length+1)
	var add int = 1

	for i := length - 1; i >= 0; i-- {
		result[i+1] = (digits[i] + add) % 10
		add = (digits[i] + add) / 10
	}

	if add > 0 {
		result[0] = add
	} else {
		result = result[1:]
	}
	return result
}

func main() {
	a := []int{1, 2, 3, 5, 7, 9}
	fmt.Println(plusOne(a))
	a1 := []int{0, 1, 9}
	fmt.Println(plusOne(a1))
}
