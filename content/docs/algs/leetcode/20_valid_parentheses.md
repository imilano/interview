---
title: 0020. Valid Parentheses
weight: 10
tags: [
	"Stack"
]
---

## Description
> Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
> 
> An input string is valid if:
> 
> - Open brackets must be closed by the same type of brackets.
> - Open brackets must be closed in the correct order.

## Solutions
简单题，直接使用栈即可。
```go
func isValid(s string) bool {
    var stack []byte
    size := len(s)
    for i := 0; i < size; i++ {
        if s[i] == '(' || s[i] == '[' || s[i] == '{' {
            stack = append(stack, s[i])
        } else {
            cap := len(stack)
            if cap == 0 {
                return false
            }
            
            tail := stack[cap-1]
            if s[i] == ')' && tail == '(' || s[i] ==']' && tail == '[' || s[i] == '}' && tail == '{' {
                stack = stack[:cap-1]
            } else {
                return false
            }
        }
    }
    
    return len(stack) == 0
}
```
