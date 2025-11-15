// L1.16
// Быстрая сортировка (quicksort)
// Реализовать алгоритм быстрой сортировки массива
// встроенными средствами языка. Можно использовать рекурсию.

// Подсказка: напишите функцию quickSort([]int) []int которая сортирует срез целых чисел.
// Для выбора опорного элемента можно взять середину или первый элемент.

// Дедлайн — 6 нояб, 02:59

// Решение:
//
// Берём элемент крайний левый, крайний правый и в середине массива
// Нормализуем их, упорядочивая и выбирая медиану pivot
// Идём в цикле с левого края вправо до
// Находим первый элемент что > pivot
// проходим по всем остальным элементам
// все элементы, что < нашего, кладём влево, которые > кладём вправо от нашего
// Далее рекурсивно делаем с левой и правой частями то же самое
//
// По хорошему, нужно было бы сделать обработку левой и правой частей через горутины.

package main

import "fmt"

var counter = 0
var atom = 0

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

// используем для сортировки алгоритм Хоара
func sort(mas []int, left, right int) error {
	if left > right || left < 0 || right < 0 {
		return fmt.Errorf("Ошибка: Неверные границы массива")
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
	fmt.Printf("Массив для упорядочивания\n%+v\n", mas[left:right+1])

	if right-left < 2 {
		if right-left == 1 {
			if mas[left] > mas[right] {
				mas[left], mas[right] = mas[right], mas[left]
			}
		}
		return nil
	}

	pivot := median(mas, left, right)
	fmt.Printf("Массив после медианы\n%+v\n", mas[left:right+1])

	if right-left == 2 {
		return nil
	}

	fmt.Println("pivot = ", pivot)

	i := left
	j := right

	for i <= j {

		fmt.Printf("i = %v, j = %v\n", i, j)

		for mas[i] < pivot {
			i++
		}
		for mas[j] > pivot {
			j--
		}
		if i <= j {

			fmt.Printf("Меняем элементы по индексам i = %v, j = %v\n", i, j)

			mas[i], mas[j] = mas[j], mas[i]
			i++
			j--
			fmt.Printf("Массив с переставленными элементами:\n%+v\n", mas[left:right+1])
		}
		// fmt.Scanln()
	}

	fmt.Printf("i = %v, j = %v\n", i, j)
	fmt.Println("Рекурсия")

	if left < j {
		fmt.Println("Сортировка левой части. Индексы:", left, "-", j)
		fmt.Printf("Массив для сортировки:\n%+v\n", mas[left:j+1])
		sort(mas, left, j)
	}

	if i < right {
		fmt.Println("Сортировка правой части. Индексы:", i, "-", right)
		fmt.Printf("Массив для сортировки:\n%+v\n", mas[i:right+1])
		sort(mas, i, right)
	}
	return nil
}

func QuickSort(mas []int) []int {
	if err := sort(mas, 0, len(mas)-1); err != nil {
		fmt.Println(err)
	}
	return mas
}

func main() {

	var mas = []int{1, 5, 33, 9, 0, -1, 0, 4, -7, 2, 5, 23, -99, 1000, -6, 3}

	fmt.Printf("\nМассив до сортировки:\n%+v\n", mas)
	fmt.Printf("\nРезультат сортировки:\n%+v\n", QuickSort(mas))
}
