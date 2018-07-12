package main

import (
	"fmt"
)

type Person struct {
	name string
}

func main() {
	p := Person{name: "Ashley"}

	nonPointer(p)
	fmt.Println(p)

	pointer(&p)
	fmt.Println(p)

	p = nonPointer2(p)
	fmt.Println(p)
}

func nonPointer(p Person) {
	p.name = "Donkey"
}

func pointer(p *Person) {
	p.name = "Moose"
}

func nonPointer2(p Person) Person {
	p.name = "Cow"
	return p
}
