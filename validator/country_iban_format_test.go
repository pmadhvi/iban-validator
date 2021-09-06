package validator

import "testing"

var country_map = map[string]string{
	"Sweden":         "SE4550000000058398257466",
	"Denmark":        "DK5000400440116243",
	"Belgium":        "BE68539007547034",
	"Brazil":         "BR9700360305000010009795493P1",
	"Georgia":        "GE29NB0000000101904917",
	"Czech Republic": "CZ6508000000192000145399",
	"Hungary":        "HU42117730161111101800000000",
	"Estonia":        "EE382200221020145685",
	"Finland":        "FI2112345600000785",
	"France":         "FR1420041010050500013M02606",
	"Germany":        "DE89370400440532013000",
	"Greece":         "GR1601101250000000012300695",
	"Greenland":      "GL8964710001000206",
	"Iceland":        "IS140159260076545510730339",
	"Italy":          "IT60X0542811101000000123456",
	"Austria":        "AT611904300234573201",
	"Latvia":         "LV80BANK0000435195001",
	"Bulgaria":       "BG80BNBG96611020345678",
	"Luxembourg":     "LU280019400644750000",
	"Norway":         "NO9386011117947",
	"Croatia":        "HR1210010051863000160",
	"Cyprus":         "CY17002001280000001200527600",
	"Poland":         "PL61109010140000071219812874",
	"Slovakia":       "SK3112000000198742637541",
	"Portugal":       "PT50000201231234567890154",
	"Qatar":          "QA58DOHB00001234567890ABCDEFG",
	"Ireland":        "IE29AIBK93115212345678",
	"Montenegro":     "ME25505000012345678951",
	"Netherland":     "NL91ABNA0417164300",
	"San Marino":     "SM86U0322509800000000270100",
	"Turkey":         "TR330006100519786457841326",
	"Saudi Arabia":   "SA0380000000608010167519",
	"Spain":          "ES9121000418450200051332",
	"Romania":        "RO49AAAA1B31007593840000",
	"Switzerland":    "CH9300762011623852957",
	"Ukraine":        "UA213996220000026007233566001",
	"Malta":          "MT84MALT011000012345MTLCAST001S",
	"United Kingdom": "GB29NWBK60161331926819",
	"Jordan":         "JO94CBJO0010000000000131000302",
	"Mauritius":      "MU17BOMM0101101030300200000MUR",
	"Belarus":        "BY13NBRB3600900000002Z00AB00",
	"Pakistan":       "PK36SCBL0000001123456702",
	"Bosnia":         "BA391290079401028494",
}

func Test_verifyCountryIbanPatternsForAllCountry(t *testing.T) {
	for country_name, iban := range country_map {
		t.Run(country_name, func(t *testing.T) {
			if err := verifyCountryIbanPattern(iban); err != nil {
				t.Errorf("verifyCountryIbanPattern() error = %v, country = %v", err, country_name)
			}
		})
	}
}
