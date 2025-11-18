// L1.18
// Конкурентный счетчик
// Реализовать структуру-счётчик, которая будет
// инкрементироваться в конкурентной среде (т.е. из нескольких горутин).
// По завершению программы структура должна выводить итоговое значение счётчика.

// Подсказка: вам понадобится механизм синхронизации, например,
// sync.Mutex или sync/Atomic для безопасного инкремента.

// Дедлайн — 16 нояб, 02:59

// Решение:

package main

import (
	"fmt"
	"sync"
	// "time"
)

var wg sync.WaitGroup

type сoncurrentCounter struct {
	counter int
	sync.Mutex
}

func main() {

	var count сoncurrentCounter

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 10; j++ {
				count.Lock()
				count.counter++
				count.Unlock()
			}

			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(count.counter)
}
