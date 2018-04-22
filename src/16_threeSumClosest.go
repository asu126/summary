package main

import (
	"fmt"
)

func qiuck_sort(arr []int, l int, r int) {
	if l < r {
		var tmp, i, j = arr[l], l, r

		for i < j {
			for i < j && arr[j] >= tmp {
				j--
			}
			if i < j {
				arr[i] = arr[j]
			}

			for i < j && arr[i] <= tmp {
				i++
			}
			if i < j {
				arr[j] = arr[i]
			}
		}
		arr[i] = tmp
		qiuck_sort(arr, l, i-1)
		qiuck_sort(arr, i+1, r)
	}
}

func absClosest(a, b int) int {
	if a >= b {
		return a - b
	} else {
		return b - a
	}
}

func threeSumClosest(nums []int, target int) int {
	var length int = len(nums)
	if length < 3 {
		return 0
	}
	qiuck_sort(nums, 0, length-1)
	// fmt.Println(nums)

	var result int = nums[0] + nums[1] + nums[2]
	var i, j, k int
	for i = 0; i < length-2; i++ {
		for j, k = i+1, length-1; j < k; {
			// fmt.Println(i, j, k)
			tmp := nums[i] + nums[j] + nums[k]

			if absClosest(target, tmp) < absClosest(target, result) {
				result = tmp
				// fmt.Println("aaa", result)
			}
			if tmp > target {
				// fmt.Println("1...")
				k--
			} else if tmp < target {
				// fmt.Println("2...")
				j++
			} else {
				// fmt.Println("3...")
				return target
			}
		}
	}

	return result
}

func main() {
	var arr []int = []int{-1, 2, 1, -4}
	fmt.Println(threeSumClosest(arr, 1))
}
