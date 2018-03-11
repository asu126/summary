package main

import "fmt"

func moveZeroes(nums []int) {
	length := len(nums)
	var index int = 0
	for i := 0; i < length; i++ {
		if nums[i] != 0 {
			nums[index] = nums[i]
			index++
		}
	}

	for ; index < length; index++ {
		nums[index] = 0
	}
}

func main() {
	arrs := []int{1, 0, 0, 0, 2, 3, 0, 0, 0}

	fmt.Println("Ints: ", arrs)

	moveZeroes(arrs)
	fmt.Println("moved: ", arrs)
}
