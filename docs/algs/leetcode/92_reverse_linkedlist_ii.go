package leetcode

/*
Given the head of a singly linked list and two integers left and right where left <= right, reverse the nodes of the list from position left to position right, and return the reversed list.
*/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || right < left {
		return nil
	}

	dummy := new(ListNode)
	dummy.Next = head
	pre := dummy
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	cur := pre.Next
	for i := left; i < right; i++ {
		t := cur.Next
		cur.Next = t.Next
		// 注意如何紧系功能头插入。下面的方法是错误的
		// pre.Next = t
		// t.Next = cur
		// pre = t
		t.Next = pre.Next
		pre.Next = t
	}
	return dummy.Next
}
