---
title: 0070. Climbing Stairs
weight: 10
---

## Descriptioin
> You are climbing a staircase. It takes n steps to reach the top.
> 
> Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

## Solutions

斐波那契数列问题，不多说。
```go
func climbStairs(n int) int {
    if n <= 2 {
        return n
    }
    FMinusOne, FMinusTwo := 2, 1
    for i := 3; i <= n; i++ {
        fn := FMinusOne + FMinusTwo
        FMinusTwo = FMinusOne
        FMinusOne = fn
        
    }
    
    return FMinusOne
}

```
