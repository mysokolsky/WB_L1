// L1.7
// Конкурентная запись в map

// Реализовать безопасную для конкуренции запись данных в структуру map.

// Подсказка: необходимость использования синхронизации
// (например, sync.Mutex или встроенная concurrent-map).

// Проверьте работу кода на гонки (util go run -race).

// Дедлайн — 19 окт, 02:59

// Решение:

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

var ints = []int{-4, -3, -2, -1, 0, 1, 2, 3, 4, 5}
var doubles = []float32{-.5, -.4, -.3, -.2, -.1, .1, .2, .3, .4, .5}
var chars = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
var strings = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}

type MyMap struct {
	data         map[int]interface{} // хранение мапы
	sync.RWMutex                     // Мьютекс для блокирования изменений данных только одной горутиной
}

func getItem() interface{} {

	var item interface{}

	randSlice := rand.Intn(4)
	randIndex := rand.Intn(10)

	switch randSlice {
	case 0:
		item = ints[randIndex]
	case 1:
		item = doubles[randIndex]
	case 2:
		item = string(chars[randIndex])
	case 3:
		item = strings[randIndex]
	default:
		panic("Error!")
	}
	return item
}

func NewMyMap() *MyMap {

	m := make(map[int]interface{})

	return &MyMap{
		data: m,
	}
}

func (mapa *MyMap) add_update(key int, value interface{}) {
	mapa.Lock()
	mapa.data[key] = value
	mapa.Unlock()
}

func (mapa *MyMap) delete(key int) {
	mapa.Lock()
	delete(mapa.data, key)
	mapa.Unlock()
}

func (mapa *MyMap) get(key int) (interface{}, bool) {
	mapa.RLock()
	value, ok := mapa.data[key]
	mapa.RUnlock()
	return value, ok
}

func writer_updater(mapa *MyMap, numWorker int) {
	defer wg.Done()
	for i := 0; i < 20; i++ {
		randValue := getItem()
		randKey := rand.Intn(20)
		mapa.add_update(randKey, randValue)
		fmt.Printf("w%v : key = %v, val = %v\n", numWorker, randKey, randValue)
		time.Sleep(100 * time.Millisecond)
	}
}

func reader(mapa *MyMap, numWorker int) {
	defer wg.Done()
	for key := 0; key < 20; key++ {
		val, ok := mapa.get(key)
		if ok {
			fmt.Printf("r%v : key = %v, val = %v\n", numWorker, key, val)
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func remover(mapa *MyMap, numWorker int) {
	defer wg.Done()
	for i := 0; i < 20; i++ {
		randKey := rand.Intn(20)
		mapa.delete(randKey)
		fmt.Printf("d%v : key = %v удалён!\n", numWorker, randKey)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {

	m := NewMyMap()

	// запись или изменение(перезапись)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go writer_updater(m, i)
	}

	// чтение
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go reader(m, i)
	}

	// удаление
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go remover(m, 1)
	}

	// вывод окончательной мапы на экран
	defer func() {
		fmt.Println("\nФинальная мапа:")
		for key, value := range m.data {
			fmt.Printf("key = %v, val = %v\n", key, value)
		}
	}()

	wg.Wait()
}
