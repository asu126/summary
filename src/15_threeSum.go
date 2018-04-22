package main

import (
	"fmt"
	"sort"
)

type ByAge []int

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i] < a[j] }

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

func threeSum(nums []int) [][]int {
	var result [][]int = make([][]int, 0, 10240)
	var length int = len(nums)
	if length < 3 {
		return result
	}
	// qiuck_sort(nums, 0, length-1)
	sort.Sort(ByAge(nums))
	// fmt.Println(nums)

	var i, j, k int
	for i = 0; i < length-2; i++ {
		for j, k = i+1, length-1; j < k; {
			// fmt.Println(i, j, k)
			tmp := nums[i] + nums[j] + nums[k]

			if tmp == 0 {
				var subRes []int = []int{nums[i], nums[j], nums[k]}
				result = append(result, subRes)
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k-1] {
					k--
				}
			} else if tmp > 0 {
				k--
			} else {
				j++
			}
		}

		for i+1 < length && nums[i] == nums[i+1] {
			i++
		}
	}

	return result
}

func main() {
	var arr []int = []int{-1, 2, 1, -4}
	fmt.Println(threeSum(arr))

	var arr1 []int = []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(arr1))

	var arr2 []int = []int{-1, 0, 1, 0}
	fmt.Println(threeSum(arr2))
	var arr3 []int = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	fmt.Println(threeSum(arr3))
}
