package human

import (
	"fmt"
	"math/rand"
	"time"
)

// Human представляет человека, расширяющего функциональность Monkey.
type Human struct {
	Monkey
	gender string // Изменил sex на gender (более корректный термин)
}

// NewHuman создает новый экземпляр Human.
func NewHuman(name string, age int, species string, bananaCount int, gender string) *Human {
	rand.Seed(time.Now().UnixNano()) // Инициализация рандома
	return &Human{
		Monkey: *NewMonkey(name, age, species, bananaCount),
		gender: gender,
	}
}

// SaySomething выводит случайную фразу в зависимости от пола.
func (h *Human) SaySomething() { // Более понятное название метода
	switch {
	case h.gender == "женщина" && rand.Intn(2) == 1:
		fmt.Printf("%s %s крикнула: Я ФЕМИНИСТКА!\n", h.Species(), h.Name())
	case h.gender == "женщина":
		fmt.Printf("%s %s промолчала.\n", h.Species(), h.Name())
	case h.gender == "мужчина":
		fmt.Printf("%s %s сказал: Эх, опять на работу\n", h.Species(), h.Name())
	default:
		fmt.Printf("%s %s сказал(а) ЭЩКЕРЕ \n", h.Species(), h.Name())
	}
}

// Work выводит сообщение о работе человека.
func (h *Human) Work() {
	fmt.Printf("%s %s сейчас работает.\n", h.Species(), h.Name())
}

// Gender возвращает пол человека.
func (h *Human) Gender() string { // Геттер в стиле Go (без Get)
	return h.gender
}

// Introduce представляет метод для представления человека.
func (h *Human) Introduce() {
	fmt.Printf("Привет, я %s, %s, мне %d лет. У меня %d бананов.\n",
		h.Name(),
		h.gender,
		h.Age(),
		h.BananaCount(),
	)
}
