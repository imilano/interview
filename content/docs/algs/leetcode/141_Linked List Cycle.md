---
title: 0141. Linked List Cycle
weight: 10
tags: [
	"LinkedList",
	"Hash Table",
	"Two Pointer"
]
---

## Description
> Given head, the head of a linked list, determine if the linked list has a cycle in it.
> 
> There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.
> 
> Return true if there is a cycle in the linked list. Otherwise, return false.

## Solutions

### Hash Table
很简单，使用哈希表记录一个节点是否被访问过两次即可。需要注意的是，题目并不保证给出的链表一定是有环的。
```go
func hasCycle(head *ListNode) bool {
    var res bool
    if head == nil || head.Next == nil {
        return res
    }
    
    dict := make(map[*ListNode]bool)
    cur := head
    for cur != nil {
        if _, ok := dict[cur]; ok {
            res = true
            break
        }
        
        dict[cur] = true
        cur = cur.Next
    }
    
    return res
}
```

### Two Pointer
使用快慢指针，如果链表有环，则快慢指针必然相交。
```go
func hasCycle(head *ListNode) bool {
    var res bool
    if head == nil || head.Next == nil {
        return res
    }
    
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
        if slow == fast {
            res = true
            break
        }
    }
    
    return res
}
```
