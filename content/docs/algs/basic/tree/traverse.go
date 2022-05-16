package tree

import (
	"fmt"

	"blog.lightsinger.top/interview/leetcode"
)

// 二叉树的遍历，包含迭代和递归的方法

//----------------- 递归 -------------------------
// 前序遍历
func preOrderTraverse(root *leetcode.TreeNode) {
	if root == nil {
		return
	}

	fmt.Println(root.Val)
	preOrderTraverse(root.Left)
	preOrderTraverse(root.Right)
}

// 中序遍历
func inOrderTraverse(root *leetcode.TreeNode) {
	if root == nil {
		return 
	}

	inOrderTraverse(root.Left)
	fmt.Println(root.Val)
	inOrderTraverse(root.Right)
}


// 后序遍历
func postOrderTraverse(root *leetcode.TreeNode) {
	if root == nil {
		return
	}

	postOrderTraverse(root.Left)
	postOrderTraverse(root.Right)
	fmt.Println(root.Val)
}


//-------------------------- 迭代法 --------------------
// 前序遍历
func preOrderIterative(root *leetcode.TreeNode) []int{
	var res []int
	if root == nil {
		return res
	}

	var stack []*leetcode.TreeNode
	stack = append(stack, root)
	for len(stack) != 0 {
		size := len(stack)
		cur := stack[size - 1]
		stack = stack[:size-1]

		res = append(res, cur.Val)
		// TODO 注意先右后左
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
	}

	return res
}


// 中序遍历
func inOrderIterative(root *leetcode.TreeNode) []int{
	var res []int
	if root == nil {
		return res
	}

	var stack []*leetcode.TreeNode
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


// 后序遍历
// 后序遍历的迭代法不如前序和中序那么好写，后序遍历可以按照根右左的顺序来进行遍历，遍历结束之后，将结果反转，即可得到左右根
func postOrderIterative(root *leetcode.TreeNode) []int{
	var res []int
	if res == nil {
		return res
	}

	var stack []*leetcode.TreeNode
	cur := root
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			res = append(res, cur.Val)
			stack = append(stack, cur)
			cur = cur.Right
		}

		size := len(stack)
		cur = stack[size-1]
		stack = stack[:size-1]

		cur = cur.Left
	}

	// reverse
	for i, j := 0, len(res) -1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}

// TODO recheck this
// 后序遍历也可以按照前序比遍历的思路进行稍微更改得到
// 先序遍历的顺序是根左右，后序遍历的顺序是左右根，只要稍微更改一下先序遍历的顺序，就能变成根右左，最后再对结果做一次反转，即可得到左右根
func postOrderIterativeFromPreOrder(root *leetcode.TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	var stack []*leetcode.TreeNode
	stack = append(stack, root)
	for len(stack) != 0 {
		size := len(stack)
		top := stack[size-1]
		stack = stack[:size-1]

		res  = append(res, top.Val)
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
	}

	// reverse 
	for i, j := 0, len(res)-1; i < j; i,j = i +1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}


// 层序遍历
func levelOrderTraverse(root *leetcode.TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	var queue []*leetcode.TreeNode
	queue = append(queue, root)
	for len(queue) != 0 {
		size := len(queue)
		var level []int
		for i := 0; i < size; i++ {
			node := queue[i]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		res = append(res, level)
	}

	return res
}
