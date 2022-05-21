---
title: 0206. Reverse Linked List
weight: 10
tags: [
	"LinkedList",
	"Stack",
	"Recursive"
]
---

## Description
> Given the head of a singly linked list, reverse the list, and return the reversed list.

## Solutions
### Iterative
迭代法，但是要注意虚拟头节点的使用。
```go
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    dummy := new(ListNode)
    cur := head
    for cur != nil {
        tail := dummy.Next
        next := cur.Next
        dummy.Next = cur
        cur.Next = tail
        cur = next
    }
    
    return dummy.Next
}
```

### Recursive
这里递归法如何反转节点也是需要十分注意的。
```go
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    res := reverseList(head.Next)
    head.Next.Next = head
    head.Next = nil
    return res
}
```
