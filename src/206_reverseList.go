/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var current, next  *ListNode = head, nil
    
    if current != nil {
        next = current.Next
        current.Next = nil 
    }

    for current != nil {
        if next == nil {
            head= current
            break
        }
        // fmt.Println(current.Val)
        temp := next.Next
        next.Next = current
        current = next
        next = temp
    }
    
    return head
}
