---
title: 0227. Basic Calculator II
weight: 10
tags: [
	"Stack"
]
---

## Description
> Given a string s which represents an expression, evaluate this expression and return its value. 
> 
> The integer division should truncate toward zero.
> 
> You may assume that the given expression is always valid. All intermediate results will be in the range of {{< katex >}}[-2^31, 2^31 - 1]{{< /katex >}}.
> 
> Note: You are not allowed to use any built-in function which evaluates strings as mathematical expressions, such as eval().

## Solutions
这个题目因为不涉及乘除法，相比第 224 题就简单很多了，老老实实使用栈来求值就好了。这里也可以先将中缀表达式转换为后缀表达式，然后再对后缀表达式进行求值。
```
声明 Q：输出队列
声明 S：操作符栈

遍历中缀表达式中的每一个 token x：
	- 如果 x 是一个操作数，则直接将 x 追加到输出队列 Q 末尾，否则往下检查；
	- 如果 x 是一个操作符：
		- 如果操作符栈 S 栈顶为一个优先级大于等于 x 的操作符，则将 S 栈顶的操作符弹出并且追加到输出队列 Q 末尾，最后将 x 压入栈顶；
		- 如果操作符栈 S 栈顶为一个优先级小于等于 x 的操作符，则直接将 x 压入栈顶即可。
最后将栈 S 中得到的元素全部依次弹出并且入队 Q 即可。
```

TODO


