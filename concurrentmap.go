package stubigo

import "sync"

type ConcurrentInterfaceArrayMap struct {
	baseMap map[string][]interface{}
	mutex   sync.RWMutex
}

func (m *ConcurrentInterfaceArrayMap) Set(key string, value []interface{}) {
	m.mutex.Lock()
	m.baseMap[key] = value
	defer m.mutex.Unlock()
}

func (m *ConcurrentInterfaceArrayMap) Get(key string) []interface{} {
	m.mutex.RLock()
	res := m.baseMap[key]
	defer m.mutex.RUnlock()
	return res
}

func NewConcurrentInterfaceArrayMap() *ConcurrentInterfaceArrayMap {
	return &ConcurrentInterfaceArrayMap{make(map[string][]interface{}), sync.RWMutex{}}
}

type ConcurrentIntMap struct {
	baseMap map[string]int
	mutex   sync.RWMutex
}

func (m *ConcurrentIntMap) Set(key string, value int) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.baseMap[key] = value
}

func (m *ConcurrentIntMap) Get(key string) int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.baseMap[key]
}

func NewConcurrentMap() *ConcurrentInterfaceArrayMap {
	return &ConcurrentInterfaceArrayMap{make(map[string][]interface{}), sync.RWMutex{}}
}

func NewConcurrentIntMap() *ConcurrentIntMap {
	return &ConcurrentIntMap{make(map[string]int), sync.RWMutex{}}
}

func (c *ConcurrentIntMap) Increase(key string) {
	c.Set(key, c.Get(key)+1)
}
