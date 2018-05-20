package main

import "fmt"
import (
	"sort"
)

type Array []int

func (a Array) Len() int           { return len(a) }
func (a Array) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Array) Less(i, j int) bool { return a[i] < a[j] }

func arrayPairSum(nums []int) int {
	var length int = len(nums)
	sort.Sort(Array(nums))
	// fmt.Println(nums)
	var result int = 0

	for i := 0; i < length; {
		result += nums[i]
		i += 2
	}

	return result
}

func main() {
	fmt.Println(arrayPairSum([]int{1, 4, 3, 2})) // 4

	fmt.Println(arrayPairSum([]int{1, 1, 3, 2})) // 3
}
