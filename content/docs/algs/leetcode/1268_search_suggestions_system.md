---
title: 1268. Search Suggestions System
weight: 10
tags: [
	"Hash Table"
]
---
## Description
> You are given an array of strings products and a string searchWord.
> 
> Design a system that suggests at most three product names from products after each character of searchWord is typed. Suggested products should have common prefix with searchWord. If there are more than three products with a common prefix return the three lexicographically minimums products.
> 
> Return a list of lists of the suggested products after each character of searchWord is typed.

## Solutions
### Hash Table
这里可以使用哈希表来解决。创建一个哈希表，其 key 为一个单词的前缀，而 value 为具有这个前缀的字符串数组。遍历 products 数组，然后数组中的每一个字符串，将该字符串对应的前缀作为 key，然后将该字符串添加到该 key 对应的字符串数组中。因为要求返回的结果要 lexicographically minimum，所以在遍历之前先对 products 进行一次排序。最后再遍历这个searchWord，从哈希表中取出对应的结果即可。
```go
func suggestedProducts(products []string, searchWord string) [][]string {
    dict := make(map[string][]string)
    sort.Strings(products)
    for _, product := range products {
        size := len(product)
        for i := 1; i <= size; i++ {
            dict[product[:i]] = append(dict[product[:i]], product)
        } 
    }
    
    var res [][]string
    size := len(searchWord)
    for i := 1; i <= size; i++ {
        if arr, ok := dict[searchWord[:i]]; ok && len(arr) >= 3 {
            res = append(res, arr[:3])
            continue
        }
        
        // 这里其实是两种情况：
        //  - 搜索词在 dict 中存在，并且对应的数组长度不足 3，那么直接添加该数组即可
        //  - 搜索词在 dict 中不存在，因为 golang map 中访问一个不存在的 key 会返回对应的零值（这里就是空数组），所以也可以直接添加
        res = append(res, dict[searchWord[:i]])
    }
    
    return res
}
```
### Pruning
在网上看到一种比较巧妙的方法，该方法的核心思想其实还是剪枝。首先先给 products 数组排个序，然后维护一个 suggested 数组，初始时suggested 数组直接拷贝 products 数组即可。然后在敲入每个字符的时候，新建一个 filter 数组，此时遍历 suggested 数组，如果单词对应位置的字符是敲入的字符的话，将单词加入 filter 数组，这样的话 filter 数组的前三个单词就是推荐的单词，取出来组成数组加入结果 res 中。然后把 suggested 数组更新为 filter 数组，这个操作就缩小了下一次查找的范围，算是一个比较巧妙的办法了。
```go
func suggestedProducts(products []string, searchWord string) [][]string {
    sort.Strings(products)
    suggested := products
    var res [][]string
    
    size := len(searchWord)
    for i := 1; i <= size; i++ {
        // 找出前缀匹配的字符加入到 filter 中
        var filter []string
        for _, word := range suggested {
            if len(word) >= i && word[:i] == searchWord[:i] {
                filter = append(filter, word)
            }
        }
        
        // 将 filter 中的前三个字符放入 res
        var tmp []string
        for _, word := range filter {
            if len(tmp) == 3 {
                break
            }
            
            tmp = append(tmp, word)
        }
        
        res = append(res, tmp)
        // 将 suggested 替换为新的 filter 
        suggested = filter
    }
    
    return res
}
```

### Trie Tree
这里还可以使用前缀树来做。不过相比上面两种方法，前缀树的解法还是稍微复杂了点，所以这里就不写了。
