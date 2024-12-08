package main

import (
	"fmt"
	"sync"
	"time"
)

type syncMap struct {
	storage map[string]int64
	mu      *sync.RWMutex
}

func (sm *syncMap) write(key string, value int64) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.storage[key] = value

}

func (sm *syncMap) read(key string) int64 {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return sm.storage[key]
}

func main() {
	values := syncMap{storage: make(map[string]int64), mu: &sync.RWMutex{}}

	go values.write("apple", 1)
	go values.write("apple", 2)
	go values.write("pineapple", 3)

	time.Sleep(1 * time.Second)

	fmt.Println(values.read("apple"))
	fmt.Println(values.read("pineapple"))
}
