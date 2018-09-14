package main

import (
	"encoding/hex"
	"fmt"
)

type Speaker interface {
	Say(string)
}

type Person struct {
	Name string
}

func (p *Person) Say(msg string) {
	fmt.Println(p.Name + " says: " + msg)
}

// START1 OMIT
// # Person implementation omitted

type Robot struct {
	Name string
}

func (r *Robot) Say(msg string) {
	fmt.Println(r.Name + " says: " + hex.EncodeToString([]byte(msg))) // HL
}

// END1 OMIT

func main() {
	// START2 OMIT
	p := Person{"John"}
	p.Say("Hello Gophers!")

	r := Robot{"R2D2"}
	r.Say("Hello Gophers!")
	// END2 OMIT
}
