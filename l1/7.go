package main

import (
	"fmt"
	"sync"
)

type concurMap struct {
	mx sync.RWMutex
	m  map[string]string
}

func NewConcurMap() *concurMap {
	return &concurMap{
		m: make(map[string]string),
	}
}

func (c *concurMap) Store(key string, value string) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.m[key] = value
}

func (c *concurMap) Load(key string) (string, bool) {
	c.mx.Lock()
	defer c.mx.Unlock()
	val, ok := c.m[key]
	return val, ok
}
func main() {
	m := NewConcurMap()
	m.Store("Hello", "World")
	if v, ok := m.Load("Hello"); ok {
		fmt.Println(v)
	}
}
