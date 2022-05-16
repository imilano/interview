package leetcode

/*
You are given the head of a singly linked-list. The list can be represented as:

L0 → L1 → … → Ln - 1 → Ln
Reorder the list to be on the following form:

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
You may not modify the values in the list's nodes. Only nodes themselves may be changed.
*/

// 分为三个阶段：
// 1. 使用快慢指针找到链表中点，并将链表从中点处断开，形成两个独立的链表
// 2. 将第二个链表翻转
// 3. 将第二个链表间隔的插入到第一个链表中
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}

	fast, slow := head, head
	// TODO  注意这里的终止条件
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	mid := slow.Next
	// TODO 这里一定要注意把链表断开，不然会出现环
	slow.Next = nil
	last := mid
	var pre *ListNode = nil
	for last != nil {
		next := last.Next
		last.Next = pre
		pre = last
		last = next
	}

	for head != nil && pre != nil {
		next := head.Next
		head.Next = pre
		pre = pre.Next
		head.Next.Next = next
		head = next
	}
}
