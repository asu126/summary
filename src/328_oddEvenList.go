/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    var count int = 0
    var pJiShu, pOuShu, pOuShuHead, next *ListNode
    next = head
    
    count++
    pJiShu = next
    next = next.Next
    
    if next != nil {
        count++
        pOuShu = next
        pOuShuHead = next
        next = next.Next
    }

    for next != nil {
        //fmt.Println(next.Val)
        count++
        if count%2 == 0{
            pOuShu.Next = next
            pOuShu = next
            next = next.Next
        }else{
            pJiShu.Next = next
            pJiShu = next
            next = next.Next
            pJiShu.Next = pOuShuHead
        }
        pOuShu.Next = nil
    }
    
    
    return head
}
