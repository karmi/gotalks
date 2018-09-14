package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
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
	p := Person{"John"}

	// 1. STDOUT
	fmt.Println(">")
	p.SayTo(os.Stdout, "Gophers!") // HL

	// 2. FILE
	f, _ := os.OpenFile("/tmp/say.txt", os.O_RDWR|os.O_CREATE, 0755)
	defer f.Close()
	p.SayTo(f, "File!") // HL

	fmt.Printf("\n$ cat %s\n", f.Name())
	out, _ := exec.Command("cat", "/tmp/say.txt").Output()
	fmt.Printf("%s\n", out)

	// 3. HTTP
	addr := "localhost:4000"
	fmt.Println("http://" + addr)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p.SayTo(w, "HTTP!") // HL
	})

	http.ListenAndServe(addr, nil)
	// END3 OMIT
}
