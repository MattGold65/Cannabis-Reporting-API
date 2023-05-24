package main

import (
	"database/sql"
	"fmt"
)

type DispoData struct {
	Data DispoSet `json:"data"`
}

type DispoSet struct {
	FilteredDispnearies []DispensaryInfo `json:"filteredDispensaries"`
}

type DispensaryInfo struct {
	DispensaryName string `json:"name"`
	LogoImage      string `json:"SpecialLogoImage"`
	TypeName       string `json:"__typename"`
	//DeliveryPickup      ActionEstimates    `json:"actionEstimates"`
	//ProductCategories   []ActiveCategories `json:"activeCategories"`
	Address     string `json:"address"`
	BannerImage string `json:"bannerImage"`
	cName       string `json:"cName"`
	//DeliveryHours       DeliveryHours      `json:"deliveryHours"`
	//DeliveryInfo        DeliveryInfo       `json:"deliveryInfo"`
	Description string `json:"description"`
	//DeliveryPickupHours effectiveHours     `json:"effectiveHours"`
	email           string `json:"email"`
	WebsiteURL      string `json:"embedBackUrl"`
	EmbeddedMenuUrl string `json:"embeddedMenuUrl"`
	//FeeTiers            []feeTiers         `json:"feeTiers"`
	GoogleAnalyticsID string `json:"googleAnalyticsID"`
	//HoursOfOperation    HoursOfOperation   `json:"hoursSettings"`
	DispensaryID string   `json:"id"`
	phone        string   `json:"phone"`
	LocationData Location `json:"location"`
}

type Location struct {
	City        string      `json:"city"`
	Street      string      `json:"ln1"`
	Street2     string      `json:"ln2"`
	State       string      `json:"state"`
	ZipCode     string      `json:"zipcode"`
	Coordinates Coordinates `json:"geometry"`
}

type Coordinates struct {
	LadAndLong []float64 `json:"coordinates"`
}

func InsertDataDispoCommon(db *sql.DB, target DispoData) {

	Dispensary := target.Data.FilteredDispnearies[0]

	stmt, err := db.Prepare("INSERT INTO DispensaryCommon (DispensaryID, DispensaryName, Website, Email, Phone, Address, State, City, Street, Street2, Zipcode, Latitude, Longitdue, type) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer stmt.Close()

	// Execute the INSERT statement with values
	_, err = stmt.Exec(Dispensary.DispensaryID, Dispensary.DispensaryName, Dispensary.WebsiteURL, Dispensary.email, Dispensary.phone, Dispensary.Address, Dispensary.LocationData.State, Dispensary.LocationData.City, Dispensary.LocationData.Street, Dispensary.LocationData.Street2, Dispensary.LocationData.ZipCode, Dispensary.LocationData.Coordinates.LadAndLong[0], Dispensary.LocationData.Coordinates.LadAndLong[1], Dispensary.TypeName)
	if err != nil {
		fmt.Println(err)
		return
	}

}

// make sure dispensaryTargets and dispensaryNames have the same ordering within the slice
func getDispensaryData(db *sql.DB, dispensaryNames []string, dispensaryTargets []DispoData) {

	for i, Dispensary := range dispensaryNames {

		DispensaryAPILink := "https://dutchie.com/graphql?operationName=ConsumerDispensaries&variables=%7B%22dispensaryFilter%22%3A%7B%22cNameOrID%22%3A%22" + Dispensary + "%22%7D%7D&extensions=%7B%22persistedQuery%22%3A%7B%22version%22%3A1%2C%22sha256Hash%22%3A%22f44ceb73f96c7fa3016ea2b9c6b68a5b92a62572317849bc4d6d9a724b93c83d%22%7D%7D"

		err := getJson(DispensaryAPILink, Dispensary+"DispensaryData", &dispensaryTargets[i])

		if err != nil {
			fmt.Println(err)
		}
	}

	for _, DispensaryMenu := range dispensaryTargets {
		InsertDataDispoCommon(db, DispensaryMenu)

	}

}

var dispoDataSlice = []DispoData{
	legalgreensbrocktonDispo, commonwealthalternativecarebrocktonrecDispo, commonwealthalternativecarebrocktonDispo, atlanticmedicinalpartnersinc1Dispo, ingoodhealthmedicalDispo, ingoodhealthbrocktonDispo, panaceawellnessquincyDispo, cannapidispensaryDispo, brocktongreenheartDispo, boterabrocktonDispo, budsgoodsandprovisionsabingtonDispo, cannavanaDispo, nativesunnorthattlebororetailrecDispo, yambaDispo, budsgoodsandprovisionswatertownDispo, siranaturalsDispo, mistymountainshopmaldenDispo, gardenremediesmelroserecDispo, gardenremediesmelroseDispo, oldeworldremediesDispo, naturesmedicinesfallriverDispo, goodchemistrylynnDispo, diemlynnDispo, calyxpeakcompaniesdbalocalcannabiscoDispo, newleafenterprisesDispo, insasalemDispo, insaavon1Dispo, rhelmcannabisDispo, theorywellnessdeliveryzone1Dispo, theorywellnessbridgewaterDispo, alternativecompassionservicesDispo, thehealthcircleDispo, greenrockcannabisDispo, releafalternativesDispo, ermontincDispo, greathitscannabisretailrectautonDispo, curaleafmahanoverDispo, panaceawellnessrecDispo, panaceawellnessDispo, quincycannabisquincyretailrecDispo, commonwealthalternativecareDispo, commonwealthalternativecare1Dispo, terpsattleboroDispo, goodchemistryDispo, goodchemistrymedDispo, bloom1Dispo, diemworcesterDispo, cookiesworcesterDispo, mayflowerworcesterDispo, budsgoodsandprovisionsworcesterDispo, resinateworcesterDispo, missionworcesterDispo, missionworcesteradultuseDispo, clearskyworcesterDispo, greencarecollectiveretailmedDispo, thevault1Dispo, naturesremedymedDispo, naturesremedy1Dispo, harmonyofmaDispo, campfirecannabis1Dispo, discerndcannabisgraftonretailDispo, curaleafmaoxfordDispo, curaleafmaoxford1Dispo, kosaDispo, greengoldDispo, societycannabiscoDispo, newenglandharvestDispo, thevaultwebsterDispo, terpscharltonDispo, temescalwellnesshudsonDispo, hudsonadultuseretailDispo, capitalcanabiscommunitydispensaryDispo, dmaholdingsllcdbagreatesthitscannabisDispo, gardenremediesmarlboroughDispo, gardenremediesrecmarlboroughDispo, greenmeadowsfarmDispo, cannabrollcdbagreenpathsouthbridgeDispo, GreenGoldMarboroughDispo, localrootssturbridgeDispo, greenngoDispo, blackstonevalleycannabisDispo, highhopesDispo, greeneraDispo, thriveofmadbaboomxshirleyDispo, ethosfitchburgDispo, atlanticmedicinalpartnersincDispo, atlanticmedicinalpartnersDispo, greenmeadowfarmfitchburgDispo, pioneercannabiscompanyDispo, naturesmedicinesuxbridgeDispo, uptopframinghamrecDispo, uptopframinghamDispo, masswellspringmaynardDispo, bountifulfarmsDispo, bleafwellnesscentreDispo, thehealingcenter1Dispo, temescalwellnessframinghamDispo, botrealtyllcdbaomgcannabisDispo, curaleafmawareDispo, localrootsfitchburgDispo, ledlloyd420Dispo, boterafranklin1Dispo, masswellspringactonDispo, ddmcannabisDispo, centralavecompassionatecareDispo, gagecannabiscoDispo, netafranklinmedDispo, netafranklinDispo, communitycarecollective1Dispo, unitedcultivationllcDispo, budbarnwinchendonDispo, siranaturalsneedhamDispo, umaflowers1Dispo, clearskybelchertownDispo, thebostongardenDispo, rediDispo, elev8cannabisatholDispo, toytownprojectllcdbatoytownhealthcareDispo, gardenremediesnewtonDispo, gardenremediesrecnewtonDispo, budsgoodsandprovisionswatertownDispo2, nativesunnorthattlebororetailrecDispo2, auraofrhodeislandDispo, ayrdispensarywatertownmedDispo, ayrdispensarywatertownDispo, communitycarecollectiveDispo, uptopwestroxburyDispo, orangecannabiscompanyDispo, EthosWatertownAdultUseDispo, EthosWatertownMedicalDispo, apothcaarlingtonDispo, releafalternativesDispo2, ocsivymaerealestatedbazaharacannabisDispo, motherearthpawtucketDispo, kgcollectivecambridgeDispo, revolutionaryclinicsfreshpondDispo, naturesremedy2Dispo, highprofileroslindaleDispo, mayflowermedicinalsallstonrecDispo, mayflowerallstonDispo, apothcajamaicaplainDispo, netabrooklineDispo, netabrooklinemedDispo, missionbrooklineDispo, redcardinal2Dispo, libertyspringfieldDispo, canacraftDispo, siranaturalsDispo2, siranaturalssomervilleadultuseDispo, medmenfenwayDispo, westernfront1Dispo, smythcannabiscoDispo, massalternativecareamherstmedDispo, massalternativecareamherstDispo, pleasantreesamherstDispo, libertysomervilleDispo, theheirloomcollectiveDispo, pureoasis1Dispo, advesamaincdbablueriverterpscambridgeDispo, insaspringfieldDispo, ayrdispensarybackbayDispo, reverie73Dispo, massalternativecareDispo, massalternativecaremedDispo, revolutionaryclinicssomervilleDispo, theorywellnesschicopeemedDispo, ethosdorchesterDispo, highprofiledorchesterDispo, lazyriverproductsDispo, mistymountainshopmaldenDispo2, jimbuddysDispo, apexnoireDispo, cannapidispensaryDispo2, ermontincDispo2, dazedcannabisDispo,
}

var legalgreensbrocktonDispo DispoData
var commonwealthalternativecarebrocktonrecDispo DispoData
var commonwealthalternativecarebrocktonDispo DispoData
var atlanticmedicinalpartnersinc1Dispo DispoData
var ingoodhealthmedicalDispo DispoData
var ingoodhealthbrocktonDispo DispoData
var panaceawellnessquincyDispo DispoData
var cannapidispensaryDispo DispoData
var brocktongreenheartDispo DispoData
var boterabrocktonDispo DispoData
var budsgoodsandprovisionsabingtonDispo DispoData
var cannavanaDispo DispoData
var nativesunnorthattlebororetailrecDispo DispoData
var yambaDispo DispoData
var budsgoodsandprovisionswatertownDispo DispoData
var siranaturalsDispo DispoData
var mistymountainshopmaldenDispo DispoData
var gardenremediesmelroserecDispo DispoData
var gardenremediesmelroseDispo DispoData
var oldeworldremediesDispo DispoData
var naturesmedicinesfallriverDispo DispoData
var goodchemistrylynnDispo DispoData
var diemlynnDispo DispoData
var calyxpeakcompaniesdbalocalcannabiscoDispo DispoData
var newleafenterprisesDispo DispoData
var insasalemDispo DispoData
var insaavon1Dispo DispoData
var rhelmcannabisDispo DispoData
var theorywellnessdeliveryzone1Dispo DispoData
var theorywellnessbridgewaterDispo DispoData
var alternativecompassionservicesDispo DispoData
var thehealthcircleDispo DispoData
var greenrockcannabisDispo DispoData
var releafalternativesDispo DispoData
var ermontincDispo DispoData
var greathitscannabisretailrectautonDispo DispoData
var curaleafmahanoverDispo DispoData
var panaceawellnessrecDispo DispoData
var panaceawellnessDispo DispoData
var quincycannabisquincyretailrecDispo DispoData
var commonwealthalternativecareDispo DispoData
var commonwealthalternativecare1Dispo DispoData
var terpsattleboroDispo DispoData
var goodchemistryDispo DispoData
var goodchemistrymedDispo DispoData
var bloom1Dispo DispoData
var diemworcesterDispo DispoData
var cookiesworcesterDispo DispoData
var mayflowerworcesterDispo DispoData
var budsgoodsandprovisionsworcesterDispo DispoData
var resinateworcesterDispo DispoData
var missionworcesterDispo DispoData
var missionworcesteradultuseDispo DispoData
var clearskyworcesterDispo DispoData
var greencarecollectiveretailmedDispo DispoData
var thevault1Dispo DispoData
var naturesremedymedDispo DispoData
var naturesremedy1Dispo DispoData
var harmonyofmaDispo DispoData
var campfirecannabis1Dispo DispoData
var discerndcannabisgraftonretailDispo DispoData
var curaleafmaoxfordDispo DispoData
var curaleafmaoxford1Dispo DispoData
var kosaDispo DispoData
var greengoldDispo DispoData
var societycannabiscoDispo DispoData
var newenglandharvestDispo DispoData
var thevaultwebsterDispo DispoData
var terpscharltonDispo DispoData
var temescalwellnesshudsonDispo DispoData
var hudsonadultuseretailDispo DispoData
var capitalcanabiscommunitydispensaryDispo DispoData
var dmaholdingsllcdbagreatesthitscannabisDispo DispoData
var gardenremediesmarlboroughDispo DispoData
var gardenremediesrecmarlboroughDispo DispoData
var greenmeadowsfarmDispo DispoData
var cannabrollcdbagreenpathsouthbridgeDispo DispoData
var GreenGoldMarboroughDispo DispoData
var localrootssturbridgeDispo DispoData
var greenngoDispo DispoData
var blackstonevalleycannabisDispo DispoData
var highhopesDispo DispoData
var greeneraDispo DispoData
var thriveofmadbaboomxshirleyDispo DispoData
var ethosfitchburgDispo DispoData
var atlanticmedicinalpartnersincDispo DispoData
var atlanticmedicinalpartnersDispo DispoData
var greenmeadowfarmfitchburgDispo DispoData
var pioneercannabiscompanyDispo DispoData
var naturesmedicinesuxbridgeDispo DispoData
var uptopframinghamrecDispo DispoData
var uptopframinghamDispo DispoData
var masswellspringmaynardDispo DispoData
var bountifulfarmsDispo DispoData
var bleafwellnesscentreDispo DispoData
var thehealingcenter1Dispo DispoData
var temescalwellnessframinghamDispo DispoData
var botrealtyllcdbaomgcannabisDispo DispoData
var curaleafmawareDispo DispoData
var localrootsfitchburgDispo DispoData
var ledlloyd420Dispo DispoData
var boterafranklin1Dispo DispoData
var masswellspringactonDispo DispoData
var ddmcannabisDispo DispoData
var centralavecompassionatecareDispo DispoData
var gagecannabiscoDispo DispoData
var netafranklinmedDispo DispoData
var netafranklinDispo DispoData
var communitycarecollective1Dispo DispoData
var unitedcultivationllcDispo DispoData
var budbarnwinchendonDispo DispoData
var siranaturalsneedhamDispo DispoData
var umaflowers1Dispo DispoData
var clearskybelchertownDispo DispoData
var thebostongardenDispo DispoData
var rediDispo DispoData
var elev8cannabisatholDispo DispoData
var toytownprojectllcdbatoytownhealthcareDispo DispoData
var gardenremediesnewtonDispo DispoData
var gardenremediesrecnewtonDispo DispoData
var budsgoodsandprovisionswatertownDispo2 DispoData
var nativesunnorthattlebororetailrecDispo2 DispoData
var auraofrhodeislandDispo DispoData
var ayrdispensarywatertownmedDispo DispoData
var ayrdispensarywatertownDispo DispoData
var communitycarecollectiveDispo DispoData
var uptopwestroxburyDispo DispoData
var orangecannabiscompanyDispo DispoData
var EthosWatertownAdultUseDispo DispoData
var EthosWatertownMedicalDispo DispoData
var apothcaarlingtonDispo DispoData
var releafalternativesDispo2 DispoData
var ocsivymaerealestatedbazaharacannabisDispo DispoData
var motherearthpawtucketDispo DispoData
var kgcollectivecambridgeDispo DispoData
var revolutionaryclinicsfreshpondDispo DispoData
var naturesremedy2Dispo DispoData
var highprofileroslindaleDispo DispoData
var mayflowermedicinalsallstonrecDispo DispoData
var mayflowerallstonDispo DispoData
var apothcajamaicaplainDispo DispoData
var netabrooklineDispo DispoData
var netabrooklinemedDispo DispoData
var missionbrooklineDispo DispoData
var redcardinal2Dispo DispoData
var libertyspringfieldDispo DispoData
var canacraftDispo DispoData
var siranaturalsDispo2 DispoData
var siranaturalssomervilleadultuseDispo DispoData
var medmenfenwayDispo DispoData
var westernfront1Dispo DispoData
var smythcannabiscoDispo DispoData
var massalternativecareamherstmedDispo DispoData
var massalternativecareamherstDispo DispoData
var pleasantreesamherstDispo DispoData
var libertysomervilleDispo DispoData
var theheirloomcollectiveDispo DispoData
var pureoasis1Dispo DispoData
var advesamaincdbablueriverterpscambridgeDispo DispoData
var insaspringfieldDispo DispoData
var ayrdispensarybackbayDispo DispoData
var reverie73Dispo DispoData
var massalternativecareDispo DispoData
var massalternativecaremedDispo DispoData
var revolutionaryclinicssomervilleDispo DispoData
var theorywellnesschicopeemedDispo DispoData
var ethosdorchesterDispo DispoData
var highprofiledorchesterDispo DispoData
var lazyriverproductsDispo DispoData
var mistymountainshopmaldenDispo2 DispoData
var jimbuddysDispo DispoData
var apexnoireDispo DispoData
var cannapidispensaryDispo2 DispoData
var ermontincDispo2 DispoData
var dazedcannabisDispo DispoData
