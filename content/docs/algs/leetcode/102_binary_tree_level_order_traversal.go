package leetcode

/*
Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).
*/

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	queue := new(Queue)
	queue.Push(root)
	for queue.Len() != 0 {
		size := queue.Len()
		var arr []int
		for size != 0 {
			ele := queue.Pop().(*TreeNode)
			size -= 1
			if ele == nil {
				continue
			}

			arr = append(arr, ele.Val)
			if ele.Left != nil {
				queue.Push(ele.Left)
			}
			if ele.Right != nil {
				queue.Push(ele.Right)
			}

		}
		res = append(res, arr)
	}

	return res
}
