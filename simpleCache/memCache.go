package simpleCache

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type memCacheValue struct {
	value      interface{}
	expireTime time.Time
	expire     time.Duration
	size       int64
}

type memCache struct {
	maxMemorySize            int64
	maxMemorySizeStr         string
	currentMemorySize        int64
	values                   map[string]*memCacheValue
	clearExpiredTimeInterval time.Duration
	locker                   sync.RWMutex
}

func NewMemCache() Cache {
	mc := &memCache{
		values:                   make(map[string]*memCacheValue),
		clearExpiredTimeInterval: time.Second,
	}

	go mc.cleanExpiredItem()

	return mc
}

// add 小粒度操作，添加缓存，如果缓存已存在，则覆盖，并计算当前值所占内存大小
func (mc *memCache) add(key string, val *memCacheValue) {
	mc.values[key] = val
	mc.currentMemorySize += val.size
}

// del 小粒度操作，删除缓存，并重新计算当前值所占内存大小
func (mc *memCache) del(key string) {
	tmp, ok := mc.get(key)
	if ok && tmp != nil {
		mc.currentMemorySize -= tmp.size
		delete(mc.values, key)
	}
}

// get 小粒度操作，获取缓存
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
		value:      value,
		expireTime: time.Now().Add(expire),
		expire:     expire,
		size:       GetValueSize(value),
	}

	// 在设置新值之前，先删除旧值，是为了保证每次操作能计算内存大小的准确性，因为如果是同名的 key，那么直接修改旧值的内存大小是不会被计算的，因此，我们把操作拆分为更小的粒度。
	mc.del(key)
	mc.add(key, v)

	if mc.currentMemorySize > mc.maxMemorySize {
		mc.del(key)
		log.Println(fmt.Sprintf("memCache: simpleCache size is over max size: %s", mc.maxMemorySizeStr))
	}

	return true
}

func (mc *memCache) Get(key string) (interface{}, bool) {
	mc.locker.RLock()
	defer mc.locker.RUnlock()

	mcVal, ok := mc.get(key)
	if ok {
		if mcVal.expire != 0 && mcVal.expireTime.Before(time.Now()) {
			mc.del(key)
			return nil, false
		}
		return mcVal.value, true
	}

	return nil, false
}

func (mc *memCache) Del(key string) bool {
	mc.locker.Lock()
	defer mc.locker.Unlock()

	mc.del(key)

	return false
}

func (mc *memCache) Exists(key string) bool {
	mc.locker.RLock()
	defer mc.locker.RUnlock()

	_, ok := mc.get(key)

	return ok
}

func (mc *memCache) Flush() bool {
	mc.locker.Lock()
	defer mc.locker.Unlock()

	mc.values = make(map[string]*memCacheValue)
	mc.currentMemorySize = 0

	return true
}

func (mc *memCache) Keys() int64 {
	mc.locker.RLock()
	defer mc.locker.RUnlock()

	return int64(len(mc.values))
}

// cleanExpiredItem 定期清除过期的缓存
func (mc *memCache) cleanExpiredItem() {
	timeTicker := time.NewTicker(mc.clearExpiredTimeInterval)
	defer timeTicker.Stop()

	for {
		select {
		case <-timeTicker.C:
			for key, item := range mc.values {
				if item.expire != 0 && item.expireTime.Before(time.Now()) {
					mc.locker.Lock()
					mc.del(key)
					mc.locker.Unlock()
				}
			}
		}
	}
}
