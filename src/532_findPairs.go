package main

import "fmt"
import (
	"sort"
)

type Array []int

func (a Array) Len() int           { return len(a) }
func (a Array) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Array) Less(i, j int) bool { return a[i] < a[j] }

func findPairs(nums []int, k int) int {
	var length int = len(nums)
	sort.Sort(Array(nums))
	// fmt.Println(nums)
	var result int = 0

	for i := 0; i < length; i++ {
		if i > 0 && nums[i] == nums[i-1] { // quchong
			continue
		}
		for j := i + 1; j < length; j++ {
			if j > i+1 && nums[j] == nums[j-1] { // quchong
				continue
			}

			tmp := nums[j] - nums[i]
			if tmp == k {
				// fmt.Println(j, i, nums[j], nums[i])
				result++
			} else if tmp > k {
				break
			}
		}
	}

	return result
}

func main() {
	var arr []int = []int{3, 1, 4, 1, 5}
	fmt.Println(findPairs(arr, 2)) // 2

	var arr1 []int = []int{1, 2, 3, 4, 5}
	fmt.Println(findPairs(arr1, 1)) // 4

	var arr2 []int = []int{1, 3, 1, 5, 4}
	fmt.Println(findPairs(arr2, 0)) //1

	var arr3 []int = []int{1, 3, 1, 5, 4, 1, 1}
	fmt.Println(findPairs(arr3, 0)) //1

	var arr4 []int = []int{1, 1, 1, 2, 2}
	fmt.Println(findPairs(arr4, 1)) //1
}
