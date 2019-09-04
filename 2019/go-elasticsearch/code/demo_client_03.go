// +build ignore

package main

import (
	"context"
	"log"
	"os"
	"strings"
)

// START1 OMIT
import (
	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi" // HL
)

// END1 OMIT

func main() {
	log.SetFlags(0)
	log.SetOutput(os.Stdout)

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	// START2 OMIT
	req := esapi.IndexRequest{ // HL
		Index:      "test",                             // HL
		DocumentID: "1",                                // HL
		Body:       strings.NewReader(`{"foo":"bar"}`), // HL
		Pretty:     true,                               // HL
	} // HL

	res, err := req.Do(context.Background(), es) // HL
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	log.Println(res)
	// END2 OMIT
}
