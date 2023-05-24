package main

import (
	"database/sql"
	"fmt"
	"time"
)

type EdiblesData struct {
	Data DataSet2 `json:"data"`
}

type DataSet2 struct {
	FilteredProducts Product2 `json:"filteredProducts"`
}

type Product2 struct {
	Products []EdiblesProducts `json:"products"`
}

type EdiblesProducts struct {
	Id                         string                 `json:"_id"`
	BrandName                  string                 `json:"brandName"`
	CBD                        string                 `json:"CBD"`
	CBDContent                 CBDContent2            `json:"CBDContent"`
	ComingSoon                 string                 `json:"comingSoon"`
	CreatedAt                  string                 `json:"createdAt"`
	DispensaryID               string                 `json:"DispensaryID"`
	EnterpriseProductId        string                 `json:"enterpriseProductId"`
	Image                      string                 `json:"Image"`
	Measurements               measurements2          `json:"measurements"`
	MedicalOnly                bool                   `json:"medicalOnly"`
	MedicalPrice               []float64              `json:"medicalPrices"`
	MedicalSpecialPrice        []float64              `json:"medicalSpecialPrices"`
	WholeSalePrice             []float64              `json:"wholesalePrices"`
	ProductName                string                 `json:"Name"`
	NonArmsLength              string                 `json:"nonArmsLength"`
	WeightOptions              []string               `json:"Options"`
	CustomerLimit              string                 `json:"limitsPerCustomer"`
	ManualInventory            string                 `json:"manualInventory"`
	POSMetaData                MetaData2              `json:"POSMetaData"`
	Price                      []float64              `json:"Prices"`
	PricingTierData            string                 `json:"pricingTierData"`
	RecOnly                    bool                   `json:"recOnly"`
	RecPrice                   []float64              `json:"recPrices"`
	RecSpecialPrice            []float64              `json:"recSpecialPrices"`
	Special                    bool                   `json:"special"`
	SpecialData                []string               `json:"specialData"`
	Status                     string                 `json:"Status"`
	StrainType                 string                 `json:"strainType"`
	SubCategory                string                 `json:"subCategory"`
	THC                        string                 `json:"THC"`
	THCContent                 THCContent2            `json:"THCContent"`
	Type                       string                 `json:"type"`
	VapeTaxApplicable          bool                   `json:"vapeTaxApplicable"`
	Weight                     float64                `json:"weight"`
	Featured                   string                 `json:"featured"`
	IsBelowThreshold           bool                   `json:"isBelowThreshold"`
	IsBelowKioskThreshold      bool                   `json:"isBelowKioskThreshold"`
	OptionsBelowThreshold      []string               `json:"optionsBelowThreshold"`
	OptionsBelowKioskThreshold []string               `json:"optionsBelowKioskThreshold"`
	CName                      string                 `json:"cName"`
	PastCNames                 []string               `json:"pastCNames"`
	BrandLogo                  string                 `json:"brandLogo"`
	BottleDepositTaxCents      string                 `json:"bottleDepositTaxCents"`
	Typename                   string                 `json:"__typename"`
	BrandInfo                  BrandInfo2             `json:"brand"`
	Effects                    Effects2               `json:"effects"`
	CannabinoidsContent        []CannabinoidsContent2 `json:"cannabinoidsV2"`
}
type THCContent2 struct {
	Unit     string    `json:"unit"`
	Range    []float64 `json:"range"`
	Typename string    `json:"__typename"`
}
type CBDContent2 struct {
	Unit     string    `json:"unit"`
	Range    []float64 `json:"range"`
	Typename string    `json:"__typename"`
}

type CannabinoidsContent2 struct {
	Value           float64          `json:"value"`
	Unit            string           `json:"PERCENTAGE"`
	CannabinoidType CannabinoidType2 `json:"cannabinoid"`
	TypeName        string           `json:"__typename"`
}

type CannabinoidType2 struct {
	Name     string `json:"name"`
	TypeName string `json:"__typename"`
}

type measurements2 struct {
	NetWeight weight2 `json:"netWeight"`
	Volume    string  `json:"volume"`
	Typename  string  `json:"__typename"`
}

type weight2 struct {
	Unit     string `json:"unit"`
	Value    []int  `json:"values"`
	Typename string `json:"__typename"`
}

type MetaData2 struct {
	CanonicalID        string                `json:"canonicalID"`
	CanonicalBrandName string                `json:"canonicalBrandName"`
	Children           []AdditionalMetaData2 `json:"children"`
	TypeName           string                `json:"__typename"`
}

type AdditionalMetaData2 struct {
	CanonicalEffectivePotencyMg float64           `json:"canonicalEffectivePotencyMg"`
	SizeOptions                 string            `json:"option"`
	InventoryQuantity           int               `json:"quantity"`
	InventoryQuantityAvailable  int               `json:"quantityAvailable"`
	InventoryKioskQuantity      int               `json:"kioskQuantityAvailable"`
	StandardEquivalent          WeightEquivalent2 `json:"standardEquivalent"`
}

type BrandInfo2 struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	ImageURL    string `json:"imageURL"`
	Name        string `json:"name"`
}

type Effects2 struct {
	RelaxedRating    float64 `json:"relaxed"`
	PainReliefRating float64 `json:"painRelief"`
	SleepyRating     float64 `json:"sleepy"`
	HappyRating      float64 `json:"happy"`
	EuphoricRating   float64 `json:"euphoric"`
}

type WeightEquivalent2 struct {
	Value    float64 `json:"value"`
	Unit     string  `json:"unit"`
	TypeName string  `json:"__typename"`
}

func InsertDataProductEdiblesCommon(db *sql.DB, target EdiblesData) {

	for i := range target.Data.FilteredProducts.Products {

		stmt, err := db.Prepare("INSERT INTO Product_EdiblesCommon (EffDate, DispensaryID, EnterpriseProductID, ProductID, BrandName, ProductName, ProductSubCategory, Weight, WeightGrams, RecOnly, RecPrice, RecSpecialPrice, MedicalOnly, MedicalPrice, MedicalSpecialPrice, WholeSalePrice, InventoryQuantity, InventoryQuantityAvail, InventoryQuantityKiosk, StrainType,  THCAmount, THCUnit, THCTypeName, CBDAMOUNT, CBDUnit, CBDTypeName, RexaledRating, PainReliefRating, SleepyRating, HappyRating, EuphoricRating, ProductImage) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
		if err != nil {
			fmt.Println(err)
			return
		}

		defer stmt.Close()

		currentDate := time.Now().Local()

		effDate := currentDate.Format("2006-01-02")

		var thcAmount, cbdAmount float64
		if len(target.Data.FilteredProducts.Products[i].THCContent.Range) > 0 {
			thcAmount = target.Data.FilteredProducts.Products[i].THCContent.Range[0]
		}
		if len(target.Data.FilteredProducts.Products[i].CBDContent.Range) > 0 {
			cbdAmount = target.Data.FilteredProducts.Products[i].CBDContent.Range[0]
		}

		for j := range target.Data.FilteredProducts.Products[i].WeightOptions {

			weight1 := target.Data.FilteredProducts.Products[i].WeightOptions[j]

			var WeightG float64

			gramCheck := CheckLastCharacter(weight1)

			if getLastTwoCharacters(weight1) == "mg" {

				strGrams := grabAllExceptLastTwo(weight1)

				WeightG, err = ConvertToFloat(strGrams)

				if err != nil {
					fmt.Println("Issue converting gram value to float64 --> MG to G")
				}

				WeightG = (WeightG / 1000)

			} else if CheckLastCharacter(gramCheck) == "g" {

				strGrams := GrabRestOfString(weight1)
				WeightG, err = ConvertToFloat(strGrams)

				if err != nil {
					fmt.Println("Issue converting gram value to float64 --> G to G")
				}

			} else if getLastTwoCharacters(weight1) == "oz" {
				if weight1 == "1/8oz" {
					WeightG = 3.5
				} else if weight1 == "1/4oz" {
					WeightG = 7.0
				} else if weight1 == "1/2oz" {
					WeightG = 14.0
				} else if weight1 == "1oz" {
					WeightG = 28.0
				}

			} else {

				fmt.Println("Could not convert " + weight1 + " to grams")
			}

			var recSpecialPrice, medSpecialPrice, recPrice, medPrice, wholePrice float64

			if len(target.Data.FilteredProducts.Products[i].RecPrice) > 0 {
				recPrice = target.Data.FilteredProducts.Products[i].RecPrice[j]

			}

			if len(target.Data.FilteredProducts.Products[i].MedicalPrice) > 0 {
				medPrice = target.Data.FilteredProducts.Products[i].MedicalPrice[j]

			}

			if len(target.Data.FilteredProducts.Products[i].WholeSalePrice) > 0 {
				wholePrice = target.Data.FilteredProducts.Products[i].WholeSalePrice[j]

			}

			if len(target.Data.FilteredProducts.Products[i].RecSpecialPrice) > 0 {
				recSpecialPrice = target.Data.FilteredProducts.Products[i].RecSpecialPrice[j]

			}

			if len(target.Data.FilteredProducts.Products[i].MedicalSpecialPrice) > 0 {

				medSpecialPrice = target.Data.FilteredProducts.Products[i].MedicalSpecialPrice[j]
			}

			var InventoryQuantity, InventoryQuantityAvailable, InventoryKioskQuantity int
			if len(target.Data.FilteredProducts.Products[i].POSMetaData.Children) > 0 {

				InventoryQuantity = target.Data.FilteredProducts.Products[i].POSMetaData.Children[j].InventoryQuantity
				InventoryQuantityAvailable = target.Data.FilteredProducts.Products[i].POSMetaData.Children[j].InventoryQuantityAvailable
				InventoryKioskQuantity = target.Data.FilteredProducts.Products[i].POSMetaData.Children[j].InventoryKioskQuantity
			}

			// Execute the INSERT statement with values
			_, err = stmt.Exec(effDate, target.Data.FilteredProducts.Products[i].DispensaryID, target.Data.FilteredProducts.Products[i].EnterpriseProductId, target.Data.FilteredProducts.Products[i].Id, target.Data.FilteredProducts.Products[i].BrandName, target.Data.FilteredProducts.Products[i].ProductName, target.Data.FilteredProducts.Products[i].SubCategory, weight1, WeightG, target.Data.FilteredProducts.Products[i].RecOnly, recPrice, recSpecialPrice, target.Data.FilteredProducts.Products[i].MedicalOnly, medPrice, medSpecialPrice, wholePrice, InventoryQuantity, InventoryQuantityAvailable, InventoryKioskQuantity, target.Data.FilteredProducts.Products[i].StrainType, thcAmount, target.Data.FilteredProducts.Products[i].THCContent.Unit, target.Data.FilteredProducts.Products[i].THCContent.Typename, cbdAmount, target.Data.FilteredProducts.Products[i].CBDContent.Unit, target.Data.FilteredProducts.Products[i].CBDContent.Typename, target.Data.FilteredProducts.Products[i].Effects.RelaxedRating, target.Data.FilteredProducts.Products[i].Effects.PainReliefRating, target.Data.FilteredProducts.Products[i].Effects.SleepyRating, target.Data.FilteredProducts.Products[i].Effects.HappyRating, target.Data.FilteredProducts.Products[i].Effects.EuphoricRating, target.Data.FilteredProducts.Products[i].Image)
			if err != nil {
				fmt.Println(err)
				return
			}

		}

	}
}

func getEdiblesData(db *sql.DB, dispensaryIDs []string, dispensaryTargets []EdiblesData) {

	for i, Dispensary := range dispensaryIDs {

		DispensaryAPILink := "https://dutchie.com/graphql?operationName=FilteredProducts&variables=%7B%22includeEnterpriseSpecials%22%3Afalse%2C%22includeCannabinoids%22%3Atrue%2C%22productsFilter%22%3A%7B%22dispensaryId%22%3A%22" + Dispensary + "%22%2C%22pricingType%22%3A%22med%22%2C%22strainTypes%22%3A%5B%5D%2C%22subcategories%22%3A%5B%5D%2C%22Status%22%3A%22Active%22%2C%22types%22%3A%5B%22Edible%22%5D%2C%22useCache%22%3Afalse%2C%22sortDirection%22%3A1%2C%22sortBy%22%3A%22weight%22%2C%22isDefaultSort%22%3Atrue%2C%22bypassOnlineThresholds%22%3Afalse%2C%22isKioskMenu%22%3Afalse%2C%22removeProductsBelowOptionThresholds%22%3Atrue%7D%2C%22page%22%3A0%2C%22perPage%22%3A1000%7D&extensions=%7B%22persistedQuery%22%3A%7B%22version%22%3A1%2C%22sha256Hash%22%3A%22cf45528052df8664bf9404f98344719a5d7c6f8fff90e8c3c8ef2eaca63b0e9e%22%7D%7D"

		err := getJson(DispensaryAPILink, Dispensary+"EdiblesData", &dispensaryTargets[i])

		if err != nil {
			fmt.Println(err)
		}
	}

	for _, EdiblesMenu := range dispensaryTargets {
		InsertDataProductEdiblesCommon(db, EdiblesMenu)

	}

}

var EdiblesDataSlice = []EdiblesData{
	legalgreensbrocktonEdibles, commonwealthalternativecarebrocktonrecEdibles, commonwealthalternativecarebrocktonEdibles, atlanticmedicinalpartnersinc1Edibles, ingoodhealthmedicalEdibles, ingoodhealthbrocktonEdibles, panaceawellnessquincyEdibles, cannapidispensaryEdibles, brocktongreenheartEdibles, boterabrocktonEdibles, budsgoodsandprovisionsabingtonEdibles, cannavanaEdibles, nativesunnorthattlebororetailrecEdibles, yambaEdibles, budsgoodsandprovisionswatertownEdibles, siranaturalsEdibles, mistymountainshopmaldenEdibles, gardenremediesmelroserecEdibles, gardenremediesmelroseEdibles, oldeworldremediesEdibles, naturesmedicinesfallriverEdibles, goodchemistrylynnEdibles, diemlynnEdibles, calyxpeakcompaniesdbalocalcannabiscoEdibles, newleafenterprisesEdibles, insasalemEdibles, insaavon1Edibles, rhelmcannabisEdibles, theorywellnessdeliveryzone1Edibles, theorywellnessbridgewaterEdibles, alternativecompassionservicesEdibles, thehealthcircleEdibles, greenrockcannabisEdibles, releafalternativesEdibles, ermontincEdibles, greathitscannabisretailrectautonEdibles, curaleafmahanoverEdibles, panaceawellnessrecEdibles, panaceawellnessEdibles, quincycannabisquincyretailrecEdibles, commonwealthalternativecareEdibles, commonwealthalternativecare1Edibles, terpsattleboroEdibles, goodchemistryEdibles, goodchemistrymedEdibles, bloom1Edibles, diemworcesterEdibles, cookiesworcesterEdibles, mayEdiblesworcesterEdibles, budsgoodsandprovisionsworcesterEdibles, resinateworcesterEdibles, missionworcesterEdibles, missionworcesteradultuseEdibles, clearskyworcesterEdibles, greencarecollectiveretailmedEdibles, thevault1Edibles, naturesremedymedEdibles, naturesremedy1Edibles, harmonyofmaEdibles, campfirecannabis1Edibles, discerndcannabisgraftonretailEdibles, curaleafmaoxfordEdibles, curaleafmaoxford1Edibles, kosaEdibles, greengoldEdibles, societycannabiscoEdibles, newenglandharvestEdibles, thevaultwebsterEdibles, terpscharltonEdibles, temescalwellnesshudsonEdibles, hudsonadultuseretailEdibles, capitalcanabiscommunitydispensaryEdibles, dmaholdingsllcdbagreatesthitscannabisEdibles, gardenremediesmarlboroughEdibles, gardenremediesrecmarlboroughEdibles, greenmeadowsfarmEdibles, cannabrollcdbagreenpathsouthbridgeEdibles, GreenGoldMarboroughEdibles, localrootssturbridgeEdibles, greenngoEdibles, blackstonevalleycannabisEdibles, highhopesEdibles, greeneraEdibles, thriveofmadbaboomxshirleyEdibles, ethosfitchburgEdibles, atlanticmedicinalpartnersincEdibles, atlanticmedicinalpartnersEdibles, greenmeadowfarmfitchburgEdibles, pioneercannabiscompanyEdibles, naturesmedicinesuxbridgeEdibles, uptopframinghamrecEdibles, uptopframinghamEdibles, masswellspringmaynardEdibles, bountifulfarmsEdibles, bleafwellnesscentreEdibles, thehealingcenter1Edibles, temescalwellnessframinghamEdibles, botrealtyllcdbaomgcannabisEdibles, curaleafmawareEdibles, localrootsfitchburgEdibles, ledlloyd420Edibles, boterafranklin1Edibles, masswellspringactonEdibles, ddmcannabisEdibles, centralavecompassionatecareEdibles, gagecannabiscoEdibles, netafranklinmedEdibles, netafranklinEdibles, communitycarecollective1Edibles, unitedcultivationllcEdibles, budbarnwinchendonEdibles, siranaturalsneedhamEdibles, umaEdibless1Edibles, clearskybelchertownEdibles, thebostongardenEdibles, rediEdibles, elev8cannabisatholEdibles, toytownprojectllcdbatoytownhealthcareEdibles, gardenremediesnewtonEdibles, gardenremediesrecnewtonEdibles, budsgoodsandprovisionswatertownEdibles2, nativesunnorthattlebororetailrecEdibles2, auraofrhodeislandEdibles, ayrdispensarywatertownmedEdibles, ayrdispensarywatertownEdibles, communitycarecollectiveEdibles, uptopwestroxburyEdibles, orangecannabiscompanyEdibles, EthosWatertownAdultUseEdibles, EthosWatertownMedicalEdibles, apothcaarlingtonEdibles, releafalternativesEdibles2, ocsivymaerealestatedbazaharacannabisEdibles, motherearthpawtucketEdibles, kgcollectivecambridgeEdibles, revolutionaryclinicsfreshpondEdibles, naturesremedy2Edibles, highprofileroslindaleEdibles, mayEdiblesmedicinalsallstonrecEdibles, mayEdiblesallstonEdibles, apothcajamaicaplainEdibles, netabrooklineEdibles, netabrooklinemedEdibles, missionbrooklineEdibles, redcardinal2Edibles, libertyspringfieldEdibles, canacraftEdibles, siranaturalsEdibles2, siranaturalssomervilleadultuseEdibles, medmenfenwayEdibles, westernfront1Edibles, smythcannabiscoEdibles, massalternativecareamherstmedEdibles, massalternativecareamherstEdibles, pleasantreesamherstEdibles, libertysomervilleEdibles, theheirloomcollectiveEdibles, pureoasis1Edibles, advesamaincdbablueriverterpscambridgeEdibles, insaspringfieldEdibles, ayrdispensarybackbayEdibles, reverie73Edibles, massalternativecareEdibles, massalternativecaremedEdibles, revolutionaryclinicssomervilleEdibles, theorywellnesschicopeemedEdibles, ethosdorchesterEdibles, highprofiledorchesterEdibles, lazyriverproductsEdibles, mistymountainshopmaldenEdibles2, jimbuddysEdibles, apexnoireEdibles, cannapidispensaryEdibles2, ermontincEdibles2, dazedcannabisEdibles,
}

var legalgreensbrocktonEdibles EdiblesData
var commonwealthalternativecarebrocktonrecEdibles EdiblesData
var commonwealthalternativecarebrocktonEdibles EdiblesData
var atlanticmedicinalpartnersinc1Edibles EdiblesData
var ingoodhealthmedicalEdibles EdiblesData
var ingoodhealthbrocktonEdibles EdiblesData
var panaceawellnessquincyEdibles EdiblesData
var cannapidispensaryEdibles EdiblesData
var brocktongreenheartEdibles EdiblesData
var boterabrocktonEdibles EdiblesData
var budsgoodsandprovisionsabingtonEdibles EdiblesData
var cannavanaEdibles EdiblesData
var nativesunnorthattlebororetailrecEdibles EdiblesData
var yambaEdibles EdiblesData
var budsgoodsandprovisionswatertownEdibles EdiblesData
var siranaturalsEdibles EdiblesData
var mistymountainshopmaldenEdibles EdiblesData
var gardenremediesmelroserecEdibles EdiblesData
var gardenremediesmelroseEdibles EdiblesData
var oldeworldremediesEdibles EdiblesData
var naturesmedicinesfallriverEdibles EdiblesData
var goodchemistrylynnEdibles EdiblesData
var diemlynnEdibles EdiblesData
var calyxpeakcompaniesdbalocalcannabiscoEdibles EdiblesData
var newleafenterprisesEdibles EdiblesData
var insasalemEdibles EdiblesData
var insaavon1Edibles EdiblesData
var rhelmcannabisEdibles EdiblesData
var theorywellnessdeliveryzone1Edibles EdiblesData
var theorywellnessbridgewaterEdibles EdiblesData
var alternativecompassionservicesEdibles EdiblesData
var thehealthcircleEdibles EdiblesData
var greenrockcannabisEdibles EdiblesData
var releafalternativesEdibles EdiblesData
var ermontincEdibles EdiblesData
var greathitscannabisretailrectautonEdibles EdiblesData
var curaleafmahanoverEdibles EdiblesData
var panaceawellnessrecEdibles EdiblesData
var panaceawellnessEdibles EdiblesData
var quincycannabisquincyretailrecEdibles EdiblesData
var commonwealthalternativecareEdibles EdiblesData
var commonwealthalternativecare1Edibles EdiblesData
var terpsattleboroEdibles EdiblesData
var goodchemistryEdibles EdiblesData
var goodchemistrymedEdibles EdiblesData
var bloom1Edibles EdiblesData
var diemworcesterEdibles EdiblesData
var cookiesworcesterEdibles EdiblesData
var mayEdiblesworcesterEdibles EdiblesData
var budsgoodsandprovisionsworcesterEdibles EdiblesData
var resinateworcesterEdibles EdiblesData
var missionworcesterEdibles EdiblesData
var missionworcesteradultuseEdibles EdiblesData
var clearskyworcesterEdibles EdiblesData
var greencarecollectiveretailmedEdibles EdiblesData
var thevault1Edibles EdiblesData
var naturesremedymedEdibles EdiblesData
var naturesremedy1Edibles EdiblesData
var harmonyofmaEdibles EdiblesData
var campfirecannabis1Edibles EdiblesData
var discerndcannabisgraftonretailEdibles EdiblesData
var curaleafmaoxfordEdibles EdiblesData
var curaleafmaoxford1Edibles EdiblesData
var kosaEdibles EdiblesData
var greengoldEdibles EdiblesData
var societycannabiscoEdibles EdiblesData
var newenglandharvestEdibles EdiblesData
var thevaultwebsterEdibles EdiblesData
var terpscharltonEdibles EdiblesData
var temescalwellnesshudsonEdibles EdiblesData
var hudsonadultuseretailEdibles EdiblesData
var capitalcanabiscommunitydispensaryEdibles EdiblesData
var dmaholdingsllcdbagreatesthitscannabisEdibles EdiblesData
var gardenremediesmarlboroughEdibles EdiblesData
var gardenremediesrecmarlboroughEdibles EdiblesData
var greenmeadowsfarmEdibles EdiblesData
var cannabrollcdbagreenpathsouthbridgeEdibles EdiblesData
var GreenGoldMarboroughEdibles EdiblesData
var localrootssturbridgeEdibles EdiblesData
var greenngoEdibles EdiblesData
var blackstonevalleycannabisEdibles EdiblesData
var highhopesEdibles EdiblesData
var greeneraEdibles EdiblesData
var thriveofmadbaboomxshirleyEdibles EdiblesData
var ethosfitchburgEdibles EdiblesData
var atlanticmedicinalpartnersincEdibles EdiblesData
var atlanticmedicinalpartnersEdibles EdiblesData
var greenmeadowfarmfitchburgEdibles EdiblesData
var pioneercannabiscompanyEdibles EdiblesData
var naturesmedicinesuxbridgeEdibles EdiblesData
var uptopframinghamrecEdibles EdiblesData
var uptopframinghamEdibles EdiblesData
var masswellspringmaynardEdibles EdiblesData
var bountifulfarmsEdibles EdiblesData
var bleafwellnesscentreEdibles EdiblesData
var thehealingcenter1Edibles EdiblesData
var temescalwellnessframinghamEdibles EdiblesData
var botrealtyllcdbaomgcannabisEdibles EdiblesData
var curaleafmawareEdibles EdiblesData
var localrootsfitchburgEdibles EdiblesData
var ledlloyd420Edibles EdiblesData
var boterafranklin1Edibles EdiblesData
var masswellspringactonEdibles EdiblesData
var ddmcannabisEdibles EdiblesData
var centralavecompassionatecareEdibles EdiblesData
var gagecannabiscoEdibles EdiblesData
var netafranklinmedEdibles EdiblesData
var netafranklinEdibles EdiblesData
var communitycarecollective1Edibles EdiblesData
var unitedcultivationllcEdibles EdiblesData
var budbarnwinchendonEdibles EdiblesData
var siranaturalsneedhamEdibles EdiblesData
var umaEdibless1Edibles EdiblesData
var clearskybelchertownEdibles EdiblesData
var thebostongardenEdibles EdiblesData
var rediEdibles EdiblesData
var elev8cannabisatholEdibles EdiblesData
var toytownprojectllcdbatoytownhealthcareEdibles EdiblesData
var gardenremediesnewtonEdibles EdiblesData
var gardenremediesrecnewtonEdibles EdiblesData
var budsgoodsandprovisionswatertownEdibles2 EdiblesData
var nativesunnorthattlebororetailrecEdibles2 EdiblesData
var auraofrhodeislandEdibles EdiblesData
var ayrdispensarywatertownmedEdibles EdiblesData
var ayrdispensarywatertownEdibles EdiblesData
var communitycarecollectiveEdibles EdiblesData
var uptopwestroxburyEdibles EdiblesData
var orangecannabiscompanyEdibles EdiblesData
var EthosWatertownAdultUseEdibles EdiblesData
var EthosWatertownMedicalEdibles EdiblesData
var apothcaarlingtonEdibles EdiblesData
var releafalternativesEdibles2 EdiblesData
var ocsivymaerealestatedbazaharacannabisEdibles EdiblesData
var motherearthpawtucketEdibles EdiblesData
var kgcollectivecambridgeEdibles EdiblesData
var revolutionaryclinicsfreshpondEdibles EdiblesData
var naturesremedy2Edibles EdiblesData
var highprofileroslindaleEdibles EdiblesData
var mayEdiblesmedicinalsallstonrecEdibles EdiblesData
var mayEdiblesallstonEdibles EdiblesData
var apothcajamaicaplainEdibles EdiblesData
var netabrooklineEdibles EdiblesData
var netabrooklinemedEdibles EdiblesData
var missionbrooklineEdibles EdiblesData
var redcardinal2Edibles EdiblesData
var libertyspringfieldEdibles EdiblesData
var canacraftEdibles EdiblesData
var siranaturalsEdibles2 EdiblesData
var siranaturalssomervilleadultuseEdibles EdiblesData
var medmenfenwayEdibles EdiblesData
var westernfront1Edibles EdiblesData
var smythcannabiscoEdibles EdiblesData
var massalternativecareamherstmedEdibles EdiblesData
var massalternativecareamherstEdibles EdiblesData
var pleasantreesamherstEdibles EdiblesData
var libertysomervilleEdibles EdiblesData
var theheirloomcollectiveEdibles EdiblesData
var pureoasis1Edibles EdiblesData
var advesamaincdbablueriverterpscambridgeEdibles EdiblesData
var insaspringfieldEdibles EdiblesData
var ayrdispensarybackbayEdibles EdiblesData
var reverie73Edibles EdiblesData
var massalternativecareEdibles EdiblesData
var massalternativecaremedEdibles EdiblesData
var revolutionaryclinicssomervilleEdibles EdiblesData
var theorywellnesschicopeemedEdibles EdiblesData
var ethosdorchesterEdibles EdiblesData
var highprofiledorchesterEdibles EdiblesData
var lazyriverproductsEdibles EdiblesData
var mistymountainshopmaldenEdibles2 EdiblesData
var jimbuddysEdibles EdiblesData
var apexnoireEdibles EdiblesData
var cannapidispensaryEdibles2 EdiblesData
var ermontincEdibles2 EdiblesData
var dazedcannabisEdibles EdiblesData
