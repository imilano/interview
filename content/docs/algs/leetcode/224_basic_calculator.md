---
title: 0224. Basic Calculator
weight: 10
tags: [
	"Stack"
]
---

## Description 
> Given a string s representing a valid expression, implement a basic calculator to evaluate it, and return the result of the evaluation.
> 
> Note: You are not allowed to use any built-in function which evaluates strings as mathematical expressions, such as eval().

## Solutions

第一个想法就是，可以先把中缀表达式转换为后缀表达式，然后再对后缀表达式进行计算，这样就会简单很多。那么问题是，如果将一个带有括号的中缀表达式转换为后缀表达式呢？这里的精简版算法如下：
```
声明 Q：输出队列
声明 S：操作符栈

遍历中缀表达式中的每一个 token x：
	- 如果 x 是一个操作数，则直接将 x 追加到输出队列 Q 末尾，否则往下检查；
	- 如果 x 是一个左括号"("，则将 x 压入操作符栈，否则往下检查；
	- 如果 x 是一个操作符：
		- 如果操作符栈 S 栈顶为一个优先级大于等于 x 的操作符，则将 S 栈顶的操作符弹出并且追加到输出队列 Q 末尾，最后将 x 压入栈顶；
		- 如果操作符栈 S 栈顶为一个优先级小于等于 x 的操作符，或者不为操作符（这里只可能是左括号"（"）,则直接将 x 压入栈顶即可。
	- 如果 x 是一个右括号，则将操作符栈 S 栈顶往下到第一个左括号之间的元素以此弹出炳且追加到输出队列末尾，然后将左括号丢弃，右括号也不用入栈。注意，如果栈到底后仍然没有找到左括号，则说明表达式不合法，左右括号不匹配。
最后将栈 S 中得到的元素全部依次弹出并且入队 Q 即可。
```

TODO
