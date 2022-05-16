package leetcode

import (
	"math"
	"sort"
)

/*
Given the root of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:

The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.
*/

func isValidBST(root *TreeNode) bool {
	var res []int
	isValidBSTDFSSolution(root, &res)

	return sort.SliceIsSorted(res, func(i, j int) bool {
		// TODO figure it out
		// 奇怪，为什么加上一个等号反而可以排除序列全相等的情况
		return res[i] <= res[j]
	})
}

// 方法 1，深度优先遍历，遍历结束就可以得到一个递归的数组，如果这个数组不是递增的，那么就说明不是一棵有效的 BST
func isValidBSTDFSSolution(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	isValidBSTDFSSolution(root.Left, res)
	if root != nil {
		*res = append(*res, root.Val)
	}
	isValidBSTDFSSolution(root.Right, res)
}

// 同样是中序遍历，但是使用一个 pre 指针指向上一个遍历到的节点值，无需再对数组进行遍历比较
func isValidBSTIterSolution(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var stack []*TreeNode
	cur := root
	var pre *TreeNode
	for len(stack) != 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		size := len(stack)
		cur = stack[size-1]
		stack = stack[:size-1]

		if pre != nil && pre.Val >= cur.Val {
			return false
		}
		pre = cur
		cur = cur.Right
	}

	return true
}

// 方法 3, 递归法，但是需要设置一个最大值和最小值
func isValidBSTRecurSolution(root *TreeNode) bool {
	mn, mx := math.MinInt, math.MaxInt
	return isValidBSTRecur(root, mn, mx)
}

func isValidBSTRecur(root *TreeNode, mn, mx int) bool {
	if root == nil {
		return true
	}

	if root.Val <= mn || root.Val >= mx {
		return false
	}

	return isValidBSTRecur(root.Left, mn, root.Val) && isValidBSTRecur(root.Right, root.Val, mx)
}
