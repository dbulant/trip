// simpletoncache
package main

import "sync"

type PersonsCache struct {
	mutex sync.RWMutex
	m     map[int64]Person
}

func NewPersonsCache() *PersonsCache {
	return &PersonsCache{m: make(map[int64]Person)}
}

func (cache *PersonsCache) GetPerson(id int64) (Person, bool) {
	cache.mutex.RLock()
	defer cache.mutex.RUnlock()
	p, found := cache.m[id]
	if found {
		return p, found
	}
	return p, found
}

func (cache *PersonsCache) InsertPerson(id int64, np Person) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	_, exists := cache.m[id]
	if !exists {
		cache.m[id] = np
	}
}

func (cache *PersonsCache) DeletePerson(id int64) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	_, exists := cache.m[id]
	if exists {
		delete(cache.m, id)
	}
}
