package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var x int
	gomaxprocs := runtime.GOMAXPROCS(0)
	for i := 0; i < gomaxprocs; i++ {
		go func() {
			for {
				x++
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("x =", x)
}
