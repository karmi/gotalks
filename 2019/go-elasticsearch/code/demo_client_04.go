// +build ignore

package main

import (
	"bytes"
	"log"
	"os"

	"github.com/elastic/go-elasticsearch"
)

// START1 OMIT
import "encoding/json"
import "github.com/tidwall/gjson"

// END1 OMIT

func main() {
	log.SetFlags(0)
	log.SetOutput(os.Stdout)

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	res, _ := es.Info()
	defer res.Body.Close()

	// START2 OMIT
	var d map[string]interface{}               // HL
	err = json.NewDecoder(res.Body).Decode(&d) // HL
	if err != nil {
		log.Fatalf("Error parsing the response: %s", err)
	}

	log.Println("encoding/json:", d["tagline"]) // HL
	// END2 OMIT

	res, _ = es.Info()
	defer res.Body.Close()

	// START3 OMIT
	var b bytes.Buffer   // HL
	b.ReadFrom(res.Body) // HL

	log.Println("tidwall/gjson:", gjson.GetBytes(b.Bytes(), "tagline")) // HL
	// END3 OMIT
}
