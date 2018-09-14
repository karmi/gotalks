package main

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strings"
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

func main() {
	// START3 OMIT
	p := Person{"Victoria"}

	var out bytes.Buffer             // HL
	p.SayTo(&out, "Hello, Gophers!") // HL

	cmd := []string{"say", "--voice", p.Name, out.String()}

	fmt.Printf("$ %s", strings.Join(cmd, " "))
	exec.Command(cmd[0], cmd[1:]...).Output() // HL
	// END3 OMIT
}
