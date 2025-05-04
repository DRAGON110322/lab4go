package monkey

import (
	"fmt"
	"math/rand"
	"github.com/DRAGON110322/lab4/animal"
)

type Monkey struct {
	animal.Animal
	bananaCount int
}

func NewMonkey(name string, age int, species string, bananas int) *Monkey {
	return &Monkey{
		Animal:      *animal.NewAnimal(name, age, species),
		bananaCount: bananas,
	}
}

func (m *Monkey) ThrowShit() {
	if rand.Intn(2) == 1 {
		fmt.Printf("%s %s попал фекалиями!\n", m.Species(), m.Name())
	} else {
		fmt.Printf("%s %s промазал!\n", m.Species(), m.Name())
	}
}
