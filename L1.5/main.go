// L1.5
// Таймаут на канал
// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала – читать эти значения.
// По истечении N секунд программа должна завершаться.

// Подсказка: используйте time.After или таймер для ограничения времени работы.

// Дедлайн — 11 окт, 02:59

// Решение:

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	ch := make(chan int, 3)

	var wg sync.WaitGroup

	numWorkers := 5
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {

		go func() {
			defer wg.Done()
			for {
				v, ok := <-ch
				if ok {
					fmt.Println(v)
				} else {
					return
				}
			}
		}()

	}

	i := 0
	for {
		i++
		select {
		case <-time.After(3 * time.Second):
			close(ch)
			return
		case ch <- i:
		}
	}

}
