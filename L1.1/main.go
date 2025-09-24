// L1.1
// Встраивание структур
// Дана структура Human (с произвольным набором полей и методов).

// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

// Подсказка: используйте композицию (embedded struct), чтобы Action имел все методы Human.

package main

import (
	"fmt"
	"math/rand/v2"
)

// Superhuman features
var SuperHumanFeatures = [...]string{
	"Fly",
	"Fire",
	"Teleportation",
	"UnderWaterBreath",
	"SuperStrong",
	"BulletProof",
}

// Professions
var Professions = [...]string{
	"Economist",
	"Developer",
	"Lawer",
	"Builder",
	"Driver",
	"Manager",
}

type Human struct {
	Name       string
	Age        uint
	Male       bool
	Partner    *Human
	Children   []Human
	Profession string
}

type SuperHuman struct {
	Human
	SuperFeature string
}

func NewHuman(name string, age uint, male bool) *Human {
	return &Human{
		Name: name,
		Age:  age,
		Male: male,
		// Children: []Human{},
	}
}

func NewSuperHuman(name string, age uint, male bool) *SuperHuman {
	human := NewHuman(name, age, male)
	fmt.Println("Появление нового сверхчеловека!")
	return &SuperHuman{Human: *human}
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

func (h *Human) GetEndVerb() string {
	if h.Male {
		return ""
	}
	return "а"
}

func (h *Human) GetProfession() {
	h.Profession = Professions[rand.IntN(len(Professions)-1)]
	fmt.Println("Персонаж", h.Name, "получил"+h.GetEndVerb(), "профессию:", h.Profession)
}

func (sh *SuperHuman) GetSuperFeature() {
	sh.SuperFeature = SuperHumanFeatures[rand.IntN(len(SuperHumanFeatures)-1)]
	fmt.Println("Персонаж", sh.Name, "обрел"+sh.GetEndVerb(), "сверхспособность:", sh.SuperFeature)
}

func (h *Human) HumanBecomeSuperHuman() *SuperHuman {
	fmt.Println("C", h.Name, "произошло экстраординарное событие!")
	fmt.Println("И", h.Name, "стал"+h.GetEndVerb(), "сверхчеловеком!")
	return &SuperHuman{Human: *h}
}

func (h *Human) UpdateAge(newAge uint) {
	oldAge := h.Age
	h.Age = newAge
	fmt.Println("У персонажа", h.Name, "изменился возраст с", oldAge, "на", h.Age)
}

func main() {

	Alex := NewHuman("Alexey", 35, true)

	Vika := NewHuman("Vika", 29, false)

	if err := Alex.Marriage(Vika); err != nil {
		fmt.Println(err)
		return
	}

	Peter, err := Vika.ChildrenBorn("Peter", true)
	if err != nil {
		fmt.Println(err)
		return
	}

	Peter.UpdateAge(18)

	SuperPeter := Peter.HumanBecomeSuperHuman()

	SuperPeter.GetSuperFeature()

	action := struct{ SuperHuman }{*SuperPeter}

	action.Print()

}
