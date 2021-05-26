package main

import "sync"

// 代表互斥锁
type FooMap struct {
	sync.Mutex
	data map[int]int
}

//代表读写锁
type BarRwMap struct {
	sync.RWMutex
	data map[int]int
}

var fooMap *FooMap

var barRwMap *BarRwMap

var syncMap *sync.Map

func init() {
	fooMap = &FooMap{data: make(map[int]int, 100)}
	barRwMap = &BarRwMap{data: make(map[int]int, 100)}
	syncMap = &sync.Map{}
}

func builtinRwMapStore(k, v int) {
	barRwMap.Lock()
	defer barRwMap.Unlock()
	barRwMap.data[k] = v
}



func builtinRwMapLookup(k int) int {
	barRwMap.RLock()
	defer barRwMap.RUnlock()
	if v, ok := barRwMap.data[k]; !ok {
		return -1
	} else {
		return v
	}
}

func builtinRwMapDelete(k int) {
	barRwMap.Lock()
	defer barRwMap.Unlock()
	if _, ok := barRwMap.data[k]; !ok {
		return
	} else {
		delete(barRwMap.data, k)
	}
}
