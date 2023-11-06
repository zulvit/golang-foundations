package main

import (
	"fmt"
	"sort"
)

/*
В этом примере анонимные функции позволяют нам явно и наглядно определить критерии сравнения прямо в месте вызова функции сортировки.
Если бы мы создавали именованные функции для каждого критерия, код стал бы менее читаемым и более разрозненным.
В данном контексте анонимные функции предоставляют чистый и выразительный способ параметризации поведения.
*/

func main() {
	persons := []Person{
		{"Alice", 44},
		{"Savva", 20},
		{"Dima", 17},
	}

	// Сортировка по имени
	sortPersons(persons, func(a, b Person) bool {
		return a.Name < b.Name
	})
	fmt.Println(persons)

	// Сортировка по возрасту
	sortPersons(persons, func(a, b Person) bool {
		return a.Age < b.Age
	})
	fmt.Println(persons)
}

func sortPersons(persons []Person, compare func(a, b Person) bool) {
	sort.Slice(persons, func(i, j int) bool {
		return compare(persons[i], persons[j])
	})
}

type Person struct {
	Name string
	Age  int
}
