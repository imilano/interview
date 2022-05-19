---
title: 0088. Merge Sorted Array
weight: 10
---

## Description

> You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.
> 
> Merge nums1 and nums2 into a single array sorted in non-decreasing order.
> 
> The final sorted array should not be returned by the function, but instead be stored inside the array nums1. To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.

## Solutions
题目意思是 nums1 的总长是 n+m， 但是只有前 m 个元素才是合法的，后 n 个元素相当于是空位置。现在需要你合并一下 nums1 和 nums2. Follow up 那里是问你是否能够就地合并。

### 归并排序
很明显应该可以想到归并排序的思路，但是问题是，归并排序需要使用额外的空间。那么如何能够不使用额外的空间也可以合并这两个数组呢（这也是 Follow up 的要求）？其实这里要学会充分利用已有的数组空间，在 LeetCode 的众多题目中，利用已有空间的题并不少。这里我们可以先把 nums1 的 m 个元素右移，然后前面就可以空出来 n 个元素，这就相当于 n 个空位，那么我们就可以借助这 n 个空位完成归并排序。
```go
func merge(nums1 []int, m int, nums2 []int, n int)  {
    // corner case
    if m == 0 && n == 0 || n == 0 {
        return
    }
    
    // 将前 m 个元素每个往后移动 n 位
    size := len(nums1)
    for i := m-1; i >= 0; i-- {
        nums1[i+n] = nums1[i]
    }
    
    // 归并排序
    // 注意 pos1 的起始坐标
    pos1, pos2, cur := size-m, 0, 0
    for pos1 < size && pos2 < n  {
        if nums1[pos1] < nums2[pos2] {
            nums1[cur] = nums1[pos1]
            pos1++
        } else {
            nums1[cur] = nums2[pos2]
            pos2++
        }
        
        cur++
    }
    
    for pos1 < size {
        nums1[cur] = nums1[pos1]
        pos1++
        cur++
    }
    
    for pos2 < n {
        nums1[cur] = nums2[pos2]
        pos2++
        cur++
    }
}
```
### 插入排序

首先，nums1 的前 m 个元素已经排好序了，那么我们只需要将 nums2 的 n 个元素插入到 nums1 中，然后从插入的位置开始进行插入排序即可。

```go
func merge(nums1 []int, m int, nums2 []int, n int)  {
    if m == 0 && n == 0 || n == 0 {
        return
    }
    
    
    size := len(nums1)
    for i,j := m, 0; i < size; i,j = i+1, j+1 {
        nums1[i] = nums2[j]
    }
    
    for i := m; i < size; i++ {
        tmp := nums1[i]
        j := i
        for j > 0 && nums1[j-1] > tmp {
            nums1[j] = nums1[j-1]
            j--
        }
        nums1[j] = tmp
    }
}
```
