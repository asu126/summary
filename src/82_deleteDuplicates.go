/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    m := make(map[int]int)
    var node *ListNode = head
    
    for node != nil {
        m[node.Val]++
        node = node.Next
    }
    
    var root, pre *ListNode = nil, nil
    node = head

    node = head
    for node != nil {
        if m[node.Val] == 1{
             temp := &ListNode{Val: node.Val, Next: nil }
            if root == nil{
                root = temp
                pre = temp
            } else {
                 pre.Next = temp
                pre = temp
            }
        }
        node = node.Next
    }
    
    return root
}
