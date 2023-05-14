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
var CannaVana ProductData
var NativeSunNorthAttleboro ProductData
var Yamba ProductData
var BudsGoodsProvisionsWatertown ProductData
var SiraNaturalsSomervilleMedical ProductData
var MistyMountainShopMalden ProductData
var GardenRemediesRecMelrose ProductData
var GardenRemediesMedMelrose ProductData
var OldeWorldRemedies ProductData
var NaturesMedicinesFallRiver ProductData
var GoodChemistryLynn ProductData
var DiemLynn ProductData
var LocalCannabisCoSwampscott ProductData
var NewLeafFallRiver ProductData
var INSAIncSalemAdultUse ProductData
var InsaAvon ProductData
var RhelmCannabis ProductData
var TheoryWellnessBostonAreaDelivery ProductData
var TheoryWellnessBridgewater ProductData
var AlternativeCompassionServices ProductData
var HealthCircle ProductData
var GreenRockCannabis ProductData
var ReleafAlternatives ProductData
var ErmontInc ProductData
var GreatestHitsCannabisTaunton ProductData
var CuraleafMAHanover ProductData
var PanaceaWellnessRec ProductData
var PanaceaWellness ProductData
var QuincyCannabisCo ProductData
var CommonwealthAlternativeCareTauntonMed ProductData
var CommonwealthAlternativeCareTauntonRec ProductData
var TERPSAttleboro ProductData

func main() {
	client = &http.Client{Timeout: 10 * time.Second}

	err := getJson(CommonwealthAltcareRecLink, "CommonwealthAltcareRec", &CommonwealthAltcareRec)
	err = getJson(LegalGreensLink, "LegalGreens", &LegalGreens)
	err = getJson(AtlanticMedicinalPartnersBrocktonRecLink, "AtlanticMedicinalPartners", &AtlanticMedicinalPartners)
	err = getJson(InGoodHealthMedicalLink, "InGoodHealthMedical", &InGoodHealthMedical)
	err = getJson(PanaceaWellnessQuincyLink, "PanaceaWellnessQuincy", &PanaceaWellnessQuincy)
	/*
		err = getJson(CommonwealthAlternativeCareBrocktonMedLink, "CommonwealthAlternativeCareBrocktonMed", &CommonwealthAlternativeCareBrocktonMed)
		err = getJson(CannapiDispensaryLink, "CannapiDispensary", &CannapiDispensary)
		err = getJson(BrocktonGreenHeartLink, "BrocktonGreenHeart", &BrocktonGreenHeart)
		err = getJson(GTEBrocktonLLCdbaBoteraBrocktonLink, "GTEBrocktonLLCdbaBoteraBrockton", &GTEBrocktonLLCdbaBoteraBrockton)
		err = getJson(BudsGoodsProvisionsAbingtonLink, "BudsGoodsProvisionsAbington", &BudsGoodsProvisionsAbington)
		err = getJson(CannaVanaLink, "CannaVana", &CannaVana)
		err = getJson(NativeSunNorthAttleboroLink, "NativeSunNorthAttleboro", &NativeSunNorthAttleboro)
		err = getJson(YambaLink, "Yamba", &Yamba)
		err = getJson(BudsGoodsProvisionsWatertownLink, "BudsGoodsProvisionsWatertown", &BudsGoodsProvisionsWatertown)
		err = getJson(SiraNaturalsSomervilleMedicalLink, "SiraNaturalsSomervilleMedical", &SiraNaturalsSomervilleMedical)
		err = getJson(MistyMountainShopMaldenLink, "MistyMountainShopMalden", &MistyMountainShopMalden)
		err = getJson(GardenRemediesRecMelroseLink, "GardenRemediesRecMelrose", &GardenRemediesRecMelrose)
		err = getJson(GardenRemediesMedMelroseLink, "GardenRemediesMedMelrose", &GardenRemediesMedMelrose)
		err = getJson(OldeWorldRemediesLink, "OldeWorldRemedies", &OldeWorldRemedies)
		err = getJson(NaturesMedicinesFallRiverLink, "NaturesMedicinesFallRiver", &NaturesMedicinesFallRiver)
		err = getJson(GoodChemistryLynnLink, "GoodChemistryLynn", &GoodChemistryLynn)
		err = getJson(DiemLynnLink, "DiemLynn", &DiemLynn)
		err = getJson(LocalCannabisCoSwampscottLink, "LocalCannabisCoSwampscott", &LocalCannabisCoSwampscott)
		err = getJson(NewLeafFallRiverLink, "NewLeafFallRiver", &NewLeafFallRiver)
		err = getJson(INSAIncSalemAdultUseLink, "INSAIncSalemAdultUse", &INSAIncSalemAdultUse)
		err = getJson(InsaAvonLink, "InsaAvon", &InsaAvon)
		err = getJson(RhelmCannabisLink, "RhelmCannabis", &RhelmCannabis)
		err = getJson(TheoryWellnessBostonAreaDeliveryLink, "TheoryWellnessBostonAreaDelivery", &TheoryWellnessBostonAreaDelivery)
		err = getJson(TheoryWellnessBridgewaterLink, "TheoryWellnessBridgewater", &TheoryWellnessBridgewater)
		err = getJson(AlternativeCompassionServicesLink, "AlternativeCompassionServices", &AlternativeCompassionServices)
		err = getJson(HealthCircleLink, "HealthCircle", &HealthCircle)
		err = getJson(GreenRockCannabisLink, "GreenRockCannabis", &GreenRockCannabis)
		err = getJson(ReleafAlternativesLink, "ReleafAlternatives", &ReleafAlternatives)
		err = getJson(ErmontIncLink, "ErmontInc", &ErmontInc)
		err = getJson(GreatestHitsCannabisTauntonLink, "GreatestHitsCannabisTaunton", &GreatestHitsCannabisTaunton)
		err = getJson(CuraleafMAHanoverLink, "CuraleafMAHanover", &CuraleafMAHanover)
		err = getJson(PanaceaWellnessRecLink, "PanaceaWellnessRec", &PanaceaWellnessRec)
		err = getJson(PanaceaWellnessLink, "PanaceaWellness", &PanaceaWellness)
		err = getJson(QuincyCannabisCoLink, "QuincyCannabisCo", &QuincyCannabisCo)
		err = getJson(CommonwealthAlternativeCareTauntonMedLink, "CommonwealthAlternativeCareTauntonMed", &CommonwealthAlternativeCareTauntonMed)
		err = getJson(CommonwealthAlternativeCareTauntonRecLink, "CommonwealthAlternativeCareTauntonRec", &CommonwealthAlternativeCareTauntonRec)
		err = getJson(TERPSAttleboroLink, "TERPSAttleboro", &TERPSAttleboro)

	*/

	db, err := sql.Open("mysql", "root:password1@tcp(127.0.0.1:3306)/CannabisDB")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Product_FlowerCommon (ProductID VARCHAR(255) NOT NULL, BrandName VARCHAR(255) NOT NULL, ProductName VARCHAR(255) NOT NULL, THCAmount FLOAT, THCUnit VARCHAR(255), THCTypeName VARCHAR(255), CBDAMOUNT FLOAT, CBDUnit VARCHAR(255), CBDTypeName VARCHAR(255), RexaledRating FLOAT, PainReliefRating FLOAT, SleepyRating FLOAT, HappyRating FLOAT, EuphoricRating FLOAT, PRIMARY KEY (ProductID))")
	if err != nil {
		panic(err.Error())
	}

	InsertDataProductFlowerCommon(db, CommonwealthAltcareRec)

	//fmt.Println(CommonwealthAltcareRec.Data.FilteredProducts.Products[0].CBDContent.Range[0])

	fmt.Println("Great Success")

	//fmt.Println(flower.Data.FilteredProducts.Products[0].ProductName)

	fmt.Println(CommonwealthAltcareRec.Data.FilteredProducts.Products[0].THCContent.Range)

	fmt.Println(LegalGreens.Data.FilteredProducts.Products[0].ProductName)

	fmt.Println(AtlanticMedicinalPartners.Data.FilteredProducts.Products[0].ProductName)

	fmt.Println(InGoodHealthMedical.Data.FilteredProducts.Products[0].ProductName)

	fmt.Println(PanaceaWellnessQuincy.Data.FilteredProducts.Products[0].ProductName)

}

func InsertDataProductFlowerCommon(db *sql.DB, target ProductData) {

	//test for now turn this into function where we can call this loop on every single dispensary json
	for i := range target.Data.FilteredProducts.Products {

		stmt, err := db.Prepare("INSERT INTO Product_FlowerCommon (ProductID, BrandName, ProductName, THCAmount, THCUnit, THCTypeName, CBDAMOUNT, CBDUnit, CBDTypeName,RexaledRating, PainReliefRating, SleepyRating, HappyRating, EuphoricRating) VALUES (?, ?,?, ?,?, ?,?, ?,?, ?,?, ?,?, ?)")
		if err != nil {
			fmt.Println(err)
			return
		}

		defer stmt.Close()

		var thcAmount, cbdAmount float64
		if len(target.Data.FilteredProducts.Products[i].THCContent.Range) > 0 {
			thcAmount = target.Data.FilteredProducts.Products[i].THCContent.Range[0]
		}
		if len(target.Data.FilteredProducts.Products[i].CBDContent.Range) > 0 {
			cbdAmount = target.Data.FilteredProducts.Products[i].CBDContent.Range[0]
		}

		// Execute the INSERT statement with values
		_, err = stmt.Exec(target.Data.FilteredProducts.Products[i].Id, target.Data.FilteredProducts.Products[i].BrandName, target.Data.FilteredProducts.Products[i].ProductName, thcAmount, target.Data.FilteredProducts.Products[i].THCContent.Unit, target.Data.FilteredProducts.Products[i].THCContent.Typename, cbdAmount, target.Data.FilteredProducts.Products[i].CBDContent.Unit, target.Data.FilteredProducts.Products[i].CBDContent.Typename, target.Data.FilteredProducts.Products[i].Effects.RelaxedRating, target.Data.FilteredProducts.Products[i].Effects.PainReliefRating, target.Data.FilteredProducts.Products[i].Effects.SleepyRating, target.Data.FilteredProducts.Products[i].Effects.HappyRating, target.Data.FilteredProducts.Products[i].Effects.EuphoricRating)
		if err != nil {
			fmt.Println(err)
			return
		}

	}
}

// find a way to combine saveJsontoFile with getJson
// everytime we grab a json from the web in main we want to write a file so that we can test the json
// will need to find a way to overwrite files rn access gets denied on rewrite attempt

//could be helpful good data with this link
//https://dutchie.com/graphql?operationName=ConsumerDispensaries&variables=%7B%22dispensaryFilter%22%3A%7B%22cNameOrID%22%3A%22commonwealth-alternative-care-brockton%22%7D%7D&extensions=%7B%22persistedQuery%22%3A%7B%22version%22%3A1%2C%22sha256Hash%22%3A%22f44ceb73f96c7fa3016ea2b9c6b68a5b92a62572317849bc4d6d9a724b93c83d%22%7D%7D
