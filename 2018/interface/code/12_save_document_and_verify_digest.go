package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"strings"
)

type Document struct {
	Name    string
	Digest  string
	Content io.Reader
}

// START1 OMIT
func Save(w io.Writer, doc *Document) error {
	bw := bytes.NewBuffer([]byte{})     // HL
	tr := io.TeeReader(doc.Content, bw) // HL
	dw := sha256.New()                  // HL

	io.Copy(dw, tr) // HL

	digest := fmt.Sprintf("%x", dw.Sum(nil))

	if digest != doc.Digest {
		return fmt.Errorf("failed to save document: digest mismatch, %q != %q", digest, doc.Digest)
	}

	_, err := io.Copy(w, bw)
	return err
}

// END1 OMIT

func main() {
	// START2 OMIT
	doc1 := &Document{Name: "Test", Content: strings.NewReader("HELLO\n"), Digest: "foobar"}
	err := Save(os.Stdout, doc1) // HL

	fmt.Fprintf(os.Stderr, "Error: %v\n", err)

	doc2 := &Document{Name: "Test", Content: strings.NewReader("HELLO\n"), Digest: "3b09aeb6f5f5336beb205d7f720371bc927cd46c21922e334d47ba264acb5ba4"}
	err = Save(os.Stdout, doc2) // HL

	fmt.Printf("Error: %v\n", err)

	// END2 OMIT
}
