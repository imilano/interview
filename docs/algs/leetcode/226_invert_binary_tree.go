package leetcode

/*
Given the root of a binary tree, invert the tree, and return its root.
*/

// 递归法
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	left, right := invertTree(root.Left), invertTree(root.Right)
	root.Left, root.Right = right, left
	return root
}

// 迭代法
// 采用中序遍历的方式
func invertTreeIterative(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]

		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return root
}
