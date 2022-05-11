---
title: 38. Count and Say
weight: 10
---
## Descrition
> The count-and-say sequence is a sequence of digit strings defined by the recursive formula:
> 
> countAndSay(1) = "1"
> countAndSay(n) is the way you would "say" the digit string from countAndSay(n-1), which is then converted into a different digit string.
> To determine how you "say" a digit string, split it into the minimal number of groups so that each group is a contiguous section all of the same character. Then for each group, say the number of characters, then say the character. To convert the saying into a digit string, replace the counts with a number and concatenate every saying.


## Solutions

统计每个字符出现的次数即可。

```go
func countAndSay(n int) string {
    if n == 0 {
        return "0"
    }   
    
    res := "1"
    for n-1 > 0 {
        cur := ""

		// 统计 res 中每个字符出现的次数，然后将再将次数和该字符拼接到 cur 上，最后将 cur 赋值给 res
        size := len(res)
        for i := 0; i < size; i++ {
            cnt := 1
            for i + 1 < size && res[i] == res[i+1] {
                cnt++
                i++
            }
            
            cur += strconv.Itoa(cnt) + string(res[i])
        }
        
        res = cur
        n--
    }
    
    return res
}
```
