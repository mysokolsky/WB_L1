// L1.3
// Работа нескольких воркеров
// Реализовать постоянную запись данных в канал (в главной горутине).

// Реализовать набор из N воркеров, которые читают данные из этого канала и выводят их в stdout.

// Программа должна принимать параметром количество воркеров и при старте создавать указанное число горутин-воркеров.

// Решение:

package main

import (
	"fmt"
	"os"
	"strconv"

	"sync"
	"time"
)

type obj struct {
	id   int
	time time.Time
}

var wg sync.WaitGroup

// продюсер пишет в канал
func producer(out chan<- *obj) {

	taskIndex := 0
	for {
		taskIndex++
		out <- &obj{
			taskIndex,
			time.Now(),
		}
		time.Sleep(300 * time.Millisecond)
	}
}

// воркер читает из канала
func worker(id_worker int, in <-chan *obj) {
	defer wg.Done()
	for value := range in {
		fmt.Printf("w_%v\t|  task_id_%v\t|  %v\n",
			id_worker,
			value.id,
			value.time.Format("2006-01-02 15:04:05.000000"))
	}
}

// главная горутина
func main() {

	defer wg.Wait() // эта штука здесь в принципе не нужна, потому что поток бесконечный

	numWorkers := 5 // если параметр при запуске не задан, то количество воркеров = 5

	if len(os.Args) > 1 {
		num, err := strconv.Atoi(os.Args[1])
		if err == nil && num > 0 {
			numWorkers = num
		} else {
			fmt.Println("Параметр не распознан как целое положительное число.")
			fmt.Printf("Количество воркеров будет по умолчанию = %v\n\n", numWorkers)
		}
	}

	ch := make(chan *obj, 3) // создали буферизированный канал на 3 объекта специально, чтоб реализовать принцип GMP-переключения горутин

	// воркеров необходимо запустить до заполнения канала, иначе до них не дойдёт программа и будет дедлок
	wg.Add(numWorkers)
	for i := 1; i <= numWorkers; i++ {
		go worker(i, ch)
	}

	producer(ch) // запускаем заполнение канала в главной горутине

}
