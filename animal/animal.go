package animal

import "fmt"

// Animal представляет базовый тип животного.
type Animal struct {
	name    string
	age     int
	species string
}

// NewAnimal создает новый экземпляр Animal.
func NewAnimal(name string, age int, species string) *Animal {
	return &Animal{
		name:    name,
		age:     age,
		species: species,
	}
}

// Eat выводит сообщение о том, что животное ест.
func (a *Animal) Eat() {
	fmt.Printf("%s %s ест.\n", a.species, a.name)
}

// MakeFamily выводит сообщение о создании семьи.
func (a *Animal) MakeFamily() {
	fmt.Printf("%s %s создает семью.\n", a.species, a.name)
}

// Sleep выводит сообщение о том, что животное спит.
func (a *Animal) Sleep() {
	fmt.Printf("%s %s спит и видит сны.\n", a.species, a.name)
}

// Name возвращает имя животного.
func (a *Animal) Name() string {
	return a.name
}

// Age возвращает возраст животного.
func (a *Animal) Age() int {
	return a.age
}

// Species возвращает вид животного.
func (a *Animal) Species() string {
	return a.species
}
