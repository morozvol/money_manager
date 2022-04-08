package objects

import (
	"context"
	"sync"
)

type CancelFuncMap struct {
	mx *sync.RWMutex
	m  map[UserChat]context.CancelFunc
}

func (c *CancelFuncMap) Load(key UserChat) (context.CancelFunc, bool) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	val, ok := c.m[key]
	return val, ok
}

func (c *CancelFuncMap) Store(key UserChat, value context.CancelFunc) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.m[key] = value
}

func NewCancelFuncMap() CancelFuncMap {
	return CancelFuncMap{mx: new(sync.RWMutex), m: make(map[UserChat]context.CancelFunc)}
}
