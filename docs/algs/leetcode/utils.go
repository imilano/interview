package leetcode

import "strconv"

//*********** compare **********
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

//************* customized structure *************
type ListNode struct {
	Val  int
	Next *ListNode
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//************** heap **************************
type MinHeapArr []int

func (h MinHeapArr) Len() int {
	return len(h)
}

func (h MinHeapArr) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeapArr) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeapArr) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeapArr) Pop() interface{} {
	size := len(*h)
	res := (*h)[size-1]
	*h = (*h)[:size-1]
	return res
}

func (h *MinHeapArr) Top() interface{} {
	// TODO
	return nil
}

//***************** stack *********************
type Stack []interface{}

func (s *Stack) Push(x interface{}) {
	*s = append(*s, x)
}

func (s *Stack) Pop() interface{} {
	if s.Len() == 0 {
		return nil
	}

	ele := (*s)[s.Len()-1]
	*s = (*s)[:s.Len()-1]
	return ele
}

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) Top() interface{} {
	if s.Len() != 0 {
		return (*s)[s.Len()-1]
	}

	return nil
}

//****************** Queue ********************
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

//****************** util ***************
func reverseStr(s string) string {
	size := len(s)
	rs := []rune(s)
	left, right := 0, size-1
	for left < right {
		rs[left], rs[right] = rs[right], rs[left]
	}

	return string(rs)
}

func arrToString(arr []int) string {
	var s string
	for _, v := range arr {
		s += strconv.Itoa(v)
	}

	return s
}
