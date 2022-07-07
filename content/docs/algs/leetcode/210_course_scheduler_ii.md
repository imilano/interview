---
title: 0210. Course Scheduler II
weight: 10
tags: [
	"BFS",
	"Topological sorting"
]
---

## Description
> There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.
> 
> For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
> Return the ordering of courses you should take to finish all courses. If there are many valid answers, return any of them. If it is impossible to finish all courses, return an empty array.

## Solutions
### Topological sorting
这里与 207 题的区别就在于，这里不仅需要判断有没有环，如果没环的话还需要给出一条路径。那么只需要在 207 题的基础上稍微做一些修改即可。这里因为只需要给出一个可行解即可，所以只需要在 BFS 的过程中每遍历到一个节点，就将这个节点的加入到 res 中即可。还需要注意的是，最后可以直接通过判断 res 的大小是否与 numCourses 相等来判断是否存在环，因为按照上面的过程，如果不存在环的话，所有的节点都会被加入到 res 中去。
```go
func findOrder(numCourses int, prerequisites [][]int) []int {
    // in 表示每个节点的入度，首先需要将所有节点的入度初始化为 0（重要）
    in := make(map[int]int)
    for i := 0; i < numCourses; i++ {
        in[i] = 0
    }
    
    // 构建邻接表
    adjTable := make(map[int][]int)
    size := len(prerequisites)
    for i := 0; i < size; i++ {
        adjTable[prerequisites[i][1]] = append(adjTable[prerequisites[i][1]], prerequisites[i][0])
        // 这里只会统计入度不为 0 的节点，如果上一步没有将所有节点的入度初始化为 0 的话，
        // 那么在第三部的 BFS 中就很可能导致这些节点的确实
        in[prerequisites[i][0]]++
    }
    
    // 找到入度为 0 的节点并入队，这些节点将作为其实节点
    var queue []int
    for key, value := range in {
        if value == 0 {
            queue = append(queue, key)
        }
    }
    
    // BFS
    var res []int
    for len(queue) != 0 {
        size := len(queue)
        for i := 0; i < size; i++ {
            // 对于每个从当前节点可达的节点，将其入度减 1。减 1 之后如果入度为 0，则将其加入队列中
            nodes := adjTable[queue[i]]
            res = append(res, queue[i])
            for _, value := range nodes {
                in[value]--
				// 这里的判断条件意味着，即使 in[value] 小于 0 也没问题
                if in[value] == 0 {
                    queue = append(queue, value)
                }
            }
        }
        
        queue = queue[size:]
    }
    
    // 统计是否有节点的入度不为 0，如果存在，则说明存在环，无法完全结束
    // for _, value := range in {
    //     if value > 0 {
    //         return false
    //     }
    // }
    // 这里只需要统计结果中的节点数等不等与 numCourse 的数量，如果不等，则说明有环，否则说明无环
    if len(res) != numCourses {
        return nil
    }
    
    // 无环，所有节点均可达
    return res
}
```
