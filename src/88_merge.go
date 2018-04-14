package main

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	var i, j, current int = m - 1, n - 1, 0

	for current = (m + n - 1); current >= 0; current-- {
		if i >= 0 && j >= 0 {
			if nums1[i] >= nums2[j] {
				nums1[current] = nums1[i]
				i--
			} else {
				nums1[current] = nums2[j]
				j--
			}
		} else if i >= 0 {
			nums1[current] = nums1[i]
			i--
		} else { // j >= 0
			nums1[current] = nums2[j]
			j--
		}

	}
}

func main() {

}
