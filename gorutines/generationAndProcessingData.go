package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(2)

	go generateNumbers(ch, &wg)
	go printNumbers(ch, &wg)

	wg.Wait()
}

func generateNumbers(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // уменьшаем счетчик группы по завершении функции
	for i := 1; i <= 5; i++ {
		ch <- i // отправляем число в канал
	}
	close(ch) // закрываем канал после завершения генерации
}

func printNumbers(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()     // уменьшаем счетчик группы по завершении функции
	for n := range ch { // читаем из канала до его закрытия
		fmt.Println(n)
	}
}
