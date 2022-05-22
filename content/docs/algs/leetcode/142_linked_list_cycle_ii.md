---
title: 0142. Linked List Cycle II
weight: 10
tags: [
	"LinkedList",
	"Hash Table",
	"Two Pointer"
]
---

## Description

> Given the head of a linked list, return the node where the cycle begins. If there is no cycle, return null.
> 
> There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to (0-indexed). It is -1 if there is no cycle. Note that pos is not passed as a parameter.
> 
> Do not modify the linked list.

## Solutions

### Hash Table
首先也是很自然的能够想到使用哈希表。
```go
func detectCycle(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return nil
    }
    
    dict := make(map[*ListNode]bool)
    cur := head
    for cur != nil {
        if _, ok := dict[cur]; ok {
            return cur
        }
        
        dict[cur] = true
        cur = cur.Next
    }
    
    return nil
}
```

### Two Pointer
快慢指针解法，参加![图示](https://interview.lightsinger.top/images/codingInterviews/entrynode_of_looplist.png)
```go
func detectCycle(head *ListNode) *ListNode {
    // corner case
    if head == nil || head.Next == nil {
        return nil
    }
    
    // 快慢指针一直往前走，直到二者相遇，或者 fast 到达无环链表末尾
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if fast == slow {
            break
        }
    }
    
    // 如果链表没有环，则返回 nil
    if fast == nil || fast.Next == nil {
        return nil
    }
    
    // fast 和 slow 的任意一个从头结点开始，然后让二者继续向前遍历，相遇节点即为入口节点
    fast = head
    for fast != slow {
        fast = fast.Next
        slow = slow.Next
    }
    
    return fast
}
```
