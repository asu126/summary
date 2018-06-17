package main

import "fmt"

import "container/heap"

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := &IntHeap{}
	heap.Init(h)
	for i := 0; i < len(nums); i++ {
		heap.Push(h, nums[i])
	}

	var count int = 0
	for h.Len() > 0 {
		tmp := heap.Pop(h).(int)

		// fmt.Println(tmp)
		count++
		if count == k {
			return tmp
		}
	}

	return 0
}

func main() {
	// 	输入: [3,2,1,5,6,4] 和 k = 2
	// 输出: 5
	array := []int{3, 2, 1, 5, 6, 4}
	fmt.Println(findKthLargest(array, 2)) // 5

	// 	输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
	// 输出: 4
	array1 := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	fmt.Println(findKthLargest(array1, 4)) // 4
}
