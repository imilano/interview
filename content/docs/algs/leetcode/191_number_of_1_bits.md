---
title: 0191. Number of 1 Bits
weight: 10
tags: [
	"Bit Operations"
]
---

## Description 
> Write a function that takes an unsigned integer and returns the number of '1' bits it has (also known as the Hamming weight).


## Solutions
太简单了，并且输入还是无符号数，没啥好说的，直接看代码。
```go
func hammingWeight(num uint32) int {
    var res int
    for num != 0 {
        res += int(num&1)
        num >>= 1
    }
    
    return res
}
```
