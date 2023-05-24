package main

import (
	"database/sql"
	"fmt"
	"time"
)

type ConcentratesData struct {
	Data DataSet3 `json:"data"`
}

type DataSet3 struct {
	FilteredProducts Product3 `json:"filteredProducts"`
}

type Product3 struct {
	Products []ConcentratesProducts `json:"products"`
}

type ConcentratesProducts struct {
	Id                         string                 `json:"_id"`
	BrandName                  string                 `json:"brandName"`
	CBD                        string                 `json:"CBD"`
	CBDContent                 CBDContent3            `json:"CBDContent"`
	ComingSoon                 string                 `json:"comingSoon"`
	CreatedAt                  string                 `json:"createdAt"`
	DispensaryID               string                 `json:"DispensaryID"`
	EnterpriseProductId        string                 `json:"enterpriseProductId"`
	Image                      string                 `json:"Image"`
	Measurements               measurements3          `json:"measurements"`
	MedicalOnly                bool                   `json:"medicalOnly"`
	MedicalPrice               []float64              `json:"medicalPrices"`
	MedicalSpecialPrice        []float64              `json:"medicalSpecialPrices"`
	WholeSalePrice             []float64              `json:"wholesalePrices"`
	ProductName                string                 `json:"Name"`
	NonArmsLength              string                 `json:"nonArmsLength"`
	WeightOptions              []string               `json:"Options"`
	CustomerLimit              string                 `json:"limitsPerCustomer"`
	ManualInventory            string                 `json:"manualInventory"`
	POSMetaData                MetaData3              `json:"POSMetaData"`
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
	THCContent                 THCContent3            `json:"THCContent"`
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
	BrandInfo                  BrandInfo3             `json:"brand"`
	Effects                    Effects3               `json:"effects"`
	CannabinoidsContent        []CannabinoidsContent3 `json:"cannabinoidsV2"`
}
type THCContent3 struct {
	Unit     string    `json:"unit"`
	Range    []float64 `json:"range"`
	Typename string    `json:"__typename"`
}
type CBDContent3 struct {
	Unit     string    `json:"unit"`
	Range    []float64 `json:"range"`
	Typename string    `json:"__typename"`
}

type CannabinoidsContent3 struct {
	Value           float64          `json:"value"`
	Unit            string           `json:"PERCENTAGE"`
	CannabinoidType CannabinoidType3 `json:"cannabinoid"`
	TypeName        string           `json:"__typename"`
}

type CannabinoidType3 struct {
	Name     string `json:"name"`
	TypeName string `json:"__typename"`
}

type measurements3 struct {
	NetWeight weight2 `json:"netWeight"`
	Volume    string  `json:"volume"`
	Typename  string  `json:"__typename"`
}

type weight3 struct {
	Unit     string `json:"unit"`
	Value    []int  `json:"values"`
	Typename string `json:"__typename"`
}

type MetaData3 struct {
	CanonicalID        string                `json:"canonicalID"`
	CanonicalBrandName string                `json:"canonicalBrandName"`
	Children           []AdditionalMetaData3 `json:"children"`
	TypeName           string                `json:"__typename"`
}

type AdditionalMetaData3 struct {
	CanonicalEffectivePotencyMg float64           `json:"canonicalEffectivePotencyMg"`
	SizeOptions                 string            `json:"option"`
	InventoryQuantity           int               `json:"quantity"`
	InventoryQuantityAvailable  int               `json:"quantityAvailable"`
	InventoryKioskQuantity      int               `json:"kioskQuantityAvailable"`
	StandardEquivalent          WeightEquivalent3 `json:"standardEquivalent"`
}

type BrandInfo3 struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	ImageURL    string `json:"imageURL"`
	Name        string `json:"name"`
}

type Effects3 struct {
	RelaxedRating    float64 `json:"relaxed"`
	PainReliefRating float64 `json:"painRelief"`
	SleepyRating     float64 `json:"sleepy"`
	HappyRating      float64 `json:"happy"`
	EuphoricRating   float64 `json:"euphoric"`
}

type WeightEquivalent3 struct {
	Value    float64 `json:"value"`
	Unit     string  `json:"unit"`
	TypeName string  `json:"__typename"`
}

func InsertDataProductConcentratesCommon(db *sql.DB, target ConcentratesData) {

	for i := range target.Data.FilteredProducts.Products {

		stmt, err := db.Prepare("INSERT INTO Product_ConcentratesCommon (EffDate, DispensaryID, EnterpriseProductID, ProductID, BrandName, ProductName, ProductSubCategory, Weight, WeightGrams, RecOnly, RecPrice, RecSpecialPrice, MedicalOnly, MedicalPrice, MedicalSpecialPrice, WholeSalePrice, InventoryQuantity, InventoryQuantityAvail, InventoryQuantityKiosk, StrainType,  THCAmount, THCUnit, THCTypeName, CBDAMOUNT, CBDUnit, CBDTypeName, RexaledRating, PainReliefRating, SleepyRating, HappyRating, EuphoricRating, ProductImage) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
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

func getConcentratesData(db *sql.DB, dispensaryIDs []string, dispensaryTargets []ConcentratesData) {

	for i, Dispensary := range dispensaryIDs {

		DispensaryAPILink := "https://dutchie.com/graphql?operationName=FilteredProducts&variables=%7B%22includeEnterpriseSpecials%22%3Afalse%2C%22includeCannabinoids%22%3Atrue%2C%22productsFilter%22%3A%7B%22dispensaryId%22%3A%22" + Dispensary + "%22%2C%22pricingType%22%3A%22rec%22%2C%22strainTypes%22%3A%5B%5D%2C%22subcategories%22%3A%5B%5D%2C%22Status%22%3A%22Active%22%2C%22types%22%3A%5B%22Concentrate%22%5D%2C%22useCache%22%3Afalse%2C%22sortDirection%22%3A1%2C%22sortBy%22%3A%22price%22%2C%22isDefaultSort%22%3Atrue%2C%22bypassOnlineThresholds%22%3Afalse%2C%22isKioskMenu%22%3Afalse%2C%22removeProductsBelowOptionThresholds%22%3Atrue%7D%2C%22page%22%3A0%2C%22perPage%22%3A1000%7D&extensions=%7B%22persistedQuery%22%3A%7B%22version%22%3A1%2C%22sha256Hash%22%3A%22cf45528052df8664bf9404f98344719a5d7c6f8fff90e8c3c8ef2eaca63b0e9e%22%7D%7D"

		err := getJson(DispensaryAPILink, Dispensary+"ConcentratesData", &dispensaryTargets[i])

		if err != nil {
			fmt.Println(err)
		}
	}

	for _, ConcentratesMenu := range dispensaryTargets {
		InsertDataProductConcentratesCommon(db, ConcentratesMenu)

	}

}

var ConcentratesDataSlice = []ConcentratesData{
	legalgreensbrocktonConcentrates, commonwealthalternativecarebrocktonrecConcentrates, commonwealthalternativecarebrocktonConcentrates, atlanticmedicinalpartnersinc1Concentrates, ingoodhealthmedicalConcentrates, ingoodhealthbrocktonConcentrates, panaceawellnessquincyConcentrates, cannapidispensaryConcentrates, brocktongreenheartConcentrates, boterabrocktonConcentrates, budsgoodsandprovisionsabingtonConcentrates, cannavanaConcentrates, nativesunnorthattlebororetailrecConcentrates, yambaConcentrates, budsgoodsandprovisionswatertownConcentrates, siranaturalsConcentrates, mistymountainshopmaldenConcentrates, gardenremediesmelroserecConcentrates, gardenremediesmelroseConcentrates, oldeworldremediesConcentrates, naturesmedicinesfallriverConcentrates, goodchemistrylynnConcentrates, diemlynnConcentrates, calyxpeakcompaniesdbalocalcannabiscoConcentrates, newleafenterprisesConcentrates, insasalemConcentrates, insaavon1Concentrates, rhelmcannabisConcentrates, theorywellnessdeliveryzone1Concentrates, theorywellnessbridgewaterConcentrates, alternativecompassionservicesConcentrates, thehealthcircleConcentrates, greenrockcannabisConcentrates, releafalternativesConcentrates, ermontincConcentrates, greathitscannabisretailrectautonConcentrates, curaleafmahanoverConcentrates, panaceawellnessrecConcentrates, panaceawellnessConcentrates, quincycannabisquincyretailrecConcentrates, commonwealthalternativecareConcentrates, commonwealthalternativecare1Concentrates, terpsattleboroConcentrates, goodchemistryConcentrates, goodchemistrymedConcentrates, bloom1Concentrates, diemworcesterConcentrates, cookiesworcesterConcentrates, mayConcentratesworcesterConcentrates, budsgoodsandprovisionsworcesterConcentrates, resinateworcesterConcentrates, missionworcesterConcentrates, missionworcesteradultuseConcentrates, clearskyworcesterConcentrates, greencarecollectiveretailmedConcentrates, thevault1Concentrates, naturesremedymedConcentrates, naturesremedy1Concentrates, harmonyofmaConcentrates, campfirecannabis1Concentrates, discerndcannabisgraftonretailConcentrates, curaleafmaoxfordConcentrates, curaleafmaoxford1Concentrates, kosaConcentrates, greengoldConcentrates, societycannabiscoConcentrates, newenglandharvestConcentrates, thevaultwebsterConcentrates, terpscharltonConcentrates, temescalwellnesshudsonConcentrates, hudsonadultuseretailConcentrates, capitalcanabiscommunitydispensaryConcentrates, dmaholdingsllcdbagreatesthitscannabisConcentrates, gardenremediesmarlboroughConcentrates, gardenremediesrecmarlboroughConcentrates, greenmeadowsfarmConcentrates, cannabrollcdbagreenpathsouthbridgeConcentrates, GreenGoldMarboroughConcentrates, localrootssturbridgeConcentrates, greenngoConcentrates, blackstonevalleycannabisConcentrates, highhopesConcentrates, greeneraConcentrates, thriveofmadbaboomxshirleyConcentrates, ethosfitchburgConcentrates, atlanticmedicinalpartnersincConcentrates, atlanticmedicinalpartnersConcentrates, greenmeadowfarmfitchburgConcentrates, pioneercannabiscompanyConcentrates, naturesmedicinesuxbridgeConcentrates, uptopframinghamrecConcentrates, uptopframinghamConcentrates, masswellspringmaynardConcentrates, bountifulfarmsConcentrates, bleafwellnesscentreConcentrates, thehealingcenter1Concentrates, temescalwellnessframinghamConcentrates, botrealtyllcdbaomgcannabisConcentrates, curaleafmawareConcentrates, localrootsfitchburgConcentrates, ledlloyd420Concentrates, boterafranklin1Concentrates, masswellspringactonConcentrates, ddmcannabisConcentrates, centralavecompassionatecareConcentrates, gagecannabiscoConcentrates, netafranklinmedConcentrates, netafranklinConcentrates, communitycarecollective1Concentrates, unitedcultivationllcConcentrates, budbarnwinchendonConcentrates, siranaturalsneedhamConcentrates, umaConcentratess1Concentrates, clearskybelchertownConcentrates, thebostongardenConcentrates, rediConcentrates, elev8cannabisatholConcentrates, toytownprojectllcdbatoytownhealthcareConcentrates, gardenremediesnewtonConcentrates, gardenremediesrecnewtonConcentrates, budsgoodsandprovisionswatertownConcentrates2, nativesunnorthattlebororetailrecConcentrates2, auraofrhodeislandConcentrates, ayrdispensarywatertownmedConcentrates, ayrdispensarywatertownConcentrates, communitycarecollectiveConcentrates, uptopwestroxburyConcentrates, orangecannabiscompanyConcentrates, EthosWatertownAdultUseConcentrates, EthosWatertownMedicalConcentrates, apothcaarlingtonConcentrates, releafalternativesConcentrates2, ocsivymaerealestatedbazaharacannabisConcentrates, motherearthpawtucketConcentrates, kgcollectivecambridgeConcentrates, revolutionaryclinicsfreshpondConcentrates, naturesremedy2Concentrates, highprofileroslindaleConcentrates, mayConcentratesmedicinalsallstonrecConcentrates, mayConcentratesallstonConcentrates, apothcajamaicaplainConcentrates, netabrooklineConcentrates, netabrooklinemedConcentrates, missionbrooklineConcentrates, redcardinal2Concentrates, libertyspringfieldConcentrates, canacraftConcentrates, siranaturalsConcentrates2, siranaturalssomervilleadultuseConcentrates, medmenfenwayConcentrates, westernfront1Concentrates, smythcannabiscoConcentrates, massalternativecareamherstmedConcentrates, massalternativecareamherstConcentrates, pleasantreesamherstConcentrates, libertysomervilleConcentrates, theheirloomcollectiveConcentrates, pureoasis1Concentrates, advesamaincdbablueriverterpscambridgeConcentrates, insaspringfieldConcentrates, ayrdispensarybackbayConcentrates, reverie73Concentrates, massalternativecareConcentrates, massalternativecaremedConcentrates, revolutionaryclinicssomervilleConcentrates, theorywellnesschicopeemedConcentrates, ethosdorchesterConcentrates, highprofiledorchesterConcentrates, lazyriverproductsConcentrates, mistymountainshopmaldenConcentrates2, jimbuddysConcentrates, apexnoireConcentrates, cannapidispensaryConcentrates2, ermontincConcentrates2, dazedcannabisConcentrates,
}

var legalgreensbrocktonConcentrates ConcentratesData
var commonwealthalternativecarebrocktonrecConcentrates ConcentratesData
var commonwealthalternativecarebrocktonConcentrates ConcentratesData
var atlanticmedicinalpartnersinc1Concentrates ConcentratesData
var ingoodhealthmedicalConcentrates ConcentratesData
var ingoodhealthbrocktonConcentrates ConcentratesData
var panaceawellnessquincyConcentrates ConcentratesData
var cannapidispensaryConcentrates ConcentratesData
var brocktongreenheartConcentrates ConcentratesData
var boterabrocktonConcentrates ConcentratesData
var budsgoodsandprovisionsabingtonConcentrates ConcentratesData
var cannavanaConcentrates ConcentratesData
var nativesunnorthattlebororetailrecConcentrates ConcentratesData
var yambaConcentrates ConcentratesData
var budsgoodsandprovisionswatertownConcentrates ConcentratesData
var siranaturalsConcentrates ConcentratesData
var mistymountainshopmaldenConcentrates ConcentratesData
var gardenremediesmelroserecConcentrates ConcentratesData
var gardenremediesmelroseConcentrates ConcentratesData
var oldeworldremediesConcentrates ConcentratesData
var naturesmedicinesfallriverConcentrates ConcentratesData
var goodchemistrylynnConcentrates ConcentratesData
var diemlynnConcentrates ConcentratesData
var calyxpeakcompaniesdbalocalcannabiscoConcentrates ConcentratesData
var newleafenterprisesConcentrates ConcentratesData
var insasalemConcentrates ConcentratesData
var insaavon1Concentrates ConcentratesData
var rhelmcannabisConcentrates ConcentratesData
var theorywellnessdeliveryzone1Concentrates ConcentratesData
var theorywellnessbridgewaterConcentrates ConcentratesData
var alternativecompassionservicesConcentrates ConcentratesData
var thehealthcircleConcentrates ConcentratesData
var greenrockcannabisConcentrates ConcentratesData
var releafalternativesConcentrates ConcentratesData
var ermontincConcentrates ConcentratesData
var greathitscannabisretailrectautonConcentrates ConcentratesData
var curaleafmahanoverConcentrates ConcentratesData
var panaceawellnessrecConcentrates ConcentratesData
var panaceawellnessConcentrates ConcentratesData
var quincycannabisquincyretailrecConcentrates ConcentratesData
var commonwealthalternativecareConcentrates ConcentratesData
var commonwealthalternativecare1Concentrates ConcentratesData
var terpsattleboroConcentrates ConcentratesData
var goodchemistryConcentrates ConcentratesData
var goodchemistrymedConcentrates ConcentratesData
var bloom1Concentrates ConcentratesData
var diemworcesterConcentrates ConcentratesData
var cookiesworcesterConcentrates ConcentratesData
var mayConcentratesworcesterConcentrates ConcentratesData
var budsgoodsandprovisionsworcesterConcentrates ConcentratesData
var resinateworcesterConcentrates ConcentratesData
var missionworcesterConcentrates ConcentratesData
var missionworcesteradultuseConcentrates ConcentratesData
var clearskyworcesterConcentrates ConcentratesData
var greencarecollectiveretailmedConcentrates ConcentratesData
var thevault1Concentrates ConcentratesData
var naturesremedymedConcentrates ConcentratesData
var naturesremedy1Concentrates ConcentratesData
var harmonyofmaConcentrates ConcentratesData
var campfirecannabis1Concentrates ConcentratesData
var discerndcannabisgraftonretailConcentrates ConcentratesData
var curaleafmaoxfordConcentrates ConcentratesData
var curaleafmaoxford1Concentrates ConcentratesData
var kosaConcentrates ConcentratesData
var greengoldConcentrates ConcentratesData
var societycannabiscoConcentrates ConcentratesData
var newenglandharvestConcentrates ConcentratesData
var thevaultwebsterConcentrates ConcentratesData
var terpscharltonConcentrates ConcentratesData
var temescalwellnesshudsonConcentrates ConcentratesData
var hudsonadultuseretailConcentrates ConcentratesData
var capitalcanabiscommunitydispensaryConcentrates ConcentratesData
var dmaholdingsllcdbagreatesthitscannabisConcentrates ConcentratesData
var gardenremediesmarlboroughConcentrates ConcentratesData
var gardenremediesrecmarlboroughConcentrates ConcentratesData
var greenmeadowsfarmConcentrates ConcentratesData
var cannabrollcdbagreenpathsouthbridgeConcentrates ConcentratesData
var GreenGoldMarboroughConcentrates ConcentratesData
var localrootssturbridgeConcentrates ConcentratesData
var greenngoConcentrates ConcentratesData
var blackstonevalleycannabisConcentrates ConcentratesData
var highhopesConcentrates ConcentratesData
var greeneraConcentrates ConcentratesData
var thriveofmadbaboomxshirleyConcentrates ConcentratesData
var ethosfitchburgConcentrates ConcentratesData
var atlanticmedicinalpartnersincConcentrates ConcentratesData
var atlanticmedicinalpartnersConcentrates ConcentratesData
var greenmeadowfarmfitchburgConcentrates ConcentratesData
var pioneercannabiscompanyConcentrates ConcentratesData
var naturesmedicinesuxbridgeConcentrates ConcentratesData
var uptopframinghamrecConcentrates ConcentratesData
var uptopframinghamConcentrates ConcentratesData
var masswellspringmaynardConcentrates ConcentratesData
var bountifulfarmsConcentrates ConcentratesData
var bleafwellnesscentreConcentrates ConcentratesData
var thehealingcenter1Concentrates ConcentratesData
var temescalwellnessframinghamConcentrates ConcentratesData
var botrealtyllcdbaomgcannabisConcentrates ConcentratesData
var curaleafmawareConcentrates ConcentratesData
var localrootsfitchburgConcentrates ConcentratesData
var ledlloyd420Concentrates ConcentratesData
var boterafranklin1Concentrates ConcentratesData
var masswellspringactonConcentrates ConcentratesData
var ddmcannabisConcentrates ConcentratesData
var centralavecompassionatecareConcentrates ConcentratesData
var gagecannabiscoConcentrates ConcentratesData
var netafranklinmedConcentrates ConcentratesData
var netafranklinConcentrates ConcentratesData
var communitycarecollective1Concentrates ConcentratesData
var unitedcultivationllcConcentrates ConcentratesData
var budbarnwinchendonConcentrates ConcentratesData
var siranaturalsneedhamConcentrates ConcentratesData
var umaConcentratess1Concentrates ConcentratesData
var clearskybelchertownConcentrates ConcentratesData
var thebostongardenConcentrates ConcentratesData
var rediConcentrates ConcentratesData
var elev8cannabisatholConcentrates ConcentratesData
var toytownprojectllcdbatoytownhealthcareConcentrates ConcentratesData
var gardenremediesnewtonConcentrates ConcentratesData
var gardenremediesrecnewtonConcentrates ConcentratesData
var budsgoodsandprovisionswatertownConcentrates2 ConcentratesData
var nativesunnorthattlebororetailrecConcentrates2 ConcentratesData
var auraofrhodeislandConcentrates ConcentratesData
var ayrdispensarywatertownmedConcentrates ConcentratesData
var ayrdispensarywatertownConcentrates ConcentratesData
var communitycarecollectiveConcentrates ConcentratesData
var uptopwestroxburyConcentrates ConcentratesData
var orangecannabiscompanyConcentrates ConcentratesData
var EthosWatertownAdultUseConcentrates ConcentratesData
var EthosWatertownMedicalConcentrates ConcentratesData
var apothcaarlingtonConcentrates ConcentratesData
var releafalternativesConcentrates2 ConcentratesData
var ocsivymaerealestatedbazaharacannabisConcentrates ConcentratesData
var motherearthpawtucketConcentrates ConcentratesData
var kgcollectivecambridgeConcentrates ConcentratesData
var revolutionaryclinicsfreshpondConcentrates ConcentratesData
var naturesremedy2Concentrates ConcentratesData
var highprofileroslindaleConcentrates ConcentratesData
var mayConcentratesmedicinalsallstonrecConcentrates ConcentratesData
var mayConcentratesallstonConcentrates ConcentratesData
var apothcajamaicaplainConcentrates ConcentratesData
var netabrooklineConcentrates ConcentratesData
var netabrooklinemedConcentrates ConcentratesData
var missionbrooklineConcentrates ConcentratesData
var redcardinal2Concentrates ConcentratesData
var libertyspringfieldConcentrates ConcentratesData
var canacraftConcentrates ConcentratesData
var siranaturalsConcentrates2 ConcentratesData
var siranaturalssomervilleadultuseConcentrates ConcentratesData
var medmenfenwayConcentrates ConcentratesData
var westernfront1Concentrates ConcentratesData
var smythcannabiscoConcentrates ConcentratesData
var massalternativecareamherstmedConcentrates ConcentratesData
var massalternativecareamherstConcentrates ConcentratesData
var pleasantreesamherstConcentrates ConcentratesData
var libertysomervilleConcentrates ConcentratesData
var theheirloomcollectiveConcentrates ConcentratesData
var pureoasis1Concentrates ConcentratesData
var advesamaincdbablueriverterpscambridgeConcentrates ConcentratesData
var insaspringfieldConcentrates ConcentratesData
var ayrdispensarybackbayConcentrates ConcentratesData
var reverie73Concentrates ConcentratesData
var massalternativecareConcentrates ConcentratesData
var massalternativecaremedConcentrates ConcentratesData
var revolutionaryclinicssomervilleConcentrates ConcentratesData
var theorywellnesschicopeemedConcentrates ConcentratesData
var ethosdorchesterConcentrates ConcentratesData
var highprofiledorchesterConcentrates ConcentratesData
var lazyriverproductsConcentrates ConcentratesData
var mistymountainshopmaldenConcentrates2 ConcentratesData
var jimbuddysConcentrates ConcentratesData
var apexnoireConcentrates ConcentratesData
var cannapidispensaryConcentrates2 ConcentratesData
var ermontincConcentrates2 ConcentratesData
var dazedcannabisConcentrates ConcentratesData
