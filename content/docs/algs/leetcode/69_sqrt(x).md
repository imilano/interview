---
title: 0069. Sqrt(x)
weight: 10
---

## Description

> Given a non-negative integer x, compute and return the square root of x.
> 
> Since the return type is an integer, the decimal digits are truncated, and only the integer part of the result is returned.
> 
> Note: You are not allowed to use any built-in exponent function or operator, such as pow(x, 0.5) or x ** 0.5.


## Solutions

使用二分法快速进行快速查找：
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
    
    
    return high - 1
}
```
