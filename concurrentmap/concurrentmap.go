package concurrentmap

import "sync"

//template type ConcurrentMap(KEY, VALUE)

type Key KEY
type Value VALUE
type cmap map[Key]Value

type ConcurrentMap struct {
	items map[Key]Value
	mutex sync.RWMutex
}

func New() *ConcurrentMap {
	return &ConcurrentMap{
		items: make(map[Key]Value),
		mutex: sync.RWMutex{},
	}
}

func (this *ConcurrentMap) Set(key Key, value Value) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.items[key] = value
}

func (this *ConcurrentMap) Remove(key Key) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	delete(this.items, key)
}

func (this *ConcurrentMap) Get(key Key) (Value, bool) {
	this.mutex.RLock()
	defer this.mutex.RUnlock()

	value, ok := this.items[key]
	return value, ok
}
