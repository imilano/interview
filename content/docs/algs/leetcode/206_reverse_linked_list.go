package leetcode

/*
Given the head of a singly linked list, reverse the list, and return the reversed list.
*/

func reverseList(head *ListNode) *ListNode {
	return reverseListSolution3(head)
}

// 迭代解法
func reverseListSolution1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var newHead *ListNode
	for head != nil {
		t := head.Next
		head.Next = newHead
		newHead = head
		head = t
	}

	return newHead
}

// 基于双指针的迭代解法
func reverseListSolution2(head *ListNode) *ListNode {
	// 下面的写法更易于理解，但是在golang中存在语法问题
	// cur, next := nil, head
	var cur *ListNode
	var next = head
	for next != nil {
		t := next.Next
		next.Next = cur
		cur = next
		next = t
	}

	return cur
}

// 递归解法
func reverseListSolution3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	ret := reverseListSolution3(head.Next)
	head.Next.Next = head
	head.Next = nil
	return ret
}
