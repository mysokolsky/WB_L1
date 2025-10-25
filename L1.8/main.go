// L1.8
// Установка бита в числе

// Дана переменная типа int64. Разработать программу,
// которая устанавливает i-й бит этого числа в 1 или 0.

// Пример: для числа 5 (0101₂) установка 1-го бита в 0 даст 4 (0100₂).

// Подсказка: используйте битовые операции (|, &^).

// Дедлайн — 24 окт, 02:59

// Решение:

package main

import (
	"errors"
	"fmt"
	"unsafe"
)

type Number struct {
	number   int64
	position uint8
	bit      bool
}

func cls() {
	fmt.Println("\033[H\033[2J")
}

func insertOneBit(number int64, position uint8) int64 {
	var sign int64 = 1
	if number < 0 {
		number = -number
		sign = -1
	}
	number |= 1 << position
	return number * sign
}

func insertZeroBit(number int64, position uint8) int64 {
	var sign int64 = 1
	if number < 0 {
		number = -number
		sign = -1
	}
	number &^= 1 << position
	return number * sign
}

func inputNumber() int64 {
	var number int64
	var err error = errors.New("not nil")
	cls()
	for err != nil {
		fmt.Println("Введите число в двоичной форме, можно с минусом. Например, -010011:")
		_, err = fmt.Scanf("%b", &number)
		if err != nil {
			cls()
			fmt.Println("Ошибка ввода")
		}
	}
	return number
}

func inputPosition() uint8 {
	var position uint8
	var err error = errors.New("not nil")
	for position > 62 || err != nil {
		fmt.Println("Введите номер перезаписываемого бита (от 0 до 62 справа налево):")
		_, err = fmt.Scanf("%d", &position)
		if err != nil || position > 62 {
			cls()
			fmt.Println("\nОшибка ввода")
		}
	}
	return position
}

func inputBit() bool {
	var bit bool
	var digit uint8
	var err error = errors.New("not nil")
	for digit > 1 || err != nil {
		fmt.Println("\nВыберите бит 0 или 1 для записи:")
		_, err = fmt.Scanf("%d", &digit)
		if err != nil || digit > 1 {
			cls()
			fmt.Println("Ошибка ввода")
		}
	}
	if digit > 0 {
		bit = true
	}
	return bit
}

func (num *Number) printOutput() {
	cls()
	println()
	pos_bit := 64 - num.position
	digit := 0
	if num.bit {
		digit = 1
	}
	fmt.Printf("     Бит - 0 или 1: %*d\n", pos_bit, digit)
	fmt.Printf("        Номер бита: %*v (%dй бит)\n", pos_bit, "v", num.position)
	fmt.Printf("             Число: %064b (в десятичной: %d)\n", num.number, num.number)
}

func (num *Number) input() {

	num.number = inputNumber()
	cls()
	fmt.Printf("\nЧисло: %064b (в десятичной: %d)\n", num.number, num.number)

	num.position = inputPosition()
	cls()
	fmt.Printf("         Номер бита: %*v (%dй бит)\n", 64-num.position, "v", num.position)
	fmt.Printf("              Число: %064b\n", num.number)

	num.bit = inputBit()
	cls()

	num.printOutput()
}

func main() {

	var num Number
	num.input()
	var result int64

	if num.bit {
		result = insertOneBit(num.number, num.position)
	} else {
		result = insertZeroBit(num.number, num.position)
	}

	fmt.Printf("Итоговый результат: ")

	width := int(unsafe.Sizeof(result) * 8) // width - ширина числа в количестве бит. Для int64 ширина width = 64
	fmt.Printf("%0*b\n", width, result)     // width подставляется вместо *

}
