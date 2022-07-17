---
title: 0559. Maximum Depth of N-ary Tree
weight: 10
tags: [
	"Tree",
	"Iterative",
	"Level Order Traversal"
]
---

## Description
> Given a n-ary tree, find its maximum depth.
> 
> The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
> 
> Nary-Tree input serialization is represented in their level order traversal, each group of children is separated by the null value (See examples).

## Solutions
### Iterative
简单的中序遍历即可。
```go
type Node struct {
    Val int
    Children []*Node
}


func maxDepth(root *Node) int {
    if root == nil {
        return 0
    }
    
    var queue []*Node = []*Node{root}
    var depth int
    for len(queue) != 0 {
        depth++
        size := len(queue)
        for i := 0; i < size; i++ {
            for _, node := range queue[i].Children {
                if node != nil {
                    queue = append(queue, node)
                } 
            }
        }
        
        queue = queue[size:]
    }
    
    return depth
}
```
