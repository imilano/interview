---
weight: 5
bookCollapseSection: true
---

这里是一些 LeetCode 的刷题题解。

## Data structure

### Stack
| Problem                                        | Difficulty | Method                | Redo | Comment                                                            |
| :--------------------------------------------- | :--------- | :-------------------- | :--- | :----------------------------------------------------------------- |
| 155. Min Stack                                 | easy       | Stack、Monotone Stack | No   |                                                                    |
| 150. Evalute Reverse Polish Notation           | medium     | Stack                 | No   | 简单题，无需多看                                                   |
| 224. Basic Calculator                          | hard       | Stack                 | YES  | 题目为难题，暂时跳过。注意这里是如何将中缀表达式转换为后缀表达式的 |
| 227. Basic Calculator                          | medium     | Stack                 | YES  | 稍微有点耗时，先跳过                                               |
| 20. Valid Parentheses                          | easy       | Stack                 | No   |                                                                    |
| 1472. Design Browser History                   | medium     | Stack                 | No   |                                                                    |
| 1047. Remove All Adjacent Duplicates in String | easy       | Stack                 | No   |                                                                    |
| 1249. Minimum Remove to Make Valid Parentheses | medium     | Stack                 | Yes  |                                                                    |
| 735. Asteroid Collision                        | medium     | Stack                 | Yes  |                                                                    |

### Heap
| Problem                                      | Difficulty | Method                  | Redo | Comment                                        |
| :------------------------------------------- | :--------- | :---------------------- | :--- | :--------------------------------------------- |
| 973. K Closest Points to Origin              | medium     | Sort、 Heap             | No   | 要看到本质问题，然后解决本质问题。             |
| 347. Top K Frequent Elements                 | medium     | Sort、 Hash Table、Heap | No   |                                                |
| 23. Merge K Sorted Lists                     | medium     | Sort、 Heap             | No   |                                                |
| 378. Kth Smallest Element in a Sorted Matrix | medium     | Heap、Max Heap          | YES  | 这个题一开始没有想到堆的解法，还是需要注意一下 |
| 767. Reorganize String                       | medium     | Heap、Sort              | YES  | 这里的解法还是比较巧妙的，需要多注意一下       |
| 1642.Furthest Building you can Reach         | medium     | Heap、Priority Queue    | YES  | 这里的解法很巧妙，需要注意一下                 |
| 215. Kth Largest Element in an Array         | medium     | Heap、Quick Select      | YES  | 这里的 Quick Select 算法需要熟练一下           |

### Queue
| Problem                              | Difficulty | Method               | Redo | Comment                                                      |
| :----------------------------------- | :--------- | :------------------- | :--- | :----------------------------------------------------------- |
| 225. Implement Stack using Queues    | easy       | Queue                | No   | 简单题，不需要多看                                           |
| 346. Moving Average from Data Stream | easy       | Queue                | No   | 简单题，不需要多看                                           |
| 281. ZigZag Iterator                 | medium     | Queue                | YES  | 注意这里的队列的使用，还是挺妙的                             |
| 1429. First Unique Number            | unknown    | Queue、 Hash Table   | No   |                                                              |
| 362. Design Hit Counter              | unknown    | Queue                | No   |                                                              |
| 264. Ugly Number II                  | medium     | Queue                | YES  | 第一次做的时候做了一个错误解法，还是要注意丑数是怎么推出来的 |
| 1642.Furthest Building you can Reach | medium     | Heap、Priority Queue | YES  | 这里的解法很巧妙，需要注意一下                               |

### Hash Table
| Problem                                                     | Difficulty | Method                         | Redo | Comment                                                                                                                                  |
| :---------------------------------------------------------- | :--------- | :----------------------------- | :--- | :--------------------------------------------------------------------------------------------------------------------------------------- |
| 1. Two Sum                                                  | easy       | Hash Table、Sort、Two Pointer  | No   |                                                                                                                                          |
| 146. LRU Cache                                              | medium     | Double Linked List、Hash Table | YES  | 在双向链表中添加 dummyHead 和 dummyTail 能够避免很多条件判断                                                                             |
| 128. Longest Consecutive Sequence                           | medium     | Hash Table                     | YES  | 注意考虑为什么这里可以把遍历过的元素删除掉                                                                                               |
| 380. Insert Delete GetRandom O(1)                           | medium     | Hash Table、Array              | YES  | 注意这里删除元素的 tricky 部分                                                                                                           |
| 1461. Check  If a String Contans All Binary Codes of Size K | medium     | Hash Table、String             | YES  | 这里首先要注意怎么一步一步从暴力算法开始进行优化，然后逐渐过渡到使用 map 来做记忆化的；另外还需要注意 golang 的格式化字符串 padding 方法 |
| 349. Intersection of Two Arrays                             | easy       | Hash Table                     | No   |                                                                                                                                          |
| 350. Intersection of Two Arrays II                          | easy       | Hash Table                     | No   |                                                                                                                                          |
| 1086. High Five                                             | easy       | Hash Table、Sort               | No   |                                                                                                                                          |
| 692. Top K Frequent Words                                   | medium     | Sort、 Hash Table              | No   |                                                                                                                                          |
| 895. Max Frequency Stack                                    | hard       | Hash Table、Stack              | YES  | 这里的解法很巧妙，需要注意一下                                                                                                           |
| 745. Prefix and Suffix Search                               | hard       | Hash Table、String             | YES  | 有时间暴力解法也是没问题的                                                                                                               |
| 1268. Search Suggestions System                             | medium     | Hash Table、String、Trie Tree  | YES  | 这里的剪枝算法需要注意一下                                                                                                               |
| 409. Longest Palindrome                                     | easy       | Hash Table、String             | YES  | 这里的解法，要注意一下                                                                                                                   |
| 820. Short Encoding of Words                                | medium     | Hash Table、String             | YES  | 这里的思路比较巧妙                                                                                                                       |
| 454. 4Sum II                                                | medium     | Hash Table、Divide             | YES  | 这里降低时间复杂度的思想比较巧妙，值得学习                                                                                               |
| 277. Find the Celebrity | medium | Hash Table、Graph | YES | 这里的图论思想很值得思考 |

### Graph
| Problem                                                     | Difficulty | Method                         | Redo | Comment                                                                                                                                  |
| :--------------- | :--------- | :-------- | :--- | :------ |
| 277. Find the Celebrity | medium | Hash Table、Graph | YES | 这里的图论思想很值得思考 |
### Linked List
| Problem                               | Difficulty | Method                               | Redo | Comment                                                      |
| :------------------------------------ | :--------- | :----------------------------------- | :--- | :----------------------------------------------------------- |
| 160. Intersection of Two Linked Lists | easy       | Two Pointer、 Hash Table             | YES  | 注意可能会出现无交点的情况                                   |
| 206. Reverse Linked List              | easy       | Stack、Recursive                     | YES  | 注意这里的递归法中，如何反转节点                             |
| 876. Middle of the Linked List        | easy       | Two Pointer                          | YES  | 注意这里是如何判断中间节点的                                 |
| 141. Linked List Cycle                | easy       | Hash Table、 Two Pointer             | YES  | 注意这里的快慢指针解法                                       |
| 142. Linked List Cycle II             | medium     | Linked List、Hash Table、Two Pointer | YES  | 注意这里的快慢指针解法                                       |
| 92. Reverse Linked List II            | medium     | Linked List、 Stack                  | YES  | 注意这里是如何使用指针来进行翻转的                           |
| 328. Odd Even Linked List             | medium     | Linked List、 Two Pointer            | YES  | 注意这里是如何连接节点的                                     |
| 146. LRU Cache                        | medium     | Double Linked List、Hash Table       | YES  | 在双向链表中添加 dummyHead 和 dummyTail 能够避免很多条件判断 |
### String
| Problem                               | Difficulty | Method             | Redo | Comment                                               |
| :------------------------------------ | :--------- | :----------------- | :--- | :---------------------------------------------------- |
| 28. Implement strStr                  | easy       | None               | No   |                                                       |
| 44. Wildcard Matching                 | hard       | Recusive           | No   |                                                       |
| 318. Maximum Product of Word Lengths  | medium     | Hash Table         | Yes  | 这里的解法很精妙，使用了一个 int 来做记录，提升了性能 |
| 1332. Remove Palindromic Subsequences | easy       | String、Palindrome | YES  | 这里是一个脑筋急转弯                                  |
| 1689. Partitioning Into Minimum Number Of Deci-Binary Numbers | medium | String、 Math | YES | 这里的解法很巧妙，如果没有答案里的特殊方法，肯定是很难解出来的 |


### Array
| Problem                                 | Difficulty | Method             | Redo  | Comment                                                                  |
| :-------------------------------------- | :--------- | :----------------- | :---- | :----------------------------------------------------------------------- |
| 36. Valid Sudoku                        | medium     | None               | YES   | 注意这里如何创建三维数组，以及坐标的转换                                 |
| 41. First Missing Positive              | Hard       | HashTable, Array   | YES   | 注意 while 循环是如何交换元素的，以及 while 循环为什么不能用 if 语句代替 |
| 54. Spiral Matrix                       | medium     |                    | Array | YES                                                                      | 注意这里的终止条件判断 |
| 73. Set Matrix Zeroes                   | medium     | 记忆化方法         | YES   | 注意这里是如何复用原数组以及采用的记忆化方法                             |
| 1480. Running Sum of 1d Array           | easy       | Array              | No    | 简单题，不用再多看了                                                     |
| 867. Transpose Matrix                   | easy       | Array、Matrix      | No    | 简单题，不需要多看了                                                     |
| 167. Two Sum II - Input Array Is Sorted | medium     | Array、Two Pointer | No    | 简单题，无需多看                                                         |

## Search

### Bianry Search

| Problem                                                     | Difficulty | Method                               | Redo | Comment                                                                                                        |
| :---------------------------------------------------------- | :--------- | :----------------------------------- | :--- | :------------------------------------------------------------------------------------------------------------- |
| 33. Search in Rotated Sorted Array                          | medium     | binary search                        | YES  |                                                                                                                |
| 34. Find First and Last Position of Element in Sorted Array | medium     | binary search                        | YES  | 注意这里的第二种解法，也就是二分查找中如何确定边界的问题                                                       |
| 162. Find Peak Element                                      | medium     | Binary search、Monotone Stack、Array | YES  | 这里二分的解法还是太 tricky 了                                                                                 |
| 278. First Bad Version                                      | easy       | Binary Search                        | No   | 简单题，直接过                                                                                                 |
| 74. Search a 2D Matrix                                      | medium     | Binary Search、Matrix                | YES  | 这题的二分解法很巧妙                                                                                           |
| 240. Search a 2D Matrix II                                  | medium     | Binary Search、Matrix                | YES  | 这个题的解法跟 74 题完全一致                                                                                   |
| 540. Single Element in a Sorted Array                       | medium     | Binary Search、Array                 | YES  | 这里的隐式二分真的太巧妙了，需要注意一下                                                                       |
| 69. Sqrt(x)                                                 | easy       | Binary Search                        | YES  | 注意这里的边界是如何区分的；开头去掉一些 corner case的话，会让主逻辑更清晰一些                                 |
| 367. Valid Perfect Square                                   | easy       | Binary Search                        | YES  | 注意这里的边界是如何区分的，以及这里与 69 题的相似性                                                           |
| 528. Random Pick with Weight                                | medium     | Binary Search                        | YES  | 这里首先需要注意的是随机数的选择方法；其次需要注意的是，sum 数组的下标需要与原数组下标匹配（针对本题答案而言） |







| Problem                                 | Difficulty | Method           | Redo |
| :-------------------------------------- | :--------- | :--------------- | :--- |
| 21. Merge Two Sorted List               | medium     | Merge Sort       | No   |
| 22. Generate Parentheses                | medium     | Prune, Recursive | YES  |
| 26. Remove Duplicates from Sorted Array | easy       |                  | No   |

## Math

### Bit Manipulation

| Problem                                          | Difficulty | Method                                | Redo | Comment                    |
| :----------------------------------------------- | :--------- | :------------------------------------ | :--- | :------------------------- |
| 29. Divide Two Integers                          | medium     | Bit manipulation                      | YES  |                            |
| 191. Number of 1 Bits                            | easy       | Bit manipulation                      | No   | 这题太简单了，没什么好说的 |
| 1342. Number of Steps to Reduce a Number to Zero | easy       | Bit manipulation                      | No   | 太简单了，不用再看         |
| 1060. Missing Number in Sorted Array             | easy       | Bit manipulation、Math、Binary Search | YES  |                            |

### Math Computation
| Problem             | Difficulty | Method | Redo | Comment                                            |
| :------------------ | :--------- | :----- | :--- | :------------------------------------------------- |
| 268. Missing Number | easy       | Math   | No   | 简单题，直接略过                                   |
| 263. Ugly Number    | easy       | Math   | YES  | 这里还是需要稍微注意一下n 的范围是怎么一步步缩小的 |
## Dynamic Programming

| Problem                                  | Difficulty | Method         | Redo | Comment                                                                                                          |
| :--------------------------------------- | :--------- | :------------- | :--- | :--------------------------------------------------------------------------------------------------------------- |
| 42. Traping Rain Water                   | medium     | DP, Stack      | YES  |                                                                                                                  |
| 474. Ones and Zeroes                     | medium     | DP             | YES  | 这里的递推式还是想不太出来，注意遍历的方向                                                                       |
| 354. Russian Doll Envelopes              | hard       | DP             | YES  | 注意这里更最长递增子序列问题的相似性                                                                             |
| 303. Range Sum Query - Immutable         | easy       | DP、Prefix Sum | YES  | 注意这里前缀和的解法                                                                                             |
| 304. Range Sum Query 2D - Immutable      | medium     | DP、Prefix Sum | YES  | 注意这里的坐标变换是如何进行的                                                                                   |
| 120. Triangle                            | medium     | DP             | YES  | 注意这里 DP 方法中复用原数组的技巧                                                                               |
| 64. Minimum Path Sum                     | medium     | DP             | YES  | 这个题跟 120 题的解法几乎完全一致 ，注意这里使用到的套路                                                         |
| 583. Delete Operations for Two Strings   | medium     | DP             | YES  | 这题的 DP 解法跟 1143 题是完全一致的，都是求解最长公共子序列                                                     |
| 1048. Longest String Chain               | medium     | DP             | YES  |                                                                                                                  |
| 5. Longest Palindrome Substring          | medium     | DP             | YES  | 要注意这里的 DP 解法；另外，单个字符串的最长的公共串/序列问题，尝试看看能不能转换为两个字符串的公共子串/序列问题 |
| 718. Maximum Length of Repeated Subarray | medium     | DP             | YES  | 注意如何将这里的解法用到第5 题中                                                                                 |
| 1062. Longest Repeating Substring        | medium     | DP             | YES  | 这个题的递推式跟 718 题完全一致                                                                                  |
## Prefix Sum
| Problem                             | Difficulty | Method         | Redo | Comment                                |
| :---------------------------------- | :--------- | :------------- | :--- | :------------------------------------- |
| 303. Range Sum Query - Immutable    | easy       | DP、Prefix Sum | YES  | 注意这里前缀和的解法                   |
| 304. Range Sum Query 2D - Immutable | medium     | DP、Prefix Sum | YES  | 注意这里的坐标变换是如何进行的         |
| 560. Subarray Sum Equals K          | medium     | Prefix Sum     | No   | 主要会前缀和的技巧，那么这里就不是很难 |
| 307. Range Query Sum - Mutable      | medium     | Prefix Sum     | No   |                                        |
| 1423. Maximum Points You Can Obrain from Cards | medium | Prefix Sum、DP | YES | 注意这个题怎么转换为前缀和来做，并且要注意 res 的初始值 |
## Sort
| Problem                         | Difficulty | Method                   | Redo | Comment                                                                          |
| :------------------------------ | :--------- | :----------------------- | :--- | :------------------------------------------------------------------------------- |
| 21. Merge Two Sorted List       | medium     | Merge Sort               | No   |                                                                                  |
| 148. Sort List                  | medium     | Merge Sort               | YES  | 注意这里的自底向上方法；以及自顶向下方法中，要注意断开 middle 和它之前节点的连接 |
| 56. Merge Intervals             | medium     | Sort                     | YES  |                                                                                  |
| 179. Largest Number             | Medium     | Sort                     | YES  | 这里的做法需要十分注意，一般可能注意不到                                         |
| 75. Sort Colors                 | medium     | Bubble Sort、 Count Sort | No   |                                                                                  |
| 215. Kth Largest Element        | medium     | Heap、Quick Select       | YES  | 这里要非常注意快速选择算法                                                       |
| 4. Median of Two Sorted Arrays  | hard       | Heap                     | YES  | 这个题太难了                                                                     |
| 973. K Closest Points to Origin | medium     | Sort、 Heap              | No   | 要看到本质问题，然后解决本质问题。                                               |
| 88. Merge Sorted Array          | easy       | Sort、 Merge Sort        | No   |                                                                                  |
| 692. Top K Frequent Words       | medium     | Sort、 Hash Table        | No   |                                                                                  |
| 1647. Minimum Deletions to Make Character Frequencies Unique | medium | Greedy algorithm、Sort | YES | 这里的贪心会比价隐晦一些 |


## Two Pointers
| Problem                        | Difficulty | Method                  | Redo | Comment                                                                   |
| :----------------------------- | :--------- | :---------------------- | :--- | :------------------------------------------------------------------------ |
| 876. Middle of the Linked List | easy       | Two Pointer             | No   | 注意这里 for 循环结束遍历的终止条件                                       |
| 27. Remove Elements            | easy       | Two Pointer             | No   |                                                                           |
| 125. Valid Palindrome          | easy       | Two Pointer             | No   | 简单题，直接跳过即可                                                      |
| 1. Two Sum                     | easy       | Two Pointer、Hash Table | No   | 注意一下这里的哈希表的解法                                                |
| 167. Two Sum II                | medium     | Two Pointer             | No   |                                                                           |
| 15. 3Sum                       | medium     | Two Pointer、Hash Table | Yes  | 注意这里是如何进行减枝的                                                  |
| 16. 3Sum Closest               | medium     | Two Pointer             | Yes  | 注意这里如何进行减枝                                                      |
| 18. 4Sum                       | medium     | Two Pointer             | YES  | 注意这里是如何去重的，在最后两个 for 循环中对重复元素去重的技巧要注意一下 |
| 11. Container with Most Water | medium | Yes | 这里移动指针的方式有一点贪心的味道 |
| 283. Moving Zeroes | easy | Two Pointer | No | |
| 26. Remove Duplicates from Sorted Array | easy | Two Pointer | No | |
| 395. Longest Substring with At Least K Repeating Characters | medium | Two Pointer | YES | 这个题暂时跳过一下 |
| 424. Longest Repeating Character Replacement | medium | Two Pointer、 Sliding Window | YES | 这里的滑动窗口/双指针的解法是非常巧妙的，值得好好学习 ，尤其是怎么来做类推，来把问题细化 |
## Backtrace
| Problem         | Difficulty | Method    | Redo | Comment                                                                                               |
| :-------------- | :--------- | :-------- | :--- | :---------------------------------------------------------------------------------------------------- |
| 51. N-Queens    | hard       | Backtrace | YES  | 惭愧，这个经典的问题自己已经忘了它的解法了                                                            |
| 52. N-Queues II | hard       | Backtrace | YES  | 这个题跟 51 题的解法是完全一致的。51 题中因为是从上到下对每一行进行的回溯，所以已经做了一个去重的操作 |

## Sliding Window
| Problem                                          | Difficulty | Method                                | Redo | Comment                                                                                                         |
| :----------------------------------------------- | :--------- | :------------------------------------ | :--- | :-------------------------------------------------------------------------------------------------------------- |
| 3. Longest Substring Without Repeating Character | medium     | sliding window、Hash Table、String    | YES  | 注意这里滑动窗口中左边界的确定，以及左边界 left 的初值和移动方向                                                |
| 1658. Minimum Operations to Reduce X to Zero     | medium     | Sliding window、Prefix Sum、Backtrace | YES  | 首先要注意这里是怎么逐步对前缀和方法进行优化的；其次要注意这里的滑动窗口解法，也算是滑动窗口中的一中 pattern 了 |
| 1695. Maximum Erasure Value                      | medium     | Sliding Window、 Hash Table           | YES  | 这里其实跟第 3 题是一样的                                                                                       |
| 643. Maximum Average Subarray I                  | easy       | Sliding Window                        | YES  | 这题没什么难度，但是要注意两个 if 语句的顺序，不能颠倒                                                          |
| 424. Longest Repeating Character Replacement | medium | Two Pointer、 Sliding Window | YES | 这里的滑动窗口/双指针的解法是非常巧妙的，值得好好学习 ，尤其是怎么来做类推，来把问题细化 |

## Greedy Algorithms
| Problem                  | Difficulty | Method           | Redo | Comment                                        |
| :----------------------- | :--------- | :--------------- | :--- | :--------------------------------------------- |
| 968. Binary Tree Cameras | hard       | greedy algorithm | YES  | 这里其实就有点像是一个递归再配合上状态机的转移 |
| 630. Course Schedule III | hard       | Greedy algorithm | YES  | 这里的关键其实还是在于最大堆与贪心算法的结合   |
| 665. Non-decreasing Array | medium | Greedy algorithm | YES | 这里的贪心有些隐晦，不易发现 |
| 1647. Minimum Deletions to Make Character Frequencies Unique | medium | Greedy algorithm、Sort | YES | 这里的贪心会比价隐晦一些 |

## Divide and Conquer
| Problem      | Difficulty | Method             | Redo | Comment                                    |
| :----------- | :--------- | :----------------- | :--- | :----------------------------------------- |
| 454. 4Sum II | medium     | Hash Table、Divide | YES  | 这里降低时间复杂度的思想比较巧妙，值得学习 |
