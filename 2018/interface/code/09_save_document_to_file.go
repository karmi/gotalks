package main

import (
	"fmt"
	"io"
	"os"
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
	f, _ := os.Create("/tmp/document.txt") // HL
	defer f.Close()

	fi, _ := f.Stat()
	fmt.Printf("File: %s [%dB]\n", fi.Name(), fi.Size())

	doc := &Document{Name: "Test", Content: strings.NewReader("HELLO")}

	Save(f, doc) // HL

	fi, _ = f.Stat()

	fmt.Printf("File: %s [%dB]\n", fi.Name(), fi.Size())
	// END2 OMIT
}
