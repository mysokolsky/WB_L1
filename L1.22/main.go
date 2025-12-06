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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/result"
)

type BigNumPair struct {
	num1 interface{}
	num2 interface{}
}

func NumToString(num interface{}) string {
	return fmt.Sprint(num)
}

func convertToBig(s string) *big.Float {

	result, _ := new(big.Float).SetString(s)

	return result
}

func Add(num1, num2 *big.Float) *big.Float {

	return new(big.Float).Add(num1, num2)
}

func main() {
	var f1 float64 = 3
	var f2 int = -1

	result := Add(s1, s2)

	fmt.Println(result)

}
