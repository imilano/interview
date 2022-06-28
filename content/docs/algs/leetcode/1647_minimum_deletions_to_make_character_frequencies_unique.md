---
title: 1647. Minimum Deletions to Make Character Frequencies Unique
weight: 10
tags: [
	"Array",
	"Greedy",
	"Sort"
]
---
## Description
> A string s is called good if there are no two different characters in s that have the same frequency.
> 
> Given a string s, return the minimum number of characters you need to delete to make s good.
> 
> The frequency of a character in a string is the number of times it appears in the string. For example, in the string "aab", the frequency of 'a' is 2, while the frequency of 'b' is 1.


## Solutions
### Greedy && Sort
题主初始看到这个题，没啥思路，但是看到 topic 是贪心和排序，心里就大概直到是咋回事了。首先肯定需要统计每个字符的出现次数，然后去做删除操作，那么怎么删除呢？先需要对每个字符的出现次数按照从大到小排个序，然后从第二大的数字开始遍历，如果这个数字大于等于前一个数字，那么就不断把这个数字减 1（这里应该也可以一步到位，不需要使用 for 循环），直到当前数字比前一个数字要小，减的过程中不断增加 res 的值，最后遍历完之后，res 的值即为结果。这里需要注意的是，为什么是从大到小进行排序，而不是从小到大进行排序呢？这里就是一个贪心的思路，如果我们从小到大进行排序的话，我们对频率数字从头到尾进行处理，那么就要要求后面的数字比前面的数字要大，如果不大的话，就要把后面的整个数字删掉。但是注意啊，因为后面的数字往往更大，我们这么一删，就相当于在结果 res 上增加了一个很大的数；而如果我们是从大到小排序，并且从大到小进行处理，那么这个时候后面的数更小，删去后面的数只会给 res 增加一个很小的值，那这显然是从大道小进行处理更划算了。
```go
func minDeletions(s string) int {
    var res int
    // corner case
    size := len(s)
    if size <= 1 {
        return res
    }
    
    // 统计每个字符的出现频率
    dict := make(map[byte]int)
    for idx, _ := range s {
        dict[s[idx]]++
    }
    
    // 对出现频率按照从大到小进行排序
    var fre []int
    for _, v := range dict {
        fre = append(fre, v)
    }
    sort.Slice(fre, func(i,j int) bool {
        return fre[i] > fre[j]
    })
    // 从第二个数开始进行遍历，如果该数大于等于前一个数，那么不断削减该数的值，直到比前一个数要小
    size = len(fre)
    for i := 1; i < size; i++ {
        // 这里要限制 fre[i] > 0 ，因为一旦频率为 0，说明这个字符串就不需要出现了，没必要再往下减了
        for fre[i-1] <= fre[i] && fre[i] > 0 {
            // 每次削减都需要让结果自增
            fre[i] -= 1
            res++
        }
    }
    
    return res
}

```

上面的 for 循环也可以一步到位，写成下面这样：
```go
func minDeletions(s string) int {
    var res int
    // corner case
    size := len(s)
    if size <= 1 {
        return res
    }
    
    // 统计每个字符的出现频率
    dict := make(map[byte]int)
    for idx, _ := range s {
        dict[s[idx]]++
    }
    
    // 对出现频率按照从大到小进行排序
    var fre []int
    for _, v := range dict {
        fre = append(fre, v)
    }
    sort.Slice(fre, func(i,j int) bool {
        return fre[i] > fre[j]
    })
    // 从第二个数开始进行遍历，如果该数大于等于前一个数，那么不断削减该数的值，直到比前一个数要小
    size = len(fre)
    for i := 1; i < size; i++ {
        // 这里要限制 fre[i] > 0 ，因为一旦频率为 0，说明这个字符串就不需要出现了，没必要再往下减了
        if fre[i] >= fre[i-1] {
            if fre[i-1] <= 1 {
                res += fre[i]
                fre[i] = 0
            } else {
                res += fre[i] - (fre[i-1]-1)
                fre[i] = fre[i-1] - 1
            }
        }
    }
    
    return res
}
```
