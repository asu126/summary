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
	var in1, in2, out1 [64]int
	var flag1, flag2 bool = false, false
	var out *ListNode = nil
	var len1, len2, len3 = 0, 0, 0
	var t *ListNode = l1
	for t != nil {
		in1[len1] = t.Val
		// fmt.Println(t.Val)
		t = t.Next
		len1++
		if len1 > 63 {
			if t.Next != nil {
				flag1 = true
				break
			}
		}
	}

	t = l2
	for t != nil {
		in2[len2] = t.Val
		// fmt.Println(t.Val)
		t = t.Next
		len2++
		if len2 > 63 {
			if t.Next != nil {
				flag2 = true
				break
			}
		}
	}

	var i, sum1, sum2, tmp int = 0, 0, 0, 1

	for i = 0; i < len1; i++ {
		sum1 = (in1[i] * tmp) + sum1
		tmp *= 10
	}

	tmp = 1
	for i = 0; i < len2; i++ {
		sum2 = (in2[i] * tmp) + sum2
		tmp *= 10
	}

	if flag2 {
		sum2 = 0
	}
	if flag1 {
		sum1 = 0
	}
	sum := sum1 + sum2
	shang := sum / 10
	yu := sum % 10
	out1[len3] = yu
	len3++
	for shang != 0 {
		yu = shang % 10
		shang = shang / 10

		out1[len3] = yu
		len3++
	}

	out = &ListNode{
		out1[0],
		nil,
	}

	index := out
	for i = 1; i < len3; i++ {
		t = &ListNode{
			out1[i],
			nil,
		}
		index.Next = t
		index = t
	}

	return out
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
		9,
		nil,
	}
	l22 := &ListNode{
		9,
		l21,
	}
	l2 := &ListNode{
		9,
		l22,
	}
	fmt.Println(l2)

	tt := addTwoNumbers(l1, l2)
	for tt != nil {
		fmt.Println(tt.Val)
		tt = tt.Next
	}
}
