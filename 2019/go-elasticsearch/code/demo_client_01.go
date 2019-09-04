// +build ignore

package main

import (
	"log"
	"os"
)

// START1 OMIT
import "github.com/elastic/go-elasticsearch"

// END1 OMIT

func main() {
	log.SetFlags(0)
	log.SetOutput(os.Stdout)

	// START2 OMIT
	// Defaults
	//
	// Note: the ELASTICSEARCH_URL environment variable is used when exported
	//
	es, err := elasticsearch.NewDefaultClient() // HL
	// END2 OMIT

	// START3 OMIT
	// Configuration
	//
	cfg := elasticsearch.Config{
		Addresses: []string{"https://example.com"},
		Username:  "foo",
		Password:  "bar",
	}

	es, err = elasticsearch.NewClient(cfg) // HL
	// END3 OMIT

	es, err = elasticsearch.NewDefaultClient()

	// START4 OMIT
	res, err := es.Cluster.Health( // HL
		es.Cluster.Health.WithLevel("indices"), // HL
		es.Cluster.Health.WithPretty(),         // HL
	) // HL
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	defer res.Body.Close()
	log.Println(res)
	// END4 OMIT
}
