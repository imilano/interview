---
title: 0752. Open the Lock
weight: 10
tags: [
	"Hash Table",
	"BFS"
]
---
## Description
> You have a lock in front of you with 4 circular wheels. Each wheel has 10 slots: '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'. The wheels can rotate freely and wrap around: for example we can turn '9' to be '0', or '0' to be '9'. Each move consists of turning one wheel one slot.
> 
> The lock initially starts at '0000', a string representing the state of the 4 wheels.
> 
> You are given a list of deadends dead ends, meaning if the lock displays any of these codes, the wheels of the lock will stop turning and you will be unable to open it.
> 
> Given a target representing the value of the wheels that will unlock the lock, return the minimum total number of turns required to open the lock, or -1 if it is impossible.


## Solutions
### BFS
这里既然是要求最小步数，那么用 BFS 会比较好，类似于层序遍历。与以前的迷宫遍历不同的是，这里是需要对字符串的每个字符来做遍历，遍历可以往前也可以往后。这里用-1 和 1 来表示往后和往前。此外还需要一个 hash table 来记录一个字符是否被遍历过，如果该字符没有被遍历过，并且也不是出现在 deadends 中的字符串，那么就可以将该字符加入到下一轮遍历的队列中。

```go
func openLock(deadends []string, target string) int {
	initialStr := "0000"
	// 如果 target 就是初始元素，则直接返回 0
	if target == initialStr {
		return 0
	}

	// 对每个 deadend 去重
	dict := make(map[string]bool)
	for _, str := range deadends {
		dict[str] = true
	}

	// 如果 初始字符串就在 deadends 中，则无法得解
	if _, ok := dict[initialStr]; ok {
		return -1
	}

	var res int
	// 将初始字符串入队
	var queue []string = []string{initialStr}
	// 记录一个字符串有没有被遍历过，如果没被遍历过，则需要加进队列中
	visited := map[string]bool{initialStr: true}
	for len(queue) != 0 {
		res++
		size := len(queue)
		// 对于当前一层的每个字符串
		for i := 0; i < size; i++ {
			strlen := len(queue[i])
			// 对于该字符串的每个字符
			for j := 0; j < strlen; j++ {
				// + 1 表示向前，-1 表示向后
				for k := -1; k <= 1; k++ {
					// 跳过 0
					if k == 0 {
						continue
					}

					// 为覆盖'9' 加 1 之后变为'0' 和 '0' 减 1 只有变为'9'的情况，
					// 统一给每个字符加上 10 来进行运算，最后让其对 10 取余
					rs := []rune(queue[i])
                    rs[j] = rune((int(rs[j]-'0') + 10 + k) % 10 + int('0'))
					str := string(rs)
					// 如果当前字符等于 target 了，则说明已经找到了最小的步数，返回结果即可
					if str == target {
						return res
					}

					// 如果 str 没有被访问过，并且没有出现在 deadends 中，则将其加入到队列中，进行下一轮遍历
					_, vis := visited[str]
					_, ok := dict[str]
					if !ok && !vis {
						queue = append(queue, str)
					}

					// 标记当前字符已经被遍历过
					visited[str] = true
				}
			}
		}

		// 缩小队列
		queue = queue[size:]
	}


	// 返回结果
	return -1
}
```
