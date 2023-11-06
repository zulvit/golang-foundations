package main

import "fmt"

func main() {
	ch := make(chan string, 4)

	ch <- "hello"
	ch <- "world"
	ch <- "lol"
	ch <- "fun"

	//ch <- "not a fun!"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
