package main

import (
	"fmt"
	"sort"
)

/*
Функция sort.Slice предоставляет обобщенный способ сортировки слайса без необходимости реализации полного
интерфейса sort.Interface. Вместо этого достаточно передать функцию сравнения.
*/

func main() {
	nums := []int{121, 5, 134, 5, 13, 45, 1, 345}
	fmt.Println("\tПеред сортировкой слайса")
	fmt.Println(nums)

	sort.Slice(nums, func(i int, j int) bool {
		return nums[i] > nums[j]
	})

	fmt.Println("\tПосле сортировки слайса")
	fmt.Println(nums, "\n")

	words := []string{"hello", "world", "car", "maaaaaaaany"}
	fmt.Println("\tПеред сортировкой слайса")
	fmt.Println(words)

	sort.Slice(words, func(i int, j int) bool {
		return len(words[i]) > len(words[j])
	})

	fmt.Println("\tПосле сортировки слайса")
	fmt.Println(words, "\n")
}
