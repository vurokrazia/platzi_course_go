package main

import "fmt"

type animal interface {
	makeSound() string
}

type dog struct{}

type cat struct{}

func (d dog) makeSound() string {
	return "Waow waow"
}

func (c cat) makeSound() string {
	return "Miau"
}

func animalMakeSound(a animal) {
	fmt.Println(a.makeSound())
}

func main() {
	cafu := dog{}
	sisa := cat{}

	animalMakeSound(cafu)
	animalMakeSound(sisa)

}
