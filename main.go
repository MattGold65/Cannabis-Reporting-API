package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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
var LegalGreens ProductData
var CommonwealthAltcareRec ProductData

func main() {

	client = &http.Client{Timeout: 10 * time.Second}

	db, err := sql.Open("mysql", "root:password1@tcp(127.0.0.1:3306)/MYSQL80")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	fmt.Println("Great Success")

	//fmt.Println(flower.Data.FilteredProducts.Products[0].ProductName)

	err = getJson(CommonwealthAltcareRecLink, "CommonwealthAltcareRec", &CommonwealthAltcareRec)

	fmt.Println(CommonwealthAltcareRec.Data.FilteredProducts.Products[0].ProductName)

	err = getJson(LegalGreensLink, "LegalGreens", &LegalGreens)

	fmt.Println(LegalGreens.Data.FilteredProducts.Products[0].ProductName)
}

// find a way to combine saveJsontoFile with getJson
// everytime we grab a json from the web in main we want to write a file so that we can test the json
// will need to find a way to overwrite files rn access gets denied on rewrite attempt
