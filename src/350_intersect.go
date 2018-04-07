package main

import (
	"fmt"
)

func intersect(nums1 []int, nums2 []int) []int {
	var ret []int
	var len1, len2 = len(nums1), len(nums2)
	if len1 >= len2 {
		ret = make([]int, 0, len1)
	} else {
		ret = make([]int, 0, len2)
	}

	var flag []bool = make([]bool, len2)
	for i := 0; i < len1; i++ {
		// fmt.Println("ss")
		for j := 0; j < len2; j++ {
			if nums1[i] == nums2[j] && !flag[j] {
				tmp := len(ret)
				ret = ret[:tmp+1]
				ret[tmp] = nums1[i]
				flag[j] = true
				break
			}
		}
	}
	return ret
}

func main() {
	a := []int{1, 2, 2, 1}

	a1 := []int{2}
	// a1 := []int{2, 2}
	fmt.Println(intersect(a, a1)) // true
	// 	[1]
	// [1,1]
}
