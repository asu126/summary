package main

import (
	"fmt"
)

type NumArray struct {
	data []int
}

func Constructor(nums []int) NumArray {
	return NumArray{data: nums}
	// fmt.Println("sss")
	// return NumArray{data: []int{1, 2, 3, 4}}
}

func (this *NumArray) SumRange(i int, j int) int {
	var sum int = 0
	for i <= j {
		sum += this.data[i]
		i++
	}
	return sum
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */

func main() {
	ints := []int{7, 2, 2, 4}
	obj := Constructor(ints)
	fmt.Println(obj.SumRange(0, 3))
}
