package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

type Document struct {
	Name    string
	Content io.Reader
}

func Save(w io.Writer, doc *Document) error {
	_, err := io.Copy(w, doc.Content)
	return err
}

func main() {
	// START2 OMIT
	doc := &Document{Name: "Test", Content: strings.NewReader("HELLO")}
	err := Save(ioutil.Discard, doc) // HL

	fmt.Printf("Error: %v\n", err)
	// END2 OMIT
}
