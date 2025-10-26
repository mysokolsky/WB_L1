// L1.14

// Определение типа переменной в runtime

// Разработать программу, которая в runtime способна определить тип переменной,
// переданной в неё (на вход подаётся interface{}).

// Типы, которые нужно распознавать: int, string, bool, chan (канал).

// Подсказка: оператор типа switch v.(type) поможет в решении.

// Дедлайн — 1 нояб, 02:59

// Решение:

package main

import "fmt"

func recognize(v interface{}) {

	switch v.(type) {
	case int:
		fmt.Printf("\nЭто переменная типа %T: %d", v, v)
	case string:
		fmt.Printf("\nЭто переменная типа %T: %s", v, v)
	case bool:
		fmt.Printf("\nЭто переменная типа %T: %t", v, v)
	case chan int:
		fmt.Printf("\nЭто канал типа: %T", v)
	}

}

func main() {
	recognize(333)
	recognize("слово")
	recognize(true)
	recognize(make(chan int))
}
