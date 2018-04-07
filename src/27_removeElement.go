package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	var i, j int = 0, 0
	for ; i < len(nums); i++ {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}

	}

	return j
}

func main() {
	a1 := []int{3, 2, 2, 3}
	a2 := []int{2, 4}
	fmt.Println(removeElement(a1, 2))
	fmt.Println(a1)
	fmt.Println(removeElement(a2, 2))
	fmt.Println(a2)
}
