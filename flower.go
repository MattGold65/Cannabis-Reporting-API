package main

import (
	"encoding/json"
	"fmt"
)

var flower productData

type productData struct {
	Data DataSet `json:"data"`
}

type DataSet struct {
	FilteredProducts Product `json:"filteredProducts"`
}

type Product struct {
	Products []flowerProducts `json:"products"`
}

type flowerProducts struct {
	Id                  string       `json:"_id"`
	BrandName           string       `json:"brandName"`
	CBD                 string       `json:"CBD"`
	CBDContent          CBDContent   `json:"CBDContent"`
	ComingSoon          string       `json:"comingSoon"`
	CreatedAt           string       `json:"createdAt"`
	DispensaryID        string       `json:"DispensaryID"`
	EnterpriseProductId string       `json:"enterpriseProductId"`
	Image               string       `json:"Image"`
	Measurements        measurements `json:"measurements"`
	MedicalOnly         bool         `json:"medicalOnly"`
	MedicalPrice        []float64    `json:"medicalPrices"`
	MedicalSpecialPrice []float64    `json:"medicalSpecialPrices"`
	WholeSalePrice      []float64    `json:"wholesalePrices"`
	ProductName         string       `json:"Name"`
	NonArmsLength       string       `json:"nonArmsLength"`
	WeightOptions       []string     `json:"Options"`
	CustomerLimit       string       `json:"limitsPerCustomer"`
	ManualInventory     string       `json:"manualInventory"`
	POSMetaData         MetaData     `json:"POSMetaData"`
	Price               []float64    `json:"Prices"`
	//PricingTierData     string     `json:"pricingTierData"`
	RecOnly         bool      `json:"recOnly"`
	RecPrice        []float64 `json:"recPrices"`
	RecSpecialPrice []float64 `json:"recSpecialPrices"`
	Special         bool      `json:"special"`
	//SpecialData                []string     `json:"specialData"`
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

func getJson(url string, target interface{}) error {
	resp, err := client.Get(url)

	if err != nil {
		return fmt.Errorf("ERROR")
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)

}

func generateProductIDs() []string {

	var productId []string

	for _, product := range flower.Data.FilteredProducts.Products {

		productId = append(productId, product.Id)

	}

	return productId

}

func generateProductBrandNames() []string {

	var productBrandNames []string

	for _, product := range flower.Data.FilteredProducts.Products {

		productBrandNames = append(productBrandNames, product.BrandName)

	}

	return productBrandNames

}

func generateProductCBDs() []string {

	var productCBDs []string

	for _, product := range flower.Data.FilteredProducts.Products {

		productCBDs = append(productCBDs, product.CBD)

	}

	return productCBDs

}

func generateProductComingSoon() []string {

	var productComingSoon []string

	for _, product := range flower.Data.FilteredProducts.Products {

		productComingSoon = append(productComingSoon, product.ComingSoon)

	}

	return productComingSoon

}

func generateProductCreatedAt() []string {

	var productCreatedAt []string

	for _, product := range flower.Data.FilteredProducts.Products {

		productCreatedAt = append(productCreatedAt, product.CreatedAt)

	}

	return productCreatedAt

}

func generateProductDispensaryID() []string {

	var productDispensaryID []string

	for _, product := range flower.Data.FilteredProducts.Products {

		productDispensaryID = append(productDispensaryID, product.DispensaryID)

	}

	return productDispensaryID

}

func generateProductEnterpriseProductID() []string {

	var productEnterpriseProductId []string

	for _, product := range flower.Data.FilteredProducts.Products {

		productEnterpriseProductId = append(productEnterpriseProductId, product.EnterpriseProductId)

	}

	return productEnterpriseProductId

}

func generateProductImage() []string {

	var productImage []string

	for _, product := range flower.Data.FilteredProducts.Products {

		productImage = append(productImage, product.Image)

	}

	return productImage

}

func generateProductMedicalOnly() []bool {

	var productMedicalOnly []bool

	for _, product := range flower.Data.FilteredProducts.Products {

		productMedicalOnly = append(productMedicalOnly, product.MedicalOnly)

	}

	return productMedicalOnly

}

func generateProductMedicalPrice() [][]float64 {

	var productMedicalPrice [][]float64

	for _, product := range flower.Data.FilteredProducts.Products {

		productMedicalPrice = append(productMedicalPrice, product.MedicalPrice)

	}

	return productMedicalPrice

}

func generateProductMedicalSpecialPrice() []float64 {

	var productMedicalSpecialPrice []float64

	for _, product := range flower.Data.FilteredProducts.Products {

		productMedicalSpecialPrice = append(productMedicalSpecialPrice, product.MedicalSpecialPrice...)

	}

	return productMedicalSpecialPrice

}

func generateProductWholesalePrice() [][]float64 {

	var productWholeSalePrice [][]float64

	for _, product := range flower.Data.FilteredProducts.Products {

		productWholeSalePrice = append(productWholeSalePrice, product.WholeSalePrice)

	}

	return productWholeSalePrice

}

func generateProductNames() []string {

	var productNames []string

	for _, product := range flower.Data.FilteredProducts.Products {

		productNames = append(productNames, product.ProductName)

	}

	return productNames

}

func generateNonArmsLength() []string {

	var productNonArmsLength []string

	for _, product := range flower.Data.FilteredProducts.Products {

		productNonArmsLength = append(productNonArmsLength, product.NonArmsLength)

	}

	return productNonArmsLength

}

func generateWeightOptions() [][]string {

	var productWeightOptions [][]string

	for _, product := range flower.Data.FilteredProducts.Products {

		productWeightOptions = append(productWeightOptions, product.WeightOptions)

	}

	return productWeightOptions

}

func generateProductTHCRange() []float64 {
	var thcPercentages []float64

	for i := 0; i < len(flower.Data.FilteredProducts.Products); i++ {
		thcPercentages = append(thcPercentages, flower.Data.FilteredProducts.Products[i].THCContent.Range...)

	}

	return thcPercentages
}
