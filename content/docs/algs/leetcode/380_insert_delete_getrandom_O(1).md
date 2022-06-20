---
title: 0380. Insert Delete GetRandom O(1)
weight: 10
tags: [
	"Hash Table",
	"Array"
]
---

## Description
> Implement the RandomizedSet class:
> 
> - `RandomizedSet()` Initializes the RandomizedSet object.
> - `bool insert(int val)` Inserts an item val into the set if not present. Returns true if the item was not present, false otherwise.
> - `bool remove(int val)` Removes an item val from the set if present. Returns true if the item was present, false otherwise.
> - `int getRandom()` Returns a random element from the current set of elements (it's guaranteed that at least one element exists when this method is called). Each element must have the same probability of being returned.
> You must implement the functions of the class such that each function works in average O(1) time complexity.
## Solutions
### Hash Table & Array
这里使用一个 Hash Table 和一个数组，Hash Table 存储的是值和该值在数组中的下标。这里稍微有点 tricky 的是，删除的时候，不是直接在 Hash Table 中删除该元素，而是在数组中将该元素和尾元素调换位置，然后删除尾元素。 这样所有操作都能达到 {{< katex >}} \Omicron(1) {{< /katex >}} 的时间复杂度。
```go
import "math/rand"

type RandomizedSet struct {
    dict map[int]int
    nums []int
}


func Constructor() RandomizedSet {
    return RandomizedSet{dict: make(map[int]int), nums: []int{}}
}


func (this *RandomizedSet) Insert(val int) bool {
    var present bool
    if _, ok := (*this).dict[val]; !ok {
        present = true
        size := len((*this).nums)
        (*this).nums = append((*this).nums, val)
        (*this).dict[val] = size
    }
    
    return present
}


func (this *RandomizedSet) Remove(val int) bool {
    var present bool
    if idx, ok := (*this).dict[val]; ok {
        present = true
        // 将待删除元素和最后一个元素交换，这样的话只需要删除最后一个元素即可，Remove 和 getRandom 都能达到 O(1) 的时间复杂度
        size := len((*this).dict)
        (*this).dict[(*this).nums[size-1]] = idx
        (*this).nums[size-1], (*this).nums[idx] = (*this).nums[idx], (*this).nums[size-1]
        (*this).nums = (*this).nums[:size-1]
        delete((*this).dict, val)
        
        // 下面的删除办法起始是不对的，下面的删除只有当待删除元素是数组中最后一个元素才正确；
        // 否则的话，在数组中删除元素之后，dict 中存储的下标是错误的下标。因为删除中间元素会导致该元素之后的元素的下标前移 1
        // idx := (*this).dict[val]
        // (*this).nums = append((*this).nums[:idx], (*this).nums[idx+1:]...)
        // delete((*this).dict, val)
    }
    
    return present
}


func (this *RandomizedSet) GetRandom() int {
    size := len((*this).nums)
    target := rand.Intn(size)
    return (*this).nums[target]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
 ```

### Hash Table
这里是直接使用 Hash Table 的解法，虽然可以通过 OJ，但是 getRandom 的时间复杂度并不是 {{< katex >}} \Omicron(1) {{< /katex >}}。
```go
import "math/rand"

type RandomizedSet struct {
    nums map[int]bool
}


func Constructor() RandomizedSet {
    return RandomizedSet{make(map[int]bool)}
}


func (this *RandomizedSet) Insert(val int) bool {
    var present bool
    if _, ok := (*this).nums[val]; !ok {
        present = true
        (*this).nums[val] = true
    }
    
    return present
}


func (this *RandomizedSet) Remove(val int) bool {
    var present bool
    if _, ok := (*this).nums[val]; ok {
        present = true
        delete((*this).nums, val)
    }
    
    return present
}


func (this *RandomizedSet) GetRandom() int {
    size := len((*this).nums)
    target := rand.Intn(size)
    var cur,res int
    for key,_ := range (*this).nums {
        if cur == target {
            res = key
            break
        }
        
        cur++
    }
    
    return res
}
```

### Double LinkedList
下面是使用双向链表的解法，也可以通过 OJ，但是 getRandom 函数的时间复杂度也不符合要求。
```go
import "math/rand"

type Node struct {
    Val int
    Next *Node
    Pre *Node
}

type RandomizedSet struct {
    dict map[int]*Node
    head *Node
    tail *Node
}

func(this *RandomizedSet) AddToHead(node *Node) {
    node.Next = this.head.Next
    this.head.Next.Pre = node
    this.head.Next = node
    node.Pre = this.head
}

func (this *RandomizedSet) RemoveNode(node *Node) {
    node.Pre.Next = node.Next
    node.Next.Pre = node.Pre
}

func Constructor() RandomizedSet {
    head,tail := new(Node), new(Node)
    head.Next, tail.Pre = tail, head
    return RandomizedSet{dict: make(map[int]*Node), head: head, tail: tail}
}


func (this *RandomizedSet) Insert(val int) bool {
    var present bool
    if _, ok := (*this).dict[val]; !ok {
        present = true
        node := new(Node)
        node.Val = val
        this.AddToHead(node)
        (*this).dict[val] = node
    }
    
    return present
}


func (this *RandomizedSet) Remove(val int) bool {
    var present bool
    if node, ok := (*this).dict[val]; ok {
        present = true
        delete((*this).dict, val)
        this.RemoveNode(node)
    }
    
    return present
}


func (this *RandomizedSet) GetRandom() int {
    size := len((*this).dict)
    target := rand.Intn(size) + 1
    cur := this.head
    var cnt,res int
    for cnt != target {
        cur = cur.Next
        cnt++
        if cnt == target {
            res = cur.Val
            break
        }
    }
    return res
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
 ```
