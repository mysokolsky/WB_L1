// L1.1
// Встраивание структур
// Дана структура Human (с произвольным набором полей и методов).

// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

// Подсказка: используйте композицию (embedded struct), чтобы Action имел все методы Human.

package main

import (
	"encoding/json"
	"fmt"
	// "github.com/davecgh/go-spew/spew"
	"github.com/kr/pretty"
	"github.com/pkg/errors"
	"math/rand/v2"
	"os"
	// "os"
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
	Name       string  // Имя
	Age        uint    // возраст
	Male       bool    // мужчина = true, женщина = false
	Parent     *Human  // родитель
	Children   []Human // дети
	Partner    *Human  // супруг
	Profession string  // профессия
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
	}
}

func NewSuperHuman(name string, age uint, male bool) *SuperHuman {
	human := NewHuman(name, age, male)
	fmt.Println("Появление нового сверхчеловека!")
	return &SuperHuman{Human: *human}
}

func prettyPrint(obj interface{}) {
	b, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		pretty.Print(obj)
		return
	}
	fmt.Println(string(b))
}

func (h *Human) Print() {
	fmt.Println("Вывод информации о персонаже:")
	prettyPrint(h)
}

func (h *Human) Marriage(partner *Human) {
	var err error
	if h == nil || partner == nil {
		err = errors.New("Бракосочетание отменяется! Один или оба партнёра не явились на церемонию (nil)")
	} else if h.Age < 18 || partner.Age < 18 {
		err = errors.New("Бракосочетание пока рано! Одному из партнёров ещё не исполнилось 18!")
	} else if h.Male && partner.Male || !h.Male && !partner.Male {
		err = errors.New("Бракосочетание невозможно! Партнёры одного пола!")
	} else if h.Partner == partner {
		err = errors.New("Бракосочетание не имеет смысла! Партнёры уже женаты!")
	} else if h.Partner != nil || partner.Partner != nil {
		err = errors.New("Бракосочетание пока невозможно! Один или оба партнёра сначала должны развестись.")
	} else if h.Parent != nil && h.Parent == partner.Parent {
		err = errors.New("Бракосочетание между родными братом и сестрой невозможно!")
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка: %+v\n", err)
		return
	}

	h.Partner = partner
	partner.Partner = h

	fmt.Printf("%v %v лет и %v %v лет теперь официально семья!\n", h.Name, h.Age, partner.Name, partner.Age)
}

func (h *Human) ChildrenBorn(name string, male bool) *Human {

	var err error

	if h == nil {
		err = errors.New("Ребёнок не может появиться из ниоткуда (nil)")
	} else if h.Partner == nil {
		err = errors.New("Один человек не может родить (нет партнёра)")
	} else if h.Male {
		err = errors.New("Мужчины не могут рожать!")
	} else if h.Age > 70 {
		err = errors.New("Рожать после 70? Вы серьёзно?")
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка: %+v\n", err)
		os.Exit(1)
	}

	newChild := NewHuman(name, 0, male)
	newChild.Parent = h
	h.Children = append(h.Children, *newChild)

	return newChild
}

func (h *Human) getEndVerb() string {
	if h.Male {
		return ""
	}
	return "а"
}

func (h *Human) GetProfession() {
	h.Profession = Professions[rand.IntN(len(Professions)-1)]
	fmt.Println("Персонаж", h.Name, "получил"+h.getEndVerb(), "профессию:", h.Profession)
}

func (sh *SuperHuman) GetSuperFeature() {
	sh.SuperFeature = SuperHumanFeatures[rand.IntN(len(SuperHumanFeatures)-1)]
	fmt.Println("Персонаж", sh.Name, "обрел"+sh.getEndVerb(), "сверхспособность:", sh.SuperFeature)
}

func (h *Human) HumanBecomeSuperHuman() *SuperHuman {
	fmt.Println("C", h.Name, "произошло невероятное событие, изменившее физические возможности!")
	fmt.Println(h.Name, "стал"+h.getEndVerb(), "сверхчеловеком!")
	newSuperHuman := &SuperHuman{Human: *h}
	newSuperHuman.Name = "super " + newSuperHuman.Name
	return newSuperHuman
}

func (h *Human) UpdateAge(newAge uint) {
	oldAge := h.Age
	h.Age = newAge
	fmt.Println("У персонажа", h.Name, "изменился возраст с", oldAge, "на", h.Age, "лет")

	if newAge > oldAge && h.Partner != nil {
		dist := newAge - oldAge
		oldAge = h.Partner.Age
		h.Partner.Age += dist
		fmt.Println("У персонажа", h.Partner.Name, "изменился возраст с", oldAge, "до", h.Partner.Age, "лет")
	}
}

type Action struct {
	SuperHuman
}

func main() {

	Alex := NewHuman("Alexey", 35, true)

	Vika := NewHuman("Vika", 29, false)

	Alex.Print()

	Alex.Marriage(Vika)

	Vika.UpdateAge(45)

	Vika.GetProfession()

	Peter := Vika.ChildrenBorn("Peter", true)

	Jessy := Vika.ChildrenBorn("Jessy", false)

	Vika.Print()

	Jessy.UpdateAge(19)
	Peter.UpdateAge(18)

	SuperPeter := Peter.HumanBecomeSuperHuman()

	SuperPeter.GetSuperFeature()

	action := Action{*SuperPeter}

	action.GetProfession()

	action.UpdateAge(56)

	action.Print()

}
