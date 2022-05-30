---
weight: 5
bookCollapseSection: true
---

这里是一些 LeetCode 的刷题题解。

## Data structure

### Stack
| Problem                | Difficulty | Method    | Redo | Comment |
| 155. Min Stack | easy | Stack、Monotone Stack | No | |
| 150. Evalute Reverse Polish Notation | medium | Stack | No | 简单题，无需多看  |
| 224. Basic Calculator | hard | Stack | YES | 题目为难题，暂时跳过。注意这里是如何将中缀表达式转换为后缀表达式的 | 
| 227. Basic Calculator | medium | Stack | YES | 稍微有点耗时，先跳过 |
| 20. Valid Parentheses | easy | Stack | No | |
| 1472. Design Browser History | medium | Stack | No | |
| 1047. Remove All Adjacent Duplicates in String | easy | Stack | No | |
| 1249. Minimum Remove to Make Valid Parentheses | medium | Stack | Yes | |
| 735. Asteroid Collision | medium | Stack | Yes | |
### Queue
| Problem                | Difficulty | Method    | Redo | Comment |
| :--------------------- | :--------- | :-------- | :--- | :------ |
| 225. Implement Stack using Queues | easy | Queue | No | 简单题，不需要多看|
| 346. Moving Average from Data Stream | easy | Queue | No | 简单题，不需要多看 | 
| 281. ZigZag Iterator | medium | Queue | YES | 注意这里的队列的使用，还是挺妙的 | 
| 1429. First Unique Number | unknown | Queue、 Hash Table | No | |
| 362. Design Hit Counter | unknown | Queue | No | |

### Hash Table
| Problem                | Difficulty | Method    | Redo | Comment |
| :--------------------- | :--------- | :-------- | :--- | :------ |
| 1. Two Sum | easy | Hash Table、Sort、Two Pointer | No | |
| 146. LRU Cache | medium | Double Linked List、Hash Table | YES | 在双向链表中添加 dummyHead 和 dummyTail 能够避免很多条件判断 |

### Linked List
| Problem                | Difficulty | Method    | Redo | Comment |
| :--------------------- | :--------- | :-------- | :--- | :------ |
| 160. Intersection of Two Linked Lists | easy | Two Pointer、 Hash Table | YES| 注意可能会出现无交点的情况|
| 206. Reverse Linked List | easy | Stack、Recursive | YES | 注意这里的递归法中，如何反转节点 |
| 876. Middle of the Linked List | easy | Two Pointer | YES | 注意这里是如何判断中间节点的 | 
| 141. Linked List Cycle | easy | Hash Table、 Two Pointer | YES | 注意这里的快慢指针解法 |
| 142. Linked List Cycle II | medium | Linked List、Hash Table、Two Pointer | YES | 注意这里的快慢指针解法 |
| 92. Reverse Linked List II | medium | Linked List、 Stack | YES | 注意这里是如何使用指针来进行翻转的 |
| 328. Odd Even Linked List | medium | Linked List、 Two Pointer | YES | 注意这里是如何连接节点的 | 
| 146. LRU Cache | medium | Double Linked List、Hash Table | YES | 在双向链表中添加 dummyHead 和 dummyTail 能够避免很多条件判断 |
### String
| Problem               | Difficulty | Method   | Redo | Comment |
| :-------------------- | :--------- | :------- | :--- | :----|
| 28. Implement strStr  | easy       | None     | No   | |
| 44. Wildcard Matching | hard       | Recusive | No   | |
| 318. Maximum Product of Word Lengths | medium | Hash Table | Yes | 这里的解法很精妙，使用了一个 int 来做记录，提升了性能| 


### Array
| Problem                    | Difficulty | Method           | Redo | Comment                                                                  |
| :------------------------- | :--------- | :--------------- | :--- | :----------------------------------------------------------------------- |
| 36. Valid Sudoku           | medium     | None             | YES  | 注意这里如何创建三维数组，以及坐标的转换                                 |
| 41. First Missing Positive | Hard       | HashTable, Array | YES  | 注意 while 循环是如何交换元素的，以及 while 循环为什么不能用 if 语句代替 |
| 54. Spiral Matrix | medium | | Array | YES | 注意这里的终止条件判断 |

## Search

### Bianry Search

| Problem                                                     | Difficulty | Method        | Redo |
| :---------------------------------------------------------- | :--------- | :------------ | :--- |
| 33. Search in Rotated Sorted Array                          | medium     | binary search | YES  |
| 34. Find First and Last Position of Element in Sorted Array | medium     | binary search | No   |







| Problem                                 | Difficulty | Method           | Redo |
| :-------------------------------------- | :--------- | :--------------- | :--- |
| 21. Merge Two Sorted List               | medium     | Merge Sort       | No   |
| 22. Generate Parentheses                | medium     | Prune, Recursive | YES  |
| 26. Remove Duplicates from Sorted Array | easy       |                  | No   |

## Math

### Bit Manipulation

| Problem                 | Difficulty | Method           | Redo | Comment |
| :---------------------- | :--------- | :--------------- | :--- | :----|
| 29. Divide Two Integers | medium     | Bit manipulation | YES  | |
| 191. Number of 1 Bits | easy | Bit manipulation | No | 这题太简单了，没什么好说的|
| 1342. Number of Steps to Reduce a Number to Zero | easy | Bit manipulation | No | 太简单了，不用再看|

### Math Computation
| Problem                 | Difficulty | Method           | Redo | Comment |
| :---------------------- | :--------- | :--------------- | :--- | :----|
| 268. Missing Number | easy | Math | No | 简单题，直接略过|	
## Dynamic Programming

| Problem                | Difficulty | Method    | Redo | Comment |
| :--------------------- | :--------- | :-------- | :--- | :------ |
| 42. Traping Rain Water | medium     | DP, Stack | YES  |         |
| 474. Ones and Zeroes | medium | DP | YES | 这里的递推式还是想不太出来，注意遍历的方向 |
| 354. Russian Doll Envelopes | hard | DP | YES | 注意这里更最长递增子序列问题的相似性|

## Sort
| Problem                | Difficulty | Method    | Redo | Comment |
| :--------------------- | :--------- | :-------- | :--- | :------ |
| 21. Merge Two Sorted List               | medium     | Merge Sort       | No   | |
| 148. Sort List | medium | Merge Sort | YES | 注意这里的自底向上方法；以及自顶向下方法中，要注意断开 middle 和它之前节点的连接|
| 56. Merge Intervals | medium | Sort | YES | |
| 179. Largest Number | Medium | Sort | YES | 这里的做法需要十分注意，一般可能注意不到|
| 75. Sort Colors | medium | Bubble Sort、 Count Sort | No | |
| 215. Kth Largest Element | medium | Heap、Quick Select | YES | 这里要非常注意快速选择算法 |
| 4. Median of Two Sorted Arrays | hard | Heap |  YES | 这个题太难了 |

## Two Pointers
| Problem                | Difficulty | Method    | Redo | Comment |
| :--------------------- | :--------- | :-------- | :--- | :------ |
| 876. Middle of the Linked List | easy | Two Pointer | No | 注意这里 for 循环结束遍历的终止条件 |
| 27. Remove Elements | easy | Two Pointer | No | |

