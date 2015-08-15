// Template ConcurrentMap type
package concurrentmap

import "sync"

//template type ConcurrentMap(Key, Value)

type Key string
type Value int
type cmap map[Key]Value

// A "thread" safe map of type Key:Value
type ConcurrentMap struct {
	items map[Key]Value
	mutex sync.RWMutex
}

// Creates a new concurrent map.
func New() *ConcurrentMap {
	return &ConcurrentMap{
		items: make(map[Key]Value),
		mutex: sync.RWMutex{},
	}
}

// Sets the given value under the specified key.
func (this *ConcurrentMap) Set(key Key, value Value) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.items[key] = value
}

// Removes an element from the map.
func (this *ConcurrentMap) Remove(key Key) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	delete(this.items, key)
}

// Retrieves an element from map under given key.
func (this *ConcurrentMap) Get(key Key) (Value, bool) {
	this.mutex.RLock()
	defer this.mutex.RUnlock()
	val, ok := this.items[key]
	return val, ok
}

// Retrieves an element from map under given key.
// If it exists, removes it from map.
func (this *ConcurrentMap) GetAndRemove(key Key) (Value, bool) {
	this.mutex.RLock()
	defer this.mutex.RUnlock()
	val, ok := this.items[key]
	if ok {
		delete(this.items, key)
	}
	return val, ok
}

// Returns the number of elements within the map.
func (this *ConcurrentMap) Count() int {
	this.mutex.RLock()
	defer this.mutex.RUnlock()

	return len(this.items)
}

// Looks up an item under specified key
func (this *ConcurrentMap) Has(key Key) bool {
	this.mutex.RLock()
	defer this.mutex.RUnlock()

	// See if element is within shard.
	_, ok := this.items[key]
	return ok
}

// Checks if map is empty.
func (this *ConcurrentMap) IsEmpty() bool {
	return this.Count() == 0
}

