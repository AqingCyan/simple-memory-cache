package cache

import (
	"fmt"
	"sync"
	"time"
)

type memCacheValue struct {
	value  interface{}
	expire time.Time
	size   int64
}

type memCache struct {
	maxMemorySize     int64
	maxMemorySizeStr  string
	currentMemorySize int64
	values            map[string]*memCacheValue
	locker            sync.RWMutex
}

func NewMemCache() Cache {
	return &memCache{}
}

func (mc *memCache) SetMaxSize(size string) bool {
	mc.maxMemorySize, mc.maxMemorySizeStr = ParseSize(size)

	fmt.Println(mc.maxMemorySize, mc.maxMemorySizeStr)

	return false
}

func (mc *memCache) Set(key string, value interface{}, expire time.Duration) bool {
	defer mc.locker.Unlock()
	v := &memCacheValue{
		value:  value,
		expire: time.Now().Add(expire),
		size:   GetValueSize(value),
	}

	mc.locker.Lock()

	mc.values[key] = v

	return false
}

func (mc *memCache) Get(key string) (interface{}, bool) {
	//TODO implement me
	return nil, false
}

func (mc *memCache) Del(key string) bool {
	//TODO implement me
	return false
}

func (mc *memCache) Exists(key string) bool {
	//TODO implement me
	return false
}

func (mc *memCache) Flush() bool {
	//TODO implement me
	return false
}

func (mc *memCache) Keys() int64 {
	//TODO implement me
	return 0
}
