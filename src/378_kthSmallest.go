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

func kthSmallest(matrix [][]int, k int) int {
	var count int = 0
	var N int = len(matrix)
	if N < 1 {
		return 0
	}

	h := &IntHeap{}
	heap.Init(h)

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			heap.Push(h, matrix[i][j])
			count++
			if count > k {
				heap.Pop(h)
				count--
			}
		}
	}

	return heap.Pop(h).(int)
}

func main() {
	// 	matrix = [
	//    [ 1,  5,  9],
	//    [10, 11, 13],
	//    [12, 13, 15]
	// ],
	// k = 8,

	// 返回 13
	array := [][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}
	fmt.Println(kthSmallest(array, 8))

	//  matrix = [
	//    [ 1,  5,  9],
	//    [10, 11, 13],
	//    [12, 13, 15]
	// ],
	// k = 8,

	// 返回 13
	array1 := [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}
	fmt.Println(kthSmallest(array1, 3))
}
