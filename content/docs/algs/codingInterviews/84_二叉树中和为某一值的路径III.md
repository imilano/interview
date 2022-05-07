---
title: 84. 二叉树中和为某一值的路径 III
weight: 10
---
## Description
> 给定一个二叉树root和一个整数值 sum ，求该树有多少路径的的节点值之和等于 sum 。
> 1.该题路径定义不需要从根节点开始，也不需要在叶子节点结束，但是一定是从父亲节点往下到孩子节点
> 2.总节点数目为n
> 3.保证最后返回的路径个数在整形范围内(即路径个数小于$2^{31}$-1)

## Solutions

注意，这里的路径不一定需要从根节点开始，也不一定需要在叶节点结束。
### Recursice

核心思想就是，每个节点都可能是一条路径的起始节点，所以每个节点都需要作为根节点来遍历一次。
```go
var res int
func FindPath( root *TreeNode ,  sum int ) int {
    // write code here
    if root == nil {
        return res
    }
    
    findPathHelper(root, sum)

	// 每个节点都需要作为根节点来进行一次查找
    FindPath(root.Left, sum)
    FindPath(root.Right, sum)
    return res
}

func findPathHelper(root *TreeNode, sum int) {
    if root == nil {
        return
    }
    // 注意这里的小技巧
    if sum == root.Val {
        res++
    }
    
    findPathHelper(root.Left, sum - root.Val)
    findPathHelper(root.Right, sum - root.Val)
}

```
