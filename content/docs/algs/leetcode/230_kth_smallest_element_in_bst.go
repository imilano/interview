package leetcode

/*
Given the root of a binary search tree, and an integer k, return the kth smallest value (1-indexed) of all the values of the nodes in the tree.
*/

func kthSmallest(root *TreeNode, k int) int {
	// return kthSmallestSolution1(root, k)
	return kthSmallestSolution3(root, k)
}

// 方法 1，DFS 遍历一次得到元素数组，然后从中取第 k 个数
func kthSmallestSolution1(root *TreeNode, k int) int {
	var res []int
	kthSmallestDFS(root, &res)
	return res[k-1]
}

func kthSmallestDFS(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	kthSmallestDFS(root.Left, res)
	*res = append(*res, root.Val)
	kthSmallestDFS(root.Right, res)
}

// 方法 2， 可以使用最小堆
// 方法 3, 中序遍历非递归
func kthSmallestSolution3(root *TreeNode, k int) int {
	var cnt int
	stack := new(Stack)
	p := root
	for p != nil || stack.Len() != 0 { // 注意这里是或，而不是与
		for p != nil {
			stack.Push(p)
			p = p.Left
		}

		p = stack.Pop().(*TreeNode)
		cnt++
		if cnt == k {
			return p.Val
		}
		p = p.Right
	}

	return 0
}
