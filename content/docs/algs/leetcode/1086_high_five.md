---
title: 1086. High Five
weight: 10
tags: [
	"Hash Table",
	"Sort"
]
---
## Description
> Given a list of the scores of different students, items, where items[i] = [IDi, scorei] represents one score from a student with IDi, calculate each student's top five average.
> 
> Return the answer as an array of pairs result, where result[j] = [IDj, topFiveAveragej] represents the student with IDj and their top five average. Sort result by IDj in increasing order.
> 
> A student's top five average is calculated by taking the sum of their top five scores and dividing it by 5 using integer division.

## Solutions
这题应该算是简单题。首先需要使用一个 map 记录每个 ID 对应的所有分数，然后对这写分数按照从大到小排序，然后取出前 5 个来计算其均值即可。
```go
func highFive(items [][]int) [][]int {
	dict := make(map[int][]int)
	for _, item := range items {
		dict[item[0]] = append(dict[item[0]],  item[1])
	}

	var res [][]int
	for key, nums := range dict {
		sort.Slice(nums, func(i,j int) bool {
			return nums[i] > nums[j]
		})

		var avg int
		for i := 0; i < 5; i++ {
			avg += nums[i]
		}

		res = append(res, []int{key, avg/5})
	}

	return res
}

```
