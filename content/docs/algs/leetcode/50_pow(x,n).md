---
title: 0050. Pow(x, n)
weight: 10
tags: [
    "Hash Table"
]
---

## Description

> Implement pow(x, n), which calculates x raised to the power n (i.e., xn).

## Solutions

### Map

可以使用快速幂以及记忆化数组，单独使用快速幂会导致超时，所以需要配合记忆化数组使用。
```go
func myPow(x float64, n int) float64 {
    if x == 0 {
        return 0
    }
    if n == 0 {
        return 1
    }
    
    var negative bool
    if n < 0 {
        negative = true
        n = -n
    }
    
    
    dict := make(map[int]float64)
    dict[0], dict[1] = 1, x
    res := helper(x, n, &dict)
    if negative {
        return 1 / res
    }
    
    return res
}

func helper(x float64, n int, dict *map[int]float64) float64 {
    if value, ok := (*dict)[n]; ok {
        return value
    }
    
    res := helper(x, n/2, dict) * helper(x, n/2, dict) * helper(x, n%2, dict)
    (*dict)[n] = res
    return res
}
```
