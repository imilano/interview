---
title: Blind 75 List
weight: 10
---
## Blind 75

### Sequence
1. Two Sum
   
   使用一个 map 记录元素值与元素下标的映射，只需要一次扫描数组，检查 target - value 是否在 map 中，同时将 value 和其下标放到 map 中 。
   
2. Best Time to Buy and Sell Stock
   
   使用一个值记录当前扫描到的最小值，然后扫描一次数组，每扫描到一个值，用当前值减去当前得到的最小值，然后得到结果，取结果的最大值。如果当前值小于最小值，则更新最小值为当前值。只需要一次扫描。

3. Contains Duplicate

	简单题，只需要一个 map 就可以了，扫描一次即可。

4. Product of Array Except Self
   创建两个数组left 和 right，数组 left[i] 表示 i 位置前的元素的累积，数组 right[i]表示 i 位置后的元素的累积。然后创建一个数组 res， res[i] = left[i] * right[i]。

5. Maximum Subarray

	Kadane 算法，最优子结构性质，当前位置的和取决于前一个位置的子数组和以及当前元素。
   ```python

   def max_subarray(arr):
      max_sum_so_far, sum_ending_here = arr[0], arr[0]
      for e in range(arr[1:]):
         sum_ending_here = max(e, sum_ending_here + e)
         max_sum_so_far = max(max_sum_so_far, sum_ending_here)

      return max_sum_so_far
   ```

   如果数组中有负数，并且允许返回长度为 0 的子数列，则该问题可以改为：
   ```python
   def max_subarray(arr):
      max_sum_so_far, sum_ending_here = 0, 0
      for e in range(arr):
         sum_ending_here = max(sum_ending_here + e, e)
         max_sum_so_far = max(max_sum_so_far, sum_ending_here)
   return max_sum_so_far
   ```

	TODO 重点关注

1. Maximum Product Subarray
   
   动态规划，这里看起来和 53 题类似，但是其实这里的 0 和负数会对结果造成比较大的影响。解决方法是使用两个数组，一个数组 greater 记录包含当前元素的最大积的值，另一个数组 lesser 记录包含当前元素的最小积的值。
   然后当前的最大最小值只会在 greater[i-1]*nums[i]、nums[i] 和 lesser[i-1]*nums[i] 中产生。

   TODO 重点关注这题

1. Find Minimum in Rotated Sorted Array

   主要还是：举出例子，注意观察；注意这个性质：旋转数组一分为二之后，其中一个一定是有序的，另一个可能有序、可能无序。
   
1. Search in Rotated Sorted Array

   二分法。谁能想到是需要根据 middle 来判断向哪一边前进呢，而且 left 和 right 的取值也很巧妙.

   对于旋转数组，需要注意的是：**因为数组原先是有序的，所以将数组一分为二之后，其中一定有一个是有序的，另一个可能有序，也可能部分有序。此时有序部分用二分查找。无序部分再一分为二，其中一个一定有序，而另一个可能有序，可能无序。循环即可。**

   在使用while 循环进行搜索的时候，一定要注意缩小范围。也就是说，一方面要让结果收敛，另一方面要保持出口单一。
   
   TODO 二分法复习

1. 3 Sum
    
    排序+双指针
    需要注意这几点：
      - 适当剪枝： 包括如何去重、为什么要进行排序、遍历多少次、为什么遇到整数之后可以直接跳出循环
      - 有序数组使用双指针
      - 是否可以更改原数组，如果可以，那么排序是否会带来不一样的效果

1. Container With Most Water
    
   贪心，怎么移动，移动的时候是移动左边还是移动右边，基于什么策略去移动，木桶效应。

### Binary

1. Sum of Two Integers
    
    二进制运算（异或模拟二进制加法和，与运算模拟进位，进位还需要左移），何时结束递归。

1. Number of 1 Bits
   
   Super easy, 后面不需要再花时间看。

1. Counting Bits
    
   想到解法很简单，但找到规律不容易。有趣的规律是：如果 n 是奇数，数字 n 中包含的 1 的个数等于数字 n/2 所包含的 1 的个数加 1；如果 n 是偶数，数字 n 包含的 1 的个数等于数字 n/2 包含的 1 的个数。

1. Missing Number
   
   很简单，只需要记住等差数列求和公式即可： n*(a1+an)/2

1. Reverse Bits

   很容易想到解法。

### Dynamic Programming
1. Climbing Stairs
    
   斐波那契数列，Fn = Fn-1 + Fn-2   

1. Coin Change
   
   TODO recheck

   注意的是，局部最优解不等于全局最优解，所以这里使用贪心算法得到一个可行解是不够的，要求出所有解，然后取其中最小值。推导式： dp[i] = min(dp[i], dp[i - coins[j]] + 1) if coins[j] <= dp[i].

1. Longest Increasing Subsequence

   TODO recheck

  这题需要重点关注一下，因为自己对这个的应用还不够熟练。

  核心代码：
   ```go
	for i := 1; i < size; i++ {
		for j := 0; j < i;j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+ 1)
			}
		}
		
		// 注意这里，最后所求的并不是dp[size-1]
		res = max(res, dp[i])
	}

	return res
}
   ```

字符串求极值一般可以考虑一下使用动态规划，动态规划想不出转移方程则可以尝试使用贪心。

1. Longest Common Subsequence
   
   TODO recheck

   这题跟 583_delete_operation_for_two_strings 是一样的解法，也跟 516_longest_palindromic_subsequence 是一样的，是典型的的二维dp问题。

   在 longest palindromic subsequence 中，一定要注意遍历的方向，dp具有最优子结构性质，而这意味着你在求解一个问题时，需要完全求解出这个问题的子问题的解。

   需要注意这几个问题：
      - 状态方程式怎么定义的
      - 为什么数组是 (m+1) * (n+1) 的
      - 为什么比较的时候是 text1[i-1] 与 text2[j-1]比较

1. Longest Palindrome Subsequence
    
    TODO recheck

    参见上一题的解答。 

1. Delete Operation for Two Strings
   
   TODO recheck
   
   参见上一题的解答

1. Word Break

   TODO  recheck

   这一题需要特别注意，尤其需要注意dp数组的定义，这里的初始化条件也需要注意，此外还需要注意子问题的求解方向。结束遍历的条件也是需要特别注意的。
1. Combination Sum
   
   Combination Sum 是一系列题，分别有leetcode 39 题 Combination Sum， 40 题的 Combination Sum II, 216 题的 Combination Sum III 和 377 题的 Combination Sum IV。

   39 题比较简单，注意边界条件即可。40 题和39 题其实是差不多的。216题也不难。

   377 题竟然是使用动态规划来解的，我肯定想不到这种解法。 
   
   TODO

   在 39 题中，需要注意的是，在回溯的过程中， 我们可以一次性添加多个元素，而不是只添加一个，这样可以减少一些回溯过程：
   ```go
    for i := start; i < size; i++ {
        for j := target/candidates[i]; j > 0; j-- {
            //一次性添加多个同一元素
            for k := 0; k < j; k++ {
                cur = append(cur, candidates[i])
            }
            combinationSumHelper(candidates, i+1, target - j * candidates[i], cur, res)
            // 复原
            cur = cur[:len(cur) - j]
        }
    }
   ```

   在第 40 题中，需要注意的是，可以直接对数组进行排序，从而就不需要再使用 map 来进行手动去重。另外还需要特别注意该题中的剪枝用法。
   
   
2. House Robber
    
   TODO recheck

   这个题需要注意初始条件是怎么算出来的。

3. House Robber II

   TODO recheck
   
   这个好取巧，直接剔除抢和不抢的情况，然后算两次取最大值。

4. Decode ways
    
    TODO 这一题还没写 

5. Unique Path

   注意排列组合计算方法。注意这里为什么需要用更小的值来作为上值，为什么选择另一个数可能会导致溢出，从而得到结果0.
   ```go
   // 计算从m里面抽n个
   res := 1
   for i := 0; i < n; i++ {
      res *= (m-i)
      res /= (i+1)
   }

   ```

6. Jump Games
    
   TODO recheck 对于动态规划，自己的解题能力还需要加强

    这里的贪心解法很有意思。



### Graph

1. Clone Graph

   TODO recheck
    
    本质上就是进行一次 DFS 或者 BFS，但是要注意去重。这里需要注意的是，map 的 key 和 value 分别对应的原节点和其克隆出来的节点吗，而不是仅仅记录一个节点是否有被 clone 过。


1. Course Schedule

   很明显的有向图环检测问题。这里并不需要我们真的建立一个图，而是可以使用邻接表（也就是二维数组）的方式来表示图，然后使用一个一维数组来表示各个节点的入度是对少。

   遍历的时候，将那些入度为0的节点放入到队列中，然后对队列进行迭代，直到队列长度为0即可。这里需要注意的是，在进行 BFS 的时候，在添加新节点的时候注意不要把那些已经遍历过的节点重新放入到队列中，这样会导致重复遍历形成环，从而导致超时。

1. Pacific Atlantic Water Flow
   对每个点都可以进行一次深度遍历，分别看这个点能否到达两大洋，如果一个点既能够到达大西洋，也能够到达太平洋，那么这个点就是结果之一。这里的一个优化方法就是，由于必须要到达两大洋，所以不需要对每个点都进行遍历，而是只需要对边上的点进行遍历即可。对边上的点进行深度遍历，如果一个点既能够到达太平洋，也能够到达大西洋，那么这个点就是我们所需要的结果。
1. Number of Islands
   
   抽象成一个图，使用深度优先遍历，遍历之后看一共有多少个连通分量即可。

   不是很难，但是要注意边界条件的设置。

1.  Longest Consecutive Sequence
   
   TODO  recheck

   这里跟图关系不大，主要还是使用 map，然后使用一个相对比较取巧的方法来计算。

1. Alien dict
   
   这题稍微看上去有点复杂，现在暂时还没做。

1. Graph Valid Tree
   
   这题很简单，就是验证给出的节点能够构成树，如果要构成树，那么就需要满足两个条件：首先就是没有环，其次就是只有一个连通分量。这个很容易解决。

1. Number of Connected Components in an Undirected Graph 
   
   这里也很简单。主要问题还是如何表示无向图和有向图的问题，解决这个问题之后，后续无非是深度遍历和广度遍历的问题。


### Interval

1. Insert Interval
   TODO recheck

   这个题需要好好注意，自己在这个题的解法上并不简介。注意边界条件。

1. Merge Intervals
   做了上一题之后，应该不太难想到这一题的思路。

1. Non-overlapping Intervals
   TODO recheck

   这题我不会，需要多注意。注意贪心策略的选择。

1. Meeting Rooms
   简单题，无需再看。

1. Meeting Rooms II
   TODO recheck

   这题比较难，最好能够动手画一下图，这样会好理解一些。 

### Linked list

链表这一块要特别注意虚拟头结点的使用，插入一个虚拟头节点能够避免很多问题。

1. Reverse a linked list
   简单题

1. Detect Cycle in a Linked List
   
   简单题，可以使用快慢指针或者 map。

1. Merge Two Sorted List

   简单题，无需再过多关注

1. Merge K Sorted List
   简单题， 只需要使用堆即可。


1. Remove Nth Node From End Of List

   简单题，无需过多关注。

1. Reorder List

   TODO recheck 

   这个也不是很难，其实还是链表的反转问题。这里一个需要注意的点是，如何找到 mid 节点，快慢指针的快指针是如何进行遍历的，以及特别需要注意的是，找到中点之后，要把前半段和后半段断开，不然很容易导致出现环，从而出现超时。

### Matrix
1. Set Matrix Zeroes
   TODO recheck
   解法比较巧妙，需要注意一下。

1. Spiral Matrix
   TODO recheck

   注意遍历方向和 coner case

1. Rotate Image
   TODO recheck

   注意矩阵主对角线旋转和副对角线旋转的公式。

   矩阵旋转：
      - 主对角线对折： arr[i][j] = arr[j][i]
      - 副对角线对折： arr[i][j] = arr[n-j-1][n-i-1]
         ```go
            for i := 0; i < n; i++ {
               for j := 0; j < n-i; j++ {
                  matrix[i][j], matrix[n-j-1][n-i-1] = matrix[n-j-1][n-i-1], matrix[i][j]
               }
            }
         ```
      - 横向对折： arr[i][j] = arr[n-i-1][j]
      - 纵向对折： arr[i][j] = arr[i][n-j-1]
  
1. Word Search
   同 Word Search II 一样，这里需要注意的是，不同情况的回溯函数之间的状态不能相互干扰。

### String

前三题都是滑动窗口题目。回文子串题注意活用中心扩散方法。
1. Longest Substring Without Repeating Characters
   TODO recheck

   注意这里是如何维持滑动窗口的。尤其要注意的是移动左指针的条件：字符出现在了 map 中但是并不一定出现在了滑动窗口中，只有字符出现在了滑动窗口中才需要滑动左指针。
1. Longest Repeating Character Replacement
   TODO recheck

   这题比较难，不过分析过程很有意思。

1. Minimum Window Substring
   TODO recheck

   这题也比较难，问题在于如何维持一个可以包含冗余字符的滑动窗口，值得再看几遍。

1. Valid Anagram
   简单题，无需关注。

1. Group Anagrams
   简单题，无需再关注。

1. Valid Parentheses
   简单题，无需关注。

1. Valid Palindrome
   简单题，无需关注。

1. Longest Palindromic Substring
   可以关注一下如何使用动态规划的方式来解决这个问题。

1. Palindromic Substrings
   TODO recheck

   这里要注意活用中心扩散算法

1. Encode and Decode Strings
   两种方法：
    - 编码时间把字符串长度编码进去，解码的时候直接根据长度进行分割即可
    - 使用一些特殊字符进行分隔
### Tree

1. Maximum Depth of Binary Tree
   简单题，无需再看。

2. Same Tree 

   TODO recheck

   简单题，但是还是需要注意一下。

3. Invert a Binary Tree
   简单题

1. Binary Tree Maximum Path Sum
   
   TODO recheck

   难题，但是分析过程还是比较有意思的。需要注意的是，res 初值不能取 0，而是要取一个最小值.


1. Serialize and Deserialize Binary Tree
   
   TODO  recheck

   其实不难，只需要使用层次遍历即可。主要的问题是，在层次遍历序列化的时候，需要保留空节点，在反序列化的时候，也需要保留空节点。

1. Subtree of Another Tree
   TODO recheck

   不是很难，关键点在于，如何写对 sameTree 这个函数。

1. Construct Binary Tree from Preorder and Inorder Traversal
   
   这题主要注意 corner case

1. Validate Binary Search Tree
   
   TODO  recheck

   这里需要注意的是如何充分利用中序遍历。这里的递归方法值得好好研究一下。

1. Kth Smallest Element in a BST
   简单题，只需要中序遍历的迭代解法即可解决。

1. Lowest Common Ancestor of BST
   TODO recheck

   关键在于抓住 BST 的性质。

1. Implement Trie(Prefix Tree)
   TODO recheck

   注意如何表示前缀树
1. Design Add and Search Words Data Structure
   TODO recheck

   虽然是前缀树的一个应用，但是需要注意如何应对通配符。
   
1. Word Search II
   这里需要注意两点：
   - 首先可以复用原数组，让其保存访问状态，这样就可以避免开辟额外的 visite 的数组来节省空间
   - 如果不复用原数组，而要额外开辟 visited 数组的话，注意每次搜索的时候 visited 的状态不能相互影响。
   
### Heap
1. Merge K Sorted List
   这个很简单，只需要好好使用堆即可。

1. Top K Frequent Elements

   可以用 Map 来做，key 表示元素，value 表示元素出现的次数，如果每个元素出现的次数是独一无二的，那么最后对 v 进行排序即可。
   也可以用自定义大根堆来做，堆的元素是二元数组，其中第一个元素表示出现的次数，第二个元素表示元素值。

1. Find Median from Data Stream
   TODO recheck

   这里的难点在于在插入时，如何保证两个堆的数量基本上是相等的。解决办法是，优先插入左边大根堆，然后将左边的大根堆堆顶元素插入右边小根堆，此时如果发现左边大根堆的数量小于右边小根堆的数量，那么将右边小根堆的堆顶元素插入左边小根堆。这样的话，就可以维持：左边的大根堆元素数量要么比右边大 1，要么二者相等。
