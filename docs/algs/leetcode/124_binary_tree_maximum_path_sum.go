package leetcode

import "math"

/*
A path in a binary tree is a sequence of nodes where each pair of adjacent nodes in the sequence has an edge connecting them. A node can only appear in the sequence at most once. Note that the path does not need to pass through the root.

The path sum of a path is the sum of the node's values in the path.

Given the root of a binary tree, return the maximum path sum of any non-empty path.
*/

func maxPathSum(root *TreeNode) int {
	res := math.MinInt
	maxPathSumSolution(root, &res)

	return res
}

func maxPathSumSolution(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}

	leftMax, rightMax := max(maxPathSumSolution(root.Left, res), 0), max(maxPathSumSolution(root.Right, res), 0)
	*res = max(*res, leftMax+rightMax+root.Val)
	return max(leftMax, rightMax) + root.Val
}
