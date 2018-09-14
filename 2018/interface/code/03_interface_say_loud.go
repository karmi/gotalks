package main

import (
	"encoding/hex"
	"fmt"
	"strings"
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

type Robot struct {
	Name string
}

func (r *Robot) Say(msg string) {
	fmt.Println(r.Name + " says: " + hex.EncodeToString([]byte(msg)))
}

// START1 OMIT
// # Person and Robot implementation omitted

func SayLoud(speaker Speaker, msg string) { // HL
	fmt.Println(strings.Repeat("!", 80))
	speaker.Say(msg)
	fmt.Println(strings.Repeat("!", 80))
	fmt.Println()
	// os.Exec "say ..."
}

// END1 OMIT

func main() {
	// START2 OMIT
	p := Person{"John"}
	SayLoud(&p, "Hello Gophers!") // HL

	r := Robot{"R2D2"}
	SayLoud(&r, "Hello Gophers!") // HL
	// END2 OMIT
}
