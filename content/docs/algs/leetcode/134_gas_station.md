---
title: 134. Gas Station
weight: 10
---
## Description

> There are n gas stations along a circular route, where the amount of gas at the ith station is gas[i].
> 
> You have a car with an unlimited gas tank and it costs cost[i] of gas to travel from the ith station to its next (i + 1)th station. You begin the journey with an empty tank at one of the gas stations.
> 
> Given two integer arrays gas and cost, return the starting gas station's index if you can travel around the circuit once in the clockwise direction, otherwise return -1. If there exists a solution, it is guaranteed to be unique

## Solutions

首先很容易想到简单模拟的暴力算法。将每一个点都作为起始点，使用 curGas 表示当前剩余的油量，然后从起始点开始遍历，每遍历到一个节点就加上该点可以获得的油量，然后减去到达该点的损耗，如果有一个点剩余油量小于 0，则说明无法到达该点，结束当前路径的遍历。然后选择下一条路径继续遍历即可。

很不幸，超时了。
```go
func canCompleteCircuit(gas []int, cost []int) int {
    res := -1
    size := len(gas)
    if size == 0 {
        return res
    }
    
    var curGas int 
    for i := 0; i < size; i++ {
        curGas = 0
        for j := i; j < i+size; j++ {
            if curGas < 0 {
                break
            }
            
            curGas += gas[j%size]
            curGas -= cost[j%size]
        }
        
        if curGas >= 0 {
            res = i
            break
        }
    }
    
    return res
}
```


这里还可以优化一下。首先很容易想到，如果整个 gas 数组的和小于 cost 数组的和，那么肯定是没有路径的。然后还有一个比较难想到的点是，如果以 A 为起点的路径无法到达 B， 那么从 A 到 B 中间的任何一个点作为起点都无法到达 B，也就是说，A 到 B 中间的任何一个点都无法作为起点。

这里可以简单证明一下，假设 A 到 B 中间存在一点 C，A 可以到达 C 但是无法到达 B。因为 A 无法到达 B，那么 A 到 B 这条路径的 gas 之和小于 cost 之和，而因为 A 可以到达 C，那么必然有 A 到 C 这条路的 gas 之和大于等于 cost 之和，因此如果以 C 作为起点，那么 C 必然无法到达 B，以 C 为起点的路径也无法环绕一圈。进而也就是说，A 到 B 之间任意一点都无法作为起点绕一圈，同理可证明 A 到 C 之前的节点。那么下一次遍历只需要从 B 的下一个节点开始即可。这样就可以降低一些总时间复杂度。

```go
func canCompleteCircuit(gas []int, cost []int) int {
    res := -1
    size := len(gas)
    if size == 0 {
        return res
    }
    
    var start int
    for start < size {
        var curGas, i int
        for  i = start; i < start + size; i++ {
            if curGas < 0 {
                break
            }
            
            curGas += gas[i%size]
            curGas -= cost[i%size]
        }
        
        if curGas >= 0 {
            res = start
            break
        }
        
        start = i
    }
    
    return res
}
```


