package leetcode

/*
Given the root of a binary tree, return the inorder traversal of its nodes' values.
*/

func inorderTraversal(root *TreeNode) []int {
	return inorderTraversalIterative(root)
}

// iterative
func inorderTraversalIterative(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	var stack []*TreeNode
	cur := root
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		size := len(stack)
		cur = stack[size-1]
		stack = stack[:size-1]

		res = append(res, cur.Val)
		cur = cur.Right
	}

	return res
}

// recursive
func inorderTraversalRecursive(root *TreeNode) []int {
	var res []int
	inorderTraversalRecursiveHelper(root, &res)
	return res
}

func inorderTraversalRecursiveHelper(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	inorderTraversalRecursiveHelper(root.Left, res)
	*res = append(*res, root.Val)
	inorderTraversalRecursiveHelper(root.Right, res)
}
