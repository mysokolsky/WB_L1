package utils

import "fmt"

var counter = 0
var atom = 0

// сортировка левого, среднего и правого элементов и возвращение среднего
func median(mas []int, left, right int) int {
	center := (left + right) / 2
	if mas[left] > mas[center] {
		mas[left], mas[center] = mas[center], mas[left]
	}
	if mas[center] > mas[right] {
		mas[right], mas[center] = mas[center], mas[right]
	}
	if mas[left] > mas[center] {
		mas[left], mas[center] = mas[center], mas[left]
	}
	pivot := mas[center]
	return pivot
}

// Быстрая сортировка по алгоритму Хоара
func QuickSort(mas []int, left, right int) error {
	if left > right || left < 0 || right < 0 {
		return fmt.Errorf("Ошибка: для сортировки переданы неверные границы массива. ЛЕВАЯ: %v, ПРАВАЯ: %v", left, right)
	}

	atom++

	fmt.Printf("\n>>>>>> Итерация №%v <<<<<<<\n", atom)

	defer func() {
		fmt.Printf("\n---------- Выход из цикла рекурсии %v -----------\n\n", counter)
		counter--
	}()
	counter++
	fmt.Printf("\n---------- Вход  в  цикл  рекурсии %v -----------\n\n", counter)

	fmt.Println("Левая граница left = ", left)
	fmt.Println("Правая граница right = ", right)
	fmt.Printf("Массив для сортировки:\n%+v\n", mas[left:right+1])

	if right-left < 2 {
		if right-left == 1 {
			if mas[left] > mas[right] {
				mas[left], mas[right] = mas[right], mas[left]
			}
		}
		return nil
	}

	pivot := median(mas, left, right)
	fmt.Printf("Массив после поиска медианы\n%+v\n", mas[left:right+1])

	if right-left == 2 {
		return nil
	}

	fmt.Println("pivot =", pivot)

	i := left
	j := right

	for i <= j {

		fmt.Printf("левый сканер на элементе %v, правый %v\n", i, j)

		for mas[i] < pivot {
			i++
		}
		for mas[j] > pivot {
			j--
		}
		if i <= j {
			if i != j {
				fmt.Printf("Меняем местами элементы с индексами: %v <--> %v\n", i, j)
				mas[i], mas[j] = mas[j], mas[i]
				fmt.Printf("Массив с переставленными элементами:\n%+v\n", mas[left:right+1])
			}
			i++
			j--

		}
		// fmt.Scanln()
	}

	fmt.Printf("Пересечение: левый сканер на элементе %v, правый %v\n", i, j)
	fmt.Println("Переходим к рекурсии левой и правой частей!")

	if left < j {
		fmt.Println("Сортировка левой части. Индексы:", left, "-", j)
		fmt.Printf("Левый подмассив для сортировки:\n%+v\n", mas[left:j+1])
		QuickSort(mas, left, j)
	}

	if i < right {
		fmt.Println("Сортировка правой части. Индексы:", i, "-", right)
		fmt.Printf("Правый подмассив для сортировки:\n%+v\n", mas[i:right+1])
		QuickSort(mas, i, right)
	}
	return nil
}
