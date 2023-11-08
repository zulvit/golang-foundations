package main

import "fmt"

func main() {
	cache := newCache()

	cache.Set("name", "Alice")
	cache.Set("age", "25")

	name := cache.Get("name")
	fmt.Println(name)

	cache.Delete("age")
}

type Cache struct {
	store map[string]string
}

func newCache() *Cache {
	return &Cache{store: make(map[string]string)}
}

func (c *Cache) Set(key string, value string) {
	c.store[key] = value
}

func (c *Cache) Get(key string) string {
	return c.store[key]
}

func (c *Cache) Delete(key string) {
	delete(c.store, key)
}
