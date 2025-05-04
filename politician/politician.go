package politician

import (
	"fmt"
	"math/rand"
	"time"
)

// Politician представляет политика, расширяющего функциональность Human.
type Politician struct {
	Human
	status string
}

// NewPolitician создает новый экземпляр Politician.
func NewPolitician(name string, age int, species string, bananaCount int, gender string, status string) *Politician {
	rand.Seed(time.Now().UnixNano())
	return &Politician{
		Human:  *NewHuman(name, age, species, bananaCount, gender),
		status: status,
	}
}

// PromiseButNotDeliver выводит типичное предвыборное обещание.
func (p *Politician) PromiseButNotDeliver() {
	genderAction := "сказал"
	if p.Gender() == "женщина" {
		genderAction = "сказала"
	}

	if p.status == "международный" {
		fmt.Printf("Политик %s %s %s: к 2030 году в каждой деревне будет скоростной интернет!\n",
			p.status, p.Name(), genderAction)
	}
}

// DistributeBuckwheat имитирует раздачу гречки перед выборами.
func (p *Politician) DistributeBuckwheat() {
	if p.status == "региональный" {
		fmt.Printf("Политик %s по имени %s заявил(а): Уважаемые жители, в преддверии выборов каждому по килограмму гречки!\n",
			p.status, p.Name())
	}
}

// BuildRoadButInWrongPlace имитирует неэффективное строительство.
func (p *Politician) BuildRoadButInWrongPlace() {
	cities := []string{"Брест", "Могилёв", "Гомель", "Гродно", "Витебск"}
	fails := []string{
		"дорогу в поле без съездов",
		"трассу, которая обрывается на половине",
		"мост ведущий в никуда",
		"асфальт прямо поверх травы и грязи",
		"тоннель, в который никто не может заехать",
	}

	fmt.Printf("В городе %s построили %s\n",
		cities[rand.Intn(len(cities))],
		fails[rand.Intn(len(fails))])
}

// CollectTaxes имитирует сбор налогов (с ошибкой в оригинальном названии).
func (p *Politician) CollectTaxes() {
	money := rand.Intn(10000) + 1000 // Диапазон 1000-11000
	fmt.Printf("Политик %s %s собрал налоги и заработал %d рублей\n",
		p.status, p.Name(), money)
}

// Status возвращает статус политика.
func (p *Politician) Status() string {
	return p.status
}

// Campaign запускает предвыборную кампанию.
func (p *Politician) Campaign() {
	fmt.Printf("%s %s начинает предвыборную кампанию!\n", p.Status(), p.Name())
	p.PromiseButNotDeliver()
	p.DistributeBuckwheat()
}
