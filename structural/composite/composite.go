package composite

import "fmt"

type Athlete struct{}

func (a *Athlete) Train() {
	fmt.Println("Training")
}

type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim func()
}

func Swim() {
	fmt.Println("Swaming")
}


type Animal struct {}

func (r *Animal)Eat() {
	println("Eating")
}

type Shark struct {
	Animal
	Swim func()
}
