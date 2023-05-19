package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"time"
)

var client *http.Client

var LegalGreensDispo DispoData
var CommonwealthAltcareRecDispo DispoData
var AtlanticMedicinalPartnersDispo DispoData
var InGoodHealthRecDispo DispoData
var InGoodHealthMedicalDispo DispoData
var PanaceaWellnessQuincyDispo DispoData
var CommonwealthAlternativeCareBrocktonMedDispo DispoData
var CannapiDispensaryDispo DispoData
var BrocktonGreenHeartDispo DispoData
var GTEBrocktonLLCdbaBoteraBrocktonDispo DispoData
var BudsGoodsProvisionsAbingtonDispo DispoData
var CannaVanaDispo DispoData
var NativeSunNorthAttleboroDispo DispoData
var YambaDispo DispoData
var BudsGoodsProvisionsWatertownDispo DispoData
var SiraNaturalsSomervilleMedicalDispo DispoData
var MistyMountainShopMaldenDispo DispoData
var GardenRemediesRecMelroseDispo DispoData
var GardenRemediesMedMelroseDispo DispoData
var OldeWorldRemediesDispo DispoData
var NaturesMedicinesFallRiverDispo DispoData
var GoodChemistryLynnDispo DispoData
var DiemLynnDispo DispoData
var LocalCannabisCoSwampscottDispo DispoData
var NewLeafFallRiverDispo DispoData
var INSAIncSalemAdultUseDispo DispoData
var InsaAvonDispo DispoData
var RhelmCannabisDispo DispoData
var TheoryWellnessBostonAreaDeliveryDispo DispoData
var TheoryWellnessBridgewaterDispo DispoData
var AlternativeCompassionServicesDispo DispoData
var HealthCircleDispo DispoData
var GreenRockCannabisDispo DispoData
var ReleafAlternativesDispo DispoData
var ErmontIncDispo DispoData
var GreatestHitsCannabisTauntonDispo DispoData
var CuraleafMAHanoverDispo DispoData
var PanaceaWellnessRecDispo DispoData
var PanaceaWellnessDispo DispoData
var QuincyCannabisCoDispo DispoData
var CommonwealthAlternativeCareTauntonMedDispo DispoData
var CommonwealthAlternativeCareTauntonRecDispo DispoData
var TERPSAttleboroDispo DispoData

var LegalGreens FlowerData
var CommonwealthAltcareRec FlowerData
var AtlanticMedicinalPartners FlowerData
var InGoodHealthRec FlowerData
var InGoodHealthMedical FlowerData
var PanaceaWellnessQuincy FlowerData
var CommonwealthAlternativeCareBrocktonMed FlowerData
var CannapiDispensary FlowerData
var BrocktonGreenHeart FlowerData
var GTEBrocktonLLCdbaBoteraBrockton FlowerData
var BudsGoodsProvisionsAbington FlowerData
var CannaVana FlowerData
var NativeSunNorthAttleboro FlowerData
var Yamba FlowerData
var BudsGoodsProvisionsWatertown FlowerData
var SiraNaturalsSomervilleMedical FlowerData
var MistyMountainShopMalden FlowerData
var GardenRemediesRecMelrose FlowerData
var GardenRemediesMedMelrose FlowerData
var OldeWorldRemedies FlowerData
var NaturesMedicinesFallRiver FlowerData
var GoodChemistryLynn FlowerData
var DiemLynn FlowerData
var LocalCannabisCoSwampscott FlowerData
var NewLeafFallRiver FlowerData
var INSAIncSalemAdultUse FlowerData
var InsaAvon FlowerData
var RhelmCannabis FlowerData
var TheoryWellnessBostonAreaDelivery FlowerData
var TheoryWellnessBridgewater FlowerData
var AlternativeCompassionServices FlowerData
var HealthCircle FlowerData
var GreenRockCannabis FlowerData
var ReleafAlternatives FlowerData
var ErmontInc FlowerData
var GreatestHitsCannabisTaunton FlowerData
var CuraleafMAHanover FlowerData
var PanaceaWellnessRec FlowerData
var PanaceaWellness FlowerData
var QuincyCannabisCo FlowerData
var CommonwealthAlternativeCareTauntonMed FlowerData
var CommonwealthAlternativeCareTauntonRec FlowerData
var TERPSAttleboro FlowerData

func main() {
	client = &http.Client{Timeout: 10 * time.Second}

	err := getJson(LegalGreensDispoLink, "LegalGreensDispo", &LegalGreensDispo)
	err = getJson(CommonwealthAltcareRecDispoLink, "CommonwealthAltcareRecDispo", &CommonwealthAltcareRecDispo)
	err = getJson(AtlanticMedicinalPartnersBrocktonRecDispoLink, "AtlanticMedicinalPartnersDispo", &AtlanticMedicinalPartnersDispo)
	err = getJson(InGoodHealthRecDispoLink, "InGoodHealthRecDispo", &InGoodHealthRecDispo)
	err = getJson(InGoodHealthMedicalDispoLink, "InGoodHealthMedicalDispo", &InGoodHealthMedicalDispo)
	err = getJson(PanaceaWellnessQuincyDispoLink, "PanaceaWellnessQuincyDispo", &PanaceaWellnessQuincyDispo)
	err = getJson(CommonwealthAlternativeCareBrocktonMedDispoLink, "CommonwealthAlternativeCareBrocktonMedDispo", &CommonwealthAlternativeCareBrocktonMedDispo)
	err = getJson(CannapiDispensaryDispoLink, "CannapiDispensaryDispo", &CannapiDispensaryDispo)
	err = getJson(BrocktonGreenHeartDispoLink, "BrocktonGreenHeartDispo", &BrocktonGreenHeartDispo)
	err = getJson(GTEBrocktonLLCdbaBoteraBrocktonDispoLink, "GTEBrocktonLLCdbaBoteraBrocktonDispo", &GTEBrocktonLLCdbaBoteraBrocktonDispo)
	err = getJson(BudsGoodsProvisionsAbingtonDispoLink, "BudsGoodsProvisionsAbingtonDispo", &BudsGoodsProvisionsAbingtonDispo)
	err = getJson(CannaVanaDispoLink, "CannaVanaDispo", &CannaVanaDispo)
	err = getJson(NativeSunNorthAttleboroDispoLink, "NativeSunNorthAttleboroDispo", &NativeSunNorthAttleboroDispo)
	err = getJson(YambaDispoLink, "YambaDispo", &YambaDispo)
	err = getJson(BudsGoodsProvisionsWatertownDispoLink, "BudsGoodsProvisionsWatertownDispo", &BudsGoodsProvisionsWatertownDispo)
	err = getJson(SiraNaturalsSomervilleMedicalDispoLink, "SiraNaturalsSomervilleMedicalDispo", &SiraNaturalsSomervilleMedicalDispo)
	err = getJson(MistyMountainShopMaldenDispoLink, "MistyMountainShopMaldenDispo", &MistyMountainShopMaldenDispo)
	err = getJson(GardenRemediesRecMelroseDispoLink, "GardenRemediesRecMelroseDispo", &GardenRemediesRecMelroseDispo)
	err = getJson(GardenRemediesMedMelroseDispoLink, "GardenRemediesMedMelroseDispo", &GardenRemediesMedMelroseDispo)
	err = getJson(OldeWorldRemediesDispoLink, "OldeWorldRemediesDispo", &OldeWorldRemediesDispo)
	err = getJson(NaturesMedicinesFallRiverDispoLink, "NaturesMedicinesFallRiverDispo", &NaturesMedicinesFallRiverDispo)
	err = getJson(GoodChemistryLynnDispoLink, "GoodChemistryLynnDispo", &GoodChemistryLynnDispo)
	err = getJson(DiemLynnDispoLink, "DiemLynnDispo", &DiemLynnDispo)
	err = getJson(LocalCannabisCoSwampscottDispoLink, "LocalCannabisCoSwampscottDispo", &LocalCannabisCoSwampscottDispo)
	err = getJson(NewLeafFallRiverDispoLink, "NewLeafFallRiverDispo", &NewLeafFallRiverDispo)
	err = getJson(INSAIncSalemAdultUseDispoLink, "INSAIncSalemAdultUseDispo", &INSAIncSalemAdultUseDispo)
	err = getJson(InsaAvonDispoLink, "InsaAvonDispo", &InsaAvonDispo)
	err = getJson(RhelmCannabisDispoLink, "RhelmCannabisDispo", &RhelmCannabisDispo)
	err = getJson(TheoryWellnessBostonAreaDeliveryDispoLink, "TheoryWellnessBostonAreaDeliveryDispo", &TheoryWellnessBostonAreaDeliveryDispo)
	err = getJson(TheoryWellnessBridgewaterDispoLink, "TheoryWellnessBridgewaterDispo", &TheoryWellnessBridgewaterDispo)
	err = getJson(AlternativeCompassionServicesDispoLink, "AlternativeCompassionServicesDispo", &AlternativeCompassionServicesDispo)
	err = getJson(HealthCircleDispoLink, "HealthCircleDispo", &HealthCircleDispo)
	err = getJson(GreenRockCannabisDispoLink, "GreenRockCannabisDispo", &GreenRockCannabisDispo)
	err = getJson(ReleafAlternativesDispoLink, "ReleafAlternativesDispo", &ReleafAlternativesDispo)
	err = getJson(ErmontIncDispoLink, "ErmontIncDispo", &ErmontIncDispo)
	err = getJson(GreatestHitsCannabisTauntonDispoLink, "GreatestHitsCannabisTauntonDispo", &GreatestHitsCannabisTauntonDispo)
	err = getJson(CuraleafMAHanoverDispoLink, "CuraleafMAHanoverDispo", &CuraleafMAHanoverDispo)
	err = getJson(PanaceaWellnessRecDispoLink, "PanaceaWellnessRecDispo", &PanaceaWellnessRecDispo)
	err = getJson(PanaceaWellnessDispoLink, "PanaceaWellnessDispo", &PanaceaWellnessDispo)
	err = getJson(QuincyCannabisCoDispoLink, "QuincyCannabisCoDispo", &QuincyCannabisCoDispo)
	err = getJson(CommonwealthAlternativeCareTauntonMedDispoLink, "CommonwealthAlternativeCareTauntonMedDispo", &CommonwealthAlternativeCareTauntonMedDispo)
	err = getJson(CommonwealthAlternativeCareTauntonRecDispoLink, "CommonwealthAlternativeCareTauntonRecDispo", &CommonwealthAlternativeCareTauntonRecDispo)
	err = getJson(TERPSAttleboroDispoLink, "TERPSAttleboroDispo", &TERPSAttleboroDispo)

	/*
			err := getJson(LegalGreensLink, "LegalGreens", &LegalGreens)
		   	err := getJson(CommonwealthAltcareRecLink, "CommonwealthAltcareRec", &CommonwealthAltcareRec)
		   	err = getJson(AtlanticMedicinalPartnersBrocktonRecLink, "AtlanticMedicinalPartners", &AtlanticMedicinalPartners)
			err = getJson(InGoodHealthRecLink, "InGoodHealthRec", &InGoodHealthRec)
		   	err = getJson(InGoodHealthMedicalLink, "InGoodHealthMedical", &InGoodHealthMedical)
		   	err = getJson(PanaceaWellnessQuincyLink, "PanaceaWellnessQuincy", &PanaceaWellnessQuincy)
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
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Product_FlowerCommon (EffDate DATE, DispensaryID VARCHAR(255), EnterpriseProductID VARCHAR(255), ProductID VARCHAR(255) NOT NULL, BrandName VARCHAR(255) NOT NULL, ProductName VARCHAR(255) NOT NULL, Weight VARCHAR(255), WeightGrams FLOAT, RecOnly BOOLEAN, RecPrice FLOAT, RecSpecialPrice Float, MedicalOnly BOOLEAN, MedicalPrice FLOAT, MedicalSpecialPrice Float, WholeSalePrice Float, InventoryQuantity int, InventoryQuantityAvail int, InventoryQuantityKiosk int, StrainType VARCHAR(255),  THCAmount FLOAT, THCUnit VARCHAR(255), THCTypeName VARCHAR(255), CBDAMOUNT FLOAT, CBDUnit VARCHAR(255), CBDTypeName VARCHAR(255), RexaledRating FLOAT, PainReliefRating FLOAT, SleepyRating FLOAT, HappyRating FLOAT, EuphoricRating FLOAT, ProductImage VARCHAR(255))")
	if err != nil {
		panic(err.Error())
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS DispensaryCommon (DispensaryID VARCHAR(255), DispensaryName VARCHAR(255), Website VARCHAR(255), Email VARCHAR(255), Phone VARCHAR(255), Address VARCHAR(255), State VARCHAR(255), City VARCHAR(255), Street VARCHAR(255), Street2 VARCHAR(255), Zipcode VARCHAR(255), Latitude FLOAT, Longitdue FLOAT, type VARCHAR(255))")
	if err != nil {
		panic(err.Error())
	}
	addPrimaryKey(db, "PrimaryKey")

	var Dispensaries []DispoData

	Dispensaries = append(Dispensaries, LegalGreensDispo, CommonwealthAltcareRecDispo, AtlanticMedicinalPartnersDispo, InGoodHealthRecDispo, InGoodHealthMedicalDispo, PanaceaWellnessQuincyDispo, CommonwealthAlternativeCareBrocktonMedDispo, CannapiDispensaryDispo, BrocktonGreenHeartDispo, GTEBrocktonLLCdbaBoteraBrocktonDispo, BudsGoodsProvisionsAbingtonDispo, CannaVanaDispo, NativeSunNorthAttleboroDispo, YambaDispo, BudsGoodsProvisionsWatertownDispo, SiraNaturalsSomervilleMedicalDispo, MistyMountainShopMaldenDispo, GardenRemediesRecMelroseDispo, GardenRemediesMedMelroseDispo, OldeWorldRemediesDispo, NaturesMedicinesFallRiverDispo, GoodChemistryLynnDispo, DiemLynnDispo, LocalCannabisCoSwampscottDispo, NewLeafFallRiverDispo, INSAIncSalemAdultUseDispo, InsaAvonDispo, RhelmCannabisDispo, TheoryWellnessBostonAreaDeliveryDispo, TheoryWellnessBridgewaterDispo, AlternativeCompassionServicesDispo, HealthCircleDispo, GreenRockCannabisDispo, ReleafAlternativesDispo, ErmontIncDispo, GreatestHitsCannabisTauntonDispo, CuraleafMAHanoverDispo, PanaceaWellnessRecDispo, PanaceaWellnessDispo, QuincyCannabisCoDispo, CommonwealthAlternativeCareTauntonMedDispo, CommonwealthAlternativeCareTauntonRecDispo, TERPSAttleboroDispo)
	for _, DispoData := range Dispensaries {
		InsertDataDispoCommon(db, DispoData)
	}

	//var products []FlowerData
	//products = append(products, LegalGreens, CommonwealthAltcareRec, AtlanticMedicinalPartners, InGoodHealthRec, InGoodHealthMedical, PanaceaWellnessQuincy, CommonwealthAlternativeCareBrocktonMed, CannapiDispensary, BrocktonGreenHeart, GTEBrocktonLLCdbaBoteraBrockton, BudsGoodsProvisionsAbington, CannaVana, NativeSunNorthAttleboro, Yamba, BudsGoodsProvisionsWatertown, SiraNaturalsSomervilleMedical, MistyMountainShopMalden, GardenRemediesRecMelrose, GardenRemediesMedMelrose, OldeWorldRemedies, NaturesMedicinesFallRiver, GoodChemistryLynn, DiemLynn, LocalCannabisCoSwampscott, NewLeafFallRiver, INSAIncSalemAdultUse, InsaAvon, RhelmCannabis, TheoryWellnessBostonAreaDelivery, TheoryWellnessBridgewater, AlternativeCompassionServices, HealthCircle, GreenRockCannabis, ReleafAlternatives, ErmontInc, GreatestHitsCannabisTaunton, CuraleafMAHanover, PanaceaWellnessRec, PanaceaWellness, QuincyCannabisCo, CommonwealthAlternativeCareTauntonMed, CommonwealthAlternativeCareTauntonRec, TERPSAttleboro)
	/*
		for _, FlowerData := range products {
			InsertDataProductFlowerCommon(db, FlowerData)
		}

	*/

}
