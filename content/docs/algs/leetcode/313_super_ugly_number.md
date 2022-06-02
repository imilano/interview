---
title: 0313. Super Ugly Number
weight: 10
tags: [
	"Heap",
	"Min Heap",
	"K Pointer"
]
---

## Description
> A super ugly number is a positive integer whose prime factors are in the array primes.
> 
> Given an integer n and an array of integers primes, return the nth super ugly number.
> 
> The nth super ugly number is guaranteed to fit in a 32-bit signed integer.


## Solutions
### K Pointer
这个题可以用第 264 题的解法来解。区别是，第 264 题中我们只有三个因子，但是这里的因子是不固定的，那么我们只需要使用一个 map 来充当 264 题中的多指针即可。

关于为什么使用多指针就可以，这里参见一个 LeetCode 上的[回答](https://leetcode.cn/problems/chou-shu-lcof/solution/chou-shu-by-leetcode-solution-0e5i/1070112):
> 这道题一开始死活不明白三指针到底是怎么用的。后来突然就想明白了：
> 
> 例如 n = 10， primes = [2, 3, 5]。 打印出丑数列表：1, 2, 3, 4, 5, 6, 8, 9, 10, 12
> 首先一定要知道，后面的丑数一定由前面的丑数乘以2，或者乘以3，或者乘以5得来。例如，8,9,10,12一定是1, 2, 3, 4, 5, 6乘以2,3,5三个质数中的某一个得到。
> 
> 这样的话我们的解题思路就是：从第一个丑数开始，一个个数丑数，并确保数出来的丑数是递增的，直到数到第n个丑数，得到答案。那么问题就是如何递增地数丑数？
> 
> 观察上面的例子，假如我们用1, 2, 3, 4, 5, 6去形成后面的丑数，我们可以将1, 2, 3, 4, 5, 6分别乘以2, 3, 5，这样得到一共6*3=18个新丑数。也就是说1, 2, 3, 4, 5, 6中的每一个丑数都有一次机会与2相乘，一次机会与3相乘，一次机会与5相乘（一共有18次机会形成18个新丑数），来得到更大的一个丑数。
> 
> 这样就可以用三个指针，
> 
> pointer2, 指向1, 2, 3, 4, 5, 6中，还没使用乘2机会的丑数的位置。该指针的前一位已经使用完了乘以2的机会。
> pointer3, 指向1, 2, 3, 4, 5, 6中，还没使用乘3机会的丑数的位置。该指针的前一位已经使用完了乘以3的机会。
> pointer5, 指向1, 2, 3, 4, 5, 6中，还没使用乘5机会的丑数的位置。该指针的前一位已经使用完了乘以5的机会。
> 下一次寻找丑数时，则对这三个位置分别尝试使用一次乘2机会，乘3机会，乘5机会，看看哪个最小，最小的那个就是下一个丑数。最后，那个得到下一个丑数的指针位置加一，因为它对应的那次乘法使用完了。
> 
> 这里需要注意下去重的问题，如果某次寻找丑数，找到了下一个丑数10，则pointer2和pointer5都需要加一，因为5乘2等于10， 5乘2也等于10，这样可以确保10只被数一次。

所以其实核心就是：
- 每个丑数都是前面的丑数乘以因子得到的
- 每个丑数都有一次乘以所有因子的机会，一旦一个丑数乘以一个因子并且乘积被我们使用了，那么这个丑数与这个因子相乘的机会就用完了。换句话说，对应的这个因子的指针就应该后移。
```go
func nthSuperUglyNumber(n int, primes []int) int {
    res := make([]int, n)
    res[0] = 1
    
	// 记录每个指针的初始位置
    dict := make(map[int]int, len(primes))
    for _, prime := range primes {
        dict[prime] = 0
    }
    
    var idx int
    for i := 1; i < n; i++ {
        r := selectMin(idx, res, dict)
        idx++
        res[idx] = r
		// 将相应位置的下标指针后移
        for prime, index := range dict {
			// 注意这里是 res[index] * prime
            if r  == res[index] * prime {
                dict[prime]++
            }
        }
    }
    
    return res[n-1]
}


// 选出一个最小的
func selectMin(idx int, res []int, dict map[int]int) int {
	// 这里注意 r 的初始化值，不能初始化为 0
    r := math.MaxInt
    for prime, index  := range dict {
        r = min(r, prime*res[index])
    }
    
    return r
}


func min(a,b int) int {
    if a < b {
        return a
    }
    
    return b
}
```
