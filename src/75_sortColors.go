package main

import (
	"fmt"
)

func sortColors(nums []int) {
	//  0、 1 和 2 分别表示红色、白色和蓝色。
	var red, white, blue int
	for i := 0; i < len(nums); i++ {
		switch nums[i] {
		case 0:
			red++
		case 1:
			white++
		case 2:
			blue++
		}
	}

	var i int = 0
	for j := 0; j < red; j++ {
		nums[i] = 0
		i++
	}
	for j := 0; j < white; j++ {
		nums[i] = 1
		i++
	}
	for j := 0; j < blue; j++ {
		nums[i] = 2
		i++
	}
}

func main() {
	// [2,0,2,1,1,0]
	var arr1 []int = []int{2, 0, 2, 1, 1, 0}
	sortColors(arr1)
	fmt.Println(arr1)

	var arr2 []int = []int{2, 0}
	sortColors(arr2)
	fmt.Println(arr2)
}
