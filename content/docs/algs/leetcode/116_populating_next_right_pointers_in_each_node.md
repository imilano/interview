---
title: 0116. Populating Next Right Pointers in Eech Node
weight: 10
---
## Description
> You are given a perfect binary tree where all leaves are on the same level, and every parent has two children. The binary tree has the following definition:
> 
> struct Node {
>   int val;
>   Node *left;
>   Node *right;
>   Node *next;
> }
> Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.
> 
> Initially, all next pointers are set to NULL.
## Solutions
简单题，只需要使用层次遍历即可。
```go
func connect(root *Node) *Node {
    if root == nil {
        return root
    }
    
    var queue []*Node
    queue = append(queue, root)
    for len(queue) != 0 {
        size := len(queue)
        var pre *Node
        for i := 0; i < size; i++ {
            if pre == nil {
                pre = queue[i]
            } else {
                pre.Next = queue[i]
                pre = queue[i]
            }
            
            if queue[i].Left != nil {
                queue = append(queue, queue[i].Left)
            }
            
            if queue[i].Right != nil {
                queue = append(queue, queue[i].Right)
            }
        }
        
        queue = queue[size:]
    }
    
    return root
}
```
