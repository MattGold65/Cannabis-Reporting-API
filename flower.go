package main

import (
	"database/sql"
	"fmt"
	"time"
)

type FlowerData struct {
	Data DataSet `json:"data"`
}

type DataSet struct {
	FilteredProducts Product `json:"filteredProducts"`
}

type Product struct {
	Products []flowerProducts `json:"products"`
}

type flowerProducts struct {
	Id                         string                `json:"_id"`
	BrandName                  string                `json:"brandName"`
	CBD                        string                `json:"CBD"`
	CBDContent                 CBDContent            `json:"CBDContent"`
	ComingSoon                 string                `json:"comingSoon"`
	CreatedAt                  string                `json:"createdAt"`
	DispensaryID               string                `json:"DispensaryID"`
	EnterpriseProductId        string                `json:"enterpriseProductId"`
	Image                      string                `json:"Image"`
	Measurements               measurements          `json:"measurements"`
	MedicalOnly                bool                  `json:"medicalOnly"`
	MedicalPrice               []float64             `json:"medicalPrices"`
	MedicalSpecialPrice        []float64             `json:"medicalSpecialPrices"`
	WholeSalePrice             []float64             `json:"wholesalePrices"`
	ProductName                string                `json:"Name"`
	NonArmsLength              string                `json:"nonArmsLength"`
	WeightOptions              []string              `json:"Options"`
	CustomerLimit              string                `json:"limitsPerCustomer"`
	ManualInventory            string                `json:"manualInventory"`
	POSMetaData                MetaData              `json:"POSMetaData"`
	Price                      []float64             `json:"Prices"`
	PricingTierData            string                `json:"pricingTierData"`
	RecOnly                    bool                  `json:"recOnly"`
	RecPrice                   []float64             `json:"recPrices"`
	RecSpecialPrice            []float64             `json:"recSpecialPrices"`
	Special                    bool                  `json:"special"`
	SpecialData                []string              `json:"specialData"`
	Status                     string                `json:"Status"`
	StrainType                 string                `json:"strainType"`
	SubCategory                string                `json:"subCategory"`
	THC                        string                `json:"THC"`
	THCContent                 THCContent            `json:"THCContent"`
	Type                       string                `json:"type"`
	VapeTaxApplicable          bool                  `json:"vapeTaxApplicable"`
	Weight                     float64               `json:"weight"`
	Featured                   string                `json:"featured"`
	IsBelowThreshold           bool                  `json:"isBelowThreshold"`
	IsBelowKioskThreshold      bool                  `json:"isBelowKioskThreshold"`
	OptionsBelowThreshold      []string              `json:"optionsBelowThreshold"`
	OptionsBelowKioskThreshold []string              `json:"optionsBelowKioskThreshold"`
	CName                      string                `json:"cName"`
	PastCNames                 []string              `json:"pastCNames"`
	BrandLogo                  string                `json:"brandLogo"`
	BottleDepositTaxCents      string                `json:"bottleDepositTaxCents"`
	Typename                   string                `json:"__typename"`
	BrandInfo                  BrandInfo             `json:"brand"`
	Effects                    Effects               `json:"effects"`
	CannabinoidsContent        []CannabinoidsContent `json:"cannabinoidsV2"`
}
type THCContent struct {
	Unit     string    `json:"unit"`
	Range    []float64 `json:"range"`
	Typename string    `json:"__typename"`
}
type CBDContent struct {
	Unit     string    `json:"unit"`
	Range    []float64 `json:"range"`
	Typename string    `json:"__typename"`
}

type CannabinoidsContent struct {
	Value           float64         `json:"value"`
	Unit            string          `json:"PERCENTAGE"`
	CannabinoidType CannabinoidType `json:"cannabinoid"`
	TypeName        string          `json:"__typename"`
}

type CannabinoidType struct {
	Name     string `json:"name"`
	TypeName string `json:"__typename"`
}

type measurements struct {
	NetWeight weight `json:"netWeight"`
	Volume    string `json:"volume"`
	Typename  string `json:"__typename"`
}

type weight struct {
	Unit     string `json:"unit"`
	Value    []int  `json:"values"`
	Typename string `json:"__typename"`
}

type MetaData struct {
	CanonicalID        string               `json:"canonicalID"`
	CanonicalBrandName string               `json:"canonicalBrandName"`
	Children           []AdditionalMetaData `json:"children"`
	TypeName           string               `json:"__typename"`
}

type AdditionalMetaData struct {
	CanonicalEffectivePotencyMg float64          `json:"canonicalEffectivePotencyMg"`
	SizeOptions                 string           `json:"option"`
	InventoryQuantity           int              `json:"quantity"`
	InventoryQuantityAvailable  int              `json:"quantityAvailable"`
	InventoryKioskQuantity      int              `json:"kioskQuantityAvailable"`
	StandardEquivalent          WeightEquivalent `json:"standardEquivalent"`
}

type BrandInfo struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	ImageURL    string `json:"imageURL"`
	Name        string `json:"name"`
}

type Effects struct {
	RelaxedRating    float64 `json:"relaxed"`
	PainReliefRating float64 `json:"painRelief"`
	SleepyRating     float64 `json:"sleepy"`
	HappyRating      float64 `json:"happy"`
	EuphoricRating   float64 `json:"euphoric"`
}

type WeightEquivalent struct {
	Value    float64 `json:"value"`
	Unit     string  `json:"unit"`
	TypeName string  `json:"__typename"`
}

func InsertDataProductFlowerCommon(db *sql.DB, target FlowerData) {

	//test for now turn this into function where we can call this loop on every single dispensary json
	for i := range target.Data.FilteredProducts.Products {

		stmt, err := db.Prepare("INSERT INTO Product_FlowerCommon (EffDate, DispensaryID, EnterpriseProductID, ProductID, BrandName, ProductName, Weight, WeightGrams, RecOnly, RecPrice, RecSpecialPrice, MedicalOnly, MedicalPrice, MedicalSpecialPrice, WholeSalePrice, InventoryQuantity, InventoryQuantityAvail, InventoryQuantityKiosk, StrainType,  THCAmount, THCUnit, THCTypeName, CBDAMOUNT, CBDUnit, CBDTypeName, RexaledRating, PainReliefRating, SleepyRating, HappyRating, EuphoricRating, ProductImage) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
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

			if CheckLastCharacter(gramCheck) == "g" {

				strGrams := GrabRestOfString(weight1)
				WeightG, err = ConvertToFloat(strGrams)

				if err != nil {
					fmt.Println("Issue converting gram value to float64")
				}

			}

			if weight1 == "1/8oz" {
				WeightG = 3.5
			} else if weight1 == "1/4oz" {
				WeightG = 7.0
			} else if weight1 == "1/2oz" {
				WeightG = 14.0
			} else if weight1 == "1oz" {
				WeightG = 28.0
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
			_, err = stmt.Exec(effDate, target.Data.FilteredProducts.Products[i].DispensaryID, target.Data.FilteredProducts.Products[i].EnterpriseProductId, target.Data.FilteredProducts.Products[i].Id, target.Data.FilteredProducts.Products[i].BrandName, target.Data.FilteredProducts.Products[i].ProductName, weight1, WeightG, target.Data.FilteredProducts.Products[i].RecOnly, recPrice, recSpecialPrice, target.Data.FilteredProducts.Products[i].MedicalOnly, medPrice, medSpecialPrice, wholePrice, InventoryQuantity, InventoryQuantityAvailable, InventoryKioskQuantity, target.Data.FilteredProducts.Products[i].StrainType, thcAmount, target.Data.FilteredProducts.Products[i].THCContent.Unit, target.Data.FilteredProducts.Products[i].THCContent.Typename, cbdAmount, target.Data.FilteredProducts.Products[i].CBDContent.Unit, target.Data.FilteredProducts.Products[i].CBDContent.Typename, target.Data.FilteredProducts.Products[i].Effects.RelaxedRating, target.Data.FilteredProducts.Products[i].Effects.PainReliefRating, target.Data.FilteredProducts.Products[i].Effects.SleepyRating, target.Data.FilteredProducts.Products[i].Effects.HappyRating, target.Data.FilteredProducts.Products[i].Effects.EuphoricRating, target.Data.FilteredProducts.Products[i].Image)
			if err != nil {
				fmt.Println(err)
				return
			}

		}

	}
}

func getFlowerData(db *sql.DB, dispensaryIDs []string, dispensaryTargets []FlowerData) {

	for i, Dispensary := range dispensaryIDs {

		DispensaryAPILink := "https://dutchie.com/graphql?operationName=FilteredProducts&variables=%7B%22includeEnterpriseSpecials%22%3Afalse%2C%22includeCannabinoids%22%3Atrue%2C%22productsFilter%22%3A%7B%22dispensaryId%22%3A%22" + Dispensary + "%22%2C%22pricingType%22%3A%22med%22%2C%22strainTypes%22%3A%5B%5D%2C%22subcategories%22%3A%5B%5D%2C%22Status%22%3A%22Active%22%2C%22types%22%3A%5B%22Flower%22%5D%2C%22useCache%22%3Afalse%2C%22sortDirection%22%3A1%2C%22sortBy%22%3Anull%2C%22isDefaultSort%22%3Atrue%2C%22bypassOnlineThresholds%22%3Afalse%2C%22isKioskMenu%22%3Afalse%2C%22removeProductsBelowOptionThresholds%22%3Atrue%7D%2C%22page%22%3A0%2C%22perPage%22%3A1000%7D&extensions=%7B%22persistedQuery%22%3A%7B%22version%22%3A1%2C%22sha256Hash%22%3A%220e884328c01ef8ed540d4bbd27101ee58fedaf5cddda6c499169209b53fdf574%22%7D%7D"

		err := getJson(DispensaryAPILink, Dispensary+"FlowerData", &dispensaryTargets[i])

		if err != nil {
			fmt.Println(err)
		}
	}

	for _, FlowerMenu := range dispensaryTargets {
		InsertDataProductFlowerCommon(db, FlowerMenu)

	}

}

var flowerDataSlice = []FlowerData{
	legalgreensbrocktonFlower, commonwealthalternativecarebrocktonrecFlower, commonwealthalternativecarebrocktonFlower, atlanticmedicinalpartnersinc1Flower, ingoodhealthmedicalFlower, ingoodhealthbrocktonFlower, panaceawellnessquincyFlower, cannapidispensaryFlower, brocktongreenheartFlower, boterabrocktonFlower, budsgoodsandprovisionsabingtonFlower, cannavanaFlower, nativesunnorthattlebororetailrecFlower, yambaFlower, budsgoodsandprovisionswatertownFlower, siranaturalsFlower, mistymountainshopmaldenFlower, gardenremediesmelroserecFlower, gardenremediesmelroseFlower, oldeworldremediesFlower, naturesmedicinesfallriverFlower, goodchemistrylynnFlower, diemlynnFlower, calyxpeakcompaniesdbalocalcannabiscoFlower, newleafenterprisesFlower, insasalemFlower, insaavon1Flower, rhelmcannabisFlower, theorywellnessdeliveryzone1Flower, theorywellnessbridgewaterFlower, alternativecompassionservicesFlower, thehealthcircleFlower, greenrockcannabisFlower, releafalternativesFlower, ermontincFlower, greathitscannabisretailrectautonFlower, curaleafmahanoverFlower, panaceawellnessrecFlower, panaceawellnessFlower, quincycannabisquincyretailrecFlower, commonwealthalternativecareFlower, commonwealthalternativecare1Flower, terpsattleboroFlower, goodchemistryFlower, goodchemistrymedFlower, bloom1Flower, diemworcesterFlower, cookiesworcesterFlower, mayflowerworcesterFlower, budsgoodsandprovisionsworcesterFlower, resinateworcesterFlower, missionworcesterFlower, missionworcesteradultuseFlower, clearskyworcesterFlower, greencarecollectiveretailmedFlower, thevault1Flower, naturesremedymedFlower, naturesremedy1Flower, harmonyofmaFlower, campfirecannabis1Flower, discerndcannabisgraftonretailFlower, curaleafmaoxfordFlower, curaleafmaoxford1Flower, kosaFlower, greengoldFlower, societycannabiscoFlower, newenglandharvestFlower, thevaultwebsterFlower, terpscharltonFlower, temescalwellnesshudsonFlower, hudsonadultuseretailFlower, capitalcanabiscommunitydispensaryFlower, dmaholdingsllcdbagreatesthitscannabisFlower, gardenremediesmarlboroughFlower, gardenremediesrecmarlboroughFlower, greenmeadowsfarmFlower, cannabrollcdbagreenpathsouthbridgeFlower, GreenGoldMarboroughFlower, localrootssturbridgeFlower, greenngoFlower, blackstonevalleycannabisFlower, highhopesFlower, greeneraFlower, thriveofmadbaboomxshirleyFlower, ethosfitchburgFlower, atlanticmedicinalpartnersincFlower, atlanticmedicinalpartnersFlower, greenmeadowfarmfitchburgFlower, pioneercannabiscompanyFlower, naturesmedicinesuxbridgeFlower, uptopframinghamrecFlower, uptopframinghamFlower, masswellspringmaynardFlower, bountifulfarmsFlower, bleafwellnesscentreFlower, thehealingcenter1Flower, temescalwellnessframinghamFlower, botrealtyllcdbaomgcannabisFlower, curaleafmawareFlower, localrootsfitchburgFlower, ledlloyd420Flower, boterafranklin1Flower, masswellspringactonFlower, ddmcannabisFlower, centralavecompassionatecareFlower, gagecannabiscoFlower, netafranklinmedFlower, netafranklinFlower, communitycarecollective1Flower, unitedcultivationllcFlower, budbarnwinchendonFlower, siranaturalsneedhamFlower, umaflowers1Flower, clearskybelchertownFlower, thebostongardenFlower, rediFlower, elev8cannabisatholFlower, toytownprojectllcdbatoytownhealthcareFlower, gardenremediesnewtonFlower, gardenremediesrecnewtonFlower, budsgoodsandprovisionswatertownFlower2, nativesunnorthattlebororetailrecFlower2, auraofrhodeislandFlower, ayrdispensarywatertownmedFlower, ayrdispensarywatertownFlower, communitycarecollectiveFlower, uptopwestroxburyFlower, orangecannabiscompanyFlower, EthosWatertownAdultUseFlower, EthosWatertownMedicalFlower, apothcaarlingtonFlower, releafalternativesFlower2, ocsivymaerealestatedbazaharacannabisFlower, motherearthpawtucketFlower, kgcollectivecambridgeFlower, revolutionaryclinicsfreshpondFlower, naturesremedy2Flower, highprofileroslindaleFlower, mayflowermedicinalsallstonrecFlower, mayflowerallstonFlower, apothcajamaicaplainFlower, netabrooklineFlower, netabrooklinemedFlower, missionbrooklineFlower, redcardinal2Flower, libertyspringfieldFlower, canacraftFlower, siranaturalsFlower2, siranaturalssomervilleadultuseFlower, medmenfenwayFlower, westernfront1Flower, smythcannabiscoFlower, massalternativecareamherstmedFlower, massalternativecareamherstFlower, pleasantreesamherstFlower, libertysomervilleFlower, theheirloomcollectiveFlower, pureoasis1Flower, advesamaincdbablueriverterpscambridgeFlower, insaspringfieldFlower, ayrdispensarybackbayFlower, reverie73Flower, massalternativecareFlower, massalternativecaremedFlower, revolutionaryclinicssomervilleFlower, theorywellnesschicopeemedFlower, ethosdorchesterFlower, highprofiledorchesterFlower, lazyriverproductsFlower, mistymountainshopmaldenFlower2, jimbuddysFlower, apexnoireFlower, cannapidispensaryFlower2, ermontincFlower2, dazedcannabisFlower,
}

var legalgreensbrocktonFlower FlowerData
var commonwealthalternativecarebrocktonrecFlower FlowerData
var commonwealthalternativecarebrocktonFlower FlowerData
var atlanticmedicinalpartnersinc1Flower FlowerData
var ingoodhealthmedicalFlower FlowerData
var ingoodhealthbrocktonFlower FlowerData
var panaceawellnessquincyFlower FlowerData
var cannapidispensaryFlower FlowerData
var brocktongreenheartFlower FlowerData
var boterabrocktonFlower FlowerData
var budsgoodsandprovisionsabingtonFlower FlowerData
var cannavanaFlower FlowerData
var nativesunnorthattlebororetailrecFlower FlowerData
var yambaFlower FlowerData
var budsgoodsandprovisionswatertownFlower FlowerData
var siranaturalsFlower FlowerData
var mistymountainshopmaldenFlower FlowerData
var gardenremediesmelroserecFlower FlowerData
var gardenremediesmelroseFlower FlowerData
var oldeworldremediesFlower FlowerData
var naturesmedicinesfallriverFlower FlowerData
var goodchemistrylynnFlower FlowerData
var diemlynnFlower FlowerData
var calyxpeakcompaniesdbalocalcannabiscoFlower FlowerData
var newleafenterprisesFlower FlowerData
var insasalemFlower FlowerData
var insaavon1Flower FlowerData
var rhelmcannabisFlower FlowerData
var theorywellnessdeliveryzone1Flower FlowerData
var theorywellnessbridgewaterFlower FlowerData
var alternativecompassionservicesFlower FlowerData
var thehealthcircleFlower FlowerData
var greenrockcannabisFlower FlowerData
var releafalternativesFlower FlowerData
var ermontincFlower FlowerData
var greathitscannabisretailrectautonFlower FlowerData
var curaleafmahanoverFlower FlowerData
var panaceawellnessrecFlower FlowerData
var panaceawellnessFlower FlowerData
var quincycannabisquincyretailrecFlower FlowerData
var commonwealthalternativecareFlower FlowerData
var commonwealthalternativecare1Flower FlowerData
var terpsattleboroFlower FlowerData
var goodchemistryFlower FlowerData
var goodchemistrymedFlower FlowerData
var bloom1Flower FlowerData
var diemworcesterFlower FlowerData
var cookiesworcesterFlower FlowerData
var mayflowerworcesterFlower FlowerData
var budsgoodsandprovisionsworcesterFlower FlowerData
var resinateworcesterFlower FlowerData
var missionworcesterFlower FlowerData
var missionworcesteradultuseFlower FlowerData
var clearskyworcesterFlower FlowerData
var greencarecollectiveretailmedFlower FlowerData
var thevault1Flower FlowerData
var naturesremedymedFlower FlowerData
var naturesremedy1Flower FlowerData
var harmonyofmaFlower FlowerData
var campfirecannabis1Flower FlowerData
var discerndcannabisgraftonretailFlower FlowerData
var curaleafmaoxfordFlower FlowerData
var curaleafmaoxford1Flower FlowerData
var kosaFlower FlowerData
var greengoldFlower FlowerData
var societycannabiscoFlower FlowerData
var newenglandharvestFlower FlowerData
var thevaultwebsterFlower FlowerData
var terpscharltonFlower FlowerData
var temescalwellnesshudsonFlower FlowerData
var hudsonadultuseretailFlower FlowerData
var capitalcanabiscommunitydispensaryFlower FlowerData
var dmaholdingsllcdbagreatesthitscannabisFlower FlowerData
var gardenremediesmarlboroughFlower FlowerData
var gardenremediesrecmarlboroughFlower FlowerData
var greenmeadowsfarmFlower FlowerData
var cannabrollcdbagreenpathsouthbridgeFlower FlowerData
var GreenGoldMarboroughFlower FlowerData
var localrootssturbridgeFlower FlowerData
var greenngoFlower FlowerData
var blackstonevalleycannabisFlower FlowerData
var highhopesFlower FlowerData
var greeneraFlower FlowerData
var thriveofmadbaboomxshirleyFlower FlowerData
var ethosfitchburgFlower FlowerData
var atlanticmedicinalpartnersincFlower FlowerData
var atlanticmedicinalpartnersFlower FlowerData
var greenmeadowfarmfitchburgFlower FlowerData
var pioneercannabiscompanyFlower FlowerData
var naturesmedicinesuxbridgeFlower FlowerData
var uptopframinghamrecFlower FlowerData
var uptopframinghamFlower FlowerData
var masswellspringmaynardFlower FlowerData
var bountifulfarmsFlower FlowerData
var bleafwellnesscentreFlower FlowerData
var thehealingcenter1Flower FlowerData
var temescalwellnessframinghamFlower FlowerData
var botrealtyllcdbaomgcannabisFlower FlowerData
var curaleafmawareFlower FlowerData
var localrootsfitchburgFlower FlowerData
var ledlloyd420Flower FlowerData
var boterafranklin1Flower FlowerData
var masswellspringactonFlower FlowerData
var ddmcannabisFlower FlowerData
var centralavecompassionatecareFlower FlowerData
var gagecannabiscoFlower FlowerData
var netafranklinmedFlower FlowerData
var netafranklinFlower FlowerData
var communitycarecollective1Flower FlowerData
var unitedcultivationllcFlower FlowerData
var budbarnwinchendonFlower FlowerData
var siranaturalsneedhamFlower FlowerData
var umaflowers1Flower FlowerData
var clearskybelchertownFlower FlowerData
var thebostongardenFlower FlowerData
var rediFlower FlowerData
var elev8cannabisatholFlower FlowerData
var toytownprojectllcdbatoytownhealthcareFlower FlowerData
var gardenremediesnewtonFlower FlowerData
var gardenremediesrecnewtonFlower FlowerData
var budsgoodsandprovisionswatertownFlower2 FlowerData
var nativesunnorthattlebororetailrecFlower2 FlowerData
var auraofrhodeislandFlower FlowerData
var ayrdispensarywatertownmedFlower FlowerData
var ayrdispensarywatertownFlower FlowerData
var communitycarecollectiveFlower FlowerData
var uptopwestroxburyFlower FlowerData
var orangecannabiscompanyFlower FlowerData
var EthosWatertownAdultUseFlower FlowerData
var EthosWatertownMedicalFlower FlowerData
var apothcaarlingtonFlower FlowerData
var releafalternativesFlower2 FlowerData
var ocsivymaerealestatedbazaharacannabisFlower FlowerData
var motherearthpawtucketFlower FlowerData
var kgcollectivecambridgeFlower FlowerData
var revolutionaryclinicsfreshpondFlower FlowerData
var naturesremedy2Flower FlowerData
var highprofileroslindaleFlower FlowerData
var mayflowermedicinalsallstonrecFlower FlowerData
var mayflowerallstonFlower FlowerData
var apothcajamaicaplainFlower FlowerData
var netabrooklineFlower FlowerData
var netabrooklinemedFlower FlowerData
var missionbrooklineFlower FlowerData
var redcardinal2Flower FlowerData
var libertyspringfieldFlower FlowerData
var canacraftFlower FlowerData
var siranaturalsFlower2 FlowerData
var siranaturalssomervilleadultuseFlower FlowerData
var medmenfenwayFlower FlowerData
var westernfront1Flower FlowerData
var smythcannabiscoFlower FlowerData
var massalternativecareamherstmedFlower FlowerData
var massalternativecareamherstFlower FlowerData
var pleasantreesamherstFlower FlowerData
var libertysomervilleFlower FlowerData
var theheirloomcollectiveFlower FlowerData
var pureoasis1Flower FlowerData
var advesamaincdbablueriverterpscambridgeFlower FlowerData
var insaspringfieldFlower FlowerData
var ayrdispensarybackbayFlower FlowerData
var reverie73Flower FlowerData
var massalternativecareFlower FlowerData
var massalternativecaremedFlower FlowerData
var revolutionaryclinicssomervilleFlower FlowerData
var theorywellnesschicopeemedFlower FlowerData
var ethosdorchesterFlower FlowerData
var highprofiledorchesterFlower FlowerData
var lazyriverproductsFlower FlowerData
var mistymountainshopmaldenFlower2 FlowerData
var jimbuddysFlower FlowerData
var apexnoireFlower FlowerData
var cannapidispensaryFlower2 FlowerData
var ermontincFlower2 FlowerData
var dazedcannabisFlower FlowerData
