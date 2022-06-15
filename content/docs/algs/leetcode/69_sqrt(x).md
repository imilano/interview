---
title: 0069. Sqrt(x)
weight: 10
tags: [
    "Binary Search",
    "Math"
]
---

## Description

> Given a non-negative integer x, compute and return the square root of x.
> 
> Since the return type is an integer, the decimal digits are truncated, and only the integer part of the result is returned.
> 
> Note: You are not allowed to use any built-in exponent function or operator, such as pow(x, 0.5) or x ** 0.5.


## Solutions

### Binary Search
使用二分法进行快速查找。开头就去除掉一些 corner case 的话，后面整体逻辑写起来就会更加的清晰。
```go
func mySqrt(x int) int {
    if x <= 1 {
      return x
    }
    
    low, high := 0, x
    for low < high {
        mid := (low+high)/2
        if mid * mid == x {
            return mid
        } else if mid * mid > x {
            high = mid
        } else {
            low = mid + 1
        }
    }
    
    // 注意这里为什么返回的是 high - 1
    return high - 1
}
```
