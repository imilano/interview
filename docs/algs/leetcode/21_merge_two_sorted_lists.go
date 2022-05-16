package leetcode

/*
You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.
*/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := new(ListNode)
	node := res
	dummyHead1, dummyHead2 := list1, list2
	for dummyHead1 != nil && dummyHead2 != nil {
		t := new(ListNode)
		if dummyHead1.Val <= dummyHead2.Val {
			t.Val = dummyHead1.Val
			dummyHead1 = dummyHead1.Next

		} else {
			t.Val = dummyHead2.Val
			dummyHead2 = dummyHead2.Next
		}
		node.Next = t
		node = node.Next
	}

	for dummyHead1 != nil {
		t := new(ListNode)
		t.Val = dummyHead1.Val
		node.Next = t
		node = node.Next
		dummyHead1 = dummyHead1.Next
	}

	for dummyHead2 != nil {
		t := new(ListNode)
		t.Val = dummyHead2.Val
		node.Next = t
		node = node.Next
		dummyHead2 = dummyHead2.Next
	}

	return res.Next
}
