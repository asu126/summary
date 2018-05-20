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

func thirdMax(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	h := &IntHeap{}
	heap.Init(h)
	for i := 0; i < len(nums); i++ {
		heap.Push(h, nums[i])
	}

	var count, head, result int = 1, 0, 0
	head = heap.Pop(h).(int)
	result = head

	for h.Len() > 0 {
		tmp := heap.Pop(h).(int)

		if tmp != result {
			count++
			result = tmp
		}
		if count == 3 {
			return result
		}
	}

	return head
}

func main() {
	array := []int{1, 3, 5}
	fmt.Println(thirdMax(array))
}
