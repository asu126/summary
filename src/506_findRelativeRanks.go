package main

import (
	"container/heap"
	"fmt"
	"strconv"
)

// An IntHeap is a min-heap of ints.
type node struct {
	val      int
	location int
}

type IntHeap []node

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i].val > h[j].val }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(node))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findRelativeRanks(nums []int) []string {
	var length int = len(nums)
	var result []string = make([]string, length)
	h := &IntHeap{}
	heap.Init(h)

	for i := 0; i < length; i++ {
		heap.Push(h, node{nums[i], i})
	}

	for i := 0; i < length; i++ {
		temp := heap.Pop(h).(node)
		switch i {
		case 0:
			result[temp.location] = "Gold Medal"
		case 1:
			result[temp.location] = "Silver Medal"
		case 2:
			result[temp.location] = "Bronze Medal"
		default:
			result[temp.location] = strconv.Itoa(i + 1)
		}
	}

	return result
}

func main() {
	var arr []int = []int{5, 4, 3, 2, 1}
	fmt.Println(findRelativeRanks(arr))

	var arr1 []int = []int{5, 4}
	fmt.Println(findRelativeRanks(arr1))
}
