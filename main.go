package main

import (
	"SimpleMemeCache/cacheServer"
	"time"
)

func main() {
	memCache := cacheServer.NewMemCache()

	memCache.SetMaxSize("300mb")

	memCache.Set("int", 1, time.Second)
	memCache.Set("string", "hello", time.Second)
	memCache.Set("data", map[string]interface{}{"a": 1}, time.Second)

	memCache.Set("int", 1)
	memCache.Set("string", "hello")
	memCache.Set("struct", struct {
		Name string
		Age  int
	}{
		Name: "test",
		Age:  18,
	})

	memCache.Set("expire", "expire", time.Second*2)

	memCache.Keys()

	memCache.Flush()
}
