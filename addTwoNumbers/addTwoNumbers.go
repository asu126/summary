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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var index, next *ListNode = nil, nil
	var sum, up int = 0, 0
	if l1 == nil || l2 == nil {
		return index
	}
	var p1, p2 *ListNode = l1, l2
	for p1 != nil && p2 != nil {
		sum = p1.Val + p2.Val + up
		tmp := &ListNode{
			sum % 10,
			nil,
		}
		up = sum / 10

		if next == nil {
			index = tmp
			next = tmp
		} else {
			next.Next = tmp
			next = tmp
		}

		p1 = p1.Next
		p2 = p2.Next
	}

	for p1 != nil {
		sum = p1.Val + up
		tmp := &ListNode{
			sum % 10,
			nil,
		}
		up = sum / 10
		next.Next = tmp
		next = tmp
		p1 = p1.Next
	}

	for p2 != nil {
		sum = p2.Val + up
		tmp := &ListNode{
			sum % 10,
			nil,
		}
		up = sum / 10
		next.Next = tmp
		next = tmp
		p2 = p2.Next
	}

	if up > 0 {
		tmp := &ListNode{
			up,
			nil,
		}
		next.Next = tmp
		next = tmp
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
		1,
		nil,
	}
	l22 := &ListNode{
		2,
		l21,
	}
	l2 := &ListNode{
		3,
		l22,
	}

	// fmt.Println(l2)
	// fmt.Println(p1)
	// fmt.Println(p2)

	tt := addTwoNumbers(l1, l2)
	for tt != nil {
		fmt.Println(tt.Val)
		tt = tt.Next
	}
}
