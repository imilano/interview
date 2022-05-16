package leetcode

/*
Given a binary search tree (BST), find the lowest common ancestor (LCA) of two given nodes in the BST.

According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”
*/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return lowestCommonAncestorSolution2(root, p, q)
}

// 递归解法
// 要注意这是 BST，这意味着存在着大小关系。
// 如果根节点的值大于两个节点的最大值，则说明最小公共祖先在左子树，继续在左子树递归
// 如果根节点的值小于两个节点的最小值，说明最小公共祖先在右子树，继续在右子树递归
func lowestCommonAncestorSolution(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > max(p.Val, q.Val) {
		return lowestCommonAncestorSolution(root.Left, p, q)
	} else if root.Val < min(p.Val, q.Val) {
		return lowestCommonAncestorSolution(root.Right, p, q)
	} else {
		return root
	}
}

// 迭代解法
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
