package main

import "fmt"

//Функция make() в Go используется для инициализации и выделения памяти для
//встроенных типов данных, таких как слайсы (slices), каналы (channels) и отображения (maps)

func main() {
	//make([]T, len, cap)
	s := make([]int, 10, 100) // Создает слайс целых чисел с длиной 10 и емкостью 100
	fmt.Println(s)

	//make(chan T, capacity)
	ch := make(chan int, 5)
	fmt.Println(ch)

	//make(map[K]V)
	m := make(map[string]int)
	fmt.Println(m)
}
