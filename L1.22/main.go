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

type BigFloatPair struct {
	num1 *big.Float
	num2 *big.Float

	result *big.Float
}

func NewBigFloatPair(n1, n2 interface{}) *BigFloatPair {

	num1 := NumToBigFloat(n1)
	num2 := NumToBigFloat(n2)

	b := BigFloatPair{num1, num2, nil}

	return &b
}

type BigFloatPairSimpleArithmetics interface {
	Add() *big.Float
	Sub() *big.Float
	Mul() *big.Float
	Div() *big.Float
}

// type Adapter struct {
// 	b *BigFloatPair
// }

// func (a *Adapter) Add() {

// }

func (b *BigFloatPair) Add() {
	b.result = Add(b.num1, b.num2)
}

func (b *BigFloatPair) Sub() {
	b.result = Sub(b.num1, b.num2)
}

func (b *BigFloatPair) Mul() {
	b.result = Mul(b.num1, b.num2)
}

func (b *BigFloatPair) Div() {
	b.result = Div(b.num1, b.num2)
}

func Add(num1, num2 *big.Float) *big.Float {

	return new(big.Float).Add(num1, num2)
}

func Sub(num1, num2 *big.Float) *big.Float {

	return new(big.Float).Sub(num1, num2)
}

func Mul(num1, num2 *big.Float) *big.Float {

	return new(big.Float).Mul(num1, num2)
}

func Div(num1, num2 *big.Float) *big.Float {

	return new(big.Float).Quo(num1, num2)

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

	b := NewBigFloatPair(n1, n2)

	b.Div()

	fmt.Println(b.result)

}
