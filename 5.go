package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	data  map[string]string
	mutex sync.Mutex
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]string),
	}
}

func (c *Cache) Set(key, value string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data[key] = value
}

func (c *Cache) Get(key string) (string, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	val, ok := c.data[key]
	return val, ok
}

func (c *Cache) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data = make(map[string]string)
	fmt.Println("Кэш очищен")
}

func main() {
	cache := NewCache()

	cache.Set("name", "Alice")
	cache.Set("city", "Moscow")
	fmt.Println("Данные в кэше добавлены")

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			cache.Clear()
		}
	}()

	// Демонстрация работы кэша
	keys := []string{"name", "city"}

	for i := 0; i < 15; i++ {
		for _, key := range keys {
			if val, ok := cache.Get(key); ok {
				fmt.Println("Из кэша:", key, "=", val)
			} else {
				fmt.Println("Кэш пуст для ключа:", key)
			}
		}
		time.Sleep(1 * time.Second)
	}

}
