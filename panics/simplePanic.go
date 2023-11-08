package main

import "fmt"

func main() {
	handlePanic()
}

func causePanic() {
	panic("something bad happened")
}

func handlePanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic")
		}
	}()
	causePanic()
	fmt.Println("This line will not return")
}
