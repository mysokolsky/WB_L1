// L1.25
// Своя функция Sleep
// Реализовать собственную функцию sleep(duration) аналогично
// встроенной функции time.Sleep, которая приостанавливает выполнение текущей горутины.

// Важно: в отличие от настоящей time.Sleep, ваша функция должна именно
// блокировать выполнение (например, через таймер или цикл),
// а не просто вызывать time.Sleep :) — это упражнение.

// Можно использовать канал + горутину, или цикл на
// проверку времени (не лучший способ, но для обучения).

// Дедлайн — 15 дек, 02:59

// Решение:

package main

import (
	"fmt"

	"sync"

	"time"
)

var wg sync.WaitGroup

// горутина
func timerExit(name string, ch chan int) {
	defer wg.Done()

	fmt.Printf("Горутина (%s) завершится через 4 секунды\n", name)
	defer fmt.Printf("Горутина (%s) завершила работу\n", name)

	for {
		value, ok := <-ch // читаем данные из канала
		if !ok {          // условие завершения горутины
			fmt.Println("Горутина: выход по окончанию таймера")
			return
		}
		fmt.Println("g:", value)
	}

}

// вызов горутины
func timerExitRun() {
	ch := make(chan int)
	timer := time.After(4 * time.Second)
	wg.Add(1)
	go timerExit("time.After", ch) // запуск горутины
loop:
	for i := 0; ; i++ {
		select {
		case <-timer: // при закрытии канала таймера
			close(ch) // закрываем канал данных
			break loop
		default:
			ch <- i*i - i // наполнение канала данными
			time.Sleep(400 * time.Millisecond)
		}
	}
}

func main() {

	timerExitRun() // time.After

	wg.Wait()
}
