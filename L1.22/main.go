// L1.22
// Большие числа и операции
// Разработать программу, которая перемножает, делит, складывает,
// вычитает две числовых переменных a, b, значения которых > 2^20 (больше 1 миллион).

// Комментарий: в Go тип int справится с такими числами, но обратите внимание на
// возможное переполнение для ещё больших значений. Для очень больших чисел можно использовать math/big.

// Дедлайн — 30 нояб, 02:59

// Решение:

package main

import (
	"fmt"
	"math/big"
)

type BigNum struct {
}

// func NumToString(num interface{}) string{

// }

func Add(num1, num2 interface{}) string {
	var result string

	return result
}

func main() {
	f1 := big.NewFloat(0)
	f1.SetString("3.14159265358979323846") // High-precision pi
	f2 := big.NewFloat(2)
	result := new(big.Float).Mul(f1, f2)
	fmt.Println(result)
}
