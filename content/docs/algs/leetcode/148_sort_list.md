---
title: 0148. Sort List
weight: 10
---

## Description
> Given the head of a linked list, return the list after sorting it in ascending order.

## Solutions


第一个想法很简单，先扫描一遍链表，把扫描到的值记录到数组，然后对数组进行排序，最后再把排序结果赋值给原链表即可。时间复杂度{{< katex >}}\Omicron(n\log n){{< /katex >}}

, 空间复杂度 $\Omicron(n)$。

方法二，可以用插入排序的思想。维持一个排好序的链表，从头结点开始扫描，每扫描到一个节点，就将其插入到这个有序链表中去。这样空间复杂度就是 $\Omicron(1)$，时间复杂度$\Omicron(n^2)$。

follow up 但是这里要求的是时间复杂度$\Omicron(n\log n)$，空间复杂度 $\Omicron(1)$，那很明显上述解法都不满足。这里竟然可以用归并排序，归并排序的时间复杂度是$\Omicron(n\logn)$，归并排序又分为自顶向下和自底向上两种，前者空间复杂度是$\Omicron(\log n)$（因为栈深度），后者可以达到$\Omicron(1)$。

这里是自顶向下的解法，需要注意的是，在 getMiddle 函数中将 middle 节点和前一个节点断开是一个非常重要的操作。
```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    return mergeSortUpDown(head)
}

func mergeSortUpDown(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    mid := getMiddle(head)
    left := mergeSortUpDown(head)
    right := mergeSortUpDown(mid)
    return mergeUpDown(left, right)
}

func mergeUpDown(left,right *ListNode) *ListNode {
    dummy := new(ListNode)
    cur := dummy
    // 因为 mid 节点和之前的节点是断开的，所以这里可以直接用 left 和 nil 进行判断
    for left != nil && right != nil {
        if left.Val < right.Val {
            cur.Next = left
            left = left.Next
        } else {
            cur.Next = right
            right = right.Next
        }
        
        cur = cur.Next 
    }
    
    // 另一半已经是有序的了，所以这里直接进行连接就好
    if left != nil {
        cur.Next = left
    } else {
        cur.Next = right
    }
    
    return dummy.Next
}

// 这里不仅仅需要得到 middle 节点，还需要把 middle 节点开始的链路和前一个节点结尾的链路断开
func getMiddle(head *ListNode) *ListNode {
    var pre *ListNode
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        pre = slow
        slow = slow.Next
        fast = fast.Next.Next
    }
    
    // 注意这里需要断开连接
    pre.Next = nil
    return slow
}
```

这里的自底向上的解法写起来相当 tricky， 不像是一个 middle 题该有的难度，所以暂时跳过了。
