package main

import "fmt"

func searchRange(nums []int, target int) []int {
	var length int = len(nums)
	var ret []int = []int{-1, -1}
	if length > 0 {
		var l int = searchLeft(nums, target, 0, length-1)

		if l != -1 {
			ret[0] = l
			var rightTarget int
			for rightTarget = l + 1; rightTarget < length; {
				if nums[l] == nums[rightTarget] {
					rightTarget++
				} else {
					break
				}
			}
			ret[1] = rightTarget - 1
		}
	}
	return ret

}

func searchLeft(nums []int, target int, l int, r int) int {
	middle := l + (r-l)/2
	// fmt.Println(l, r, middle)
	if l > r {
		if middle < len(nums) && nums[middle] == target {
			return l
		} else {
			return -1
		}
	}

	if nums[middle] >= target {
		return searchLeft(nums, target, l, middle-1)
	} else {
		return searchLeft(nums, target, middle+1, r)
	}

	return -1

}

func main() {
	// 给出 [5, 7, 7, 8, 8, 10] 和目标值 8，
	// 返回 [3, 4]。
	a := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(searchRange(a, 8))

	a1 := []int{5, 7, 7, 8, 10}
	fmt.Println(searchRange(a1, 8))

	a2 := []int{}
	fmt.Println(searchRange(a2, 8))

	a3 := []int{1, 2, 3}
	fmt.Println(searchRange(a3, 2))

	fmt.Println(searchRange(a3, 9))
}
