package main

import "fmt"

func search(nums []int, target int) int {
	var i int
	for i = 0; i < len(nums); i++ {
		if nums[i] == target {
			break
		}
	}
	if i == len(nums) {
		return -1
	}
	return i
}

func main() {
	a := []int{7, 8, 1, 2, 3, 5, 6}

	fmt.Println(search(a, 8)) // true
}
