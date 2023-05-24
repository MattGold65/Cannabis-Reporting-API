package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

func addPrimaryKey(db *sql.DB, columnName string) error {
	// Prepare the ALTER TABLE query
	query := fmt.Sprintf("ALTER TABLE Product_EdiblesCommon ADD COLUMN `%s` INT AUTO_INCREMENT PRIMARY KEY", columnName)

	// Execute the ALTER TABLE query
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func CheckLastCharacter(str string) string {
	length := len(str)
	if length > 0 {
		lastChar := str[length-1]
		return string(lastChar)
	}
	return ""
}

func GrabRestOfString(str string) string {
	length := len(str)
	if length > 1 {
		restOfString := str[:length-1]
		return restOfString
	}
	return ""
}

func getLastTwoCharacters(str string) string {
	length := len(str)
	if length < 2 {
		return "String is too short."
	} else {
		lastTwoCharacters := str[length-2:]
		return lastTwoCharacters
	}
}

func grabAllExceptLastTwo(str string) string {
	length := len(str)
	if length < 2 {
		return "String is too short."
	} else {
		allExceptLastTwo := str[:length-2]
		return allExceptLastTwo
	}
}

func ConvertToFloat(str string) (float64, error) {
	floatValue, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, err
	}
	return floatValue, nil
}

func GenerateFlowerProductLinkName(target FlowerData) []string {
	var LinkReadyProducts []string
	for i := range target.Data.FilteredProducts.Products {
		currentProduct := target.Data.FilteredProducts.Products[i].ProductName
		//replace "-" link with " "
		currentProduct = strings.ReplaceAll(currentProduct, "-", "")
		currentProduct = strings.ReplaceAll(currentProduct, ":", "-")
		currentProduct = strings.ReplaceAll(currentProduct, ".", "-")
		currentProduct = strings.ReplaceAll(currentProduct, "  ", " ")

		//remove unwanted characters from the string
		//change this to a slice and iterate through the slice and remove each char individually
		removeChars := ":/?#[]@&='"
		for _, char := range removeChars {

			currentProduct = strings.ReplaceAll(currentProduct, string(char), "")

		}

		//replace " " with "-"
		currentProduct = strings.ReplaceAll(currentProduct, " ", "-")

		//append front of menu tag
		currentProduct = "?dtche%5Bproduct%5D=" + currentProduct

		currentProduct = strings.ToLower(currentProduct)

		LinkReadyProducts = append(LinkReadyProducts, currentProduct)

	}

	return LinkReadyProducts

}
