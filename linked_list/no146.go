package linked_list

import "container/list"

type LRUCache struct {
	capacity int
	ll       *list.List
	cache    map[int]*list.Element
}

type entry struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		ll:       list.New(),
		cache:    make(map[int]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	if ele, hit := this.cache[key]; hit {
		this.ll.MoveToFront(ele)
		return ele.Value.(*entry).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if ele, ok := this.cache[key]; ok {
		this.ll.MoveToFront(ele)
		ele.Value.(*entry).value = value
		return
	}

	ele := this.ll.PushFront(&entry{key, value})
	this.cache[key] = ele
	if this.capacity < this.ll.Len() {
		e := this.ll.Back()
		this.ll.Remove(e)
		delete(this.cache, e.Value.(*entry).key)
	}
}
