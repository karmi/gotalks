package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
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
	var b bytes.Buffer // HL

	var f, _ = os.Create("/tmp/document.txt") // HL
	defer f.Close()

	doc := &Document{Name: "Test", Content: strings.NewReader("HELLO")}

	var mw = io.MultiWriter(&b, f) // HL

	Save(mw, doc) // HL

	fi, _ := f.Stat()

	fmt.Printf("%T: %q\n", b, b.String())
	fmt.Printf("File: %s [%dB]\n", fi.Name(), fi.Size())
	// END2 OMIT
}
