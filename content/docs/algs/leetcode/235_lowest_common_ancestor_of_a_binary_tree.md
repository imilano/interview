---
title: 0235. Lowest Common Ancestor of Binary Search Tree
weight: 10
tags: [
	"Binary Tree",
	"Binary Search Tree",
	"Recursive",
	"Iterative"
]
---
## Description
> Given a binary search tree (BST), find the lowest common ancestor (LCA) node of two given nodes in the BST.
> 
> According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”
## Solutions
### Recursive 
这题很简单，因为 BST 本身就具备有序的性质，所以我们只需要抓住这个性质来进行遍历即可。首先我们假设 p 比 q 要小，那么对于递归遍历到的每个节点 root， 如果 root 的值大于最大 q 的值，那么说明最小公共节点一并比 root 要小，我们需要往root 的左子树进行查找；如果 root 的值小于最小值 p 的值，那么说明最小公共祖先一定比 p 要大，那么我们需要往 p 的右节点进行查找；否则，说明 root 的值刚好结余 p 和 q 之间，此时 root 就是我们所查找的最小公共组节点，返回 root 即可。
```go
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}


func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    if p.Val > q.Val {
        return lowestCommonAncestor(root, q, p)
    }
    
    if root.Val > q.Val {
        return lowestCommonAncestor(root.Left, p, q)
    } else if root.Val < p.Val {
        return lowestCommonAncestor(root.Right, p, q)
    } else {
        return root
    }
}
```

### Recursive
迭代解法也比较简单，思路跟上面一样，我这里就不多说了，直接上代码。
```go
func lowestCommonAncestorSolution2(root, p, q *TreeNode) *TreeNode {
	node := root
	for {
		if node.Val < min(q.Val, p.Val) {
			node = node.Right
		} else if node.Val > max(p.Val, q.Val) {
			node = node.Left
		} else {
			return node
		}
	}
}
```
