/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type MinStack struct {
	Head *Node
	// Next *Node
}

/** initialize your data structure here. */
func Constructor() MinStack {
	s := MinStack{nil}
	return s
}

func (this *MinStack) Push(x int) {
	no := &Node{
		x,
		nil,
	}

	tmp := this.Head
	if this.Head == nil {
		this.Head = no
		return
	}
	for tmp.next != nil {
		tmp = tmp.next
	}
	tmp.next = no
}

func (this *MinStack) Pop() {
	tmp := this.Head
	var before *Node = nil
	if tmp == nil {
		return
	}
	if tmp.next == nil {
		this.Head = nil
		return
	}
	for tmp.next != nil {
		before = tmp
		tmp = tmp.next
	}
	before.next = nil
}

func (this *MinStack) Top() int {
	tmp := this.Head
	var t int = 0

	for tmp != nil {
		t = tmp.data
		tmp = tmp.next
	}
	return t
}

func (this *MinStack) GetMin() int {
	tmp := this.Head
	var min int

	if tmp == nil {
		min = 0
	} else {
		min = tmp.data
	}

	for tmp.next != nil {
		tmp = tmp.next
		if min > tmp.data {
			min = tmp.data
		}
	}

	return min
}

func (this *MinStack) Print() {
	tmp := this.Head

	for tmp != nil {
		fmt.Println(tmp.data)
		tmp = tmp.next
	}
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	obj := Constructor()
	obj.Push(2)
	obj.Push(0)
	obj.Push(3)
	obj.Push(0)
	obj.Print()
	fmt.Println("------GetMin-------")
	fmt.Println(obj.GetMin())
	fmt.Println("------pop-------")
	obj.Pop()
	fmt.Println("------GetMin-------")
	fmt.Println(obj.GetMin())
	fmt.Println("------pop-------")
	obj.Pop()
	fmt.Println("------GetMin-------")
	fmt.Println(obj.GetMin())
	fmt.Println("------pop-------")
	obj.Pop()
	// fmt.Println("------top-------")
	// fmt.Println(obj.Top())

	// fmt.Println("------****-------")
	// obj.Print()
	// fmt.Println("------top-------")
	// fmt.Println(obj.Top())
	fmt.Println("------GetMin-------")
	fmt.Println(obj.GetMin())
}
