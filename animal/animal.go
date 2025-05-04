package animal

import "fmt"

type Animal struct {
	name    string
	age     int
	species string
}

func NewAnimal(name string, age int, species string) *Animal {
	return &Animal{
		name:    name,
		age:     age,
		species: species,
	}
}

func (a *Animal) Eat() {
	fmt.Printf("%s %s ест.\n", a.species, a.name)
}

func (a *Animal) Name() string {
	return a.name
}
