---
title: 1249. Minimum Remove to Make Valid Parentheses
weight: 10
tags: [
	"Stack"
]
---
## Description
> Given a string s of '(' , ')' and lowercase English characters.
> 
> Your task is to remove the minimum number of parentheses ( '(' or ')', in any positions ) so that the resulting parentheses string is valid and return any valid string.
> 
> Formally, a parentheses string is valid if and only if:
> 
> - It is the empty string, contains only lowercase characters, or
> - It can be written as AB (A concatenated with B), where A and B are valid strings, or
> - It can be written as (A), where A is a valid string.

## Solutions
括号合法性检测还是需要用到栈，这里用一个栈来保存那些不合法的括号（也就是最终字符串中不能包含的括号）。第一次检测时，遍历每个字符，如果该字符是括号，那么将其入栈，然后检测栈顶的两个字符是否能组成合法的一对括号，如果可以，则将该对括号出栈，然后继续遍历。遍历之后有可能会出现一种情况，那就是输入字符串全为括号，并且该字符串中不包含合法括号，那么此时栈长度就是输入字符串的长度，此时需要返回空字符串，这算是一个 corner case，需要特别处理。最后再遍历以此字符串，将出现栈中的字符排除到最终的输出字符串中即可。
```go
func minRemoveToMakeValid(s string) string {
    var stack []int
    
    for idx, _ := range s {
        if s[idx] != '(' && s[idx] != ')' {
            continue
        }
        
        stack = append(stack, idx)
        size := len(stack)
        if size >= 2 && s[stack[size-1]] ==')' && s[stack[size-2]] == '(' {
            stack = stack[:size-2]
        }
    }
    

    var res []rune
    size := len(stack)
    if size == len(s) {
        return ""
    }
    for idx, r := range s {
        if size != 0 && stack[0] == idx {
            stack = stack[1:]
            size--
            continue
        }
        
        res = append(res, r)
    }
    
    return string(res)
}


```
