package main

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	result := make([]int, 2)
	var i, j int = 0, len(numbers) - 1
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			result[0] = i + 1
			result[1] = j + 1
			return result
		} else if sum > target {
			j--
		} else {
			i++
		}
	}

	return nil
}

func main() {
	var arr []int = []int{2, 7}
	fmt.Println(twoSum(arr, 9))
}
