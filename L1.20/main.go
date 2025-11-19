// L1.20
// Разворот слов в предложении
// Разработать программу, которая переворачивает порядок слов в строке.

// Пример: входная строка:

// «snow dog sun», выход: «sun dog snow».

// Считайте, что слова разделяются одиночным пробелом.
// Постарайтесь не использовать дополнительные срезы, а выполнять операцию «на месте».

// Дедлайн — 25 нояб, 02:59

// Решение:

package main

import "fmt"
import "github.com/mysokolsky/WB_L1/utils" // необходимо загрузить утилиты в консоли из папки проекта командой go get github.com/mysokolsky/WB_L1/utils

func main() {

	str := "собака лает очень громко!"
	fmt.Printf("%+v\n", str)

	strExt := []rune(str)

	utils.FlipInt32Slice(strExt, 0, len(strExt)-1)

	utils.FlipWords(strExt)

	fmt.Printf("%+v", string(strExt))

}
