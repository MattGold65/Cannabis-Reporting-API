package main

import (
	"database/sql"
	"fmt"
	"time"
)

type prerollsData struct {
	Data DataSet4 `json:"data"`
}

type DataSet4 struct {
	FilteredProducts Product4 `json:"filteredProducts"`
}

type Product4 struct {
	Products []prerollsProducts `json:"products"`
}

type prerollsProducts struct {
	Id                         string                 `json:"_id"`
	BrandName                  string                 `json:"brandName"`
	CBD                        string                 `json:"CBD"`
	CBDContent                 CBDContent4            `json:"CBDContent"`
	ComingSoon                 string                 `json:"comingSoon"`
	CreatedAt                  string                 `json:"createdAt"`
	DispensaryID               string                 `json:"DispensaryID"`
	EnterpriseProductId        string                 `json:"enterpriseProductId"`
	Image                      string                 `json:"Image"`
	Measurements               measurements4          `json:"measurements"`
	MedicalOnly                bool                   `json:"medicalOnly"`
	MedicalPrice               []float64              `json:"medicalPrices"`
	MedicalSpecialPrice        []float64              `json:"medicalSpecialPrices"`
	WholeSalePrice             []float64              `json:"wholesalePrices"`
	ProductName                string                 `json:"Name"`
	NonArmsLength              string                 `json:"nonArmsLength"`
	WeightOptions              []string               `json:"Options"`
	CustomerLimit              string                 `json:"limitsPerCustomer"`
	ManualInventory            string                 `json:"manualInventory"`
	POSMetaData                MetaData4              `json:"POSMetaData"`
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
	THCContent                 THCContent4            `json:"THCContent"`
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
	BrandInfo                  BrandInfo4             `json:"brand"`
	Effects                    Effects4               `json:"effects"`
	CannabinoidsContent        []CannabinoidsContent4 `json:"cannabinoidsV2"`
}
type THCContent4 struct {
	Unit     string    `json:"unit"`
	Range    []float64 `json:"range"`
	Typename string    `json:"__typename"`
}
type CBDContent4 struct {
	Unit     string    `json:"unit"`
	Range    []float64 `json:"range"`
	Typename string    `json:"__typename"`
}

type CannabinoidsContent4 struct {
	Value           float64          `json:"value"`
	Unit            string           `json:"PERCENTAGE"`
	CannabinoidType CannabinoidType4 `json:"cannabinoid"`
	TypeName        string           `json:"__typename"`
}

type CannabinoidType4 struct {
	Name     string `json:"name"`
	TypeName string `json:"__typename"`
}

type measurements4 struct {
	NetWeight weight4 `json:"netWeight"`
	Volume    string  `json:"volume"`
	Typename  string  `json:"__typename"`
}

type weight4 struct {
	Unit     string `json:"unit"`
	Value    []int  `json:"values"`
	Typename string `json:"__typename"`
}

type MetaData4 struct {
	CanonicalID        string                `json:"canonicalID"`
	CanonicalBrandName string                `json:"canonicalBrandName"`
	Children           []AdditionalMetaData4 `json:"children"`
	TypeName           string                `json:"__typename"`
}

type AdditionalMetaData4 struct {
	CanonicalEffectivePotencyMg float64           `json:"canonicalEffectivePotencyMg"`
	SizeOptions                 string            `json:"option"`
	InventoryQuantity           int               `json:"quantity"`
	InventoryQuantityAvailable  int               `json:"quantityAvailable"`
	InventoryKioskQuantity      int               `json:"kioskQuantityAvailable"`
	StandardEquivalent          WeightEquivalent4 `json:"standardEquivalent"`
}

type BrandInfo4 struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	ImageURL    string `json:"imageURL"`
	Name        string `json:"name"`
}

type Effects4 struct {
	RelaxedRating    float64 `json:"relaxed"`
	PainReliefRating float64 `json:"painRelief"`
	SleepyRating     float64 `json:"sleepy"`
	HappyRating      float64 `json:"happy"`
	EuphoricRating   float64 `json:"euphoric"`
}

type WeightEquivalent4 struct {
	Value    float64 `json:"value"`
	Unit     string  `json:"unit"`
	TypeName string  `json:"__typename"`
}

func InsertDataProductprerollsCommon(db *sql.DB, target prerollsData) {

	for i := range target.Data.FilteredProducts.Products {

		stmt, err := db.Prepare("INSERT INTO Product_prerollsCommon (EffDate, DispensaryID, EnterpriseProductID, ProductID, BrandName, ProductName, ProductSubCategory, Weight, WeightGrams, RecOnly, RecPrice, RecSpecialPrice, MedicalOnly, MedicalPrice, MedicalSpecialPrice, WholeSalePrice, InventoryQuantity, InventoryQuantityAvail, InventoryQuantityKiosk, StrainType,  THCAmount, THCUnit, THCTypeName, CBDAMOUNT, CBDUnit, CBDTypeName, RexaledRating, PainReliefRating, SleepyRating, HappyRating, EuphoricRating, ProductImage) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
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

func getprerollsData(db *sql.DB, dispensaryIDs []string, dispensaryTargets []prerollsData) {

	for i, Dispensary := range dispensaryIDs {

		DispensaryAPILink := "https://dutchie.com/graphql?operationName=FilteredProducts&variables=%7B%22includeEnterpriseSpecials%22%3Afalse%2C%22includeCannabinoids%22%3Atrue%2C%22productsFilter%22%3A%7B%22dispensaryId%22%3A%22" + Dispensary + "%22%2C%22pricingType%22%3A%22med%22%2C%22strainTypes%22%3A%5B%5D%2C%22subcategories%22%3A%5B%5D%2C%22Status%22%3A%22Active%22%2C%22types%22%3A%5B%22Pre-Rolls%22%5D%2C%22useCache%22%3Afalse%2C%22sortDirection%22%3A1%2C%22sortBy%22%3A%22weight%22%2C%22isDefaultSort%22%3Atrue%2C%22bypassOnlineThresholds%22%3Afalse%2C%22isKioskMenu%22%3Afalse%2C%22removeProductsBelowOptionThresholds%22%3Atrue%7D%2C%22page%22%3A0%2C%22perPage%22%3A1000%7D&extensions=%7B%22persistedQuery%22%3A%7B%22version%22%3A1%2C%22sha256Hash%22%3A%22cf45528052df8664bf9404f98344719a5d7c6f8fff90e8c3c8ef2eaca63b0e9e%22%7D%7D"

		err := getJson(DispensaryAPILink, Dispensary+"prerollsData", &dispensaryTargets[i])

		if err != nil {
			fmt.Println(err)
		}
	}

	for _, prerollsMenu := range dispensaryTargets {
		InsertDataProductprerollsCommon(db, prerollsMenu)

	}

}

var prerollsDataSlice = []prerollsData{
	legalgreensbrocktonprerolls, commonwealthalternativecarebrocktonrecprerolls, commonwealthalternativecarebrocktonprerolls, atlanticmedicinalpartnersinc1prerolls, ingoodhealthmedicalprerolls, ingoodhealthbrocktonprerolls, panaceawellnessquincyprerolls, cannapidispensaryprerolls, brocktongreenheartprerolls, boterabrocktonprerolls, budsgoodsandprovisionsabingtonprerolls, cannavanaprerolls, nativesunnorthattlebororetailrecprerolls, yambaprerolls, budsgoodsandprovisionswatertownprerolls, siranaturalsprerolls, mistymountainshopmaldenprerolls, gardenremediesmelroserecprerolls, gardenremediesmelroseprerolls, oldeworldremediesprerolls, naturesmedicinesfallriverprerolls, goodchemistrylynnprerolls, diemlynnprerolls, calyxpeakcompaniesdbalocalcannabiscoprerolls, newleafenterprisesprerolls, insasalemprerolls, insaavon1prerolls, rhelmcannabisprerolls, theorywellnessdeliveryzone1prerolls, theorywellnessbridgewaterprerolls, alternativecompassionservicesprerolls, thehealthcircleprerolls, greenrockcannabisprerolls, releafalternativesprerolls, ermontincprerolls, greathitscannabisretailrectautonprerolls, curaleafmahanoverprerolls, panaceawellnessrecprerolls, panaceawellnessprerolls, quincycannabisquincyretailrecprerolls, commonwealthalternativecareprerolls, commonwealthalternativecare1prerolls, terpsattleboroprerolls, goodchemistryprerolls, goodchemistrymedprerolls, bloom1prerolls, diemworcesterprerolls, cookiesworcesterprerolls, mayprerollsworcesterprerolls, budsgoodsandprovisionsworcesterprerolls, resinateworcesterprerolls, missionworcesterprerolls, missionworcesteradultuseprerolls, clearskyworcesterprerolls, greencarecollectiveretailmedprerolls, thevault1prerolls, naturesremedymedprerolls, naturesremedy1prerolls, harmonyofmaprerolls, campfirecannabis1prerolls, discerndcannabisgraftonretailprerolls, curaleafmaoxfordprerolls, curaleafmaoxford1prerolls, kosaprerolls, greengoldprerolls, societycannabiscoprerolls, newenglandharvestprerolls, thevaultwebsterprerolls, terpscharltonprerolls, temescalwellnesshudsonprerolls, hudsonadultuseretailprerolls, capitalcanabiscommunitydispensaryprerolls, dmaholdingsllcdbagreatesthitscannabisprerolls, gardenremediesmarlboroughprerolls, gardenremediesrecmarlboroughprerolls, greenmeadowsfarmprerolls, cannabrollcdbagreenpathsouthbridgeprerolls, GreenGoldMarboroughprerolls, localrootssturbridgeprerolls, greenngoprerolls, blackstonevalleycannabisprerolls, highhopesprerolls, greeneraprerolls, thriveofmadbaboomxshirleyprerolls, ethosfitchburgprerolls, atlanticmedicinalpartnersincprerolls, atlanticmedicinalpartnersprerolls, greenmeadowfarmfitchburgprerolls, pioneercannabiscompanyprerolls, naturesmedicinesuxbridgeprerolls, uptopframinghamrecprerolls, uptopframinghamprerolls, masswellspringmaynardprerolls, bountifulfarmsprerolls, bleafwellnesscentreprerolls, thehealingcenter1prerolls, temescalwellnessframinghamprerolls, botrealtyllcdbaomgcannabisprerolls, curaleafmawareprerolls, localrootsfitchburgprerolls, ledlloyd420prerolls, boterafranklin1prerolls, masswellspringactonprerolls, ddmcannabisprerolls, centralavecompassionatecareprerolls, gagecannabiscoprerolls, netafranklinmedprerolls, netafranklinprerolls, communitycarecollective1prerolls, unitedcultivationllcprerolls, budbarnwinchendonprerolls, siranaturalsneedhamprerolls, umaprerollss1prerolls, clearskybelchertownprerolls, thebostongardenprerolls, rediprerolls, elev8cannabisatholprerolls, toytownprojectllcdbatoytownhealthcareprerolls, gardenremediesnewtonprerolls, gardenremediesrecnewtonprerolls, budsgoodsandprovisionswatertownprerolls2, nativesunnorthattlebororetailrecprerolls2, auraofrhodeislandprerolls, ayrdispensarywatertownmedprerolls, ayrdispensarywatertownprerolls, communitycarecollectiveprerolls, uptopwestroxburyprerolls, orangecannabiscompanyprerolls, EthosWatertownAdultUseprerolls, EthosWatertownMedicalprerolls, apothcaarlingtonprerolls, releafalternativesprerolls2, ocsivymaerealestatedbazaharacannabisprerolls, motherearthpawtucketprerolls, kgcollectivecambridgeprerolls, revolutionaryclinicsfreshpondprerolls, naturesremedy2prerolls, highprofileroslindaleprerolls, mayprerollsmedicinalsallstonrecprerolls, mayprerollsallstonprerolls, apothcajamaicaplainprerolls, netabrooklineprerolls, netabrooklinemedprerolls, missionbrooklineprerolls, redcardinal2prerolls, libertyspringfieldprerolls, canacraftprerolls, siranaturalsprerolls2, siranaturalssomervilleadultuseprerolls, medmenfenwayprerolls, westernfront1prerolls, smythcannabiscoprerolls, massalternativecareamherstmedprerolls, massalternativecareamherstprerolls, pleasantreesamherstprerolls, libertysomervilleprerolls, theheirloomcollectiveprerolls, pureoasis1prerolls, advesamaincdbablueriverterpscambridgeprerolls, insaspringfieldprerolls, ayrdispensarybackbayprerolls, reverie73prerolls, massalternativecareprerolls, massalternativecaremedprerolls, revolutionaryclinicssomervilleprerolls, theorywellnesschicopeemedprerolls, ethosdorchesterprerolls, highprofiledorchesterprerolls, lazyriverproductsprerolls, mistymountainshopmaldenprerolls2, jimbuddysprerolls, apexnoireprerolls, cannapidispensaryprerolls2, ermontincprerolls2, dazedcannabisprerolls,
}

var legalgreensbrocktonprerolls prerollsData
var commonwealthalternativecarebrocktonrecprerolls prerollsData
var commonwealthalternativecarebrocktonprerolls prerollsData
var atlanticmedicinalpartnersinc1prerolls prerollsData
var ingoodhealthmedicalprerolls prerollsData
var ingoodhealthbrocktonprerolls prerollsData
var panaceawellnessquincyprerolls prerollsData
var cannapidispensaryprerolls prerollsData
var brocktongreenheartprerolls prerollsData
var boterabrocktonprerolls prerollsData
var budsgoodsandprovisionsabingtonprerolls prerollsData
var cannavanaprerolls prerollsData
var nativesunnorthattlebororetailrecprerolls prerollsData
var yambaprerolls prerollsData
var budsgoodsandprovisionswatertownprerolls prerollsData
var siranaturalsprerolls prerollsData
var mistymountainshopmaldenprerolls prerollsData
var gardenremediesmelroserecprerolls prerollsData
var gardenremediesmelroseprerolls prerollsData
var oldeworldremediesprerolls prerollsData
var naturesmedicinesfallriverprerolls prerollsData
var goodchemistrylynnprerolls prerollsData
var diemlynnprerolls prerollsData
var calyxpeakcompaniesdbalocalcannabiscoprerolls prerollsData
var newleafenterprisesprerolls prerollsData
var insasalemprerolls prerollsData
var insaavon1prerolls prerollsData
var rhelmcannabisprerolls prerollsData
var theorywellnessdeliveryzone1prerolls prerollsData
var theorywellnessbridgewaterprerolls prerollsData
var alternativecompassionservicesprerolls prerollsData
var thehealthcircleprerolls prerollsData
var greenrockcannabisprerolls prerollsData
var releafalternativesprerolls prerollsData
var ermontincprerolls prerollsData
var greathitscannabisretailrectautonprerolls prerollsData
var curaleafmahanoverprerolls prerollsData
var panaceawellnessrecprerolls prerollsData
var panaceawellnessprerolls prerollsData
var quincycannabisquincyretailrecprerolls prerollsData
var commonwealthalternativecareprerolls prerollsData
var commonwealthalternativecare1prerolls prerollsData
var terpsattleboroprerolls prerollsData
var goodchemistryprerolls prerollsData
var goodchemistrymedprerolls prerollsData
var bloom1prerolls prerollsData
var diemworcesterprerolls prerollsData
var cookiesworcesterprerolls prerollsData
var mayprerollsworcesterprerolls prerollsData
var budsgoodsandprovisionsworcesterprerolls prerollsData
var resinateworcesterprerolls prerollsData
var missionworcesterprerolls prerollsData
var missionworcesteradultuseprerolls prerollsData
var clearskyworcesterprerolls prerollsData
var greencarecollectiveretailmedprerolls prerollsData
var thevault1prerolls prerollsData
var naturesremedymedprerolls prerollsData
var naturesremedy1prerolls prerollsData
var harmonyofmaprerolls prerollsData
var campfirecannabis1prerolls prerollsData
var discerndcannabisgraftonretailprerolls prerollsData
var curaleafmaoxfordprerolls prerollsData
var curaleafmaoxford1prerolls prerollsData
var kosaprerolls prerollsData
var greengoldprerolls prerollsData
var societycannabiscoprerolls prerollsData
var newenglandharvestprerolls prerollsData
var thevaultwebsterprerolls prerollsData
var terpscharltonprerolls prerollsData
var temescalwellnesshudsonprerolls prerollsData
var hudsonadultuseretailprerolls prerollsData
var capitalcanabiscommunitydispensaryprerolls prerollsData
var dmaholdingsllcdbagreatesthitscannabisprerolls prerollsData
var gardenremediesmarlboroughprerolls prerollsData
var gardenremediesrecmarlboroughprerolls prerollsData
var greenmeadowsfarmprerolls prerollsData
var cannabrollcdbagreenpathsouthbridgeprerolls prerollsData
var GreenGoldMarboroughprerolls prerollsData
var localrootssturbridgeprerolls prerollsData
var greenngoprerolls prerollsData
var blackstonevalleycannabisprerolls prerollsData
var highhopesprerolls prerollsData
var greeneraprerolls prerollsData
var thriveofmadbaboomxshirleyprerolls prerollsData
var ethosfitchburgprerolls prerollsData
var atlanticmedicinalpartnersincprerolls prerollsData
var atlanticmedicinalpartnersprerolls prerollsData
var greenmeadowfarmfitchburgprerolls prerollsData
var pioneercannabiscompanyprerolls prerollsData
var naturesmedicinesuxbridgeprerolls prerollsData
var uptopframinghamrecprerolls prerollsData
var uptopframinghamprerolls prerollsData
var masswellspringmaynardprerolls prerollsData
var bountifulfarmsprerolls prerollsData
var bleafwellnesscentreprerolls prerollsData
var thehealingcenter1prerolls prerollsData
var temescalwellnessframinghamprerolls prerollsData
var botrealtyllcdbaomgcannabisprerolls prerollsData
var curaleafmawareprerolls prerollsData
var localrootsfitchburgprerolls prerollsData
var ledlloyd420prerolls prerollsData
var boterafranklin1prerolls prerollsData
var masswellspringactonprerolls prerollsData
var ddmcannabisprerolls prerollsData
var centralavecompassionatecareprerolls prerollsData
var gagecannabiscoprerolls prerollsData
var netafranklinmedprerolls prerollsData
var netafranklinprerolls prerollsData
var communitycarecollective1prerolls prerollsData
var unitedcultivationllcprerolls prerollsData
var budbarnwinchendonprerolls prerollsData
var siranaturalsneedhamprerolls prerollsData
var umaprerollss1prerolls prerollsData
var clearskybelchertownprerolls prerollsData
var thebostongardenprerolls prerollsData
var rediprerolls prerollsData
var elev8cannabisatholprerolls prerollsData
var toytownprojectllcdbatoytownhealthcareprerolls prerollsData
var gardenremediesnewtonprerolls prerollsData
var gardenremediesrecnewtonprerolls prerollsData
var budsgoodsandprovisionswatertownprerolls2 prerollsData
var nativesunnorthattlebororetailrecprerolls2 prerollsData
var auraofrhodeislandprerolls prerollsData
var ayrdispensarywatertownmedprerolls prerollsData
var ayrdispensarywatertownprerolls prerollsData
var communitycarecollectiveprerolls prerollsData
var uptopwestroxburyprerolls prerollsData
var orangecannabiscompanyprerolls prerollsData
var EthosWatertownAdultUseprerolls prerollsData
var EthosWatertownMedicalprerolls prerollsData
var apothcaarlingtonprerolls prerollsData
var releafalternativesprerolls2 prerollsData
var ocsivymaerealestatedbazaharacannabisprerolls prerollsData
var motherearthpawtucketprerolls prerollsData
var kgcollectivecambridgeprerolls prerollsData
var revolutionaryclinicsfreshpondprerolls prerollsData
var naturesremedy2prerolls prerollsData
var highprofileroslindaleprerolls prerollsData
var mayprerollsmedicinalsallstonrecprerolls prerollsData
var mayprerollsallstonprerolls prerollsData
var apothcajamaicaplainprerolls prerollsData
var netabrooklineprerolls prerollsData
var netabrooklinemedprerolls prerollsData
var missionbrooklineprerolls prerollsData
var redcardinal2prerolls prerollsData
var libertyspringfieldprerolls prerollsData
var canacraftprerolls prerollsData
var siranaturalsprerolls2 prerollsData
var siranaturalssomervilleadultuseprerolls prerollsData
var medmenfenwayprerolls prerollsData
var westernfront1prerolls prerollsData
var smythcannabiscoprerolls prerollsData
var massalternativecareamherstmedprerolls prerollsData
var massalternativecareamherstprerolls prerollsData
var pleasantreesamherstprerolls prerollsData
var libertysomervilleprerolls prerollsData
var theheirloomcollectiveprerolls prerollsData
var pureoasis1prerolls prerollsData
var advesamaincdbablueriverterpscambridgeprerolls prerollsData
var insaspringfieldprerolls prerollsData
var ayrdispensarybackbayprerolls prerollsData
var reverie73prerolls prerollsData
var massalternativecareprerolls prerollsData
var massalternativecaremedprerolls prerollsData
var revolutionaryclinicssomervilleprerolls prerollsData
var theorywellnesschicopeemedprerolls prerollsData
var ethosdorchesterprerolls prerollsData
var highprofiledorchesterprerolls prerollsData
var lazyriverproductsprerolls prerollsData
var mistymountainshopmaldenprerolls2 prerollsData
var jimbuddysprerolls prerollsData
var apexnoireprerolls prerollsData
var cannapidispensaryprerolls2 prerollsData
var ermontincprerolls2 prerollsData
var dazedcannabisprerolls prerollsData
