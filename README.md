# Simple Memory Cache

一个简单的内存缓存，支持过期时间，支持缓存大小限制，支持并发安全。

按以下接口实现

```go
type Cache interface {
	// size: 缓存大小限制，1kb 100kb 1mb 2mb 1gb
	SetMaxSize(size string) bool
	
	// 将 value 写入缓存
	Set(key string, value interface{}, expire time.Duration) bool
	
	// 从缓存中读取 key 对应的值
	Get(key string) (interface{}, bool)
	
	// 删除缓存中 key 对应的值
	Del(key string) bool
	
	// 判断 key 是否存在
	Exists(key string) bool
	
	// 清空缓存
	Flush() bool
	
	// 缓存中的 key 数量
	Keys() int64
}
```

使用示例

```go
cache := NewMemCache()

cache.SetMaxSize("10mb")

cache.Set("int", 1)
cache.Set("string", "hello")
cache.Set("struct", struct {
	Name string
	Age  int
}{
	Name: "test", 
	Age:  18,
})
cache.Set("expire", "expire", time.Second*2)

cache.Flush()

cache.Keys()
```