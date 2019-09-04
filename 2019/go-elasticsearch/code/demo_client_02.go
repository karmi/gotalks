// +build ignore

package main

import (
	"log"
	"os"
	"strings"

	"github.com/elastic/go-elasticsearch"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(os.Stdout)

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	// START1 OMIT
	res, err := es.Index( // HL
		"test",                             // HL
		strings.NewReader(`{"foo":"bar"}`), // HL
		es.Index.WithDocumentID("1"),       // HL
		es.Index.WithPretty(),              // HL
	) // HL
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	defer res.Body.Close()
	log.Println(res)
	// END1 OMIT
}
