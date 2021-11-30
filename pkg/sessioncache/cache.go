package sessioncache

import (
	"fmt"
	"time"
)

type Cache interface {
	SetValue(k string, x interface{},d time.Duration) error
	Get(k string) (interface{},bool)
	StartService(config string) error
}

type Instance func() Cache

var adapters = make(map[string]Instance)

func Register(name string,adapter Instance ) {
	if adapter == nil {
		panic("cache: register adapter is nil" )
	}
	if _, ok := adapters[name]; ok {
		panic("cache: Register called twice for adapter " + name)
	}
	adapters[name] = adapter
}

func NewCache(adapterName, config string) (adapter Cache, err error) {
	instanceFunc, ok := adapters[adapterName]
	if !ok {
		err = fmt.Errorf("cache: unknown adapter name %q (forgot to import?)", adapterName)
		return
	}
	adapter = instanceFunc()
	err = adapter.StartService(config)
	if err != nil {
		adapter = nil
	}
	return
}






