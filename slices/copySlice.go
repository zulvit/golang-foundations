package main

import (
	"fmt"
)

//при срезе с массива или слайса мы не создаем новый слайс. Также, если мы присвоим одной переменной
//значение слайса другой переменной - они обе будут указывать на один массив

func main() {
	nameSlice1 := []string{"z", "u", "l", "v", "i", "t"}
	secondSlice1 := nameSlice1
	fmt.Printf(
		"Slice nameSlice1 - %s\n"+
			"Slice cap - %d\n"+
			"Slice len - %d\n"+
			"Slice adress - %p\n\n", nameSlice1, cap(nameSlice1), len(nameSlice1), &nameSlice1[0])

	fmt.Printf(
		"Slice nameSlice1 - %s\n"+
			"Slice cap - %d\n"+
			"Slice len - %d\n"+
			"Slice adress - %p\n\n", secondSlice1, cap(secondSlice1), len(secondSlice1), &secondSlice1[0])

	secondSlice1[0] = "n"
	fmt.Println(nameSlice1, secondSlice1)

	nameSlice2 := []string{"z", "u", "l", "v", "i", "t"}
	secondSlice2 := make([]string, len(nameSlice2))
	copy(secondSlice2, nameSlice2)

	fmt.Printf(
		"Slice nameSlice1 - %s\n"+
			"Slice cap - %d\n"+
			"Slice len - %d\n"+
			"Slice adress - %p\n\n", nameSlice2, cap(nameSlice2), len(nameSlice2), &nameSlice2[0])
	fmt.Printf(
		"Slice nameSlice1 - %s\n"+
			"Slice cap - %d\n"+
			"Slice len - %d\n"+
			"Slice adress - %p\n\n", secondSlice2, cap(secondSlice2), len(secondSlice2), &secondSlice2[0])
	secondSlice2[0] = "n"
	fmt.Println(nameSlice2, secondSlice2)
}
