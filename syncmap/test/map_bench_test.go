package test

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)



type FooMap struct {
	sync.Mutex
	data map[int]int
}

type BarRwMap struct {
	sync.RWMutex
	data map[int]int
}

var fooMap *FooMap
var barRwMap *BarRwMap
var syncMap *sync.Map

func init()  {
	fooMap = &FooMap{data: make(map[int]int, 100)}
	barRwMap = &BarRwMap{data: make(map[int]int,100)}
	syncMap = &sync.Map{}
}


func builtinMapStore(k, v int) {
	fooMap.Lock()
	defer fooMap.Unlock()
	fooMap.data[k] = v
}

func builtinMapLookup(k int) int {
	fooMap.Lock()
	defer fooMap.Unlock()
	if v, ok := fooMap.data[k]; !ok {
		return -1
	} else {
		return v
	}
}

func builtinMapDelete(k int) {
	fooMap.Lock()
	defer fooMap.Unlock()
	if _, ok := fooMap.data[k]; !ok {
		return
	} else {
		delete(fooMap.data, k)
	}
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

func BuiltinRwMapDelete(k int) {
	barRwMap.Lock()
	defer barRwMap.Unlock()
	if _, ok := barRwMap.data[k]; !ok {
		return
	} else {
		delete(barRwMap.data, k)
	}
}



func BenchmarkBuiltinRwMapDeleteParalell(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		for pb.Next() {
			k := r.Intn(100000000)
			BuiltinRwMapDelete(k)
		}
	})
}

func BenchmarkBuiltinMapDeleteParalell(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		for pb.Next() {
			k := r.Intn(100000000)
			builtinMapDelete(k)
		}
	})
}

func BenchmarkBuiltinSyncMapDeleteParalell(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		for pb.Next() {
			k := r.Intn(100000000)
			syncMap.Delete(k)
		}
	})
}