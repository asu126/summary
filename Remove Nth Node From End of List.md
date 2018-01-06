# Remove Nth Node From End of List

Given a linked list, remove the nth node from the end of list and return its head.

For example,
```
   Given linked list: 1->2->3->4->5, and n = 2.

   After removing the second node from the end, the linked list becomes 1->2->3->5.
```

Note:
Given n will always be valid.
Try to do this in one pass.

题解：这一题也是`链表指针`题目，首先，读清楚题目：要求`删除倒数`第N个节点。由于链表是单向的，所以，我们的在遍历的过程中知道需要删除那个节点。
为此：
  1. 通过便利找出长度，进而得到删除从头开始的第几个元素；
  2. 删除开头的第几个元素需要知道他的前一个节点，所以before 节点至关重要，另外，也要考虑删除的是不是头结点；
  3. 最后一点，题目已经清楚的说了对于Int 无需考虑输入的合法性。

[code](src/removeNthNodeList.go)
