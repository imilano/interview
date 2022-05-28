---
title: 0735. Asteroid Collisiion
weight: 10
tags: [
	"Stack"
]
---
## Description
> We are given an array asteroids of integers representing asteroids in a row.
> 
> For each asteroid, the absolute value represents its size, and the sign represents its direction (positive meaning right, negative meaning left). Each asteroid moves at the same speed.
> 
> Find out the state of the asteroids after all collisions. If two asteroids meet, the smaller one will explode. If both are the same size, both will explode. Two asteroids moving in the same direction will never meet

## Solutions
仔细分析一下，会发现只有一种情况下才会出现相互抵消的可能：前一个元素非空且前一个元素为正，后一个元素为负。因为我们需要不断的知道前一个元素是什么，所以可以使用栈来解。

具体解法就是，不断的进行遍历，如果当前遍历到的元素为负数且栈非空，栈顶元素为正，那么就需要考虑从栈中弹出元素，其它情况下，都需要把元素直接入栈。那么怎么弹出呢？**注意这里的弹出顺序是很重要的，另外弹出时候需要保证栈非空且栈顶元素为正**：
- 首先需要考虑栈顶元素比当前遍历元素要小的情况，此时需要把栈顶元素不断出栈。注意我们这里说的小，都是指小于当前遍历元素的绝对值。
- 然后上面的循环可能会出现两种情况，这两种情况下，我们都需要把当前遍历元素入栈，然后开启下一轮循环：
  - 当前遍历元素比栈中所有元素都大，此时栈中没有元素了；
  - 当前栈顶元素出现了负数，那么此时栈顶元素和遍历元素同号，此时无法抵消，也需要入栈。
- 上面的循环还可能会导致栈顶元素与当前遍历元素相等，此时弹出栈顶元素，然后开启下一轮循环。
```go
func asteroidCollision(asteroids []int) []int {
    size := len(asteroids)
    if size <= 0 {
        return nil
    }
    
    var stack []int
    stack = append(stack, asteroids[0])
    for i := 1; i < size; i++ {
        idx := len(stack) - 1
        // 只有栈顶为正数且当前遍历元素为负数的情况才需要考虑弹栈
        if idx >= 0 && asteroids[i] < 0 && stack[idx] > 0 {
            // 如果当前遍历到的元素一直比栈顶元素要大，那么将栈顶元素不断出栈
            // 需要保证栈非空且栈顶元素为正数
            for idx >= 0 && stack[idx] > 0 && stack[idx] < abs(asteroids[i]) {
                stack = stack[:idx]
                idx--
            }
            
            // 上面的循环中可能出现栈中所有元素都要比当前遍历元素小，或者栈非空但是栈顶元素跟当前遍历元素同号且相等的情况，此时需要将当前遍历元素入栈；然后继续下一轮遍历
            // 
            if idx < 0 || idx >= 0 && stack[idx] < 0 {
                stack = append(stack, asteroids[i])
                continue
            }
            
            // 上面的循环中可能出现栈顶元素与当前遍历到的元素相等的情况，此时将栈顶元素出栈，然后继续下一轮遍历
            // 需要保证栈非空且栈顶元素为正
            if idx >= 0 && stack[idx] > 0 && abs(asteroids[i]) == stack[idx] {
                stack = stack[:idx]
                idx--
                continue
            }

            // 上面的循环中可能出现栈顶元素比当前遍历元素要大的情况，此时当前遍历元素被抵消，继续下一轮遍历
            // 需要保证栈非空且栈顶元素为正
            if idx >= 0 && stack[idx] > 0 && stack[idx] > abs(asteroids[i]) {
                continue
            }
            
        } else {
            // 其余情况都不存在相互抵消的可能，所以直接入栈即可
            stack = append(stack, asteroids[i])
        }
    }
    
    return stack
}


func abs(a int) int {
    if a < 0 {
        return -a
    }
    
    return a
}
```
