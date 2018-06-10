package main

import "fmt"

func findDuplicate(nums []int) int {
	var length int = len(nums)
	var array []int = make([]int, length+1)

	for i := 0; i < length; i++ {
		if nums[i] <= length && nums[i] >= 1 {
			if array[nums[i]] == 0 {
				array[nums[i]] = nums[i]
			} else {
				return nums[i]
			}
		}
	}

	return 0
}

func main() {
	array := []int{1, 3, 4, 2, 2}
	fmt.Println(findDuplicate(array)) // 2

	array1 := []int{3, 1, 3, 4, 2}
	fmt.Println(findDuplicate(array1)) // 3
}
