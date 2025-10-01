// L1.4
// Завершение по Ctrl+C
// Программа должна корректно завершаться по нажатию Ctrl+C (SIGINT).

// Выберите и обоснуйте способ завершения работы всех горутин-воркеров при получении сигнала прерывания.

// Подсказка: можно использовать контекст (context.Context) или канал для оповещения о завершении.

// Решение
// Дедлайн — 6 окт, 02:59

package main

import (
	"fmt"
	"os"
	"os/signal"
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
func producer(out chan<- *obj, ch_interrupt <-chan os.Signal) {

	taskIndex := 0
	for {
		taskIndex++

		select {

		case <-ch_interrupt:
			fmt.Println("\nПолучен сигнал остановки программы!")
			close(out)
			return

		case out <- &obj{
			taskIndex,
			time.Now(),
		}:
			time.Sleep(300 * time.Millisecond)
		}

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

	ch_interrupt := make(chan os.Signal, 1) // специальный канал для ловли сигнала interrupt от ОС

	signal.Notify(ch_interrupt, os.Interrupt) // ловим сигнал пользователя при нажатии на CTRL+C. Вообще, лучше ещё ловить и сигнал SIGTERM

	// воркеров необходимо запустить до заполнения канала, иначе до них не дойдёт программа и будет дедлок
	wg.Add(numWorkers)
	for i := 1; i <= numWorkers; i++ {
		go worker(i, ch)
	}

	producer(ch, ch_interrupt) // запускаем заполнение канала в главной горутине и там же будем отслеживать сигнал комбинации клавиш CTRL+C

}
