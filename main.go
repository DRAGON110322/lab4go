package main

import (
	"github.com/DRAGON110322/lab4/animal"
	"github.com/DRAGON110322/lab4/monkey"
	"github.com/DRAGON110322/lab4/human"
	"github.com/DRAGON110322/lab4/politician"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	a := animal.NewAnimal("Лось", 5, "Лось")
	a.Eat()

	m := monkey.NewMonkey("Абу", 3, "Обезьяна", 5)
	m.ThrowShit()

	h := human.NewHuman("Иван", 30, "Человек", 0, "мужчина")
	h.Work()

	p := politician.NewPolitician("Петров", 50, "Человек", 0, "мужчина", "региональный")
	p.Campaign()
}
