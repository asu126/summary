/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var carry,sum int = 0, 0
    var root,current *ListNode;

    for l1 != nil && l2 != nil {
        sum = l1.Val + l2.Val + carry
        carry = sum/10
       // fmt.Println(sum%10)
        temp := &ListNode{sum%10, nil}
        if root == nil {
            root  = temp
            current = root
        }else{
            current.Next = temp
            current = temp
        }
        l1 = l1.Next
        l2 = l2.Next
    }
    
    for l1 != nil  {
        sum = l1.Val + carry
        carry = sum/10
        temp := &ListNode{sum%10, nil}
        if root == nil {
            root  = temp
            current = root
        }else{
            current.Next = temp
            current = temp
        }
        l1 = l1.Next
    }
    
    for l2 != nil  {
        sum = l2.Val + carry
        carry = sum/10
        temp := &ListNode{sum%10, nil}
       if root == nil {
            root  = temp
            current = root
        }else{
            current.Next = temp
            current = temp
        }
        l2 = l2.Next
    }
    
    if carry != 0 {
       temp := &ListNode{carry, nil}
       if root == nil {
            root  = temp
            current = root
        }else{
            current.Next = temp
            current = temp
        }
    }
    return root
}
