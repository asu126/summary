/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    var array []*ListNode = make([]*ListNode, 0, 100)
    
    var next *ListNode = head
    for next != nil {
        array =append(array, next)
        next = next.Next
    }
    
    var length int =len(array)
    
    if length <= 1 {
        return head
    } else{
        return array[length/2]
    }

}