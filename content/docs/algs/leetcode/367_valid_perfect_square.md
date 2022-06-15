---
title: 0362. Design Hit Counter
weight: 10
tags: [
	"Binary Search",
	"Math"
]
---
## Description
> Given a positive integer num, write a function which returns True if num is a perfect square else False.
> 
> Follow up: Do not use any built-in library function such as `sqrt`.

## Solutions
### Binary Search
这里跟 69 题其实是完全一样的解法，需要注意的是，在开始的时候去除掉一些 corner case 能够让后面的整体逻辑更清晰一些。
```go
func isPerfectSquare(num int) bool {
    if num <= 1 {
        return true
    }
    left, right := 0, num
    for left < right {
        mid := left + (right - left)/2
        res := mid * mid
        if res == num {
            return true
        }
        if res > num {
            right = mid
        } else {
            left = mid + 1
        }
    }
    
    return false
}
```
