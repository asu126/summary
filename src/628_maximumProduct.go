package main

import (
	"fmt"
	"sort"
)

type IncSort []int

func (a IncSort) Len() int           { return len(a) }
func (a IncSort) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a IncSort) Less(i, j int) bool { return a[i] < a[j] }

func maximumProduct(nums []int) int {
	var length int = len(nums)
	if length < 3 {
		return 0
	}
	sort.Sort(IncSort(nums))
	var max1 int = nums[length-1] * nums[length-2] * nums[length-3]
	var max2 int = nums[length-1] * nums[0] * nums[1]
	if max1 >= max2 {
		return max1
	} else {
		return max2
	}
}

func main() {
	var arr []int = []int{-1, 2, 1, -4}
	fmt.Println(maximumProduct(arr)) // 8

	var arr1 []int = []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(maximumProduct(arr1)) // 8

	var arr2 []int = []int{-1, 0, 1, 0}
	fmt.Println(maximumProduct(arr2)) // 0

	var arr3 []int = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	fmt.Println(maximumProduct(arr3)) // 0
}
