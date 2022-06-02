
## Description
> Given an array of strings words and an integer k, return the k most frequent strings.
> 
> Return the answer sorted by the frequency from highest to lowest. Sort the words with the same frequency by their lexicographical order.
## Solutions

### Hash Table & Sort
首先还是创建一个 dict 统计一下各个单词出现的次数， 然后用单词以及该单词出现的次数组成 pair 对，把这些 pair 对组成数组中，然后根据每个单词的出现频率对这个数组进行排序，如果单词频率相同，则根据单词的字母顺序进行排序。最后放回前 k 个即可。
```go
func topKFrequent(words []string, k int) []string {
    dict := make(map[string]int)
    for _, word := range words {
        dict[word]++
    }
    
    type Pair struct {
        str string
        cnt int
    }
    var pairs []Pair
    for word, cnt := range dict {
        pairs = append(pairs, Pair{word, cnt})
    }
    
    sort.Slice(pairs, func(i, j int) bool {
        if pairs[i].cnt == pairs[j].cnt {
            return pairs[i].str < pairs[j].str
        }
        
        return pairs[i].cnt > pairs[j].cnt
    })
    
    var res []string
    for i := 0; i < k; i++ {
        res = append(res, pairs[i].str)
    }
    
    return res
}
```

