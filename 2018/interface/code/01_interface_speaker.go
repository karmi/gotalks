package main

import (
	"fmt"
)

// START1 OMIT
type Speaker interface {
	Say(string) // HL
}

type Person struct {
	Name string
}

func (p *Person) Say(msg string) { // HL
	fmt.Println(p.Name + " says: " + msg)
}

// END1 OMIT

func main() {
	// START2 OMIT
	p := Person{"John"}
	p.Say("Hello Gophers!")
	// END2 OMIT
}
