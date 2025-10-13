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

// func autoReturn(name string) {
// 	defer wg.Done()

// 	fmt.Printf("Горутина №1 (%s) завершится через 4 секунды\n", name)
// 	time.Sleep(4000 * time.Millisecond)
// 	fmt.Printf("Горутина №1 (%s) завершает работу\n", name)
// }

// func closeCh(name string, ch2 chan int) {
// 	defer wg.Done()

// 	fmt.Printf("Горутина №2 (%s) завершится через 4 секунды\n", name)
// 	for value := range ch2 {
// 		fmt.Println("g2:", value)
// 	}
// 	fmt.Printf("Горутина №2 (%s) завершает работу\n", name)
// }

// func signalCh(name string, ch3 chan int, quitCh chan struct{}) {
// 	defer wg.Done()

// 	fmt.Printf("Горутина №3 (%s) завершится через 4 секунды\n", name)
// 	for {
// 		select {
// 		case value := <-ch3:
// 			fmt.Println("g3:", value)
// 		case <-quitCh:
// 			close(ch3)
// 			fmt.Printf("Горутина №3 (%s) завершает работу\n", name)
// 			return
// 		}
// 	}

// }

func contextExit(name string, ch4 chan int, ctx context.Context) {
	defer wg.Done()
	fmt.Printf("Горутина №4 (%s) завершится через 4 секунды\n", name)
	defer fmt.Printf("Горутина №4 (%s) завершает работу\n", name)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Получен сигнал отмены по контексту")
			return
		case value := <-ch4:
			fmt.Println("g4:", value)
		}
	}
}

func main() {

	// wg.Add(1)
	// go autoReturn("автовыход через Return")

	// ch2 := make(chan int)
	// wg.Add(1)
	// go closeCh("закрытие канала", ch2)
	// for i := 0; ; i++ {
	// 	ch2 <- i
	// 	time.Sleep(400 * time.Millisecond)
	// 	if i > 10 {
	// 		close(ch2)
	// 		break
	// 	}
	// }

	// ch3 := make(chan int)
	// quitCh := make(chan struct{})
	// wg.Add(1)
	// go signalCh("сигнальный канал", ch3, quitCh)
	// for i := 0; ; i++ {
	// 	ch3 <- i * i
	// 	time.Sleep(400 * time.Millisecond)
	// 	if i > 10 {
	// 		close(quitCh) // закрываем сигнальный канал
	// 		break
	// 	}
	// }

	ch4 := make(chan int)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(4*time.Second))
	defer cancel()
	wg.Add(1)
	go contextExit("context", ch4, ctx)
loop:
	for i := 0; i < 100; i++ {
		select {
		case <-ctx.Done():
			close(ch4)
			break loop
		}
		ch4 <- i*i + i
		time.Sleep(400 * time.Millisecond)
	}

	wg.Wait()
}
