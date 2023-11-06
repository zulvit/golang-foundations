package main

import "fmt"

func main() {
	ints := []int{1, 2, 34, 5, 4, 1, 412, 4}
	fmt.Println(ints)
	fmt.Println(len(ints))
	fmt.Println(cap(ints))
	fmt.Printf("Adress of ints capacity: %p\n", &ints)

	for i := 0; i < 10; i++ {
		ints = append(ints, i)
	}

	fmt.Println(ints)
	fmt.Println(len(ints))
	fmt.Println(cap(ints))
	fmt.Printf("Adress of ints capacity: %p\n", &ints)

	i := ints[:3]
	fmt.Println(i)
	fmt.Printf("Adress of ints capacity: %p\n", &ints)
}
