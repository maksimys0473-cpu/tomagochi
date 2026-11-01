package service

import "fmt"

type Pet struct {
	Animal string
	Name   string
	Healt  int
	Hungry int
	Happy  int
	Energy int
	Age    int
}

func (p *Pet) Feed() {
	if p.Hungry >= 100 {
		fmt.Println("питоиец не голоден")
		p.Hungry = 100
	} else {
		p.Hungry += 40
		p.Healt += 5
		p.Age++
		if p.Hungry >= 100 {
			p.Hungry = 100
		}
	}
}

func (p *Pet) Sleep() {
	if p.Energy >= 100 {
		p.Energy = 100
		p.Age++
		fmt.Println("питоиец не хочет отдыхать")
	} else {
		p.Energy += 70
		p.Healt += 5
		if p.Energy >= 100 {
			p.Energy = 100
		}
	}
}

func (p *Pet) Decay() {
	p.Hungry -= 1
	if p.Hungry <= 0 {
		p.Healt -= 1
		p.Happy -= 1
	}

}
func (p Pet) String() string {
	return fmt.Sprintf("имя:%v щастье%v енергия:%v тип:%v голод:%v здоровье:%v возраст:%v\n",
		p.Name, p.Happy, p.Energy, p.Animal, p.Hungry, p.Healt, p.Age)
}

func (p *Pet) Play() {
	if p.Happy >= 100 {
		p.Happy = 100
		p.Age++
		fmt.Println("питоиец не хочет отдыхать")
	} else {
		p.Happy += 30
		p.Healt += 5
		if p.Energy >= 100 {
			p.Energy = 100
		}
	}
}
