package sessioncache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

type MemCache struct {
	*cache.Cache
}

func NewMemCache() Cache{
	return &MemCache{cache.New(60*time.Minute,5*time.Minute)}
}

func (m *MemCache) SetValue(k string,v interface{},d time.Duration) error{
	m.Set(k,v,d)
	return nil
}

func (m *MemCache) StartService(config string) (error) {
	return nil
}

func init() {
	Register("memory",NewMemCache)
}

