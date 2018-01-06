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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	var index, now, berore *ListNode = head, head, head
	var i, length int = 1, 0

	for now != nil {
		length++
		now = now.Next
	}

	// remove th form head
	remTh := length - n

	now = head
	//remove first
	if remTh == 0 {
		if now.Next == nil {
			return nil
		} else {
			index = now.Next
			return index
		}
	}

	// remove not first
	now = head.Next
	for now != nil {
		if i == remTh {
			if now.Next != nil {
				berore.Next = now.Next
			} else {

				berore.Next = nil
			}
			return index
		} else {
			berore = now
			now = now.Next
		}
		i++
	}
	return index
}

func main() {
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

	tt := removeNthFromEnd(l2, 1)
	for tt != nil {
		fmt.Println(tt.Val)
		tt = tt.Next
	}
}
