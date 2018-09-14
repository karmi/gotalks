package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type Document struct {
	Name    string
	Digest  string
	Content io.Reader
}

// Implement http.ResponseWriter

type RWriter struct {
	w io.Writer
}

func (w RWriter) Header() http.Header { return http.Header{} }

func (w RWriter) WriteHeader(s int) {}

func (w RWriter) Write(b []byte) (int, error) {
	return w.w.Write(b)
}

// START1 OMIT
func Save(w io.Writer, doc *Document) error {
	maxBytes := int64(1 * 1024 * 1042)                                                // 1MB
	br := http.MaxBytesReader(RWriter{w: w}, ioutil.NopCloser(doc.Content), maxBytes) // HL
	bw := bytes.NewBuffer([]byte{})                                                   // HL

	_, err := io.Copy(bw, br) // HL
	if err != nil {
		return fmt.Errorf("document content too big: > %dMB", maxBytes/1024/1024)
	}

	_, err = io.Copy(w, bw)
	return err
}

// END1 OMIT

func main() {
	// START2 OMIT
	c, _ := ioutil.ReadFile("/usr/share/dict/words")
	fmt.Printf("Content size: ~ %dMB\n", len(c)/1024/1024)

	doc := &Document{Name: "Test", Content: bytes.NewBuffer(c)}
	err := Save(ioutil.Discard, doc) // HL

	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	// END2 OMIT
}
