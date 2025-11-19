package utils

import (
	"fmt"
)

var counter = 0
var atom = 0

// Служебная функция. Нужна для Quick Sort
// Сортировка крайнего левого, среднего и крайнего правого элементов и возвращение среднего
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

// Метод сортировки Quick Sort по алгоритму Хоара
// Функция принимает указатель на слайс, а так же диапазон для сортировки в виде крайнего левого и крайнего правого индексов внутри слайса.
// И сортирует элементы внутри слайса по возрастанию
// Возвращает ошибку, если проблема с диапазоном входных крайних индексов. Если вернула nil, значит слайс удалось отсортировать
func QuickSort(mas []int, left, right int) error {
	if left > right || left < 0 || right < 0 || right >= len(mas) {
		return fmt.Errorf("Ошибка: для сортировки переданы неверные границы массива. ЛЕВАЯ: %v, ПРАВАЯ: %v", left, right)
	}

	atom++

	fmt.Printf("\n>>>>>> Итерация №%v <<<<<<<\n", atom)

	defer func() {
		fmt.Printf("\n---------- Выход из цикла рекурсии %v -----------\n\n", counter)
		if counter == 1 {
			println("\nРяд чисел отсортирован.")
		}
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
		// fmt.Scanln() // отладочная пауза по Enter(Command)
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

// Бинарный поиск числа в отсортированном ряде чисел
// На вход принимает указатель на слайс mas, служебную переменную offset первого индекса ряда (обычно это 0), и искомое число target
// На выход отправляет индекс найденного числа (первого попавшегося если такое число в слайсе не одно, а не первого от начала слайса) или -1 если число не нашлось
func BinSearch(mas []int, target int) int {
	println("\n===============================================================")
	println("\nПоиск числа", target, "(target) в сортированном ряде чисел.")

	left := 0
	right := len(mas) - 1

	iter := 0

	for {

		iter++

		fmt.Printf("\n ---- Итерация №%v ---- \n", iter)
		// если target не найден, то:
		if left > right {
			return -1
		}

		middle := (right + left) / 2

		println("Ряд чисел:")
		fmt.Printf("%+v\n", mas[left:right+1])
		println("Индекс левой границы left =", left)
		println("Индекс правой границы right =", right)
		println("Индекс среднего элемента middle:", middle)

		if mas[middle] == target {
			fmt.Printf("middle = target: %v = %v \n\n", mas[middle], target)
			return middle
		} else if mas[middle] < target {
			left = middle + 1
			fmt.Printf("middle < target: %v < %v \n", mas[middle], target)
		} else {
			right = middle - 1
			fmt.Printf("middle > target: %v > %v \n", mas[middle], target)
		}
	}

}

// Зеркальное отражение слайса.
// Первые элементы становятся последними и наоборот
func flipInt32Slice(mas []rune, left, right int) {

	i := left
	j := right
	for i < j {
		mas[i], mas[j] = mas[j], mas[i]
		i++
		j--
	}
}

// разворот слов разделённых пробелами, например
// "собака лает очень громко!" -> "громко! очень лает собака"
func flipWords(mas []rune) {

	prevSpace := 0

	for i := 0; i < len(mas); i++ {
		if mas[i] == ' ' {
			flipInt32Slice(mas, prevSpace, i-1)
			prevSpace = i + 1
		}
	}
	flipInt32Slice(mas, prevSpace, len(mas)-1) // последнее слово
}
