// L1.1
// Встраивание структур
// Дана структура Human (с произвольным набором полей и методов).

// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

// Подсказка: используйте композицию (embedded struct), чтобы Action имел все методы Human.

package main

import "fmt"

type Human struct {
	Name     string
	Age      uint
	Male     bool
	Partner  *Human
	Children []Human
}

type Action struct {
	Human
}

func NewHuman(name string, age uint, male bool) *Human {
	return &Human{
		Name: name,
		Age:  age,
		Male: male,
		// Children: []Human{},
	}
}

func (h *Human) Print() {
	fmt.Printf("%+v", h)
}

func (h *Human) Marriage(partner *Human) error {
	if h == nil || partner == nil {
		return fmt.Errorf("Бракосочетание отменяется! Один или оба партнёра не явились на церемонию (nil)")
	}
	if h.Age < 18 || partner.Age < 18 {
		return fmt.Errorf("Бракосочетание пока рано! Одному из партнёров ещё не исполнилось 18!")
	}
	if h.Male && partner.Male || !h.Male && !partner.Male {
		return fmt.Errorf("Бракосочетание невозможно! Партнёры одного пола!")
	}
	if h.Partner == partner {
		return fmt.Errorf("Бракосочетание не имеет смысла! Партнёры уже женаты!")
	}
	if h.Partner != nil || partner.Partner != nil {
		return fmt.Errorf("Бракосочетание пока невозможно! Один или оба партнёра сначала должны развестись.")
	}

	h.Partner = partner
	partner.Partner = h

	fmt.Printf("Наши поздравления новой семье!\n%v, %v и %v, %v теперь официально женаты!", h.Name, h.Age, partner.Name, partner.Age)
	return nil
}

func (h *Human) ChildrenBorn(name string, male bool) (*Human, error) {

	if h == nil {
		return nil, fmt.Errorf("Ребёнок не может появиться из ниоткуда (nil)")
	}
	if h.Partner == nil {
		return nil, fmt.Errorf("Один человек не может родить (нет партнёра)")
	}
	if h.Male {
		return nil, fmt.Errorf("Мужчины не могут родить!")
	}

	if h.Age > 70 {
		return nil, fmt.Errorf("Рожать после 70? Вы серьёзно?")
	}

	newChild := NewHuman(name, 0, male)
	h.Children = append(h.Children, *newChild)

	return newChild, nil
}

func main() {

	Alex := NewHuman("Alexey", 35, true)

	Vika := NewHuman("Vika", 29, false)

	var act Action

}
