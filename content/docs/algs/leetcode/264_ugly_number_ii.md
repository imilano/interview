---
title: 0264. Ugly Number II
weight: 10
tags: [
	"Queue",
]
---
## Description
> An ugly number is a positive integer whose prime factors are limited to 2, 3, and 5.
> 
> Given an integer n, return the nth ugly number.

## Solutions
下面这种是错误的解法，下面的逻辑意味着丑数只能从 2 的幂、3 的幂以及 5 的幂中出现，但是 6 并不属于上述任何一种，但是 6 也是幂。
```go
func nthUglyNumber(n int) int {
    if n <= 6 {
        return n
    }
    
    res := 1
    i2,i3,i5 := 1,1,1
    for i := 2; i < n; i++ {
        res = min(i2*2, min(i3*3, i5*5))
        // 错误做法。按照下面的逻辑，丑数只能从 2 的幂，3 的幂以及 5 的幂中出现，这个逻辑是不对的。
        if res /2 == i2 {
            i2 *= 2
        }
        
        if res / 3 == i3 {
            i3 *= 3
        }
        
        if res / 5 == i5 {
            i5 *= 5
        }
    }
    
    return res
}

func min(a,b int) int {
    if a < b {
        return a
    }
    
    return b
}
```

下面这个才是正确的解法，这个解法是从已有丑数中找出下一个最小的丑数。
```go
func nthUglyNumber(n int) int {
    if n <= 6 {
        return n
    }
    
    var i2, i3, i5 int
    res := make([]int, n)
    res[0] = 1
    for i := 1; i < n; i++ {
		// 取出当前队列的最小值
        res[i] = min(res[i2]*2, min(res[i3]*3, res[i5]*5))
		// 如果是这个队列贡献了这个值，那么更新队列元素
        if res[i] == res[i2]*2 {
            i2++
        }
        if res[i] == res[i3] * 3 {
            i3++
        }
        if res[i] == res[i5] * 5 {
            i5++
        }
    }
    
    return res[n-1]
}

func min(a,b int) int {
    if a < b {
        return a
    }
    
    return b
}
```
