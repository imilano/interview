---
title: 0160. Intersection of Two Linked Lists
weight: 10
tags: [
	"LinkedList",
	"Hash Table",
	"Two Pointer"
]
---

## Description
> Given the heads of two singly linked-lists headA and headB, return the node at which the two lists intersect. If the two linked lists have no intersection at all, return null.

## Solutions

### Hash Table
首先很容易想到哈希表的解法。
```go
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
        return nil
    }
    
    dict := make(map[*ListNode]bool)
    dummyA , dummyB := headA, headB
    for dummyA != nil {
        dict[dummyA] = true
        dummyA = dummyA.Next
    }
    
    for dummyB != nil {
        if _, ok := dict[dummyB]; ok {
            return dummyB
        }
        
        dummyB = dummyB.Next
    }
    
    
    return nil
}
```

### Two Pointer
其次可以想到的一个接法是，A 和 B 都不断往后走，直到走到尾部，此时让 A 从 B 开头继续走，而让 B 从 A 开头继续走，如果二者有相交节点的话，必然会在相交节点相遇。如果二者无相交节点的话，那么二者应该会同时为空。
```go
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
        return nil
    }
    
    dummyA, dummyB := headA, headB
    for {
        // 并不能确定一定有环，所以当二者同时为空时，说明无环
        if dummyA == nil && dummyB == nil {
            return nil
        }
        
        // A 继续从 B 头部开始遍历
        if dummyA == nil {
            dummyA = headB
        }
        
        // B 继续从 A 头部开始遍历
        if dummyB == nil {
            dummyB = headA
        }
        
        // 有重叠
        if dummyA == dummyB {
            return dummyA
        }
        
        dummyA = dummyA.Next
        dummyB = dummyB.Next
    }
    
    return nil
}
```
