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

type ConcurrentCounter struct {
	counter int
	mute    sync.Mutex
}

func main() {

	var count ConcurrentCounter

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 10; j++ {
				count.mute.Lock()
				count.counter++
				count.mute.Unlock()
			}
			// time.Sleep(100 * time.Millisecond)
			wg.Done()
		}()
	}
	wg.Wait()
	defer fmt.Println(count.counter)
}
