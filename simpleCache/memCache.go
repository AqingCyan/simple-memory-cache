package simpleCache

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

func (mc *memCache) add(key string, val *memCacheValue) {
	mc.values[key] = val
	mc.currentMemorySize += val.size
}

func (mc *memCache) del(key string) {
	tmp, ok := mc.get(key)
	if ok && tmp != nil {
		mc.currentMemorySize -= tmp.size
		delete(mc.values, key)
	}
}

func (mc *memCache) get(key string) (*memCacheValue, bool) {
	val, ok := mc.values[key]
	return val, ok
}

func (mc *memCache) SetMaxSize(size string) bool {
	mc.maxMemorySize, mc.maxMemorySizeStr = ParseSize(size)

	fmt.Println(mc.maxMemorySize, mc.maxMemorySizeStr)

	return true
}

func (mc *memCache) Set(key string, value interface{}, expire time.Duration) bool {
	mc.locker.Lock()
	defer mc.locker.Unlock()

	v := &memCacheValue{
		value:  value,
		expire: time.Now().Add(expire),
		size:   GetValueSize(value),
	}

	mc.del(key)
	mc.add(key, v)

	if mc.currentMemorySize > mc.maxMemorySize {
		mc.del(key)
		panic(fmt.Sprintf("memCache: simpleCache size is over max size: %s", mc.maxMemorySizeStr))
	}

	return true
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
