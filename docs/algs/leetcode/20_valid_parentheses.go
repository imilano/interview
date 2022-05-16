package leetcode

/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
*/

func isValid(s string) bool {
	size := len(s)
	if size == 0 {
		return true
	}
	if size == 1 {
		return false
	}

	stack := new(Stack)
	for i := 0; i < size; i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack.Push(s[i])
			continue
		}

		ele := stack.Top()
		if ele == nil {
			return false
		} else {
			ele = ele.(uint8)
		}
		if (s[i] == '}' && ele != uint8('{')) || (s[i] == ']' && ele != uint8('[')) || (s[i] == ')' && ele != uint8('(')) {
			return false
		} else {
			stack.Pop()
		}
	}

	return stack.Len() == 0
}

// for testing
func ValidParentheses(s string) bool {
	return isValid(s)
}
