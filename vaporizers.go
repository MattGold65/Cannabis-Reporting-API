package main

import (
	"database/sql"
	"fmt"
	"time"
)

type vapeorizersData struct {
	Data DataSet5 `json:"data"`
}

type DataSet5 struct {
	FilteredProducts Product5 `json:"filteredProducts"`
}

type Product5 struct {
	Products []vapeorizersProducts `json:"products"`
}

type vapeorizersProducts struct {
	Id                         string                 `json:"_id"`
	BrandName                  string                 `json:"brandName"`
	CBD                        string                 `json:"CBD"`
	CBDContent                 CBDContent5            `json:"CBDContent"`
	ComingSoon                 string                 `json:"comingSoon"`
	CreatedAt                  string                 `json:"createdAt"`
	DispensaryID               string                 `json:"DispensaryID"`
	EnterpriseProductId        string                 `json:"enterpriseProductId"`
	Image                      string                 `json:"Image"`
	Measurements               measurements5          `json:"measurements"`
	MedicalOnly                bool                   `json:"medicalOnly"`
	MedicalPrice               []float64              `json:"medicalPrices"`
	MedicalSpecialPrice        []float64              `json:"medicalSpecialPrices"`
	WholeSalePrice             []float64              `json:"wholesalePrices"`
	ProductName                string                 `json:"Name"`
	NonArmsLength              string                 `json:"nonArmsLength"`
	WeightOptions              []string               `json:"Options"`
	CustomerLimit              string                 `json:"limitsPerCustomer"`
	ManualInventory            string                 `json:"manualInventory"`
	POSMetaData                MetaData5              `json:"POSMetaData"`
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
	THCContent                 THCContent5            `json:"THCContent"`
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
	BrandInfo                  BrandInfo5             `json:"brand"`
	Effects                    Effects5               `json:"effects"`
	CannabinoidsContent        []CannabinoidsContent5 `json:"cannabinoidsV2"`
}
type THCContent5 struct {
	Unit     string    `json:"unit"`
	Range    []float64 `json:"range"`
	Typename string    `json:"__typename"`
}
type CBDContent5 struct {
	Unit     string    `json:"unit"`
	Range    []float64 `json:"range"`
	Typename string    `json:"__typename"`
}

type CannabinoidsContent5 struct {
	Value           float64          `json:"value"`
	Unit            string           `json:"PERCENTAGE"`
	CannabinoidType CannabinoidType5 `json:"cannabinoid"`
	TypeName        string           `json:"__typename"`
}

type CannabinoidType5 struct {
	Name     string `json:"name"`
	TypeName string `json:"__typename"`
}

type measurements5 struct {
	NetWeight weight5 `json:"netWeight"`
	Volume    string  `json:"volume"`
	Typename  string  `json:"__typename"`
}

type weight5 struct {
	Unit     string `json:"unit"`
	Value    []int  `json:"values"`
	Typename string `json:"__typename"`
}

type MetaData5 struct {
	CanonicalID        string                `json:"canonicalID"`
	CanonicalBrandName string                `json:"canonicalBrandName"`
	Children           []AdditionalMetaData5 `json:"children"`
	TypeName           string                `json:"__typename"`
}

type AdditionalMetaData5 struct {
	CanonicalEffectivePotencyMg float64           `json:"canonicalEffectivePotencyMg"`
	SizeOptions                 string            `json:"option"`
	InventoryQuantity           int               `json:"quantity"`
	InventoryQuantityAvailable  int               `json:"quantityAvailable"`
	InventoryKioskQuantity      int               `json:"kioskQuantityAvailable"`
	StandardEquivalent          WeightEquivalent5 `json:"standardEquivalent"`
}

type BrandInfo5 struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	ImageURL    string `json:"imageURL"`
	Name        string `json:"name"`
}

type Effects5 struct {
	RelaxedRating    float64 `json:"relaxed"`
	PainReliefRating float64 `json:"painRelief"`
	SleepyRating     float64 `json:"sleepy"`
	HappyRating      float64 `json:"happy"`
	EuphoricRating   float64 `json:"euphoric"`
}

type WeightEquivalent5 struct {
	Value    float64 `json:"value"`
	Unit     string  `json:"unit"`
	TypeName string  `json:"__typename"`
}

func InsertDataProductvapeorizersCommon(db *sql.DB, target vapeorizersData) {

	for i := range target.Data.FilteredProducts.Products {

		stmt, err := db.Prepare("INSERT INTO Product_vapeorizersCommon (EffDate, DispensaryID, EnterpriseProductID, ProductID, BrandName, ProductName, ProductSubCategory, Weight, WeightGrams, RecOnly, RecPrice, RecSpecialPrice, MedicalOnly, MedicalPrice, MedicalSpecialPrice, WholeSalePrice, InventoryQuantity, InventoryQuantityAvail, InventoryQuantityKiosk, StrainType,  THCAmount, THCUnit, THCTypeName, CBDAMOUNT, CBDUnit, CBDTypeName, RexaledRating, PainReliefRating, SleepyRating, HappyRating, EuphoricRating, ProductImage) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
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

func getvapeorizersData(db *sql.DB, dispensaryIDs []string, dispensaryTargets []vapeorizersData) {

	for i, Dispensary := range dispensaryIDs {

		DispensaryAPILink := "https://dutchie.com/graphql?operationName=FilteredProducts&variables=%7B%22includeEnterpriseSpecials%22%3Afalse%2C%22includeCannabinoids%22%3Atrue%2C%22productsFilter%22%3A%7B%22dispensaryId%22%3A%22" + Dispensary + "%22%2C%22pricingType%22%3A%22rec%22%2C%22strainTypes%22%3A%5B%5D%2C%22subcategories%22%3A%5B%5D%2C%22Status%22%3A%22Active%22%2C%22types%22%3A%5B%22Vaporizers%22%5D%2C%22useCache%22%3Afalse%2C%22sortDirection%22%3A1%2C%22sortBy%22%3A%22price%22%2C%22isDefaultSort%22%3Atrue%2C%22bypassOnlineThresholds%22%3Afalse%2C%22isKioskMenu%22%3Afalse%2C%22removeProductsBelowOptionThresholds%22%3Atrue%7D%2C%22page%22%3A0%2C%22perPage%22%3A1000%7D&extensions=%7B%22persistedQuery%22%3A%7B%22version%22%3A1%2C%22sha256Hash%22%3A%22cf45528052df8664bf9404f98344719a5d7c6f8fff90e8c3c8ef2eaca63b0e9e%22%7D%7D"
		err := getJson(DispensaryAPILink, Dispensary+"vapeorizersData", &dispensaryTargets[i])

		if err != nil {
			fmt.Println(err)
		}
	}

	for _, vapeorizersMenu := range dispensaryTargets {
		InsertDataProductvapeorizersCommon(db, vapeorizersMenu)

	}

}

var vapeorizersDataSlice = []vapeorizersData{
	legalgreensbrocktonvapeorizers, commonwealthalternativecarebrocktonrecvapeorizers, commonwealthalternativecarebrocktonvapeorizers, atlanticmedicinalpartnersinc1vapeorizers, ingoodhealthmedicalvapeorizers, ingoodhealthbrocktonvapeorizers, panaceawellnessquincyvapeorizers, cannapidispensaryvapeorizers, brocktongreenheartvapeorizers, boterabrocktonvapeorizers, budsgoodsandprovisionsabingtonvapeorizers, cannavanavapeorizers, nativesunnorthattlebororetailrecvapeorizers, yambavapeorizers, budsgoodsandprovisionswatertownvapeorizers, siranaturalsvapeorizers, mistymountainshopmaldenvapeorizers, gardenremediesmelroserecvapeorizers, gardenremediesmelrosevapeorizers, oldeworldremediesvapeorizers, naturesmedicinesfallrivervapeorizers, goodchemistrylynnvapeorizers, diemlynnvapeorizers, calyxpeakcompaniesdbalocalcannabiscovapeorizers, newleafenterprisesvapeorizers, insasalemvapeorizers, insaavon1vapeorizers, rhelmcannabisvapeorizers, theorywellnessdeliveryzone1vapeorizers, theorywellnessbridgewatervapeorizers, alternativecompassionservicesvapeorizers, thehealthcirclevapeorizers, greenrockcannabisvapeorizers, releafalternativesvapeorizers, ermontincvapeorizers, greathitscannabisretailrectautonvapeorizers, curaleafmahanovervapeorizers, panaceawellnessrecvapeorizers, panaceawellnessvapeorizers, quincycannabisquincyretailrecvapeorizers, commonwealthalternativecarevapeorizers, commonwealthalternativecare1vapeorizers, terpsattleborovapeorizers, goodchemistryvapeorizers, goodchemistrymedvapeorizers, bloom1vapeorizers, diemworcestervapeorizers, cookiesworcestervapeorizers, mayvapeorizersworcestervapeorizers, budsgoodsandprovisionsworcestervapeorizers, resinateworcestervapeorizers, missionworcestervapeorizers, missionworcesteradultusevapeorizers, clearskyworcestervapeorizers, greencarecollectiveretailmedvapeorizers, thevault1vapeorizers, naturesremedymedvapeorizers, naturesremedy1vapeorizers, harmonyofmavapeorizers, campfirecannabis1vapeorizers, discerndcannabisgraftonretailvapeorizers, curaleafmaoxfordvapeorizers, curaleafmaoxford1vapeorizers, kosavapeorizers, greengoldvapeorizers, societycannabiscovapeorizers, newenglandharvestvapeorizers, thevaultwebstervapeorizers, terpscharltonvapeorizers, temescalwellnesshudsonvapeorizers, hudsonadultuseretailvapeorizers, capitalcanabiscommunitydispensaryvapeorizers, dmaholdingsllcdbagreatesthitscannabisvapeorizers, gardenremediesmarlboroughvapeorizers, gardenremediesrecmarlboroughvapeorizers, greenmeadowsfarmvapeorizers, cannabrollcdbagreenpathsouthbridgevapeorizers, GreenGoldMarboroughvapeorizers, localrootssturbridgevapeorizers, greenngovapeorizers, blackstonevalleycannabisvapeorizers, highhopesvapeorizers, greeneravapeorizers, thriveofmadbaboomxshirleyvapeorizers, ethosfitchburgvapeorizers, atlanticmedicinalpartnersincvapeorizers, atlanticmedicinalpartnersvapeorizers, greenmeadowfarmfitchburgvapeorizers, pioneercannabiscompanyvapeorizers, naturesmedicinesuxbridgevapeorizers, uptopframinghamrecvapeorizers, uptopframinghamvapeorizers, masswellspringmaynardvapeorizers, bountifulfarmsvapeorizers, bleafwellnesscentrevapeorizers, thehealingcenter1vapeorizers, temescalwellnessframinghamvapeorizers, botrealtyllcdbaomgcannabisvapeorizers, curaleafmawarevapeorizers, localrootsfitchburgvapeorizers, ledlloyd420vapeorizers, boterafranklin1vapeorizers, masswellspringactonvapeorizers, ddmcannabisvapeorizers, centralavecompassionatecarevapeorizers, gagecannabiscovapeorizers, netafranklinmedvapeorizers, netafranklinvapeorizers, communitycarecollective1vapeorizers, unitedcultivationllcvapeorizers, budbarnwinchendonvapeorizers, siranaturalsneedhamvapeorizers, umavapeorizerss1vapeorizers, clearskybelchertownvapeorizers, thebostongardenvapeorizers, redivapeorizers, elev8cannabisatholvapeorizers, toytownprojectllcdbatoytownhealthcarevapeorizers, gardenremediesnewtonvapeorizers, gardenremediesrecnewtonvapeorizers, budsgoodsandprovisionswatertownvapeorizers2, nativesunnorthattlebororetailrecvapeorizers2, auraofrhodeislandvapeorizers, ayrdispensarywatertownmedvapeorizers, ayrdispensarywatertownvapeorizers, communitycarecollectivevapeorizers, uptopwestroxburyvapeorizers, orangecannabiscompanyvapeorizers, EthosWatertownAdultUsevapeorizers, EthosWatertownMedicalvapeorizers, apothcaarlingtonvapeorizers, releafalternativesvapeorizers2, ocsivymaerealestatedbazaharacannabisvapeorizers, motherearthpawtucketvapeorizers, kgcollectivecambridgevapeorizers, revolutionaryclinicsfreshpondvapeorizers, naturesremedy2vapeorizers, highprofileroslindalevapeorizers, mayvapeorizersmedicinalsallstonrecvapeorizers, mayvapeorizersallstonvapeorizers, apothcajamaicaplainvapeorizers, netabrooklinevapeorizers, netabrooklinemedvapeorizers, missionbrooklinevapeorizers, redcardinal2vapeorizers, libertyspringfieldvapeorizers, canacraftvapeorizers, siranaturalsvapeorizers2, siranaturalssomervilleadultusevapeorizers, medmenfenwayvapeorizers, westernfront1vapeorizers, smythcannabiscovapeorizers, massalternativecareamherstmedvapeorizers, massalternativecareamherstvapeorizers, pleasantreesamherstvapeorizers, libertysomervillevapeorizers, theheirloomcollectivevapeorizers, pureoasis1vapeorizers, advesamaincdbablueriverterpscambridgevapeorizers, insaspringfieldvapeorizers, ayrdispensarybackbayvapeorizers, reverie73vapeorizers, massalternativecarevapeorizers, massalternativecaremedvapeorizers, revolutionaryclinicssomervillevapeorizers, theorywellnesschicopeemedvapeorizers, ethosdorchestervapeorizers, highprofiledorchestervapeorizers, lazyriverproductsvapeorizers, mistymountainshopmaldenvapeorizers2, jimbuddysvapeorizers, apexnoirevapeorizers, cannapidispensaryvapeorizers2, ermontincvapeorizers2, dazedcannabisvapeorizers,
}

var legalgreensbrocktonvapeorizers vapeorizersData
var commonwealthalternativecarebrocktonrecvapeorizers vapeorizersData
var commonwealthalternativecarebrocktonvapeorizers vapeorizersData
var atlanticmedicinalpartnersinc1vapeorizers vapeorizersData
var ingoodhealthmedicalvapeorizers vapeorizersData
var ingoodhealthbrocktonvapeorizers vapeorizersData
var panaceawellnessquincyvapeorizers vapeorizersData
var cannapidispensaryvapeorizers vapeorizersData
var brocktongreenheartvapeorizers vapeorizersData
var boterabrocktonvapeorizers vapeorizersData
var budsgoodsandprovisionsabingtonvapeorizers vapeorizersData
var cannavanavapeorizers vapeorizersData
var nativesunnorthattlebororetailrecvapeorizers vapeorizersData
var yambavapeorizers vapeorizersData
var budsgoodsandprovisionswatertownvapeorizers vapeorizersData
var siranaturalsvapeorizers vapeorizersData
var mistymountainshopmaldenvapeorizers vapeorizersData
var gardenremediesmelroserecvapeorizers vapeorizersData
var gardenremediesmelrosevapeorizers vapeorizersData
var oldeworldremediesvapeorizers vapeorizersData
var naturesmedicinesfallrivervapeorizers vapeorizersData
var goodchemistrylynnvapeorizers vapeorizersData
var diemlynnvapeorizers vapeorizersData
var calyxpeakcompaniesdbalocalcannabiscovapeorizers vapeorizersData
var newleafenterprisesvapeorizers vapeorizersData
var insasalemvapeorizers vapeorizersData
var insaavon1vapeorizers vapeorizersData
var rhelmcannabisvapeorizers vapeorizersData
var theorywellnessdeliveryzone1vapeorizers vapeorizersData
var theorywellnessbridgewatervapeorizers vapeorizersData
var alternativecompassionservicesvapeorizers vapeorizersData
var thehealthcirclevapeorizers vapeorizersData
var greenrockcannabisvapeorizers vapeorizersData
var releafalternativesvapeorizers vapeorizersData
var ermontincvapeorizers vapeorizersData
var greathitscannabisretailrectautonvapeorizers vapeorizersData
var curaleafmahanovervapeorizers vapeorizersData
var panaceawellnessrecvapeorizers vapeorizersData
var panaceawellnessvapeorizers vapeorizersData
var quincycannabisquincyretailrecvapeorizers vapeorizersData
var commonwealthalternativecarevapeorizers vapeorizersData
var commonwealthalternativecare1vapeorizers vapeorizersData
var terpsattleborovapeorizers vapeorizersData
var goodchemistryvapeorizers vapeorizersData
var goodchemistrymedvapeorizers vapeorizersData
var bloom1vapeorizers vapeorizersData
var diemworcestervapeorizers vapeorizersData
var cookiesworcestervapeorizers vapeorizersData
var mayvapeorizersworcestervapeorizers vapeorizersData
var budsgoodsandprovisionsworcestervapeorizers vapeorizersData
var resinateworcestervapeorizers vapeorizersData
var missionworcestervapeorizers vapeorizersData
var missionworcesteradultusevapeorizers vapeorizersData
var clearskyworcestervapeorizers vapeorizersData
var greencarecollectiveretailmedvapeorizers vapeorizersData
var thevault1vapeorizers vapeorizersData
var naturesremedymedvapeorizers vapeorizersData
var naturesremedy1vapeorizers vapeorizersData
var harmonyofmavapeorizers vapeorizersData
var campfirecannabis1vapeorizers vapeorizersData
var discerndcannabisgraftonretailvapeorizers vapeorizersData
var curaleafmaoxfordvapeorizers vapeorizersData
var curaleafmaoxford1vapeorizers vapeorizersData
var kosavapeorizers vapeorizersData
var greengoldvapeorizers vapeorizersData
var societycannabiscovapeorizers vapeorizersData
var newenglandharvestvapeorizers vapeorizersData
var thevaultwebstervapeorizers vapeorizersData
var terpscharltonvapeorizers vapeorizersData
var temescalwellnesshudsonvapeorizers vapeorizersData
var hudsonadultuseretailvapeorizers vapeorizersData
var capitalcanabiscommunitydispensaryvapeorizers vapeorizersData
var dmaholdingsllcdbagreatesthitscannabisvapeorizers vapeorizersData
var gardenremediesmarlboroughvapeorizers vapeorizersData
var gardenremediesrecmarlboroughvapeorizers vapeorizersData
var greenmeadowsfarmvapeorizers vapeorizersData
var cannabrollcdbagreenpathsouthbridgevapeorizers vapeorizersData
var GreenGoldMarboroughvapeorizers vapeorizersData
var localrootssturbridgevapeorizers vapeorizersData
var greenngovapeorizers vapeorizersData
var blackstonevalleycannabisvapeorizers vapeorizersData
var highhopesvapeorizers vapeorizersData
var greeneravapeorizers vapeorizersData
var thriveofmadbaboomxshirleyvapeorizers vapeorizersData
var ethosfitchburgvapeorizers vapeorizersData
var atlanticmedicinalpartnersincvapeorizers vapeorizersData
var atlanticmedicinalpartnersvapeorizers vapeorizersData
var greenmeadowfarmfitchburgvapeorizers vapeorizersData
var pioneercannabiscompanyvapeorizers vapeorizersData
var naturesmedicinesuxbridgevapeorizers vapeorizersData
var uptopframinghamrecvapeorizers vapeorizersData
var uptopframinghamvapeorizers vapeorizersData
var masswellspringmaynardvapeorizers vapeorizersData
var bountifulfarmsvapeorizers vapeorizersData
var bleafwellnesscentrevapeorizers vapeorizersData
var thehealingcenter1vapeorizers vapeorizersData
var temescalwellnessframinghamvapeorizers vapeorizersData
var botrealtyllcdbaomgcannabisvapeorizers vapeorizersData
var curaleafmawarevapeorizers vapeorizersData
var localrootsfitchburgvapeorizers vapeorizersData
var ledlloyd420vapeorizers vapeorizersData
var boterafranklin1vapeorizers vapeorizersData
var masswellspringactonvapeorizers vapeorizersData
var ddmcannabisvapeorizers vapeorizersData
var centralavecompassionatecarevapeorizers vapeorizersData
var gagecannabiscovapeorizers vapeorizersData
var netafranklinmedvapeorizers vapeorizersData
var netafranklinvapeorizers vapeorizersData
var communitycarecollective1vapeorizers vapeorizersData
var unitedcultivationllcvapeorizers vapeorizersData
var budbarnwinchendonvapeorizers vapeorizersData
var siranaturalsneedhamvapeorizers vapeorizersData
var umavapeorizerss1vapeorizers vapeorizersData
var clearskybelchertownvapeorizers vapeorizersData
var thebostongardenvapeorizers vapeorizersData
var redivapeorizers vapeorizersData
var elev8cannabisatholvapeorizers vapeorizersData
var toytownprojectllcdbatoytownhealthcarevapeorizers vapeorizersData
var gardenremediesnewtonvapeorizers vapeorizersData
var gardenremediesrecnewtonvapeorizers vapeorizersData
var budsgoodsandprovisionswatertownvapeorizers2 vapeorizersData
var nativesunnorthattlebororetailrecvapeorizers2 vapeorizersData
var auraofrhodeislandvapeorizers vapeorizersData
var ayrdispensarywatertownmedvapeorizers vapeorizersData
var ayrdispensarywatertownvapeorizers vapeorizersData
var communitycarecollectivevapeorizers vapeorizersData
var uptopwestroxburyvapeorizers vapeorizersData
var orangecannabiscompanyvapeorizers vapeorizersData
var EthosWatertownAdultUsevapeorizers vapeorizersData
var EthosWatertownMedicalvapeorizers vapeorizersData
var apothcaarlingtonvapeorizers vapeorizersData
var releafalternativesvapeorizers2 vapeorizersData
var ocsivymaerealestatedbazaharacannabisvapeorizers vapeorizersData
var motherearthpawtucketvapeorizers vapeorizersData
var kgcollectivecambridgevapeorizers vapeorizersData
var revolutionaryclinicsfreshpondvapeorizers vapeorizersData
var naturesremedy2vapeorizers vapeorizersData
var highprofileroslindalevapeorizers vapeorizersData
var mayvapeorizersmedicinalsallstonrecvapeorizers vapeorizersData
var mayvapeorizersallstonvapeorizers vapeorizersData
var apothcajamaicaplainvapeorizers vapeorizersData
var netabrooklinevapeorizers vapeorizersData
var netabrooklinemedvapeorizers vapeorizersData
var missionbrooklinevapeorizers vapeorizersData
var redcardinal2vapeorizers vapeorizersData
var libertyspringfieldvapeorizers vapeorizersData
var canacraftvapeorizers vapeorizersData
var siranaturalsvapeorizers2 vapeorizersData
var siranaturalssomervilleadultusevapeorizers vapeorizersData
var medmenfenwayvapeorizers vapeorizersData
var westernfront1vapeorizers vapeorizersData
var smythcannabiscovapeorizers vapeorizersData
var massalternativecareamherstmedvapeorizers vapeorizersData
var massalternativecareamherstvapeorizers vapeorizersData
var pleasantreesamherstvapeorizers vapeorizersData
var libertysomervillevapeorizers vapeorizersData
var theheirloomcollectivevapeorizers vapeorizersData
var pureoasis1vapeorizers vapeorizersData
var advesamaincdbablueriverterpscambridgevapeorizers vapeorizersData
var insaspringfieldvapeorizers vapeorizersData
var ayrdispensarybackbayvapeorizers vapeorizersData
var reverie73vapeorizers vapeorizersData
var massalternativecarevapeorizers vapeorizersData
var massalternativecaremedvapeorizers vapeorizersData
var revolutionaryclinicssomervillevapeorizers vapeorizersData
var theorywellnesschicopeemedvapeorizers vapeorizersData
var ethosdorchestervapeorizers vapeorizersData
var highprofiledorchestervapeorizers vapeorizersData
var lazyriverproductsvapeorizers vapeorizersData
var mistymountainshopmaldenvapeorizers2 vapeorizersData
var jimbuddysvapeorizers vapeorizersData
var apexnoirevapeorizers vapeorizersData
var cannapidispensaryvapeorizers2 vapeorizersData
var ermontincvapeorizers2 vapeorizersData
var dazedcannabisvapeorizers vapeorizersData
