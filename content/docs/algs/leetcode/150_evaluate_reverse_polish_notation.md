---
title: 0150. Evaluate Reverse Polish Notation
weight: 10
tags: [
    "Stack"
]
---
## Description
> Evaluate the value of an arithmetic expression in Reverse Polish Notation.
> 
> Valid operators are +, -, *, and /. Each operand may be an integer or another expression.
> 
> Note that division between two integers should truncate toward zero.
> 
> It is guaranteed that the given RPN expression is always valid. That means the expression would always evaluate to a result, and there will not be any division by zero operation.



## Solutions
逆波兰表达式求值，使用栈能够很好的解决，在此几部多说了，直接上代码:
```go
// using stack
func evalRPN(tokens []string) int {
    var res int
    size := len(tokens)
    if size <= 0 {
        return res
    }
    
    var stack []int
    for _, str := range tokens {
        num, isNum := strconv.Atoi(str)
        if len(stack) == 0 || isNum == nil {
            stack = append(stack, num)
            continue
        }
        
        size := len(stack)
        op1, op2 := stack[size-2], stack[size-1]
        stack = stack[:size-2]
        switch str {
            case "+":
                stack = append(stack, op1 + op2)
            case "-" :
                stack = append(stack, op1 - op2)
            case "*":
                stack = append(stack, op1*op2)
            case "/": 
                stack = append(stack, op1/op2)
        }
    }
    
    return stack[0]
}
```
