package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func getJson(url string, dispensaryName string, target interface{}) error {
	resp, err := client.Get(url)

	if err != nil {
		return fmt.Errorf("ERROR")
	}

	defer resp.Body.Close()

	jsonData := json.NewDecoder(resp.Body).Decode(target)

	saveJsontoFile(dispensaryName, target)

	return jsonData

}

func saveJsontoFile(filename string, target interface{}) {

	//err := getJson(LegalGreensLink, &flower)

	jsonMarshal, err := json.Marshal(target)
	jsonString := string(jsonMarshal)

	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("jsonFiles/"+filename+".json", []byte(jsonString), 0666)
	if err != nil {
		log.Fatal(err)
	}

}
