package main

import (
	"fmt"
)

func intersection(nums1 []int, nums2 []int) []int {
	var ret []int
	var len1, len2 = len(nums1), len(nums2)
	if len1 >= len2 {
		ret = make([]int, 0, len1)
	} else {
		ret = make([]int, 0, len2)
	}

	for i := 0; i < len1; i++ {
		fmt.Println("ss")
		for j := 0; j < len2; j++ {
			if nums1[i] == nums2[j] && !isExist(ret, nums1[i]) {
				tmp := len(ret)
				ret = ret[:tmp+1]
				ret[tmp] = nums1[i]
			}
		}
	}
	return ret
}

func isExist(nums []int, n int) bool {
	for i := 0; i < len(nums); i++ {
		if nums[i] == n {
			return true
		}
	}

	return false
}

func main() {
	a := []int{}

	a1 := []int{1}
	fmt.Println(intersection(a, a1)) // true
}
