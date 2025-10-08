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

			for {
				v, ok := <-ch
				if !ok {
					break
				}
				fmt.Println(v)
			}
			wg.Done()
		}()

	}

	timer := time.After(3 * time.Second)
	j := 0
loop:
	for {
		j++
		select {
		case <-timer:
			close(ch)
			break loop
		case ch <- j:
			time.Sleep(100 * time.Millisecond)
		}
	}

	wg.Wait()

}
