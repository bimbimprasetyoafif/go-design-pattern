package main

import "fmt"

type AnimalType string

const (
	Cat AnimalType = "cat"
	Dog AnimalType = "dog"
)

type IAnimal interface {
	Talk() string
	FavFood() string
}

type Animal struct {
	Sound string
	Food  string
}

func (a Animal) Talk() string {
	return a.Sound
}

func (a Animal) FavFood() string {
	return a.Food
}

type AnimalCat struct {
	Animal
}

func newCat() IAnimal {
	return &AnimalCat{
		Animal{
			Sound: "meong",
			Food:  "ikan",
		},
	}
}

type AnimalDog struct {
	Animal
}

func newDog() IAnimal {
	return &AnimalDog{
		Animal{
			Sound: "guk guk",
			Food:  "tulang",
		},
	}
}

func adopt(animalType AnimalType) IAnimal {
	switch animalType {
	case Cat:
		return newCat()
	case Dog:
		return newDog()
	}
	return nil
}

func main() {
	cat := adopt(Cat)
	dog := adopt(Dog)

	fmt.Println(dog.Talk())
	fmt.Println(cat.Talk())

	fmt.Println(dog.FavFood())
	fmt.Println(cat.FavFood())
}
