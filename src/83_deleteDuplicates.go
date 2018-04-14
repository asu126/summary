/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    
    var root, next *ListNode= head, head.Next
    
    for next != nil{
        if head.Val == next.Val {
            head.Next = next.Next
            next = head.Next
        }else{
            head = next
            next = next.Next
        }
    }
    
    return root
}
