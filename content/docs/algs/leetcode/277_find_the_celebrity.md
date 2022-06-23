---
title: 0277. Find the Celebrity
weight: 10
tags: [
	"Hash Table",
	"Matrix"
]
---
## Description
> Suppose you are at a party with n people (labeled from 0 to n - 1) and among them, there may exist one celebrity. The definition of a celebrity is that all the other n - 1 people know him/her but he/she does not know any of them.
> 
> Now you want to find out who the celebrity is or verify that there is not one. The only thing you are allowed to do is to ask questions like: "Hi, A. Do you know B?" to get information of whether A knows B. You need to find out the celebrity (or verify there is not one) by asking as few questions as possible (in the asymptotic sense).
> 
> You are given a helper function bool knows(a, b)which tells you whether A knows B. Implement a function int findCelebrity(n). There will be exactly one celebrity if he/she is in the party. Return the celebrity's label if there is a celebrity in the party. If there is no celebrity, return -1.

## Solutions
### Hash Table
题主初始是这么想的：创建一个长度为 n 的 map，表示每个人都是名人，然后开始遍历矩阵，每遍历到一个位置，就 matrix[i][j]，如果 i 认识 j，那么说明 i 肯定不是名人，把 i 从 map 中删去。最后 map 中如果还剩下一个元素，那么最后剩下的那个元素就是名人，否则说明没有名人。
```go
func findCelebrity( n int) int {
	// 假定每个人都是名人
	dict := make(map[int]int)
	for i := 0; i < n; i++ {
		dict[i] = true
	}

	for i := 0; i < n ;i++ {
		for j := 0; j < n; j++  {
			if knows(i,j) {
				// 如果 i 认识 j， 那么 i 不是名人；如果 i 还在 map 中，那么就把 i 从 map 中删去
				if _, ok := dict[i]; ok {
					delete(dict, i)
				}
			}
		}
	}

	// 名人只会有一个
	if len(dict) != 1 {
		return -1
	}

	for key, _ := range dict {
		return key
	}
}

func knows(i,j int) bool {
	// do something
	return true
}
```

题主网上看到一种比较妙的解法，核心大意是：初始假定候选人 res 为 0 号选手，然后遍历一次所有选手，如果 res 认识 j，那么说明 res 肯定不是名人，而 j 可能是名人，则将 res 更新为 j，等遍历完之后，再来判断 res 到底是不是名人。这时有需要从头开始遍历一次，如果 res 认识 i 或者 i 不认识 res，那么都说明 res 不是名人，**此时说明没有名人**，返回 -1 即可；否则如果全部遍历完之后，res 有效的话，说明 res 就是名人，则返回 res 即可。
> 解释一下上面加粗的部分，为什么认为这个时候就没有名人了呢？
> 其实可以从树的角度来进行思考（做个类比），上面的第一次遍历中，其实已经确保了一个事实：如果这是一个有向连通图，那么 res 表示了图中的一个节点，此时 res 必然不是名人；反之，说明 res 此时表示了那个名人节点（任意节点均有路径与其相连，但是它找不到一条路径与其它节点相连（除自身））。那么第二次循环所做的努力就是，排除上述的有向连通图的可能。
```go
func findCelebrity(n int) int {
	// 找出那个可能的名人节点
	var res int
	for i := 1; i < n; i++ {
		if knows(res, i) {
			res = i
		}
	}

	// 排除这是一个连通图的可能
	for i := 0; i < n; i++ {
		// 除自身外，如果有人不认识 res 或者 res 认识它，都说明这是一个连通图，则返回-1
		if i != res && (knows(res, i) || !knows(i, res)) {
			return -1
		}
	}

	// res 是一个有效节点，返回即可
	return res
}

```
