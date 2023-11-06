package main

import "fmt"

func change(num *int) {
	*num = 12
}

func main() {
	a := 5
	fmt.Println(a)
	change(&a)
	fmt.Println(a)
}
