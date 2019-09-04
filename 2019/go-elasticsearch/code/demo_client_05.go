// +build ignore

package main

import (
	"log"
	"os"

	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
)

// START1 OMIT
import "github.com/elastic/go-elasticsearch/esutil"

// END1 OMIT

func main() {
	log.SetFlags(0)
	log.SetOutput(os.Stdout)

	var (
		res *esapi.Response
		err error
	)

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	// START2 OMIT
	res, _ = es.Index(
		"test",
		esutil.NewJSONReader(map[string]interface{}{"foo": "bar"}), // HL
	)
	// END2 OMIT

	defer res.Body.Close()
	log.Println(res)

	// START3 OMIT
	type Doc struct { // HL
		Title string `json:"title"` // HL
	} // HL

	doc := Doc{Title: "Test"} // HL

	res, _ = es.Index(
		"test",
		esutil.NewJSONReader(&doc), // HL
	)
	// END3 OMIT

	defer res.Body.Close()
	log.Println(res)
}
