package cache

import (
	"fmt"
	"time"
)

type memCache struct {
	maxMemorySize     int64
	maxMemorySizeStr  string
	currentMemorySize int64
}

// NewMemCache 创建一个新的内存缓存
func NewMemCache() Cache {
	return &memCache{}
}

// SetMaxSize size: 缓存大小限制，1kb 100kb 1mb 2mb 1gb
func (mc *memCache) SetMaxSize(size string) bool {
	mc.maxMemorySize, mc.maxMemorySizeStr = ParseSize(size)

	fmt.Println(mc.maxMemorySize, mc.maxMemorySizeStr)

	return false
}

// Set 将 value 写入缓存
func (mc *memCache) Set(key string, value interface{}, expire time.Duration) bool {
	//TODO implement me
	return false
}

// Get 从缓存中读取 key 对应的值
func (mc *memCache) Get(key string) (interface{}, bool) {
	//TODO implement me
	return nil, false
}

// Del 删除缓存中 key 对应的值
func (mc *memCache) Del(key string) bool {
	//TODO implement me
	return false
}

// Exists 判断 key 是否存在
func (mc *memCache) Exists(key string) bool {
	//TODO implement me
	return false
}

// Flush 清空缓存
func (mc *memCache) Flush() bool {
	//TODO implement me
	return false
}

// Keys 缓存中的 key 数量
func (mc *memCache) Keys() int64 {
	//TODO implement me
	return 0
}
