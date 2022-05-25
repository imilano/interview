---
title: 1429. First Unique Number
weight: 10
tags: [
	"Queue",
	"Two Pointer"
]
---
## Description
> You have a queue of integers, you need to retrieve the first unique integer in the queue.
> 
> Implement the FirstUnique class:
> 
> 	- FirstUnique(int[] nums) Initializes the object with the numbers in the queue.
>	- int showFirstUnique() returns the value of the first unique integer of the queue, and returns -1 if there is no such integer.
> 	- void add(int value) insert value to the queue.

## Solutions
可以这么做： 维护一个哈希表，哈希表的 key 为数字，value 为一个Pair， Pair 中存储 key 在数组中的下标以及该数字的出现次数。然后再创建一个 queue， 将所有出现次数为 1 的数字按照在队列中的出现顺序入队。 调用 showFirstUnique 的时候， 从队列中取出元素，检查在哈希表中该元素的出现次数是否为 1，如果不是，则持续出对队，直到找到第一个出现次数为 1 的元素，然后将该元素输出。add 的时候，将该元素在哈希表中的对应次数加 1，并检查队头元素是否为该元素，如果是， 则将该元素出队即可。

TODO

