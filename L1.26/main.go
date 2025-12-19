// L1.26
// Уникальные символы в строке
// Разработать программу, которая проверяет, что все символы
// в строке встречаются один раз (т.е. строка состоит из уникальных символов).

// Вывод: true, если все символы уникальны, false, если есть повторения.
// Проверка должна быть регистронезависимой, т.е. символы в разных регистрах считать одинаковыми.

// Например: "abcd" -> true, "abCdefAaf" -> false (повторяются a/A), "aabcd" -> false.

// Подумайте, какой структурой данных удобно воспользоваться для проверки условия.

// Дедлайн — 20 дек, 02:59
// Решение:

package main

import "strings"
import "fmt"

func main() {
	str := "строка без povtareniy"
	editLowerStr := []rune(strings.ToLower(str))
	// editStr := []rune(str) // создаём копию строки str типа []rune

	m := make(map[rune]struct{}, len(editLowerStr))

	for _, value := range editLowerStr {
		_, ok := m[value]
		if ok {
			println("false")
			fmt.Printf("дубликат: '%v'\n", string(value))
			return
		} else {
			m[value] = struct{}{}
		}
	}
	println("true")
	println("дублирующих символов нет")
}
