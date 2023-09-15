package cacheServer

import (
	"SimpleMemeCache/simpleCache"
	"time"
)

type CacheServer struct {
	memCache simpleCache.Cache
}

func (cs *CacheServer) SetMaxSize(size string) bool {
	return cs.memCache.SetMaxSize(size)
}

func (cs *CacheServer) Set(key string, value interface{}, expire ...time.Duration) bool {
	expireTs := time.Second * 0
	if len(expire) > 0 {
		expireTs = expire[0]
	}

	return cs.memCache.Set(key, value, expireTs)
}

func (cs *CacheServer) Get(key string) (interface{}, bool) {
	return cs.memCache.Get(key)
}

func (cs *CacheServer) Del(key string) bool {
	return cs.memCache.Del(key)
}

func (cs *CacheServer) Exists(key string) bool {
	return cs.memCache.Exists(key)
}

func (cs *CacheServer) Flush() bool {
	return cs.memCache.Flush()
}

func (cs *CacheServer) Keys() int64 {
	return cs.memCache.Keys()
}

func NewMemCache() *CacheServer {
	return &CacheServer{
		memCache: simpleCache.NewMemCache(),
	}
}
