// L1.3
// Работа нескольких воркеров
// Реализовать постоянную запись данных в канал (в главной горутине).

// Реализовать набор из N воркеров, которые читают данные из этого канала и выводят их в stdout.

// Программа должна принимать параметром количество воркеров и при старте создавать указанное число горутин-воркеров.

// Решение:

package main

import (
	"fmt"
	"sync"
)

type Job string

var jobs = [...]Job{Job("One"), Job("Two"), Job("Three"), Job("Four"), Job("Five")}

var wg sync.WaitGroup

func producer(out chan<- Job) {
	defer wg.Done()
	for _, value := range jobs {
		out <- value
	}
	// close(out)
}

func worker(in <-chan Job) {
	defer wg.Done()
	for value := range in {
		fmt.Println(string(value))
	}
}

func main() {

	fmt.Println("Super!")

	ch := make(chan Job, 10) // создали буферизированный канал на 10 объектов Job

	// ch <- Job("Это строка")
	// ch <- Job("Это вторая строка")
	// ch <- Job("Прикол")

	// close(ch)

	// for {
	// 	s, ok := <-ch
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println(string(s))
	// 	fmt.Println(len(ch))
	// 	fmt.Println(cap(ch))
	// }
	wg.Add(1)
	go producer(ch)

	for i := 1; i < 5; i++ {
		wg.Add(1)
		go worker(ch)
	}

	wg.Wait()

}
