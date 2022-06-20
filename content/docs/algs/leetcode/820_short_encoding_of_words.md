---
title: 0820. Short Encoding of Words
weight: 10
tags: [
	"String",
	"Hash Table"
]
---
## Description


## Solutions
### Hash Table
> 这道题给了我们一个单词数组，让我们对其编码，不同的单词之间加入#号，每个单词的起点放在一个坐标数组内，终点就是#号，能合并的单词要进行合并，问输入字符串的最短长度。题意不难理解，难点在于如何合并单词，我们观察题目的那个例子，me和time是能够合并的，只要标清楚其实位置，time的起始位置是0，me的起始位置是2，那么根据#号位置的不同就可以顺利的取出me和time。需要注意的是，如果me换成im，或者tim的话，就不能合并了，因为我们是要从起始位置到#号之前所有的字符都要取出来。

这里使用 Hash Table 来做，将所有的单词先放到这个Hash Table 中。原理是对于每个单词，我们遍历其所有的后缀，比如time，那么就遍历ime，me，e，然后看 Hash Table 中是否存在这些后缀，有的话就删掉，那么 Hash Table 中的 me 就会被删掉，这样保证了留下来的单词不可能再合并了，最后再加上每个单词的长度到结果 res，并且同时要加上 # 号的长度。

```go
func minimumLengthEncoding(words []string) int {
    // 先将每个单词放进 hash 表中
    dict := make(map[string]bool)
    for _, word := range words {
        dict[word] = true
    }
    
    // 对于 words 中的每个单词，找到其每个后缀，查看该后缀是否在 dict 中，如果在的话，则从 dict 中将该后缀删除
    for _, word := range words {
        size := len(word)
        for i := 1; i < size; i++ {
            if _, ok := dict[word[i:]]; ok {
                delete(dict, word[i:])
            }
        }
    }
    
    // 最后 dict 中剩下来的字符串都是不能合并的字符串，因为每个字符串和周围的字符串都要有一个#分隔，最后一个位置也需要一个#分隔，
    // 所以可以遍历 dict 中剩下的每个字符串，将其长度加 1 放到 res 中
    var res int
    for word, _ := range dict {
        res += len(word) + 1
    }
    
    return res
}
```
