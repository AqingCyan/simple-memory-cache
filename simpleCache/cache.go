package simpleCache

import "time"

type Cache interface {
	// SetMaxSize size: 缓存大小限制，1kb 100kb 1mb 2mb 1gb
	SetMaxSize(size string) bool

	// Set 将 value 写入缓存
	Set(key string, value interface{}, expire time.Duration) bool

	// Get 从缓存中读取 key 对应的值
	Get(key string) (interface{}, bool)

	// Del 删除缓存中 key 对应的值
	Del(key string) bool

	// Exists 判断 key 是否存在
	Exists(key string) bool

	// Flush 清空缓存
	Flush() bool

	// Keys 缓存中的 key 数量
	Keys() int64
}
