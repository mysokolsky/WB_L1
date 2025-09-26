// Задание L1.2:
// Конкурентное возведение в квадрат

// Написать программу, которая конкурентно рассчитает
// значения квадратов чисел, взятых из массива [2,4,6,8,10],
// и выведет результаты в stdout.

// Подсказка: запусти несколько горутин, каждая из которых возводит число в квадрат.

// Дедлайн — 1 окт, 02:59

// Решение:

package main

import (
	"fmt"

	log "github.com/mysokolsky/gologen"
)

func square(num uint) uint {
	return num * num
}

func main() {
	defer log.Flush()

	mas := [...]uint{2, 4, 6, 8, 10}

	for _, value := range mas {

		fmt.Println(square(value))
	}

}
