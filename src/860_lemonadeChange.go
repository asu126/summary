package main

import (
	"fmt"
)

func lemonadeChange(bills []int) bool {
	var length int = len(bills)

	if length <= 0 {
		return true
	}

	var i int = 0
	var result bool = true
	var money5, money10, money20 int = 0, 0, 0

	for ; i < length; i++ {
		if !result {
			break
		}

		if bills[i] == 5 {
			money5 += 1
		} else if bills[i] == 10 {
			if money5 >= 1 {
				money5 -= 1
				money10 += 1
			} else {
				result = false
			}
		} else if bills[i] == 20 {
			if money5 >= 1 && money10 >= 1 {
				money5 -= 1
				money10 -= 1
				money20 += 1
			} else if money5 >= 3 {
				money5 -= 3
				money20 += 1
			} else {
				result = false
			}
		} else {
			result = false
		}

	}

	return result
}

func main() {
	ints := []int{5, 10}
	fmt.Println(lemonadeChange(ints))

	fmt.Println(lemonadeChange([]int{5, 5, 5, 10, 20}))

	fmt.Println(lemonadeChange([]int{5, 5, 10}))

	fmt.Println(lemonadeChange([]int{10, 10}))
	fmt.Println(lemonadeChange([]int{5, 5, 10, 10, 20}))

	fmt.Println(lemonadeChange([]int{5, 5, 10, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 20, 5, 20, 5}))
}
