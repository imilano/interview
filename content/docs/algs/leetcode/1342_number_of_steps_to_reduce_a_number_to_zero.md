---
title: 1342. Number of Steps to Reduce a Number to Zero
weight: 10
tags: [
	"Math",
	"Bit Manipulation"
]
---
## Description
> Given an integer num, return the number of steps to reduce it to zero.
> 
> In one step, if the current number is even, you have to divide it by 2, otherwise, you have to subtract 1 from it.

## Solutions
这题太简单了，没什么好说的，直接上代码。
```go
func numberOfSteps(num int) int {
    var res int
    for num != 0 {
        if num%2 == 0 {
            num /= 2
        } else {
            num -= 1
        }
        
        res++
    }
    
    return res
}
```
