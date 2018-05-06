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

func findPoisonedDuration(timeSeries []int, duration int) int {
	var length int = len(timeSeries)
	qiuck_sort(timeSeries, 0, length-1)

	var allTime = 0
	if length < 1 {
		return allTime
	}

	for i := 1; i < length; i++ {
		if timeSeries[i]-timeSeries[i-1] >= duration {
			allTime += duration
		} else {
			allTime += (timeSeries[i] - timeSeries[i-1])
		}
	}

	allTime += duration
	return allTime
}

func main() {
	var arr []int = []int{1, 4}
	fmt.Println(findPoisonedDuration(arr, 2)) // 4
	var arr1 []int = []int{1, 2}
	fmt.Println(findPoisonedDuration(arr1, 2)) //3

	var arr2 []int = []int{}
	fmt.Println(findPoisonedDuration(arr2, 2)) //0

	var arr3 []int = []int{1, 1, 1, 5}
	fmt.Println(findPoisonedDuration(arr3, 2)) //4
}
