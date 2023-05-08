package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

/*
Next steps:
1) test every field among multiple dispensaries and fix any bugs with unmarshalling
2) figure out a way to iterate through all the products on a page
3) figure out a way to iterate though all the pages. Hopefully this can be done by manipulating the links and iterating though until a request is made unsuccessfully
4) establish a database draft
5) start making other products compatible
*/

var client *http.Client

var LegalGreensLink = "https://dutchie.com/graphql?operationName=FilteredProducts&variables=%7B%22includeEnterpriseSpecials%22%3Afalse%2C%22includeCannabinoids%22%3Atrue%2C%22productsFilter%22%3A%7B%22dispensaryId%22%3A%22605d194989577500ba7990a4%22%2C%22pricingType%22%3A%22rec%22%2C%22strainTypes%22%3A%5B%5D%2C%22subcategories%22%3A%5B%5D%2C%22Status%22%3A%22Active%22%2C%22types%22%3A%5B%22Flower%22%5D%2C%22useCache%22%3Afalse%2C%22sortDirection%22%3A1%2C%22sortBy%22%3A%22price%22%2C%22isDefaultSort%22%3Atrue%2C%22bypassOnlineThresholds%22%3Afalse%2C%22isKioskMenu%22%3Afalse%2C%22removeProductsBelowOptionThresholds%22%3Atrue%7D%2C%22page%22%3A0%2C%22perPage%22%3A1000%7D&extensions=%7B%22persistedQuery%22%3A%7B%22version%22%3A1%2C%22sha256Hash%22%3A%220e884328c01ef8ed540d4bbd27101ee58fedaf5cddda6c499169209b53fdf574%22%7D%7D"

func main() {

	client = &http.Client{Timeout: 10 * time.Second}

	db, err := sql.Open("mysql", "root:password1@tcp(127.0.0.1:3306)/MYSQL80")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	fmt.Println("Great Success")

	err = getJson(LegalGreensLink, &flower)

	fmt.Println(flower.Data.FilteredProducts.Products[0].ProductName)

	//find a way to overwrite the json file
	saveJsontoFile("LegalGreensTest")

}

// find a way to combine saveJsontoFile with getJson
// everytime we grab a json from the web in main we want to write a file so that we can test the json
// will need to find a way to overwrite files rn access gets denied on rewrite attempt
func saveJsontoFile(filename string) {

	err := getJson(LegalGreensLink, &flower)

	jsonMarshal, err := json.Marshal(flower)
	jsonString := string(jsonMarshal)

	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(filename+".json", []byte(jsonString), 066)
	if err != nil {
		log.Fatal(err)
	}

}
