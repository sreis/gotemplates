/*
 A thread-safe concurrent-map template.
*/
package cache

import (
	"sync"
	"time"
)

//template type Cache(Key, Value)

type entry struct {
	expiresAt uint32
	value     Value
}

type cache map[Key]*entry

// A "thread" safe expirable cache  of type Key:Value
type Cache struct {
	items cache
	mutex sync.RWMutex
}

// Creates a new concurrent cache.
func New() *Cache {
	return &Cache{
		items: make(cache),
		mutex: sync.RWMutex{},
	}
}

// Sets the given value under the specified key.
func (c *Cache) Set(key Key, value Value, expireSeconds int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	now := uint32(time.Now().Unix())
	expiresAt := now + uint32(expireSeconds)

	c.items[key] = &entry{
		expiresAt: expiresAt,
		value:     value,
	}
}

// Retrieves an element from map under given key.
func (c *Cache) Get(key Key) (Value, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	if v, ok := c.items[key]; ok {

		now := uint32(time.Now().Unix())

		if v.expiresAt < now {
			delete(c.items, key)
			// we have expired the key
			return *new(Value), false
		}

		return v.value, true
	}

	return *new(Value), false
}

// Returns the number of elements within the cache.
func (c *Cache) Count() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return len(c.items)
}
