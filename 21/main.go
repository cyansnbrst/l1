package main

import (
	"fmt"
)

type animal interface {
	makeSound()
}

type dog struct{}

func (d *dog) makeSound() {
	fmt.Println("bark!")
}

type cat struct{}

func (c *cat) makeSound() {
	fmt.Println("meow!")
}

type phoenix struct{}

func (p *phoenix) performAction(action string) {
	switch action {
	case "sound":
		fmt.Println("roar!")
	case "fly":
		fmt.Println("a phoenix flies")
	default:
		fmt.Println("a phoenix do something")
	}
}

type animalAdapter struct {
	ph *phoenix
}

func (a *animalAdapter) makeSound() {
	a.ph.performAction("sound")
}

func main() {
	var animals []animal
	animals = append(animals, &dog{})
	animals = append(animals, &cat{})

	phoenix := &phoenix{}
	adapter := &animalAdapter{ph: phoenix}
	animals = append(animals, adapter)

	for _, animal := range animals {
		animal.makeSound()
	}
}
