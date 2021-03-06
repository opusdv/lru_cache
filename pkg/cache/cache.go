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
		el.Value.(*Item).Value = value
		return false
	}

	if c.queue.Len() >= c.capacity {
		if element := c.queue.Back(); element != nil {
			key := element.Value.(*Item).Key
			c.Remove(key)
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
	if key == "" {
		return false
	}
	fmt.Printf("Remove element %s\n", key)
	c.queue.Remove(c.items[key])
	delete(c.items, key)
	fmt.Println(c.queue.Len())
	fmt.Println(c.items)
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
