package cache

import (
	"container/list"
	"fmt"
)

type LRUCache interface {
	Add(key, value string) bool
	Get(key string) (value string, ok bool)
	Remove(key string) (ok bool)
}

type Item struct {
	Key   string
	Value string
}

type LRU struct {
	capacity int
	items    map[string]*list.Element
	queue    *list.List
}

func NewLRUCache(n int) LRUCache {
	return &LRU{
		capacity: n,
		items:    make(map[string]*list.Element, n),
		queue:    list.New(),
	}
}

func (c *LRU) Add(key, value string) bool {
	if el, ok := c.items[key]; ok {
		c.queue.MoveToFront(el)
		el.Value.(*Item).Value = value
		return false
	}

	if c.queue.Len() >= c.capacity {
		if element := c.queue.Back(); element != nil {
			item := c.queue.Remove(element).(*Item)
			c.Remove(item.Key)
		}
	}

	item := &Item{
		Key:   key,
		Value: value,
	}

	el := c.queue.PushFront(item)
	c.items[item.Key] = el

	return true
}

func (c *LRU) Remove(key string) (ok bool) {
	fmt.Printf("Remove element %s\n", key)
	delete(c.items, key)
	return true
}

func (c *LRU) Get(key string) (value string, ok bool) {
	element, exists := c.items[key]
	if !exists {
		return "", false
	}
	c.queue.MoveToFront(element)
	return element.Value.(*Item).Value, true
}