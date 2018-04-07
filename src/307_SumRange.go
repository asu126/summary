package main

import (
	"fmt"
)

type NumArray struct {
	data []int
}

func Constructor(nums []int) NumArray {
	return NumArray{data: nums}
}

func (this *NumArray) SumRange(i int, j int) int {
	var sum int = 0
	for i <= j {
		sum += this.data[i]
		i++
	}
	return sum
}

func (this *NumArray) Update(i int, val int) {
	if i < len(this.data) {
		this.data[i] = val
	}
	// fmt.Println(this.data)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */

func main() {
	ints := []int{7, 2, 2, 4}
	obj := Constructor(ints)
	fmt.Println(obj.SumRange(0, 3))

	obj.Update(0, 0)
	fmt.Println(obj.SumRange(0, 3))
}
