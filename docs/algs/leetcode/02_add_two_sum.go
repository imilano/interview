package leetcode

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// handler nil
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1 == nil && l2 == nil {
		return nil
	}

	// find the longer number
	var longerNumber *ListNode
	dummyL1, dummyL2 := l1, l2
	for dummyL1 != nil && dummyL2 != nil {
		dummyL1 = dummyL1.Next
		dummyL2 = dummyL2.Next
	}
	if dummyL1 == nil {
		longerNumber = l2
	} else {
		longerNumber = l1
	}

	// add from the shorter one to longer one
	dummyHead := longerNumber
	dummyL1, dummyL2 = l1, l2
	var carry int
	for dummyL1 != nil && dummyL2 != nil {
		sum := (dummyL1.Val + dummyL2.Val + carry) % 10
		carry = (dummyL1.Val + dummyL2.Val + carry) / 10

		longerNumber.Val = sum
		if longerNumber.Next != nil {
			longerNumber = longerNumber.Next
		}
		dummyL1 = dummyL1.Next
		dummyL2 = dummyL2.Next
	}

	// add the reset longer one
	if dummyL1 == nil && dummyL2 != nil {
		for dummyL2 != nil {
			sum := (dummyL2.Val + carry) % 10
			carry = (dummyL2.Val + carry) / 10

			dummyL2 = dummyL2.Next
			longerNumber.Val = sum
			if longerNumber.Next != nil {
				longerNumber = longerNumber.Next
			}
		}
	}

	if dummyL2 == nil && dummyL1 != nil {
		for dummyL1 != nil {
			sum := (dummyL1.Val + carry) % 10
			carry = (dummyL1.Val + carry) / 10

			dummyL1 = dummyL1.Next
			longerNumber.Val = sum
			if longerNumber.Next != nil {
				longerNumber = longerNumber.Next
			}
		}
	}

	// if carry is not 0
	if carry != 0 {
		longerNumber.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
	}

	return dummyHead
}
