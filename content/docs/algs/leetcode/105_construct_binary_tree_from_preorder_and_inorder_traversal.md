---
title: 0105. Construct Bianry Tree From Preorder and Inorder Traversal
weight: 10
---

## Description
> Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree and inorder is the inorder traversal of the same tree, construct and return the binary tree.
## Solutions

```go
func buildTree(preorder []int, inorder []int) *TreeNode {
	pre_size, in_size := len(preorder), len(inorder)
    return buildTreeSolution(preorder, 0, pre_size-1, inorder, 0, in_size -1)
}

func buildTreeSolution(preorder []int, pleft, pright int, inorder []int, ileft, iright int) *TreeNode {
	if pleft > pright || ileft > iright {
		return nil
	}

	var rootIndex int
	for i := ileft; i <= iright; i++ {
		if inorder[i] == preorder[pleft] {
			rootIndex = i
			break
		}
	}

	root := new(TreeNode)
	root.Val = preorder[pleft]
	root.Left = buildTreeSolution(preorder, pleft+1, pleft + rootIndex - ileft, inorder, ileft, rootIndex-1)
	root.Right = buildTreeSolution(preorder, pleft + rootIndex-ileft+1, pright, inorder, rootIndex+1, iright)
	return root
}
```
