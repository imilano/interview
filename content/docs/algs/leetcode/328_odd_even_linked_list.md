---
title: 0328. Odd Even Linked List
weight: 10
tags: [
	"LinkedList",
	"Two Pointer"
]
---

## Description
> Given the head of a singly linked list, group all the nodes with odd indices together followed by the nodes with even indices, and return the reordered list.
> 
> The first node is considered odd, and the second node is even, and so on.
> 
> Note that the relative order inside both the even and odd groups should remain as it was in the input.
> 
> You must solve the problem in O(1) extra space complexity and O(n) time complexity.

## Solutions

双指针。使用两个奇偶指针分别指向奇偶节点的起始位置，另外需要一个单独的指针 even_head 保存偶节点的起点位置。然后把奇节点的下一个位置指向偶节点的下一个位置（必然是奇节点），奇节点指针后移一步；再把偶节点的下一个位置指向奇节点的下一个问题（必然是偶节点），然后偶节点后移一步。最后把分开的偶节点连在奇节点后即可。
```go
func oddEvenList(head *ListNode) *ListNode {
	// corner case
    if head == nil || head.Next == nil || head.Next.Next == nil {
        return head
    }
    
	// 奇偶链表尾结点
    odd, even := head, head.Next
	// even_head 指向偶链表起始节点
    even_head := even
	// 注意这里的结束条件
    for even != nil && even.Next != nil {
		// 奇链表的下一个是偶链表尾结点的下一个
        odd.Next = even.Next
		// 奇链表尾结点后移
        odd = even.Next
		// 偶链表的下一个是奇链表尾结点的下一个
        even.Next = odd.Next
		// 偶链表尾结点后移
        even = odd.Next
    }
    
	// 把偶链表链接到奇链表后面
    odd.Next = even_head
    return head
}
```
