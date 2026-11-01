package main

import (
	"fmt"
	"tamagochi/service"
)

// структура петомца для игры тамогочи

func main() {
	var NamePet string
	var AnimalPet string
	var choice int

	fmt.Println("Введите имя петомца")
	fmt.Scan(&NamePet)

	fmt.Println("какое животное хотите?")
	fmt.Scan(&AnimalPet)

	Pet1 := service.Pet{

		Animal: AnimalPet,
		Name:   NamePet,
		Happy:  70,
		Energy: 40,
		Healt:  100,
		Hungry: 100,
		Age:    0,
	}

	for {
		// fmt.Printf("имя:%v щастье%v енергия:%v тип:%v голод:%v здоровье:%v возраст:%v\n", Pet1.Name, Pet1.Happy, Pet1.Energy, Pet1.Animal, Pet1.Hungry, Pet1.Healt, Pet1.Age)

		fmt.Println(Pet1)

		go Pet1.Decay()

		fmt.Println("1.покормить")
		fmt.Println("2.поспать")
		fmt.Println("3.поиграть")
		fmt.Scan(&choice)
		if choice == 1 {
			Pet1.Feed()
		} else if choice == 2 {
			Pet1.Sleep()
		} else if choice == 3 {
			Pet1.Play()
		}

	}
}
