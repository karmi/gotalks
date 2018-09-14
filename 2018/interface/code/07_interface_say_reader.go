package main

import (
	"io"
	"os"
	"strings"
)

// START1 OMIT
type Speaker interface {
	Say(string)
	SayTo(io.Writer, string)
	Statement() io.Reader // HL
}

// END1 OMIT

type Person struct {
	Name string
}

// START2 OMIT
func (p *Person) Say(msg string) { io.Copy(os.Stdout, p.Statement(msg)) }

func (p *Person) SayTo(w io.Writer, msg string) { io.Copy(w, p.Statement(msg)) }

func (p *Person) Statement(msg string) io.Reader { // HL
	return strings.NewReader(p.Name + " says: " + msg + "\n")
}

// END2 OMIT

func main() {
	// START3 OMIT
	p := Person{"John"}

	p.Say("Hello, streams!")
	// END3 OMIT
}
