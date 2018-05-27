/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    for head != nil && head.Val == val {
        head = head.Next
    }
    
    if head == nil {
        return head
    }
    var current , next *ListNode =head, head.Next
    
    
    for next != nil {
        if next.Val == val {
            next = next.Next
            current.Next = next
        } else {
            current = next
            next = next.Next
        }
    }
   
    return head
}
