package leetcode

/*
Given the root of a binary tree, return the preorder traversal of its nodes' values.
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	// return recur(root)
	return iterative(root)
}

// iterative
func iterative(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) != 0 {
		size := len(stack)
		top := stack[size-1]
		stack = stack[:size-1]

		res = append(res, top.Val)
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
	}

	return res
}

func recur(root *TreeNode) []int {
	var res []int
	recursive(root, &res)
	return res
}

func recursive(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	*res = append(*res, root.Val)
	recursive(root.Left, res)
	recursive(root.Right, res)
}
