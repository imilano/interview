package leetcode

import "math"

/*
Given two strings s and t of lengths m and n respectively, return the minimum window substring of s such that every character in t (including duplicates) is included in the window. If there is no such substring, return the empty string "".

The testcases will be generated such that the answer is unique.

A substring is a contiguous sequence of characters within the string.
*/

// 这里应该很容易可以看出需要使用滑动窗口来做，但是与一般的滑动窗口不同的是，这里的滑动窗口中可以包含一些不在字符串 t 中的字符。
// 另外还需要注意的是，结果字符串中的字符顺序并不需要和 t 中的字符顺序相同。
// 首相应该很容易能够想到的是，需要一个 map 来扫一遍 t，记录 t 中出现的字符的个数。而问题的难点在于如何维持窗口，并保证t 中的字符都出现在窗口内。
// 其实在维护这个滑动窗口的过程中，我们可以完全无视那些没有出现在 t 中的字符，这些字符在窗口有边界扩张的时候可以跳过，在左边界收缩的时候也可以跳过。
// 那么我们先扫描一遍 t，统计每个字符出现的个数，将其存到字典 dict 中。然后初始化一个 left = 0，表示滑动窗口左边界；我们还需要一个 minLen，初始化为最大值，
// 表示包含 t 中的左右字符的滑动窗口的最小长度；我们还需要一个变量 cnt， 表示当前滑动窗口中已经匹配到的 t 中字符个数。
// 接着开始遍历 s 中的每个字符，这相当于扩展滑动窗口有边界。对于每个遍历到的字符，在 dict 中减去该字符的出现次数，减去之后如果该字符的出现次数大于等于 0，则说明这个出现在滑动窗口中的字符 c 是出现在 t 中的，则 cnt 加 1。
// 这里忽略了那些没有出现在滑动窗口中但是没有出现在 t 中的字符，因为减去 1 之后，dict 对该字符的计数小于 0，并不会对 cnt 产生影响。
// 一旦 cnt 等于 t 长度，说明此时 t 中的所有字符都已经出现在了滑动窗口中，那么接下来就应该收缩窗口，找到能包含 t 中所有字符的最小窗口。需要收缩有边界吗？答案是不需要，因为我们在 cnt 等于 t 长度的时候才停下来，这个时候滑动窗口刚刚才匹配了 t 中最后一个字符，所以右边是没有冗余字符的，不需要收缩右边界。
// 但是对于左边界，是有可能会出现冗余的，于是我们就需要收缩左边界。那么怎么收缩呢？首先需要明确的是，我们收缩的目的是为了去除左边界中的冗余字符，那么就是说，在收缩左边界的过程中，滑动窗口中已经匹配的字符数 cnt 应该不受影响，其值应该正好等于 t 的长度。
// 收缩左边界的过程中，如果发现 minLen 比 i - left + 1还要长了，那么说明我们抛弃了更多的冗余字符，那么更新结果。在收缩的过程中，因为左边界移动了，所以要把 dict 中字符补充回来，所以需要对字符的出现次数加 1；如果加 1 之后，该字符的出现次数大于等于 1 了，那么说明这个字符是出现在 t 中的字符，这次收缩影响到了滑动窗口中的 t 中字符出现次数，则 cnt 需要减去 1.
func minWindow(s string, t string) string {
    m,n := len(s), len(t)
    var res string
    minLen, left, cnt := math.MaxInt, 0, 0
    
    dict := make(map[byte]int)
    for idx, _  := range t {
        dict[t[idx]]++
    }
    
    for i := 0; i < m; i++ {
        dict[s[i]]--
        if dict[s[i]] >= 0 {
            cnt++
        }
        
        for cnt == n {
			// 需要下面的 if，因为 cnt==n 会出现多次，我们不希望大的窗口覆盖掉我们的小窗口
			// res = s[left:i+1]
            if minLen > i - left + 1 {
                minLen = i - left + 1
                res = s[left:i+1]
            }
            
            dict[s[left]]++
            if dict[s[left]] > 0 {
                cnt--
            }
            left++
        }
    }
    
    return res
}
