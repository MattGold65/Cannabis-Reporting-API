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
var AtlanticMedicinalPartners ProductData
var InGoodHealthMedical ProductData
var PanaceaWellnessQuincy ProductData
var CommonwealthAlternativeCareBrocktonMed ProductData
var CannapiDispensary ProductData
var BrocktonGreenHeart ProductData
var GTEBrocktonLLCdbaBoteraBrockton ProductData
var BudsGoodsProvisionsAbington ProductData

func main() {
	client = &http.Client{Timeout: 10 * time.Second}

	err := getJson(CommonwealthAltcareRecLink, "CommonwealthAltcareRec", &CommonwealthAltcareRec)
	err = getJson(LegalGreensLink, "LegalGreens", &LegalGreens)
	err = getJson(AtlanticMedicinalPartnersBrocktonRecLink, "AtlanticMedicinalPartners", &AtlanticMedicinalPartners)
	err = getJson(InGoodHealthMedicalLink, "InGoodHealthMedical", &InGoodHealthMedical)
	err = getJson(PanaceaWellnessQuincyLink, "PanaceaWellnessQuincy", &PanaceaWellnessQuincy)

	err = getJson(CommonwealthAlternativeCareBrocktonMedLink, "CommonwealthAlternativeCareBrocktonMed", &CommonwealthAlternativeCareBrocktonMed)
	err = getJson(CannapiDispensaryLink, "CannapiDispensary", &CannapiDispensary)
	err = getJson(BrocktonGreenHeartLink, "BrocktonGreenHeart", &BrocktonGreenHeart)
	err = getJson(GTEBrocktonLLCdbaBoteraBrocktonLink, "GTEBrocktonLLCdbaBoteraBrockton", &GTEBrocktonLLCdbaBoteraBrockton)
	err = getJson(BudsGoodsProvisionsAbingtonLink, "BudsGoodsProvisionsAbington", &BudsGoodsProvisionsAbington)

	db, err := sql.Open("mysql", "root:password1@tcp(127.0.0.1:3306)/MYSQL80")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	fmt.Println("Great Success")

	//fmt.Println(flower.Data.FilteredProducts.Products[0].ProductName)

	fmt.Println(CommonwealthAltcareRec.Data.FilteredProducts.Products[0].ProductName)

	fmt.Println(LegalGreens.Data.FilteredProducts.Products[0].ProductName)

	fmt.Println(AtlanticMedicinalPartners.Data.FilteredProducts.Products[0].ProductName)

	fmt.Println(InGoodHealthMedical.Data.FilteredProducts.Products[0].ProductName)

	fmt.Println(PanaceaWellnessQuincy.Data.FilteredProducts.Products[0].ProductName)

}

// find a way to combine saveJsontoFile with getJson
// everytime we grab a json from the web in main we want to write a file so that we can test the json
// will need to find a way to overwrite files rn access gets denied on rewrite attempt

//could be helpful good data with this link
//https://dutchie.com/graphql?operationName=ConsumerDispensaries&variables=%7B%22dispensaryFilter%22%3A%7B%22cNameOrID%22%3A%22commonwealth-alternative-care-brockton%22%7D%7D&extensions=%7B%22persistedQuery%22%3A%7B%22version%22%3A1%2C%22sha256Hash%22%3A%22f44ceb73f96c7fa3016ea2b9c6b68a5b92a62572317849bc4d6d9a724b93c83d%22%7D%7D
