package leetcode

/*
Given the head of a linked list, remove the nth node from the end of the list and return its head.
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	return removeNthFromEndSolution2(head, n)
}

// 很简单的递归解法，如果遍历到倒数第n个数，则直接跳过即可。
func removeNthFromEndSolution1(head *ListNode, n *int) *ListNode {
	if head == nil || head.Next == nil {
		*n -= 1
		if *n == 0 {
			return nil
		}
		return head
	}

	next := removeNthFromEndSolution1(head.Next, n)
	*n -= 1
	if *n == 0 {
		return next
	}

	head.Next = next
	return head
}

// 遍历解法, 双指针
func removeNthFromEndSolution2(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}

	pre, cur := head, head
	for i := 0; i < n; i++ {
		cur = cur.Next
	}
	if cur == nil {
		return head.Next
	}

	for cur.Next != nil {
		cur = cur.Next
		pre = pre.Next
	}

	pre.Next = pre.Next.Next
	return head
}
