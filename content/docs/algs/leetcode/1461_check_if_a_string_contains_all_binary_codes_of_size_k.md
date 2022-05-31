---
title: 1461. Check If a String Contains All Binary Codes of Size K
weight: 10
tags: [
	"Hash Table",
	"String"
]
---
## Description
> Given a binary string s and an integer k, return true if every binary code of length k is a substring of s. Otherwise, return false.

## Solutions

### Brute Force
下面是题主一开始想出来的方法，可惜最后几个用例过于变态，所以超时了。
```go
func hasAllCodes(s string, k int) bool {
    size := len(s)
    if size < k {
        return false
    }
    
    // golang 格式化打印字符串，%04b 这个格式化字符串能够将数字转换为二进制表示，并且当表示字符串长度小于 4 时，能够自动在前面增加 0 作为 padding。
    fmtStr := "%0" + strconv.Itoa(k) + "b"
    maxEle := int(math.Pow(2, float64(k))) - 1
    for i := 0; i <= maxEle; i++ {
        r := fmt.Sprintf(fmtStr, i)
        if !strAnd(s, r)  {
            return false
        }
    }
    
    return true
}


// 比较 op1 中是否包含 op2 这个字符串
func strAnd(op1, op2 string) bool {
    s1,s2 := len(op1), len(op2)
    for i := 0; i <= s1 - s2; i++ {
        if op1[i:i+s2] == op2 {
            return true
        }
    }
    
    return false
}
```

### Hash Table
很明显，上面的方法中 strAnd 部分的时间复杂度比较高，那么如何对这部分进行优化呢？strAnd 部分的时间复杂度较高的原因之一是，在使用不同的 op2 的时候，op1 都会被从头遍历，也就是说，op1 在 strAnd 中是不变的部分，重复对 op1 进行遍历是没必要的，我们可以直接使用一个 map 来存储 op1 中所有长度为 k 的子串，这样的话就能够省去很多的重复遍历。优化后的方法如下：
```go
func hasAllCodes(s string, k int) bool {
    size := len(s)
    if size < k {
        return false
    }
    
    // 一次性找出 s 中符号要求的字符串，这样可以避免 strAnd 中的重复比较
    dict := make(map[string]bool)
    for i := 0; i <= size-k; i++ {
        dict[s[i:i+k]] = true
    }
    
    // golang 格式化打印字符串，%04b 这个格式化字符串能够将数字转换为二进制表示，并且当表示字符串长度小于 4 时，能够自动在前面增加 0 作为 padding。
    fmtStr := "%0" + strconv.Itoa(k) + "b"
    maxEle := int(math.Pow(2, float64(k))) - 1
    for i := 0; i <= maxEle; i++ {
        r := fmt.Sprintf(fmtStr, i)
        if _, ok := dict[r]; !ok {
            return false
        }
        // if !strAnd(s, r)  {
        //     return false
        // }
    }
    
    return true
}


// 比较 op1 中是否包含 op2 这个字符串
func strAnd(op1, op2 string) bool {
    s1,s2 := len(op1), len(op2)
    for i := 0; i <= s1 - s2; i++ {
        if op1[i:i+s2] == op2 {
            return true
        }
    }
    
    return false
}
```

上面的解法还可以再优化,因为 map 自带去重功能，而一个字符串如果要包含所有长度为 k 的二进制字符串，那么 dict 的长度是需要大于等于 2^k 的。
```go
func hasAllCodes(s string, k int) bool {
    size := len(s)
    if size < k {
        return false
    }
    
    // 一次性找出 s 中符号要求的字符串，这样可以避免 strAnd 中的重复比较
    dict := make(map[string]bool)
    for i := 0; i <= size-k; i++ {
        dict[s[i:i+k]] = true
    }
    
    return  len(dict) >= (1 << k) 
    // 实际上到这里已经不需要下面的循环来进行比较了。因为 map 自带去重功能，而一个字符串如果要包含所有长度为 k 的二进制字符串，那么 dict 的长度是需要大于等于 2^k 的。
    // golang 格式化打印字符串，%04b 这个格式化字符串能够将数字转换为二进制表示，并且当表示字符串长度小于 4 时，能够自动在前面增加 0 作为 padding。
//     fmtStr := "%0" + strconv.Itoa(k) + "b"
//     maxEle := int(math.Pow(2, float64(k))) - 1
//     for i := 0; i <= maxEle; i++ {
//         r := fmt.Sprintf(fmtStr, i)
//         if _, ok := dict[r]; !ok {
//             return false
//         }
//         // if !strAnd(s, r)  {
//         //     return false
//         // }
//     }
    
//     return true
}


// 比较 op1 中是否包含 op2 这个字符串
func strAnd(op1, op2 string) bool {
    s1,s2 := len(op1), len(op2)
    for i := 0; i <= s1 - s2; i++ {
        if op1[i:i+s2] == op2 {
            return true
        }
    }
    
    return false
}
```
