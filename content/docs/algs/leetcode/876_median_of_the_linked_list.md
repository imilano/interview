m--
title: 0876. Median of the Linked List
weight: 10
tags: [
	"LinkedList",
	"Two Pointer"
]
---

## Description

> Given the head of a singly linked list, return the middle node of the linked list.
> 
> If there are two middle nodes, return the second middle node.


## Solutions

使用快慢指针法，最后返回慢指针指向的节点即可。
```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    slow, fast := head, head
	// 注意这里的判断条件
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    
    return slow
}
```
