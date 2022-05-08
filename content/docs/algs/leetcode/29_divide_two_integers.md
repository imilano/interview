---
title: 29. Divide Two Integers
weight: 10
---

## Description

> Given two integers dividend and divisor, divide two integers without using multiplication, division, and mod operator.
>
> The integer division should truncate toward zero, which means losing its fractional part. For example, 8.345 would be truncated to 8, and -2.7335 would be truncated to -2.
> 
> Return the quotient after dividing dividend by divisor.
> 
> Note: Assume we are dealing with an environment that could only store integers within the 32-bit signed integer range: [−231, 231 − 1]. For this problem, if the quotient is strictly greater than 231 - 1, then return 231 - 1, and if the quotient is strictly less than -231, then return -231.

## Solutions

### Subtraction
俗话说得好，"计算中的除法都是通过减法来完成的"。 :)

很不幸，下面这个解法超时了. :)
```go
func divide(dividend int, divisor int) int {
    var quotient int
    if dividend == 0 {
        return quotient
    }
    negative := true
    if dividend > 0 && divisor > 0 || dividend < 0 && divisor < 0 {
        negative = false
    }
    

    
    dividend, divisor = abs(dividend), abs(divisor)
    for dividend > 0 {
        dividend -= divisor
        if dividend >= 0 {
            quotient++
        }
        
        if negative && dividend > 0 && quotient >= int(math.MaxInt32) + 1 {
            return int(math.MinInt32)
        } 
        
        if !negative && dividend > 0 && quotient >= int(math.MaxInt32) {
            return int(math.MaxInt32)
        }
    }
    
    if negative {
        return -quotient
    }
    
    return quotient
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    
    return a
}
```
