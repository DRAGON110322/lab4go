package main

import (
	"example.com/animal"
	"example.com/monkey"
	"example.com/human"
	"example.com/politician"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Пример использования
	a := animal.NewAnimal("Инокентий", 18, "Лось")
	a.Eat()

	m := monkey.NewMonkey("Игнат", 19, "Обезьяна", 4)
	m.ThrowShit()

	h := human.NewHuman("Анна", 25, "Человек", 0, "женщина")
	h.SaySomething()

	p := politician.NewPolitician("Петров", 50, "Человек", 0, "мужчина", "региональный")
	p.Campaign()
}
