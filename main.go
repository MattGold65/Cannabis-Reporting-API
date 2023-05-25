package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"time"
)

var client *http.Client

var entries = []string{
	"legal-greens-brockton", "commonwealth-alternative-care-brockton-rec", "commonwealth-alternative-care-brockton", "atlantic-medicinal-partners-inc1", "in-good-health-medical", "in-good-health-brockton", "panacea-wellness-quincy", "cannapi-dispensary", "brockton-green-heart", "botera-brockton", "buds-goods-and-provisions-abington", "cannavana", "native-sun-north-attleboro-retail-rec", "yamba", "buds-goods-and-provisions-watertown", "sira-naturals", "misty-mountain-shop-malden", "garden-remedies-melrose-rec", "garden-remedies-melrose", "olde-world-remedies", "natures-medicines-fall-river", "good-chemistry-lynn", "diem-lynn", "calyxpeak-companies-dba-local-cannabis-co", "new-leaf-enterprises", "insa-salem", "insa-avon1", "rhelm-cannabis", "theory-wellness-delivery-zone-1", "theory-wellness-bridgewater", "alternative-compassion-services", "the-health-circle", "green-rock-cannabis", "releaf-alternatives", "ermont-inc", "great-hits-cannabis-retail-rec-tauton", "curaleaf-ma-hanover", "panacea-wellness-rec", "panacea-wellness", "quincy-cannabis-quincy-retail-rec", "commonwealth-alternative-care", "commonwealth-alternative-care1", "terps-attleboro", "good-chemistry", "good-chemistry-med", "bloom1", "diem-worcester", "cookies-worcester", "mayflower-worcester", "buds-goods-and-provisions-worcester", "resinate-worcester", "mission-worcester", "mission-worcester-adult-use", "clear-sky-worcester", "green-care-collective-retail-med", "the-vault1", "natures-remedy-med", "natures-remedy1", "harmony-of-ma", "campfire-cannabis1", "discernd-cannabis-grafton-retail", "curaleaf-ma-oxford", "curaleaf-ma-oxford1", "kosa", "green-gold", "society-cannabis-co", "new-england-harvest", "the-vault-webster", "terps-charlton", "temescal-wellness-hudson", "hudson-adult-use-retail", "capital-canabis-community-dispensary", "dma-holdings-llc-dba-greatest-hits-cannabis", "garden-remedies-marlborough", "garden-remedies-rec-marlborough", "green-meadows-farm", "cannabro-llc-dba-green-path-southbridge", "GreenGold-Marborough", "local-roots-sturbridge", "green-n-go", "black-stone-valley-cannabis", "high-hopes", "green-era", "thrive-of-ma-dba-boom-x-shirley", "ethos-fitchburg", "atlantic-medicinal-partners-inc", "atlantic-medicinal-partners", "green-meadow-farm-fitchburg", "pioneer-cannabis-company", "natures-medicines-uxbridge", "uptop-framingham-rec", "uptop-framingham", "mass-wellspring-maynard", "bountiful-farms", "b-leaf-wellness-centre", "the-healing-center1", "temescal-wellness-framingham", "bot-realty-llc-dba-omg-cannabis", "curaleaf-ma-ware", "local-roots-fitchburg", "ledlloyd420", "botera-franklin1", "mass-wellspring-acton", "ddm-cannabis", "central-ave-compassionate-care", "gage-cannabis-co", "neta-franklin-med", "neta-franklin", "community-care-collective1", "united-cultivation-llc", "bud-barn-winchendon", "sira-naturals-needham", "uma-flowers1", "clear-sky-belchertown", "the-boston-garden", "redi", "elev8-cannabis-athol", "toy-town-project-llc-dba-toy-town-health-care", "garden-remedies-newton", "garden-remedies-rec-newton", "buds-goods-and-provisions-watertown", "native-sun-north-attleboro-retail-rec", "aura-of-rhode-island", "ayr-dispensary-watertown-med", "ayr-dispensary-watertown", "community-care-collective", "uptop-west-roxbury", "orange-cannabis-company", "Ethos-Watertown-Adult-Use", "Ethos-Watertown-Medical", "apothca-arlington", "releaf-alternatives", "ocs-ivy-mae-real-estate-dba-zahara-cannabis", "mother-earth-pawtucket", "kg-collective-cambridge", "revolutionary-clinics-fresh-pond", "natures-remedy2", "high-profile-roslindale", "mayflower-medicinals-allston-rec", "mayflower-allston", "apothca-jamaica-plain", "neta-brookline", "neta-brookline-med", "mission-brookline", "red-cardinal2", "liberty-springfield", "cana-craft", "sira-naturals", "sira-naturals-somerville-adult-use", "medmen-fenway", "western-front1", "smyth-cannabis-co", "mass-alternative-care-amherst-med", "mass-alternative-care-amherst", "pleasantrees-amherst", "liberty-somerville", "the-heirloom-collective", "pure-oasis1", "advesa-ma-inc-dba-blue-river-terps-cambridge", "insa-springfield", "ayr-dispensary-back-bay", "reverie-73", "mass-alternative-care", "mass-alternative-care-med", "revolutionary-clinics-somerville", "theory-wellness-chicopee-med", "ethos-dorchester", "high-profile-dorchester", "lazy-river-products", "misty-mountain-shop-malden", "jimbuddys", "apex-noire", "cannapi-dispensary", "ermont-inc", "dazed-cannabis",
}

var dispensaryIDs = []string{
	"605d194989577500ba7990a4", "6196936e56927100a51e6ede", "6129656b926cf500b7692e3e", "61a93ae4b7cb5900a1cb4281", "60ae731c6fc0b200b5fc3199", "603fb6799c1fa700d1f00332", "63fe436c9cab490070c0f32d", "60872cce441a7700bbd93cab", "6220e024c2de06165d4f04bc", "606357337bf59f00d7d0ebdb", "605a3668ec6fe800b7dfc26d", "5e6a6e46ae8f1700737ef4a1", "636516602f74a400cf2eca24", "604a6f7da6080c5d5563512c", "605a36a99d444100aadb7258", "5e7102bd1e3f82007828115b", "639c73b365caa500cc033ad5", "5f29fbd9fdcba06d60a7c5a5", "5e84b94edc7d0600a189a630", "60c92ada2a83f900db2e53f7", "qCz2peDnERM7553ig", "61a83037d089a300af603d81", "609a91adc37dbe00d1836766", "63597d106e6b8100ef2c6765", "62262e8cb1d475009f22e6ec", "5e6a68a6c18a320081c19a04", "6358653d24cb9300e5170460", "61df00d78116a200a23fc8e4", "5ff675751d596900de321597", "nhc3TYevNCuWB5J5R", "5f9855060284f200bffa1366", "5ed532ad36873f014ab73c4f", "6327ecd675164e01073019e8", "61310e409383cd00ced35523", "5ee940133d5ebe0118fc5f6d", "63c165817c726d00b5dd9485", "6074bdee5bc6a000c2801f6b", "5f492d39d48d6e00b6f7fe32", "kJgMn5xRoSjQaX7N2", "62be222ecb29952e8ba1079c", "5fa58bdd885d7100e1f7ec20", "61b7dc66925234275c7ca250", "61fc7321a6ed9c32ae45e6e3", "SCMabGYMyTMFByhdA", "Mymud8mGNaay8ad6d", "5ec95249c26f7b010a1640d1", "5ec2e958d4ae44013c0988f7", "610c74538f866600a91fc6e4", "60a2aace3b946a00a0e4a04b", "605a36211965683355ba8fe5", "5ef25a345f1aa00153304038", "h9fToyXu9Yu4wjoaK", "5f4932cd99ce2a00ac3b1188", "611c5d0373fee600a2b4f370", "62ab9bf862097e3d8be939bd", "618943ccc4d1f40080cb7438", "5fc7cb93869a9979d8a1bcd4", "djCG6qPhqD3zCJDAo", "6079d5254fbfa900a178b376", "5f47e4d4a05d180107986512", "62bb21e6d01e380087479d37", "6074be1b6e932000cbd31f39", "61005ae72bf50200da88996c", "623a0700e7d0f5009e5dbdbd", "5e3860ae5196d000725c9505", "60f06debec6f72465aa377d1", "60b95a396ddea700a66fa93f", "5f481cb77cbf950101492f06", "6275060766226b5e2d5a95bb", "61c1edd24d760f00a201f1b8", "636516c60abf3700c85c4352", "6137cf6b73cf4b00cc090134", "6193dc69eb404a0085b7a36e", "5e84b9215dbbc400a9dc2ccb", "5ed90fe02128c000af476f49", "5fa438215264a30102844ad9", "6333a7b6ab2f5300baca1f32", "60d3997057277400c3861425", "60d399423d36bb00d86596c0", "6151db521a147500d9302cda", "6091a06067436c00c0c045ec", "60afce2c86ea5f00bb833a9c", "616db3ac0752e400c6835ef4", "628d2602e9a8cb00a1aa5610", "5f7512baedc62900f134f89b", "5ee12bbf3d5a58010523af1e", "60a6e0a482f81600cf6173e2", "631210c0442b2800c5fb503d", "60ff09d5d4805300daf77d9f", "DpDWtpBYPwTgoeAdB", "63e41bdbdb7d460085634140", "61cdd394c93e36008b99271d", "5f43f9e68e1e290102170f13", "5f495c86af752600f34cc41c", "61b8e5fe6c3d5500829f155d", "60afc34350c2ee00b4834af2", "61c1ee2a5231d800946a0ac9", "6298e911c71b9d00804d7b42", "6074be6ff1f75200d1d93e3c", "60d39915da210b00b4671cb2", "2kNKKbMPTMmjRrR2R", "6112e20c1645f700a29cfb3a", "5f444577646bda00c99b0ed7", "5fb2a31a53be060104faad54", "kXTKyCwwLCDmky6Fv", "5e5d2ffcf079ab006c066ba2", "62585605ac5744009c951359", "6256dd0c1f44fa00a4faf8ec", "615c82b4070b9800d2ce92d5", "61b229390d9929008a16f286", "61c212981e89214c4c58d5d2", "5e71246c9c49db006ac54bdb", "61786e7816b797007ae4349a", "62461dd181d049009d0f9248", "6033f43986b8664204ffca4e", "6010b8bd37b04900bd51dfcc", "6268731c52e81900b21a7546", "620e8fed165a2b009a6fb199", "5e84b98b32957600bbff8993", "5ed91052f0a83e00f97c808e", "605a36a99d444100aadb7258", "636516602f74a400cf2eca24", "62be2003916c750095e616d5", "628fd97bf8b71200b05109b9", "622f77708bdb3f009f3a0946", "611a9e6eadb4c100be3a2f05", "61ce325d2f99fc2ea466f426", "5ecfd56d2abef500daca91c2", "5fc9737249b48500f00081ff", "XtHagRnjRdT8BJrXw", "5f15cd1149f2b100e76a3c73", "61310e409383cd00ced35523", "61d86faab6601000a24c0d55", "6304f1dad20be000cf6d95bf", "62fc42aaa2826300c0f47e0f", "4Qqqwdo92xn5qoupj", "5ee26c349578ec00ff76dbb7", "61aa5d230c15bb0083774b58", "632cc1684eb7b700d47b969f", "60a2ab1e05ff7100d2c116e8", "615c79b260bf4100cb06c0d1", "6256dda39f4805009d5bbd30", "625855d980dcc900a3e3890c", "60a443eb5be74d00a8bd97c2", "608c16b8fdcb5400bac018d2", "5f98867881b29700c8e58d76", "6244bae360a8c5009db0d4bb", "5e7102bd1e3f82007828115b", "631911acaeaa7b00e6a024d5", "613b98cde6576f00ce20f1bb", "61fc2234e325be4351c24dd7", "60f881018b6bc800c70b8393", "5f511afe0fd39d00d2a25d49", "5f3c562257321d01080d02d6", "615df5eac4d44700b8254cd8", "y3wvjnoqRGzpM2DjS", "5ea20c7b1e09a7071bfc266e", "4xjpSMy8hRkEyvAFG", "631793ccf1d2fa00e6aa161a", "5e6a6905d035f6006cc5359a", "622f7732c7431600a56a1015", "6192b70e0a6d9300803a1c3e", "3jiajRb46dunW7apx", "5eee33d3a899f53d0b36aec0", "yswxv7Lt423nPZm6W", "5f932aa55f94c10102a710c8", "6008d7aab75c4900ab2a9f1c", "61aa5d8db160eb009934c31d", "5f5fc6c5c1d5c36b8f5d0574", "639c73b365caa500cc033ad5", "616596b70848f500a7ee5b20", "637e4cda7d6b3100b87e95d7", "60872cce441a7700bbd93cab", "5ee940133d5ebe0118fc5f6d", "613fb7462dc78900a9bd305d",
}

func main() {
	client = &http.Client{Timeout: 10 * time.Second}

	db, err := sql.Open("mysql", "root:password1@tcp(127.0.0.1:3306)/CannabisDB")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	/*
		_, err = db.Exec("CREATE TABLE IF NOT EXISTS DispensaryCommon (DispensaryID varCHAR(255), DispensaryName varCHAR(255), Website varCHAR(255), Email varCHAR(255), Phone varCHAR(255), Address varCHAR(255), State varCHAR(255), City varCHAR(255), Street varCHAR(255), Street2 varCHAR(255), Zipcode varCHAR(255), Latitude FLOAT, Longitdue FLOAT, type varCHAR(255))")
		if err != nil {
			panic(err.Error())
		}

		getDispensaryData(db, entries, dispoDataSlice)
		addDispoPrimaryKey(db, "PrimaryKey")

	*/

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Product_FlowerCommon (EffDate DATE, DispensaryID varCHAR(255), EnterpriseProductID varCHAR(255), ProductID varCHAR(255) NOT NULL, BrandName varCHAR(255) NOT NULL, ProductName varCHAR(255) NOT NULL, Weight varCHAR(255), WeightGrams FLOAT, RecOnly BOOLEAN, RecPrice FLOAT, RecSpecialPrice Float, MedicalOnly BOOLEAN, MedicalPrice FLOAT, MedicalSpecialPrice Float, WholeSalePrice Float, InventoryQuantity int, InventoryQuantityAvail int, InventoryQuantityKiosk int, StrainType varCHAR(255),  THCAmount FLOAT, THCUnit varCHAR(255), THCTypeName varCHAR(255), CBDAMOUNT FLOAT, CBDUnit varCHAR(255), CBDTypeName varCHAR(255), RexaledRating FLOAT, PainReliefRating FLOAT, SleepyRating FLOAT, HappyRating FLOAT, EuphoricRating FLOAT, ProductImage varCHAR(255))")
	if err != nil {
		panic(err.Error())
	}

	getFlowerData(db, dispensaryIDs, flowerDataSlice)
	addFlowerPrimaryKey(db, "PrimaryKey")

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Product_EdiblesCommon (EffDate DATE, DispensaryID varCHAR(255), EnterpriseProductID varCHAR(255), ProductID varCHAR(255) NOT NULL, BrandName varCHAR(255) NOT NULL, ProductName varCHAR(255) NOT NULL, ProductSubCategory  varCHAR(255) NOT NULL, Weight varCHAR(255), WeightGrams FLOAT, RecOnly BOOLEAN, RecPrice FLOAT, RecSpecialPrice Float, MedicalOnly BOOLEAN, MedicalPrice FLOAT, MedicalSpecialPrice Float, WholeSalePrice Float, InventoryQuantity int, InventoryQuantityAvail int, InventoryQuantityKiosk int, StrainType varCHAR(255),  THCAmount FLOAT, THCUnit varCHAR(255), THCTypeName varCHAR(255), CBDAMOUNT FLOAT, CBDUnit varCHAR(255), CBDTypeName varCHAR(255), RexaledRating FLOAT, PainReliefRating FLOAT, SleepyRating FLOAT, HappyRating FLOAT, EuphoricRating FLOAT, ProductImage varCHAR(255))")
	if err != nil {
		panic(err.Error())
	}

	getEdiblesData(db, dispensaryIDs, EdiblesDataSlice)
	addEdiblesPrimaryKey(db, "PrimaryKey")

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Product_ConcentratesCommon (EffDate DATE, DispensaryID varCHAR(255), EnterpriseProductID varCHAR(255), ProductID varCHAR(255) NOT NULL, BrandName varCHAR(255) NOT NULL, ProductName varCHAR(255) NOT NULL, ProductSubCategory  varCHAR(255) NOT NULL, Weight varCHAR(255), WeightGrams FLOAT, RecOnly BOOLEAN, RecPrice FLOAT, RecSpecialPrice Float, MedicalOnly BOOLEAN, MedicalPrice FLOAT, MedicalSpecialPrice Float, WholeSalePrice Float, InventoryQuantity int, InventoryQuantityAvail int, InventoryQuantityKiosk int, StrainType varCHAR(255),  THCAmount FLOAT, THCUnit varCHAR(255), THCTypeName varCHAR(255), CBDAMOUNT FLOAT, CBDUnit varCHAR(255), CBDTypeName varCHAR(255), RexaledRating FLOAT, PainReliefRating FLOAT, SleepyRating FLOAT, HappyRating FLOAT, EuphoricRating FLOAT, ProductImage varCHAR(255))")
	if err != nil {
		panic(err.Error())
	}

	getConcentratesData(db, dispensaryIDs, ConcentratesDataSlice)
	addConcentratesPrimaryKey(db, "PrimaryKey")

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Product_prerollsCommon (EffDate DATE, DispensaryID varCHAR(255), EnterpriseProductID varCHAR(255), ProductID varCHAR(255) NOT NULL, BrandName varCHAR(255) NOT NULL, ProductName varCHAR(255) NOT NULL, ProductSubCategory  varCHAR(255) NOT NULL, Weight varCHAR(255), WeightGrams FLOAT, RecOnly BOOLEAN, RecPrice FLOAT, RecSpecialPrice Float, MedicalOnly BOOLEAN, MedicalPrice FLOAT, MedicalSpecialPrice Float, WholeSalePrice Float, InventoryQuantity int, InventoryQuantityAvail int, InventoryQuantityKiosk int, StrainType varCHAR(255),  THCAmount FLOAT, THCUnit varCHAR(255), THCTypeName varCHAR(255), CBDAMOUNT FLOAT, CBDUnit varCHAR(255), CBDTypeName varCHAR(255), RexaledRating FLOAT, PainReliefRating FLOAT, SleepyRating FLOAT, HappyRating FLOAT, EuphoricRating FLOAT, ProductImage varCHAR(255))")
	if err != nil {
		panic(err.Error())
	}

	getprerollsData(db, dispensaryIDs, prerollsDataSlice)
	addprerollsPrimaryKey(db, "PrimaryKey")

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Product_vapeorizersCommon (EffDate DATE, DispensaryID varCHAR(255), EnterpriseProductID varCHAR(255), ProductID varCHAR(255) NOT NULL, BrandName varCHAR(255) NOT NULL, ProductName varCHAR(255) NOT NULL, ProductSubCategory  varCHAR(255) NOT NULL, Weight varCHAR(255), WeightGrams FLOAT, RecOnly BOOLEAN, RecPrice FLOAT, RecSpecialPrice Float, MedicalOnly BOOLEAN, MedicalPrice FLOAT, MedicalSpecialPrice Float, WholeSalePrice Float, InventoryQuantity int, InventoryQuantityAvail int, InventoryQuantityKiosk int, StrainType varCHAR(255),  THCAmount FLOAT, THCUnit varCHAR(255), THCTypeName varCHAR(255), CBDAMOUNT FLOAT, CBDUnit varCHAR(255), CBDTypeName varCHAR(255), RexaledRating FLOAT, PainReliefRating FLOAT, SleepyRating FLOAT, HappyRating FLOAT, EuphoricRating FLOAT, ProductImage varCHAR(255))")
	if err != nil {
		panic(err.Error())
	}

	getvapeorizersData(db, dispensaryIDs, vapeorizersDataSlice)
	addvapeorizersPrimaryKey(db, "PrimaryKey")

}
