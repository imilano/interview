---
title: 118. Pascal's Triangle
weight: 10
---
## Description
> Given an integer numRows, return the first numRows of Pascal's triangle.
> 
> In Pascal's triangle, each number is the sum of the two numbers directly above it

## Solutions

这里有很多 trick，还是需要注意一下。
```go
func generate(numRows int) [][]int {
    var res [][]int
    for i := 0; i < numRows; i++ {
        var cur []int
        for j := 0; j < i+1; j++ {
            cur = append(cur, 1)
        }
        res = append(res, cur)
        
        for j := 1; j < i; j++ {
            res[i][j] = res[i-1][j-1] + res[i-1][j]
        }
    }
    
    return res
}
```
