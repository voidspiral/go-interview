package concurrency

import (
	"sync"
)

/**
 * @Description 实现并发安全map
 **/
//
//type RwMutexImpConcurrencyMap interface {
//	Get(key string) (any, bool)
//	Set(key string, value any)
//	Delete(key string)
//	DeleteTimeout(key string, timeout int)
//	Len() int
//}
//
//type ConcurrencyMap struct {
//	rw    sync.RWMutex
//	items map[string]any
//}
//
//func NewConcurrencyMap() *ConcurrencyMap {
//	return &ConcurrencyMap{
//		rw:    sync.RWMutex{},
//		items: make(map[string]any),
//	}
//}
//
//func (m *ConcurrencyMap) Get(key string) (any, bool) {
//	m.rw.RLock()
//	defer m.rw.Unlock()
//	if v, ok := m.items[key]; ok {
//		return v, true
//	}
//	return nil, false
//}
//
//func (m *ConcurrencyMap) Set(key string, value any) {
//	m.rw.Lock()
//	defer m.rw.Unlock()
//	m.items[key] = value
//}
//
//func (m *ConcurrencyMap) Delete(key string) {
//	m.rw.Lock()
//	defer m.rw.Unlock()
//	delete(m.items, key)
//}
//func (m *ConcurrencyMap) Len() int {
//	m.rw.RLock()
//	defer m.rw.RUnlock()
//	return len(m.items)
//}

type ConcurrentMap struct {
    sync.RWMutex
    items map[string]interface{}
}

func (m *ConcurrentMap) Get(key string) (interface{}, bool) {
    m.RLock()
    defer m.RUnlock()
    val, ok := m.items[key]
    return val, ok
}

func (m *ConcurrentMap) Set(key string, value interface{}) {
    m.Lock()
    defer m.Unlock()
    m.items[key] = value
}

func (m *ConcurrentMap) Delete(key string) {
    m.Lock()
    defer m.Unlock()
    delete(m.items, key)
}

func (m *ConcurrentMap) Len() int {
    m.RLock()
    defer m.RUnlock()
    return len(m.items)
}
