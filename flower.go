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
