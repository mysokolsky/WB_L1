// L1.12

// Собственное множество строк

// Имеется последовательность строк:
// ("cat", "cat", "dog", "cat", "tree").
// Создать для неё собственное множество.

// Ожидается: получить набор уникальных слов.
// Для примера, множество = {"cat", "dog", "tree"}.

// Дедлайн — 1 нояб, 02:59

// Решение:

package main

import "fmt"

func main() {

	array := [...]string{"cat", "cat", "dog", "cat", "tree"}

	m := make(map[string]struct{})

	for _, value := range array {
		m[value] = struct{}{}
	}

	fmt.Printf("\nПоследовательность:\n%+v\n", array)

	fmt.Printf("\nУникальные слова:\n")

	// вывод уникальных слов
	for key := range m {
		fmt.Printf("%v ", key)
	}

}
