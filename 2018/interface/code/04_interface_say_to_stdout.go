package main

import (
	"fmt"
	"io"
	"os"
)

// START1 OMIT
type Speaker interface {
	Say(string)
	SayTo(io.Writer, string) // HL
}

// END1 OMIT

type Person struct {
	Name string
}

// START2 OMIT
func (p *Person) Say(msg string) {
	fmt.Println(p.statement(msg))
}

func (p *Person) SayTo(w io.Writer, msg string) { // HL
	fmt.Fprint(w, p.statement(msg))
}

func (p *Person) statement(msg string) string { return p.Name + " says: " + msg + "\n" }

// END2 OMIT

// type Robot struct {
// 	Name string
// }

// func (r *Robot) Say(msg string) {
// 	fmt.Println(r.statement(msg))
// }

// func (r *Robot) SayTo(w io.Writer, msg string) {
// 	fmt.Fprint(w, r.statement(msg)+"\n")
// }

// func (r *Robot) statement(msg string) string {
// 	return r.Name + " says: " + hex.EncodeToString([]byte(msg))
// }

func main() {
	// START3 OMIT
	p := Person{"John"}
	p.Say("Hello!")
	p.SayTo(os.Stdout, "Gophers!")
	// END3 OMIT
}
