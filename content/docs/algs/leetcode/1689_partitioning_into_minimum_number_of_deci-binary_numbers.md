---
title: 1689. Partitioning Into Minimum Number of Deci-Binary Numbers
weight: 10
tags: [
	"String",
	"Math"
]
---

## Description
> A decimal number is called deci-binary if each of its digits is either 0 or 1 without any leading zeros. For example, 101 and 1100 are deci-binary, while 112 and 3001 are not.
> 
> Given a string n that represents a positive decimal integer, return the minimum number of positive deci-binary numbers needed so that they sum up to n.


## Solutions
这题题主以为是要用动态规划之类的解法来做，结果想了好一会都没想出来，然后到网上去搜大神们答案，于是搜到了[这个](https://leetcode.com/problems/partitioning-into-minimum-number-of-deci-binary-numbers/discuss/970318/JavaC++Python-Just-return-max-digit)。看完答案，我只想说，人和人真的，天壤之别...

这里的解法浓缩成一句话就是：答案就是 n 这个字符串中最大的数字（指 0-9 之间）。翻译一下大佬的证明就是：
> 假设 n 中最大的数字是 x。以内 deci-binary 中只能包含 0 和 1，所以我们至少需要 x 个数字加起来才能达到数字 x（这里的意思就是，假设一个数字是 5，那么因为 deci-binary 中最大的数只能为 1，所以达到和为 5 的话，至少也需要 5 个数字，这样应该稍微好理解一些）。
> 
> 以 n = 135 为例，我们初始化 5 个 deci-binary 数（因为 135 中最大的数字是 5），每个数长度为 3（长度要与字符串 n 的长度相同，故而 3）：
> - a1 = 000
> - a2 = 000
> - a3 = 000
> - a4 = 000
> - a5 = 000
>  
> 对于第一个数字（所谓的第一个数字，应该是指字符串 n 的第一个数字），我们把{{< katex >}}a_x {{< /katex >}} 前 n[0] 个数字初始化为 1；对于第二个数字，我们把前 n[1]个数字初始化为 1；同理对于第三个数字，我们把前 n[2] 个数字初始化为 1，最后有：
> - a1 = 111
> - a2 = 011
> - a3 = 011
> - a4 = 001
> - a5 = 001
>
> 可以发现，把这些数字加起来，可以得到 n，即 111 + 22 + 22 + 1 + 1 = 135.
> 
> 于是得证，只需要 x 个数字即可得到 n 这个数。
> 
> 时间复杂度 {{< katex >}} \Omircon(L) {{< /katex >}}， 空间复杂度 {{< katex >}} \Omircon(1) {{< /katex >}}。

代码如下：
```go
func minPartitions(n string) int {
    var res int
    size := len(n)
    for i := 0; i < size; i++ {
        res = max(res, int(n[i] - '0'))
    }
    
    return res
}

func max(a,b int) int {
    if a < b {
        return b
    }
    
    return a
}
```
