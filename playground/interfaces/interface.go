package main

import "fmt"

type greeter interface {
	hello()
}

type Pirate struct{}

func (p Pirate) hello() {
	fmt.Println("Ahoy, Maty")
}

type Bro struct{}

func (b Bro) hello() {
	fmt.Println("Sup, Bro")
}

func greet(g greeter) {
	g.hello()
}

func main() {
	bro := new(Bro)
	pirate := Pirate{}

	greet(bro)

	greet(pirate)
}
