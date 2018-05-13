package main

import (
	"fmt"
)

func findDisappearedNumbers(nums []int) []int {
	var length int = len(nums)
	var restlt []int = make([]int, 0, length)

	for i := 0; i < length; i++ {
		if nums[i] < 1 || nums[i] > length {
			return restlt
		}
		nums[i] -= 1
	}

	for i := 0; i < length; i++ {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				break
			}

			var tmp int = nums[i]
			nums[i] = nums[tmp]
			nums[tmp] = tmp
			// fmt.Println(nums)
		}

	}
	for i := 0; i < length; i++ {
		if nums[i] != i {
			restlt = append(restlt, i+1)
			// fmt.Println(nums[i])
		}

	}

	return restlt
}

func main() {
	var arr []int = []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(arr)) // [2,3]

	var arr1 []int = []int{3, 4, 2, 1, 3, 6, 4}
	fmt.Println(findDisappearedNumbers(arr1)) // 2
}
