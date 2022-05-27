---
title: 1209. Remove All Adjacent Duplicates in String
weight: 10
tags: [
	"Stack"
]
---
## Description
> You are given a string s and an integer k, a k duplicate removal consists of choosing k adjacent and equal letters from s and removing them, causing the left and the right side of the deleted substring to concatenate together.
> 
> We repeatedly make k duplicate removals on s until we no longer can.
> 
> Return the final string after all such duplicate removals have been made. It is guaranteed that the answer is unique.

## Solutions
第一想法还是使用跟 1047 题差不多同样的解法来解，但是可惜超时了。这里把代码贴出来：
```go
func removeDuplicates(s string, k int) string {
    size := len(s)
    if size < k {
        return s
    }
    
    var stack []rune
    for _, r := range s {
        size = len(stack)
        if size < k-1 || getSub(r, k-1) != string(stack[size-k+1:]) {
            stack = append(stack, r)
            continue
        }
        
        stack = stack[:size-k+1]
    }
    
    return string(stack)
}

func getSub(r rune, k int) string {
    var res string
    for i := 0; i < k; i++ {
        res += string(r)
    } 
    
    return res
}
```

看到网上的一种解法，还是使用栈，但是这里栈中并不直接存储字符，而是存储一个 pair 对，这个 pair 里包含两个元素，一个是对应的字符，另一个是该字符出现的次数。初始栈为空，然后开始遍历字符串，如果栈为空或者当前遍历到的元素跟栈顶元素不一致，则将当前遍历到的元素入栈；如果当前遍历到的元素与栈顶元素一致，则更新该栈顶元素字符的出现次数，如果出现次数到达了 k，那么就将该字符出栈。最后根据栈中的 pair 对再重新组装一次字符串即可。这个方法确实挺妙的.
```go
type Pair struct {
    r rune
    count int
}

func removeDuplicates(s string, k int) string {
    size := len(s)
    if size < k {
        return s
    }
    
    var stack []Pair
    for _, r := range s {
        size = len(stack)
        // 如果栈为空或者当前遍历到的元素与栈顶元素不同，则入栈
        if size == 0 || stack[size-1].r != r {
            stack = append(stack, Pair{r:r, count:1})
            continue
        }

        // 如果当前遍历到的元素与栈顶元素相同，则判断是否需要将该元素出栈
        if stack[size-1].r == r && stack[size-1].count < k {
            // 如果已经有了 k 个，则将该元素出栈
            stack[size-1].count++
            if stack[size-1].count == k {
                stack = stack[:size-1]
            }
        }
    }
    
    var res string
    for _, pair := range stack {
        res += getSub(pair.r, pair.count)
    }
    return res
}

func getSub(r rune, k int) string {
    var res string
    for i := 0; i < k; i++ {
        res += string(r)
    } 
    
    return res
}
```
