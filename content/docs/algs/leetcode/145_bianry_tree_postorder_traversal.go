package leetcode

/*
Given the root of a binary tree, return the postorder traversal of its nodes' values.
*/

func postorderTraversal(root *TreeNode) []int {
	return postorderTraversalRecursive(root)
}

// recursie
func postorderTraversalRecursive(root *TreeNode) []int {
	var res []int
	postorderTraversalRecursiveHelper(root, &res)
	return res
}

func postorderTraversalRecursiveHelper(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	postorderTraversalRecursiveHelper(root.Left, res)
	postorderTraversalRecursiveHelper(root.Right, res)
	*res = append(*res, root.Val)
}

// iterative
func postOrderTraversalIterative(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) != 0 {
		size := len(stack)
		cur := stack[size-1]
		stack = stack[:size-1]
		res = append(res, cur.Val)

		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
	}

	// reverse
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}
