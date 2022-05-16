package leetcode

import (
	"encoding/json"
	"math"
)

/*
Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later in the same or another computer environment.

Design an algorithm to serialize and deserialize a binary tree. There is no restriction on how your serialization/deserialization algorithm should work. You just need to ensure that a binary tree can be serialized to a string and this string can be deserialized to the original tree structure.

Clarification: The input/output format is the same as how LeetCode serializes a binary tree. You do not necessarily need to follow this format, so please be creative and come up with different approaches yourself.
*/

// 第一种方法，分别求出该树的中序遍历序列和先序序列，这样就转变成 leetcode 的根据这两个序列复原二叉树的问题。
// 第二种方法，采用层次遍历的方法来序列化，然后根据层次遍历的思想进行反序列化
// 第三种方法，递归方法，使用先序序列进行序列化，反序列化时候也按照先序序列进行反序列化。

type Codec struct {
}

// func Constructor() Codec {
// 	return Codec{}
// }

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var queue []*TreeNode
	var nodes []int
	var res string
	if root == nil {
		return res
	}
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if node != nil {
			nodes = append(nodes, node.Val)
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		} else {
			nodes = append(nodes, math.MinInt)
		}
	}

	b, _ := json.Marshal(nodes)
	return string(b)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	var nums []int
	json.Unmarshal([]byte(data), &nums)
	root := new(TreeNode)
	root.Val = nums[0]
	nums = nums[1:]

	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]

		// 因为我们保存了空节点值，所以如果从 nums 中取到的元素是整数最小值的话，那么我们直接将其忽略即可。
		// 因为序列化的时候保留了空节点，而反序列化的时候也保留了空节点，他们的顺序是一一对应的，故而不会出现问题。
		if len(nums) == 0 {
			break
		}
		value := nums[0]
		nums = nums[1:]
		if value != math.MinInt {
			t := new(TreeNode)
			t.Val = value
			queue = append(queue, t)
			node.Left = t
		}

		if len(nums) == 0 {
			break
		}
		value = nums[0]
		nums = nums[1:]
		if value != math.MinInt {
			t := new(TreeNode)
			t.Val = value
			queue = append(queue, t)
			node.Right = t
		}
	}

	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
