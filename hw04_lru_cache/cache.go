package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	Cache // Remove me after realization.

	capacity int
	queue    List
	items    map[Key]*ListItem
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

func (c *lruCache) Set(key Key, value interface{}) bool {
	if element, exists := c.items[key]; exists == true {
		c.queue.Front(element)
		element.Value.
		return true
	}
	if c.queue.Len() == c.capacity {
		c.purge()
	}
	item := &Item{
		Key: key;
		Value: value,
	}
	element = c.queue.PushFront(item)
	c.items[item.Key] = element
	return true
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	return
}

func (c *lruCache) Clear() {

}