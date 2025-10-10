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
	"fmt"
	"sync"
	"time"

	"golang.org/x/text/number"
)

var wg sync.WaitGroup

func autoReturn(name string) {
	defer wg.Done()

	fmt.Printf("Горутина №1 (%s) завершится через 4 секунды", name)
	time.Sleep(4000 * time.Millisecond)
	fmt.Printf("Горутина №1 (%s) завершает работу", name)
}

func main() {

	wg.Add(1)
	go autoReturn("автовыход через Return")

	wg.Wait()
}
