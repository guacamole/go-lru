package main

import (
	"container/list"
	"fmt"
)


//["LRUCache","get","put","get","put","put","get","get"]
//   [[2],     [2],[2,6],[1],[1,5],[1,2],[1],[2]]
func main() {

	lru := NewLRU(2)
	fmt.Println(lru.Get(2))
	lru.Put(Input{2,6})
	fmt.Println(lru.Get(1))
	lru.Put(Input{1,5})
	lru.Put(Input{1,2})
	fmt.Println(lru.Get(1))
	fmt.Println(lru.Get(2))
}



type LRU struct {
	store map[int]int
	ll *list.List
	capacity int
}

type Input struct {
	Key int
	Value int
}

func NewLRU(capacity int) *LRU {
	return &LRU{
		store:    map[int]int{},
		ll:       list.New(),
		capacity: capacity,
	}
}

func (l *LRU) Get(key int) int {
	v, ok := l.store[key]
	if !ok {
		return -1
	}

	i := 0
	node := l.ll.Front()
	for i < l.ll.Len() {
		kv := node.Value.(Input)
		if kv.Key == key {
			l.ll.MoveToFront(node)
			break
		}
		node = node.Next()
		i++
	}

	return v
}

func (l *LRU) Put(kv Input) {
	i := 0
	node := l.ll.Front()
	for i < l.ll.Len() {
		kv1,ok := node.Value.(Input)
		if !ok{break}
		if kv1.Key == kv.Key {
			node.Value = kv
			l.ll.MoveToFront(node)
			l.store[kv.Key] = kv.Value
			return
		}
		node = node.Next()
		i++
	}

	if l.isFull() {
		fmt.Println("full")
		removeLast := l.ll.Back().Value.(Input)
		l.ll.Remove(l.ll.Back())
		delete(l.store, removeLast.Key)
	}

	if l.ll.Front() == nil {
		l.ll.PushFront(kv)
	} else {
		_ = l.ll.InsertBefore(kv, l.ll.Front())
		//l.ll.MoveToFront(el)
	}

	l.store[kv.Key] = kv.Value
	fmt.Println("null")
}

func (l *LRU) isFull() bool {

	return l.capacity == l.ll.Len()
}