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

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var index, next *ListNode = nil, nil

	var p1, p2 *ListNode = l1, l2
	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			if index == nil {
				index = p1
				next = p1
			} else {
				next.Next = p1
			}

			next = p1
			p1 = p1.Next
		} else {
			if index == nil {
				index = p2
				next = p2
			} else {
				next.Next = p2
			}

			next = p2
			p2 = p2.Next
		}
	}

	for p1 != nil {
		if index == nil {
			index = p1
			next = p1
		} else {
			next.Next = p1
		}

		next = p1
		p1 = p1.Next
	}

	for p2 != nil {
		if index == nil {
			index = p2
			next = p2
		} else {
			next.Next = p2
		}

		next = p2
		p2 = p2.Next
	}

	return index
}

func main() {
	// l11 := &ListNode{
	// 	3,
	// 	nil,
	// }
	// l12 := &ListNode{
	// 	4,
	// 	l11,
	// }
	l1 := &ListNode{
		1,
		nil,
	}
	fmt.Println(l1)

	l21 := &ListNode{
		3,
		nil,
	}
	l22 := &ListNode{
		2,
		l21,
	}
	l2 := &ListNode{
		1,
		l22,
	}

	// fmt.Println(l2)
	// fmt.Println(p1)
	// fmt.Println(p2)

	tt := mergeTwoLists(l1, l2)
	for tt != nil {
		fmt.Println(tt.Val)
		tt = tt.Next
	}
}
