package main

import (
	"fmt"
)

func distributeCandies(candies []int) int {
	m := make(map[int]int)
	var length int = len(candies)

	for i := 0; i < length; i++ {
		m[candies[i]]++
	}

	if len(m) <= length/2 {
		return len(m)
	} else {
		return length / 2
	}
}

func main() {
	var arr []int = []int{1, 1, 2, 2, 3, 3}
	fmt.Println(distributeCandies(arr)) //3

	var arr1 []int = []int{1, 1, 2, 3}
	fmt.Println(distributeCandies(arr1)) // 2

	var arr2 []int = []int{1, 1, 1, 1}
	fmt.Println(distributeCandies(arr2)) // 1
}
