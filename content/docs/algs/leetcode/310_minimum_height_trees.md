---
title: 0310. Minimum Height Tree
weight: 10
tags: [
	"Graph Theory",
	"BFS",
	"DFS"
]
---

## Description
> A tree is an undirected graph in which any two vertices are connected by exactly one path. In other words, any connected graph without simple cycles is a tree.
> 
> Given a tree of n nodes labelled from 0 to n - 1, and an array of n - 1 edges where edges[i] = [ai, bi] indicates that there is an undirected edge between the two nodes ai and bi in the tree, you can choose any node of the tree as the root. When you select a node x as the root, the result tree has height h. Among all possible rooted trees, those with minimum height (i.e. min(h))  are called minimum height trees (MHTs).
> 
> Return a list of all MHTs' root labels. You can return the answer in any order.
> 
> The height of a rooted tree is the number of edges on the longest downward path between the root and a leaf.

## Solutions
### BFS
题主最先想出的解法就是把每个节点都视作根节点，然后从根节点开始遍历，查找其最长路径长度。提交之后超时了：
```go
func findMinHeightTrees(n int, edges [][]int) []int {
    res := make(map[int][]int)
    graph := make(map[int][]int)
    curMin := math.MaxInt
    for _, edge := range edges {
        graph[edge[0]] = append(graph[edge[0]], edge[1])
        graph[edge[1]] = append(graph[edge[1]], edge[0])
    }
    for i := 0; i < n; i++ {
        visited := make([]bool, n)
        var queue []int = []int{i}
        var cnt int
        for len(queue) != 0 {
            cnt++
            if cnt > curMin {
                break
            }
            size := len(queue)
            for j := 0; j < size; j++ {
                nodes := graph[queue[j]]
                for _, node := range nodes {
                    if visited[node] {
                        continue
                    }
                    
                    queue = append(queue, node)
                }
                
                visited[queue[j]] = true
            }
            
            queue = queue[size:]
        }
        
        curMin = min(curMin, cnt)
        res[cnt] = append(res[cnt], i)
    }
    
    mht := math.MaxInt
    for cnt, _ := range res {
        if cnt < mht {
            mht = cnt
        }
    }
    
    return res[mht]
}

func min(a,b int) int {
    if a < b {
        return a
    }
    
    return b
}
```
其实不难发现这个题超时的可能性太大了，因为上面的解法把每个节点都作为根节点，然后对以该节点作为根节点的子树进行遍历。

题主想了好久都没想到解法，发呆发了好久。最后突然觉得这题跟之前做过的一个题目好像（但题主没找到这个题目）。大致解法如下:

既然是求高度最小的数，那么也就意味着要尽可能缩小树的直径。怎么缩小树的直径呢，当然是每一轮都尽可能减少节点的数量。那么我们可以先统计每个节点的入度，然后统计每个节点可达的节点，因为这是无向图，所以入度要双向统计。然后将入度为 1 的节点入队，开始 BFS。每轮 BFS 都将当前节点可达的节点中将入度减 1，如果该可达节点的入度减为 1 的话，则将该节点加入队列中。那么什么时候停止呢，~~直到队列中节点数小于等于 2 的时候停止~~直到 n 的值小于等于 2 停止，此时的 BFS 轮数即为最小高度，而剩下的节点即为根节点。

```go
func findMinHeightTrees(n int, edges [][]int) []int {
    // n 为 1 时，只有一个节点可以作为根节点
    if n == 1 {
        return []int{0}
    }
    graph := make(map[int]map[int]bool, n)
    for i := 0; i < n; i++ {
        graph[i] = make(map[int]bool)
    }
    
    // 无向图，双向统计
    for _, edge := range edges {
        // in[edge[1]]++
        // 记录从 edge [0] 可以到达哪些节点
        graph[edge[0]][edge[1]] = true
        graph[edge[1]][edge[0]] = true
    }
    
    // 对于叶节点，则入队
    var queue []int
    for i := 0; i < n; i++ {
        if len(graph[i]) == 1 {
            queue = append(queue, i)
        }
    }
    
    // 注意这里是 n > 2， 而不是 len(queue) > 2。
    for n > 2 {
        size := len(queue)
        n -= size
        for i := 0; i < size; i++ {
            // 找到从 queue[i] 可达的所有节点
            nodes := graph[queue[i]]
            for node, _ := range nodes {
                // 从 queue[i] 可达的所有节点中，删除其与 queue[i] 的关系。删除之后如果变成了叶节点，则将该节点入队
                delete(graph[node], queue[i])
                if len(graph[node]) == 1 {
                    queue = append(queue, node)
                }
            }
        }
        
        queue = queue[size:]
    }
    
    return  queue
}
```
