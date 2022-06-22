---
title: 0016. 3Sum Closest
weight: 10
tags: [
	"Two Pointer",
	"Sort"
]
---

## Description
> Given an integer array nums of length n and an integer target, find three integers in nums such that the sum is closest to target.
> 
> Return the sum of the three integers.
> 
> You may assume that each input would have exactly one solution.

## Solutions
### Two Pointer
这个解法是题主一开始想到的解法，时间复杂度比较高，提交的结果也不如人意，不过还是可以 AC 的。
```go
func threeSumClosest(nums []int, target int) int {
    var res int
    sort.Ints(nums)
    size, diff := len(nums), math.MaxInt
    for idx, _ := range nums {
        left, right := idx + 1, size - 1 
        for left < right {
            sum :=  nums[left] + nums[right] + nums[idx]
            if sum == target {
                return target
            }
            
			// 注意这这里要使用绝对差
            if abs(target - sum) < diff {
                diff = abs(target - sum)
                res = sum
            }
            
            if sum > target {
                right--
            } else {
                left++
            }
        }
    }
    
    return res
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    
    return a
}
```

下面是上面的方法优化之后的版本：
```go
func abs(a int) int {
	if a >= 0 {
		return a
	}

	return -a
}


// optimized
func threeSumClosest(nums []int, target int) int {
	res := math.MaxInt16
	if len(nums) <= 2 {
		return 0
	}

	sort.Ints(nums)
	for i := 0 ; i< len(nums)-2;i++ {
		left,right := i+1,len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				return target
			}

			if sum > target {
				right--
			} else {
				left++
			}

			// 这里的会比较巧妙
			if abs(target-sum) < abs(target-res) {
				res = sum
			}
		}
	}
	
	return res
}
```
