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

// Благодаря RLock мы можем паралелльно читать с нескольких горутин, но запись заблокирована
func (c *concurMap) Load(key string) (string, bool) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	val, ok := c.m[key]
	return val, ok
}

//Переопределим функции мапы добавив мьютекс для операций Записи и получения значения по ключу

func main() {
	m := NewConcurMap()
	m.Store("Hello", "World")
	if v, ok := m.Load("Hello"); ok {
		fmt.Println(v)
	}
}
