---
title: 0102. Binary Tree Level Order Traversal
weight: 10
tags: [
	"Binary Tree",
	"BFS"
]
---

## Description
> Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).

## Solutions
层次遍历，使用队列来做即可。
```go
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

type Queue []interface{}


func (q *Queue) Push(x interface{}) {
	*q = append(*q, x)
}

func (q *Queue) Pop() interface{} {
	size := q.Len()
	if size == 0 {
		return nil
	}

	ele := (*q)[0]
	*q = (*q)[1:]
	return ele
}
func (q *Queue) Len() int {
	return len(*q)
}

```
也可以用一个数组来简单的模拟队列：
```go
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    
    var res [][]int
    var queue []*TreeNode
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
        
        queue = queue[size:]
        res = append(res, level)
    }
    
    return res
}
```
