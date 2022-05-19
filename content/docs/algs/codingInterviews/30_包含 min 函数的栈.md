---
title: 30. 包含 min 函数的栈
weight: 10
---
## Description
> 定义栈的数据结构，请在该类型中实现一个能够得到栈中所含最小元素的 min 函数，输入操作时保证 pop、top 和 min 函数操作时，栈中一定有元素。
>
> 此栈包含的方法有：
> push(value):将value压入栈中
> pop():弹出栈顶元素
> top():获取栈顶元素
> min():获取栈中最小元素
> 
> 数据范围：操作数量满足 {{< katex >}}$0 \le n \le 300${{< /katex >}}  ，输入的元素满足 {{< katex >}}$|val| \le 10000${{< /katex >}} 
> 进阶：栈的各个操作的时间复杂度是 {{< katex >}}$\Omicron(1)${{< /katex >}}  ，空间复杂度是 {{< katex >}}$\Omicron(n)${{< /katex >}}。 

## Solutions

### Stack
使用两个栈，一个栈stack1用于正常压入和弹出，另一个栈stack2用于实现单调栈。当压入一个元素时，如果当前元素比 stack2 的栈顶元素还要大，则将当前元素压入 stack1 的同时，还将当前元素压入 stack2，如果当前元素不比 stack2 的栈顶元素要大，则将 stack2 的栈顶元素重复入栈。当弹出元素时，将 stack1 和 stack2 的栈顶元素同步弹出。当获取最小值时，只需要取出 stack2 的栈顶元素即可。
```go
// 用于实现正常的 push 与 pop 操作
var stack1 []int
// 用于存储最小值
var stack2 []int

func Push(node int) {
    // write code here
    stack1 = append(stack1, node)
    size := len(stack2)
    if size == 0 || stack2[size-1] > node {
        stack2 = append(stack2, node)
    } else {
        stack2 = append(stack2, stack2[size-1])
    }    
    
}
func Pop() {
    // write code here
    size := len(stack1)
    stack1 = stack1[:size-1]
    stack2 = stack2[:size-1]
}
func Top() int {
    // write code here
    size := len(stack1)
    return stack1[size-1]
}
func Min() int {
    // write code here
    size := len(stack2)
    return stack2[size-1]
}
```
