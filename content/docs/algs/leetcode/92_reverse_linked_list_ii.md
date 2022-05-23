---
title: 0092. Reverse Linked List II
weight: 10
tags: [
	"Linked List",
	"Recursive",
	"Stack"
]
---

## Description
> Given the head of a singly linked list and two integers left and right where left <= right, reverse the nodes of the list from position left to position right, and return the reversed list.
## Solutions

使用头插法来翻转节点：先找到待翻转节点的前一个节点pre，pre 的下一个几点就是要翻转的第一个节点 cur，使用一个节点 t 表示 cur 的下一个节点，防止断链。然后将 cur 连接到 t 的下一个节点上，然后将 t 的下一个节点设置为 cur，然后再将 pre 指向 t，这样就完成了一个节点的翻转，然后继续翻转下面的节点即可。
```go
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    if head == nil || right < left {
		return nil
	}

    // 试用 dummy 防止出现需要翻转第一个节点的情况
	dummy := new(ListNode)
	dummy.Next = head
    // 找到要翻转的节点的前一个节点
	pre := dummy
	for i := 0; i < left-1;i++ {
		pre = pre.Next
	}

    // 翻转链表
	cur := pre.Next
	for i := left; i < right; i++ {
        // 注意这里的翻转方式
		t := cur.Next
		cur.Next = t.Next
        t.Next = pre.Next
        pre.Next = t
		// pre.Next = t
		// t.Next = cur
		// pre = t
	}
	return dummy.Next
}
```

