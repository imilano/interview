---
title: 0263. Ugly Number
weight: 10
tags: [
	"Math",
]
---
## Description
> An ugly number is a positive integer whose prime factors are limited to 2, 3, and 5.
> 
> Given an integer n, return true if n is an ugly number.

## Solutions
如果一个数是丑数，那么它的因子必然只有 2、3、5 这三个，那么如果 n 能够被它的这些因子整除，那就不断的缩小 n，最后检查 n 不断被整除后的 n 其是否等于 1 即可。
```go
func isUgly(n int) bool {
    if n <= 0 {
        return false
    }
    if n <= 6 {
        return true
    }
    
    for n % 2 == 0 {
        n /= 2
    }
    
    for n %3 == 0 {
        n /= 3
    }
    
    for n % 5 == 0 {
        n /= 5
    }
    
    return n == 1
}
```


