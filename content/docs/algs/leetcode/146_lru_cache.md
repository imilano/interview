---
title: 0146. LRU Cache
weight: 10
tags: [
    "Hash Table",
    "LinkedList"
]
---
## Description
> Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.
> 
> Implement the LRUCache class:
> 
> LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
> int get(int key) Return the value of the key if the key exists, otherwise return -1.
> void put(int key, int value) Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache. If the number of keys exceeds the capacity from this operation, evict the least recently used key.
> The functions get and put must each run in O(1) average time complexity.

## Solutions

### Hash Table & Array
下面的解法严格来说并不符合题目要求的时间复杂度。
```go

import "time"

// 一个Cell
type Cell struct {
	Key int
	Value int
	RecentUse int64
}

// 一个LRU cache队列，通过map中的key快速映射到Data中相应节点的index
type LRUCache struct {
	Cap int
	Data []*Cell
	Map map[int]int
}


func Constructor(capacity int) LRUCache {
	return LRUCache{
		Cap: capacity,
		Data: make([]*Cell,0),
		Map:	make(map[int]int),
	}
}


//func (this *LRUCache) Get(key int) int {
//	value := -1
//
//	if v,ok := this.Map[key]; ok && v != -1 {
//		value = this.Data[v-1].Value
//		this.Data[v-1].RecentUse = time.Now().UnixNano()
//	}
//
//	return value
//}

func (this *LRUCache) Get(key int) int {
	value := -1

	if v,ok := this.Map[key]; ok && v != -1  {
		value = this.Data[v].Value  // if Data[v], will out of index
		this.Data[v].RecentUse = time.Now().UnixNano()
	}

	return value
}


func (this *LRUCache) Put(key int, value int)  {
	// If already exist, modify value and time
	index,ok := this.Map[key]
	
	// If the node already exist
	if ok && index != -1 {
		this.Data[index].Value = value
		this.Data[index].RecentUse = time.Now().UnixNano()
	} else {
		// If the value  doesn't exist, then insert new node or replace with new one
		cell := &Cell{
			Key: key,
			Value: value,
			RecentUse: time.Now().UnixNano(),
		}
	
	
		// If  reach out its capacity, then replace the oldest one
		if this.Cap == len(this.Data) {
			min := time.Now().UnixNano()
			var idx int
	
			for i,v := range this.Data {
				if v.RecentUse < min {
					min = v.RecentUse
					idx = i
				}
			}
			
			this.Map[this.Data[idx].Key] = -1
			this.Data[idx] = cell
			this.Map[key] = idx
			//this.Data[idx].RecentUse = time.Now().UnixNano()
			//this.Data[idx].Value = value
			//this.Data[idx].Key= key
			//this.Map[key] = idx
		} else {
			// It is not full, so we insert a node
	
			this.Data = append(this.Data, cell)
			this.Map[key] = len(this.Data) -1
		}
	}
}
```

### Hash Table & Double LinkedList
这里使用双向链表和 Hash Table 来做，这样的话，get 和 put 操作都能得到 {{< katex >}} \Omicron(1) {{< /katex >}} 的复杂度。这里特别需要注意的是，**双向链表最好可以加上一个 head 和 tail，这样在做删除操作的时候就可以避免很多条件判断**，从而节省时间。
```go
type Node struct {
    Key int
    Val int
    Next *Node
    Pre *Node
}

type LRUCache struct {
    size int
    capacity int
    head *Node
    tail *Node
    dict map[int]*Node
}

func (this *LRUCache) MoveToHead(node *Node) {
    this.RemoveNode(node)
    this.AddToHead(node)
}

func (this *LRUCache) RemoveNode(node *Node) {
    node.Pre.Next = node.Next
    node.Next.Pre = node.Pre
}

func (this *LRUCache) AddToHead(node *Node) {
    node.Next = this.head.Next
    this.head.Next.Pre = node
    node.Pre = this.head
    this.head.Next = node
}

func (this *LRUCache) RemoveTail() *Node {
    tail := this.tail.Pre
    this.RemoveNode(tail)
    return tail
    // 下面这种写法不对，因为删除之后 this.tail.Pre 指向的就不是同一个值了。
    // this.RemoveNode(this.tail.Pre)
    // return this.tail.Pre
}


func Constructor(capacity int) LRUCache {
    head, tail := new(Node), new(Node)
    head.Next = tail
    tail.Pre = head
    return LRUCache{
        head: head,
        tail: tail,
        capacity: capacity,
        size: 0,
        dict: make(map[int]*Node),
    }
}


func (this *LRUCache) Get(key int) int {
    if _, ok := (*this).dict[key]; !ok {
        return -1
    }
    
    node := (*this).dict[key]
    this.MoveToHead(node)
    return node.Val
}

// - 已经存在，则直接返回
// - 否则直接插入，然后检查是否需要删除最后一个元素
func (this *LRUCache) Put(key int, value int)  {
    if _, ok := (*this).dict[key]; ok {
        node := (*this).dict[key]
        node.Val = value
        this.MoveToHead(node)
    } else {
            node := new(Node)
            node.Key = key
            node.Val = value
            this.AddToHead(node)
            (*this).dict[key] = node
            (*this).size++
            if this.size > this.capacity {
                removed := this.RemoveTail()
                delete((*this).dict, removed.Key)
                (*this).size--
            } 
    }
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
 ```
