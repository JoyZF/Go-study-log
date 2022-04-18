package main

import (
	"fmt"
	"github.com/coocood/freecache"
	"math"
	"runtime/debug"
	"time"
)

func main() {

	// In bytes, where 1024 * 1024 represents a single Megabyte, and 100 * 1024*1024 represents 100 Megabytes.
	cacheSize := 100 * 1024 * 1024
	cache := freecache.NewCache(cacheSize)
	debug.SetGCPercent(20)
	key := []byte("abc")
	val := []byte("def")
	expire := 60 // expire in 60 seconds
	start := time.Now()
	for i := 0; i < math.MaxUint32; i++ {
		cache.Set(key, val, expire)
	}
	fmt.Println(start.Sub(time.Now()))

}



