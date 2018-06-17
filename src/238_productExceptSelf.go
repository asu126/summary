package main

import "fmt"

func productExceptSelf(nums []int) []int {
	var length int = len(nums)
	var B []int = make([]int, length)

	if length <= 1 {
		return B
	}

	var i int = length - 1
	B[i] = 1
	i--
	for ; i >= 0; i-- {
		B[i] = B[i+1] * nums[i+1]
	}

	var tmp int = nums[0]
	nums[0] = 1
	for i = 1; i < length; i++ {
		var current int = nums[i]
		nums[i] = nums[i-1] * tmp
		tmp = current
	}

	for i = 0; i < length; i++ {
		B[i] = B[i] * nums[i]
	}

	return B
}

func main() {
	// 输入: [1,2,3,4]
	// 输出: [24,12,8,6]
	var array1 []int = []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(array1))

	// 输入: [2,3]
	// 输出: [3,2]
	var array2 []int = []int{2, 3}
	fmt.Println(productExceptSelf(array2))
}
