package monkey

import (
	"fmt"
	"math/rand"
	"time"
)

// Monkey представляет обезьяну, расширяющую базовый тип Animal.
type Monkey struct {
	Animal
	bananaCount int  // Изменил название для ясности
}

// NewMonkey создает новый экземпляр Monkey.
func NewMonkey(name string, age int, species string, bananaCount int) *Monkey {
	rand.Seed(time.Now().UnixNano()) // Инициализация рандома здесь
	return &Monkey{
		Animal:      *NewAnimal(name, age, species),
		bananaCount: bananaCount,
	}
}

// ThrowShit имитирует бросание обезьяной фекалий.
func (m *Monkey) ThrowShit() {
	if rand.Intn(2) == 1 {  // Упростил условие
		fmt.Printf("%s %s попал(а) фекалиями! Наконец-то!\n", m.Species(), m.Name())
	} else {
		fmt.Printf("%s %s промазал(а) фекалиями и загрустил(а).\n", m.Species(), m.Name())
	}
}

// BananaCount возвращает количество бананов у обезьяны.
func (m *Monkey) BananaCount() int {
	return m.bananaCount
}

// AddBanana добавляет бананы обезьяне.
func (m *Monkey) AddBanana(count int) {
	m.bananaCount += count
	fmt.Printf("%s получил(а) %d бананов. Теперь их %d.\n", m.Name(), count, m.bananaCount)
}
