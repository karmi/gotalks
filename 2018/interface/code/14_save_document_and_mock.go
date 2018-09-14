package main

import (
	"fmt"
	"io"
	"strings"
)

type Document struct {
	Name    string
	Content io.Reader
}

func Save(w io.Writer, doc *Document) error {
	_, err := io.Copy(w, doc.Content) // HL
	return err
}

// START1 OMIT

type ErrorWriter struct{} // HL

func (w ErrorWriter) Write(b []byte) (int, error) {
	return 0, fmt.Errorf("MOCK ERROR") // HL
}

// END1 OMIT

func main() {
	// START2 OMIT
	doc := &Document{Name: "Test", Content: strings.NewReader("HELLO")}
	err := Save(ErrorWriter{}, doc) // HL

	if err != nil {
		fmt.Printf("FAIL: %v\n", err)
	} else {
		fmt.Println("OK")
	}

	// END2 OMIT
}
