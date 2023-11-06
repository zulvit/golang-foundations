package main

import "fmt"

func main() {
	quackers := []Quacker{Duck{}, Drake{}, Duckling{}}
	for _, q := range quackers {
		q.Quack()
	}
}

type (
	Quacker interface {
		Quack()
	}
	Duck     struct{}
	Drake    struct{}
	Duckling struct{}
)

func (Duck) Quack() {
	fmt.Println("Duck: quack")
}

func (Drake) Quack() {
	fmt.Println("Drake:...")
}

func (Duckling) Quack() {
	fmt.Println("Duckling: quack")
}
