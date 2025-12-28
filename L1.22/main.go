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

type BigFloat struct {
	value *big.Float
}

func (b *BigFloat) Add(num *big.Float) *big.Float {
	return new(big.Float).Add(b.value, num)
}
func (b *BigFloat) Sub(num *big.Float) *big.Float {
	return new(big.Float).Sub(b.value, num)
}
func (b *BigFloat) Mul(num *big.Float) *big.Float {
	return new(big.Float).Mul(b.value, num)
}
func (b *BigFloat) Div(num *big.Float) *big.Float {
	return new(big.Float).Quo(b.value, num)
}

func NumToString(num interface{}) string {
	return fmt.Sprint(num)
}

func StringToBig(s string) *big.Float {

	result, _ := new(big.Float).SetString(s)

	return result
}

func NumToBigFloat(num interface{}) *big.Float {

	str := NumToString(num)
	b := StringToBig(str)

	return b
}

func main() {
	var n1 float64 = 3
	var n2 int = -1

	num1 := NumToBigFloat(n1)
	num2 := NumToBigFloat(n2)

	result :=

		fmt.Println(b.result)

}
