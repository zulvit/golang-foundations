package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	counter := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go counter.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(counter.Value("somekey"))
}

// SafeCounter защищает счетчик от одновременного доступа
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

// Value возвращает текущее значение счетчика для данного ключа
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock, чтобы только одна горутина могла читать c.v одновременно
	defer c.mux.Unlock() // Unlock будет вызван после завершения функции
	return c.v[key]
}
