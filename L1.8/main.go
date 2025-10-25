// L1.8
// Установка бита в числе

// Дана переменная типа int64. Разработать программу,
// которая устанавливает i-й бит этого числа в 1 или 0.

// Пример: для числа 5 (0101₂) установка 1-го бита в 0 даст 4 (0100₂).

// Подсказка: используйте битовые операции (|, &^).

// Дедлайн — 24 окт, 02:59

// Решение:

// Пример запуска программы из консоли с автоматическим вводом:
// go run . <<< "-011111101001011 12 0"

package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
	"unsafe"
)

// структура для хранения числа, номера бита, нуля или единицы(в формате bool)
type Number struct {
	number   *int64
	position *uint8
	bit      *bool
}

// чистка консоли
func cls() {
	fmt.Println("\033[H\033[2J")
}

// печать ошибки
func printError() {
	cls()
	fmt.Println("Ошибка ввода")
}

// замена бита в числе на 1
func insertOneBit(number int64, position uint8) int64 {
	var sign int64 = 1
	if number < 0 {
		number = -number
		sign = -1
	}
	number |= 1 << position
	return number * sign
}

// замена бита в числе на 0
func insertZeroBit(number int64, position uint8) int64 {
	var sign int64 = 1
	if number < 0 {
		number = -number
		sign = -1
	}
	number &^= 1 << position
	return number * sign
}

// ввод числа
func inputNumber() int64 {
	var number int64
	var err error = errors.New("not nil")
	cls()
	for err != nil {
		fmt.Println("Введите число в двоичной форме, можно с минусом. Например, -010011:")
		_, err = fmt.Scanf("%b", &number)
		if err != nil {
			printError()
		}
	}
	return number
}

// ввод номера бита в числе
func inputPosition() uint8 {
	var position uint8
	var err error = errors.New("not nil")
	for position > 62 || err != nil {
		fmt.Println("Введите номер перезаписываемого бита (от 0 до 62 справа налево):")
		_, err = fmt.Scanf("%d", &position)
		if err != nil || position > 62 {
			printError()
		}
	}
	return position
}

// ввод бита
func inputBit() bool {
	var bit bool
	var digit uint8
	var err error = errors.New("not nil")
	for digit > 1 || err != nil {
		fmt.Println("Выберите бит 0 или 1 для записи:")
		_, err = fmt.Scanf("%d", &digit)
		if err != nil || digit > 1 {
			printError()
		}
	}
	if digit > 0 {
		bit = true
	}
	return bit
}

// вывод в консоль введённых данных
func (num *Number) printOutput() {
	cls()
	println()

	bitStr := "?"
	posStr := "?"
	numStr := "?"

	var positionBit uint8
	if num.position != nil {
		positionBit = 64 - *num.position
		posStr = fmt.Sprintf("%*v (%dй бит)", positionBit, "v", *num.position)
		bitStr = fmt.Sprintf("%*v", positionBit, "?")
	}

	if num.bit != nil {
		digit := 0
		if *num.bit {
			digit = 1
		}
		bitStr = fmt.Sprintf("%*d", positionBit, digit)
	}

	numStr = fmt.Sprintf("%064b (в десятичной: %d)", *num.number, *num.number)

	// выводим данные ввода
	fmt.Printf("     Бит - 0 или 1: %s\n", bitStr)
	fmt.Printf("        Номер бита: %s\n", posStr)
	fmt.Printf("             Число: %s\n", numStr)

	// печатаем линию-разделитель
	lenLine := utf8.RuneCountInString(numStr + "Итоговый результат: ")
	for i := 0; i < lenLine; i++ {
		fmt.Printf("─")
	}
	println()

}

func (num *Number) input() {

	// ввод числа
	number := inputNumber()
	num.number = &number
	num.printOutput()

	// ввод номера бита в числе для замены
	position := inputPosition()
	num.position = &position
	num.printOutput()

	// ввод бита для вставки
	bit := inputBit()
	num.bit = &bit
	num.printOutput()
}

func main() {

	num := Number{}
	num.input()

	var result int64

	if num.bit != nil {
		if *num.bit {
			result = insertOneBit(*num.number, *num.position)
		} else {
			result = insertZeroBit(*num.number, *num.position)
		}

		fmt.Printf("Итоговый результат: ")

		width := int(unsafe.Sizeof(result) * 8)                        // width - ширина числа в количестве бит. Для int64 ширина width = 64
		fmt.Printf("%0*b (в десятичной: %d)\n", width, result, result) // width подставляется вместо *
	}

}
