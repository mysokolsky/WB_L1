// L1.6
// Остановка горутины
// Реализовать все возможные способы остановки выполнения горутины.

// Классические подходы:
// выход по условию,
// через канал уведомления,
// через контекст,
// прекращение работы runtime.Goexit()
// и др.

// Продемонстрируйте каждый способ в отдельном фрагменте кода.

// Дедлайн — 14 окт, 02:59

// Решение:

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// горутина №1
func autoReturn(name string) {
	defer wg.Done()

	fmt.Printf("Горутина №1 (%s) завершится через 4 секунды\n", name)
	defer fmt.Printf("Горутина №1 (%s) завершила работу\n", name)

	fmt.Println("g1: работа горутины...") // безканальная работа
	time.Sleep(4000 * time.Millisecond)
	fmt.Println("Горутина №1: авто выход по завершению работы функции")

}

// горутина №2
func closeCh(name string, ch2 chan int) {
	defer wg.Done()

	fmt.Printf("Горутина №2 (%s) завершится через 4 секунды\n", name)
	defer fmt.Printf("Горутина №2 (%s) завершила работу\n", name)

	for value := range ch2 { // читаем данные из канала
		fmt.Println("g2:", value)
	} // условие завершения горутины

	fmt.Println("Горутина №2: выход по закрытию канала")

}

// горутина №3
func signalCh(name string, ch3 chan int, quitCh chan struct{}) {
	defer wg.Done()

	fmt.Printf("Горутина №3 (%s) завершится через 4 секунды\n", name)
	defer fmt.Printf("Горутина №3 (%s) завершила работу\n", name)

	for {
		select {
		case value := <-ch3: // читаем данные из канала
			fmt.Println("g3:", value)
		case <-quitCh: // условие завершения горутины
			fmt.Println("Горутина №3: выход по закрытию сигнального канала")
			close(ch3)
			return
		}
	}

}

// горутина №4
func contextExit(name string, ch4 chan int, ctx context.Context) {
	defer wg.Done()

	fmt.Printf("Горутина №4 (%s) завершится через 4 секунды\n", name)
	defer fmt.Printf("Горутина №4 (%s) завершила работу\n", name)

	for {
		select {
		case <-ctx.Done(): // условие завершения горутины
			fmt.Println("Горутина №4: выход по сигналу отмены контекста")
			return
		case value := <-ch4: // читаем данные из канала
			fmt.Println("g4:", value)
		}
	}
}

// горутина №5
func timerExit(name string, ch5 chan int) {
	defer wg.Done()

	fmt.Printf("Горутина №5 (%s) завершится через 4 секунды\n", name)
	defer fmt.Printf("Горутина №5 (%s) завершила работу\n", name)

	for {
		value, ok := <-ch5 // читаем данные из канала
		if !ok {           // условие завершения горутины
			fmt.Println("Горутина №5: выход по окончанию таймера")
			return
		}
		fmt.Println("g5:", value)
	}

}

// вызов горутины №1
func autoReturnRun() {
	wg.Add(1)
	go autoReturn("автовыход через Return") // запуск горутины
}

// вызов горутины №2
func closeChRun() {
	ch2 := make(chan int)
	wg.Add(1)
	go closeCh("закрытие канала", ch2) // запуск горутины
	for i := 0; ; i++ {
		ch2 <- i // наполнение канала данными
		time.Sleep(400 * time.Millisecond)
		if i > 10 {
			close(ch2)
			break
		}
	}
}

// вызов горутины №3
func signalChRun() {
	ch3 := make(chan int)
	quitCh := make(chan struct{})
	wg.Add(1)
	go signalCh("сигнальный канал", ch3, quitCh) // запуск горутины
	for i := 0; ; i++ {
		ch3 <- i * i // наполнение канала данными
		time.Sleep(400 * time.Millisecond)
		if i > 10 {
			close(quitCh) // закрываем сигнальный канал
			break
		}
	}
}

// вызов горутины №4
func contextExitRun() {
	ch4 := make(chan int)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(4*time.Second)) // завершаем контекст при достижении момента времени "сейчас + 4 секунды"
	defer cancel()
	wg.Add(1)
	go contextExit("context", ch4, ctx) // запуск горутины
loop:
	for i := 0; i < 5; i++ {
		select {
		case <-ctx.Done():
			close(ch4)
			break loop
		default:
			ch4 <- i*i + i // наполнение канала данными
			time.Sleep(400 * time.Millisecond)
		}

	}
}

// вызов горутины №5
func timerExitRun() {
	ch5 := make(chan int)
	timer := time.After(4 * time.Second)
	wg.Add(1)
	go timerExit("time.After", ch5) // запуск горутины
loop2:
	for i := 0; ; i++ {
		select {
		case <-timer: // при закрытии канала таймера
			close(ch5) // закрываем канал данных
			break loop2
		default:
			ch5 <- i*i - i // наполнение канала данными
			time.Sleep(400 * time.Millisecond)
		}
	}
}

func main() {

	autoReturnRun()  // 1 - автовыход через Return
	closeChRun()     // 2 - закрытие канала
	signalChRun()    // 3 - сигнальный канал
	contextExitRun() // 4 - context с дедлайном
	timerExitRun()   // 5 - time.After

	wg.Wait()
}
