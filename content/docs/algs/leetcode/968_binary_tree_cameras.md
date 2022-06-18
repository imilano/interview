---
title: 0968. Binary Tree Cameras
weight: 10
tags: [
	"Tree",
	"Binary Tree",
	"Greedy Algorithms"
]
---

## Description
> You are given the root of a binary tree. We install cameras on the tree nodes where each camera at a node can monitor its parent, itself, and its immediate children.
> 
> Return the minimum number of cameras needed to monitor all nodes of the tree.

## Solutions
### Greedy Algorithms
这种 hard 题，题主肯定是不会做的啦（囧�� ），所以只好求助于网上大神了：

> 这里先考虑把相机放在什么位置上能看到的节点最多（这样的话相机数量就会最少）？能放在叶节点吗？显然不能，叶节点最多只能看到两个节点；能放到根节点吗？根节点最多也只能看到 3 个。最优解是放在叶节点的父节点上，这样最多就可以看到四个节点。所以策略是先找到叶节点，然后在其父节点上放相机，同时标记父节点的父节点为被拍到了的状态。这样就有三种不同的状态，用 0 表示当前节点是叶节点，用 1 表示当前节点是叶节点的父节点并且放置了相机，用 2 表示当前节点是叶节点的爷爷节点，并且被相机拍到了。这里使用一个全局变量 res 记录相机个数。在递归过程中，若当前节点不存在，则返回 2，空节点也可以看做被相机拍到了。否则对相机左右节点递归调用，若二者中有一个返回 0，则当前节点至少有一个节点是叶节点，需要在当前位置放置一个相机，res 自增 1，并返回 1；如果左右节点的返回值中有一个为 1，说明左右节点中至少有一个已经放上了相机，当前节点已经被拍到了，返回 2。若都不是，则说明当前节点是叶节点，返回 0。在主函数中，若对根节点调用递归的返回值是 0，说明这个树只有一个节点或者根节点的左右节点没有子节点或者根节点就是叶节点，此时没有办法，只能在根节点上也放一个相机，否则不用加。

其实上面的的解法，就是遍历再加上一个状态表征，理解了之后，发现其实也没有那么难（不是）。
详细代码如下：
```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minCameraCover(root *TreeNode) int {
    var res int
    r := helper(root, &res)
    // 如果返回值是 0，则说明当前根节点需要放置一个相机；否则不需要
    if r < 1 {
        res += 1
    }
    
    return res
}


func helper(root *TreeNode, res *int) int {
    // 空节点可以当做祖父节点
    if root == nil {
        return 2
    }
    
    left, right := helper(root.Left, res), helper(root.Right, res)
    // 如果左右节点有一个是叶节点，那么当前节点就是叶节点的父节点，需要放置一个相机
    if left == 0 || right == 0 {
        *res++
        return 1
    }
    
    // 如果左右节点有一个节点是叶节点的父节点，那么当前节点就是叶节点的祖父节点，那么当前节点不需要放置相机;
    // 返回 2 表征自己当前是祖父节点
    if left == 1 || right == 1 {
        return 2
    }
    
    // 否则，说明当前节点是叶节点，则返回 0
    return 0
}
```
