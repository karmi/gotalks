package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// START1 OMIT
type Document struct {
	Name    string
	Content io.Reader
}

func Save(w io.Writer, doc *Document) error {
	_, err := io.Copy(w, doc.Content) // HL
	return err
}

// END1 OMIT

func main() {
	// START2 OMIT
	var b bytes.Buffer

	doc := &Document{Name: "Test", Content: strings.NewReader("HELLO")}

	Save(&b, doc) // HL

	fmt.Printf("%T: %q\n", b, b.String())
	// END2 OMIT
}
