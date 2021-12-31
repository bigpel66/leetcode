// 146. LRU Cache
//
// https://leetcode.com/problems/lru-cache/
//
// 	Pairs of integer key and integer value could be cached.
// 	It supports the O(1) to find the key.
// 	When touching the key, it is automatically fixed by  LRU.
const notFound int = -1

type node struct {
	key        int
	value      int
	prev, next *node
}

type LRUCache struct {
	capacity   int
	cache      map[int]*node
	head, tail *node
}

// Empty Head and Tail will be created to make works easier.
func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*node, capacity),
		head:     &node{},
		tail:     &node{},
	}
	cache.head.next = cache.tail
	cache.tail.prev = cache.head
	return cache
}

// if Cache Miss, return -1.
// else, fix the order of the entry.
func (this *LRUCache) Get(key int) int {
	if res := this.cache[key]; res != nil {
		this.pop(res)
		this.pushFront(res)
		return res.value
	}
	return notFound
}

// if Key already exist, value will be updated.
// else, new entry will be created on the cache.
// Nothing will be deleted when cache has enough space, or the last entry will be deleted.
func (this *LRUCache) Put(key int, value int) {
	if res := this.cache[key]; res != nil {
		delete(this.cache, key)
		this.pop(res)
	}
	n := &node{key: key, value: value}
	if len(this.cache) >= this.capacity {
		delete(this.cache, this.tail.prev.key)
		this.pop(this.tail.prev)
	}
	this.cache[key] = n
	this.pushFront(n)
}

func (this *LRUCache) pushFront(n *node) {
	n.next = this.head.next
	this.head.next.prev = n
	this.head.next = n
	n.prev = this.head
}

func (this *LRUCache) pop(n *node) {
	n.prev.next = n.next
	n.next.prev = n.prev
}
