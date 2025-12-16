// L1.23
// Удаление элемента слайса

// Удалить i-ый элемент из слайса.

// Продемонстрируйте корректное удаление без утечки памяти.

// Подсказка: можно сдвинуть хвост слайса на место удаляемого
// элемента (copy(slice[i:], slice[i+1:])) и уменьшить длину слайса на 1.

// Дедлайн — 5 дек, 02:59

// Решение:

package main

import "fmt"

func dropSliceIndex(index int, mas []int) {
	if index <0 || index>len(mas)-1 {
		fmt.Errorf("error: wrong index to delete")
		return
	}
	if len(mas) == 0 {
		fmt.Errorf("error: empty array")
		return
	}

	for i:=index;i<len(mas);i++ {
		if i<len(mas)-1 {
		mas[i] = mas[i+1]
	}

	mas.
	}

}


func main() {
var mas = []int{0,1,2,3,4,5,6,7,8,9,10}


}
