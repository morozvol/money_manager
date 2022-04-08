package util

import (
	"context"
	"github.com/morozvol/money_manager/internal/tgbot"
	"sync"
)

type CancelFuncMap struct {
	mx *sync.RWMutex
	m  map[tgbot.UserChat]context.CancelFunc
}

func (c *CancelFuncMap) Load(key tgbot.UserChat) (context.CancelFunc, bool) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	val, ok := c.m[key]
	return val, ok
}

func (c *CancelFuncMap) Store(key tgbot.UserChat, value context.CancelFunc) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.m[key] = value
}

func NewCancelFuncMap() CancelFuncMap {
	return CancelFuncMap{mx: new(sync.RWMutex), m: make(map[tgbot.UserChat]context.CancelFunc)}
}
