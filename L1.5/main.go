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
	"os"
	"strconv"
	"sync"
	"time"
)

// функция, которая читает из консоли первый параметр при запуске программы, переводя его в количество секунд до завершения программы
func timeout() float64 {
	var timeout float64 = 3 // по умолчанию таймаут до выхода из программы 3 секунды
	if len(os.Args) > 1 {
		seconds, err := strconv.ParseFloat(os.Args[1], 64)
		if err == nil && seconds > 0 {
			timeout = seconds
		} else {
			fmt.Println("Параметр не распознан как целое положительное число.")
			fmt.Printf("Таймаут по умолчанию = %v секунд\n\n", timeout)
		}
	}
	return timeout
}

func main() {

	timeoutDuration := time.Duration(timeout() * float64(time.Second)) // перевели задержку из секунд в наносекунды, как этого требует Duration

	fmt.Printf("\nДо завершения программы осталось: %v секунд\n\n", timeoutDuration.Seconds())

	ch := make(chan int, 3) // основной канал для записи и чтения данных

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

	timerChan := time.After(timeoutDuration)  // функция просто возвращает канал из которого можно только читать.
	ticker := time.NewTicker(1 * time.Second) // будем каждую секунду оповещать о количестве оставшегося времени до завершения программы
	j := 0
	passedSeconds := 0 // счётчик количества секунд от старта программы
loop: // метка для выхода из цикла по брейку
	for {
		j++
		select {
		case <-timerChan: // при закрытии канала timerChan срабатывает это событие
			ticker.Stop() // останавливаем тикер
			close(ch)     // закрываем канал
			fmt.Printf("\nДостигнуто предельное значение времени ожидания работы программы = %v секунд\n", timeoutDuration.Seconds())
			fmt.Println("Завершаем работу.")
			break loop
		case ch <- j:
			time.Sleep(100 * time.Millisecond) // задержка выполнения
		case <-ticker.C:
			passedSeconds++
			expireSeconds := timeoutDuration.Seconds() - float64(passedSeconds)
			fmt.Printf("\nОсталось %v секунд до завершения программы.\n\n", expireSeconds)
		}
	}

	wg.Wait()

}
