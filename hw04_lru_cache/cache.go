package hw04lrucache

import (
	"sync"
)

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	m        sync.Mutex
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	l.m.Lock()
	defer l.m.Unlock()
	el, ok := l.items[key]
	if ok {
		el.Value = cacheItem{key: key, value: value}
		l.queue.MoveToFront(el)
		l.items[key] = l.queue.Front()
		return ok
	}

	if l.queue.Len() == l.capacity {
		cItem := l.queue.Back().Value.(cacheItem)
		delete(l.items, cItem.key)
		l.queue.Remove(l.queue.Back())
	}

	li := l.queue.PushFront(cacheItem{key: key, value: value})
	l.items[key] = li
	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	l.m.Lock()
	defer l.m.Unlock()
	el, inCash := l.items[key]
	if !inCash {
		return nil, inCash
	}

	l.queue.MoveToFront(el)
	l.items[key] = l.queue.Front()

	cItem := l.queue.Front().Value.(cacheItem)
	return cItem.value, inCash
}

func (l *lruCache) Clear() {
	l.m.Lock()
	defer l.m.Unlock()
	l.items = make(map[Key]*ListItem, l.capacity)
	l.queue = NewList()
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
