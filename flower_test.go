package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
	_ "time"
)

// Take written JSON files from main and test those instead of making a new get request
// iterate through testFiles folder and return error with the name of the file that failed

var testflower productData

func TestProductIDs(t *testing.T) {

	jsonFile, err := os.Open("LegalGreensTest.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	err = json.NewDecoder(jsonFile).Decode(&testflower)

	if err != nil {
		t.Fail()
		log.Fatal(err)
	}

	defer jsonFile.Close()

}
