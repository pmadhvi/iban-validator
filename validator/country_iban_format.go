package validator

import (
	"regexp"
)

// countryIbanFormats is map of 2-digit countrycode with their corresponding iban regex pattern to match aginst
// this is used for validating the iban for a country
var countryIbanFormats = map[string]*regexp.Regexp{
	"AE": regexp.MustCompile("^AE[0-9]{2}[0-9]{3}[0-9]{16}$"),                      //UAE
	"AT": regexp.MustCompile("^AT[0-9]{2}[0-9]{5}[0-9]{11}$"),                      //Austria
	"BA": regexp.MustCompile("^BA[0-9]{2}[0-9]{3}[0-9]{3}[0-9]{8}[0-9]{2}$"),       //Bosnia
	"BE": regexp.MustCompile("^BE[0-9]{2}[0-9]{3}[0-9]{7}[0-9]{2}$"),               //Belgium
	"BG": regexp.MustCompile("^BG[0-9]{2}[A-Z]{4}[0-9]{4}[A-Z0-9]{8}[0-9]{2}$"),    //Bulgaria
	"BR": regexp.MustCompile("^BR[0-9]{2}[0-9]{8}[0-9]{5}[0-9]{10}[A-Z][A-Z0-9]$"), //Brazil
	"BY": regexp.MustCompile("^BY[0-9]{2}[A-Z]{4}[A-Z0-9]{20}$"),                   //Belarus
	"CH": regexp.MustCompile("^CH[0-9]{2}[0-9]{5}[A-Z0-9]{12}$"),                   //Switzerland
	"CY": regexp.MustCompile("^CY[0-9]{2}[0-9]{3}[0-9]{5}[A-Z0-9]{16}$"),           //Cyprus
	"CZ": regexp.MustCompile("^CZ[0-9]{2}[0-9]{4}[0-9]{6}[0-9]{10}$"),              //Czech Republic
	"DE": regexp.MustCompile("^DE[0-9]{2}[0-9]{8}[0-9]{10}$"),                      //Germany
	"DK": regexp.MustCompile("^DK[0-9]{2}[0-9]{4}[0-9]{10}$"),                      //Denmark
	"EE": regexp.MustCompile("^EE[0-9]{2}[0-9]{2}[0-9]{2}[0-9]{11}[0-9]$"),         //Estonia
	"ES": regexp.MustCompile("^ES[0-9]{2}[0-9]{4}[0-9]{4}[0-9]{10}[0-9]{2}$"),      //Spain
	"FI": regexp.MustCompile("^FI[0-9]{2}[0-9]{6}[0-9]{7}[0-9]$"),                  //Finland
	"FR": regexp.MustCompile("^FR[0-9]{2}[0-9]{5}[0-9]{5}[A-Z0-9]{11}[0-9]{2}$"),   //France
	"GB": regexp.MustCompile("^GB[0-9]{2}[A-Z]{4}[0-9]{6}[0-9]{8}$"),               //United Kingdom
	"GE": regexp.MustCompile("^GE[0-9]{2}[A-Z]{2}[0-9]{16}$"),                      //Georgia
	"GR": regexp.MustCompile("^GR[0-9]{2}[0-9]{3}[0-9]{4}[A-Z0-9]{16}$"),           //Greece
	"GL": regexp.MustCompile("^GL[0-9]{2}[0-9]{4}[0-9]{9}[0-9]$"),                  //Greenland
	"HR": regexp.MustCompile("^HR[0-9]{2}[0-9]{7}[0-9]{10}$"),                      //Crotia
	"HU": regexp.MustCompile("^HU[0-9]{2}[0-9]{3}[0-9]{4}[0-9]{16}[0-9]$"),         //Hungary
	"IE": regexp.MustCompile("^IE[0-9]{2}[A-Z]{4}[0-9]{6}[0-9]{8}$"),               //Ireland
	"IS": regexp.MustCompile("^IS[0-9]{2}[0-9]{4}[0-9]{2}[0-9]{6}[0-9]{10}$"),      //Iceland
	"IT": regexp.MustCompile("^IT[0-9]{2}[A-Z]{1}[0-9]{5}[0-9]{5}[A-Z0-9]{12}$"),   //Italy
	"JO": regexp.MustCompile("^JO[0-9]{2}[A-Z]{4}[0-9]{4}[A-Z0-9]{18}$"),           //Jordan
	"LU": regexp.MustCompile("^LU[0-9]{2}[0-9]{3}[0-9]{13}$"),                      //Luxembourg
	"LV": regexp.MustCompile("^LV[0-9]{2}[A-Z]{4}[A-Z0-9]{13}$"),                   //Latvia
	"MT": regexp.MustCompile("^MT[0-9]{2}[A-Z]{4}[0-9]{5}[A-Z0-9]{18}$"),           //Malta
	"MU": regexp.MustCompile("^MU[0-9]{2}[A-Z]{4}[A-Z0-9]{22}$"),                   //Mauritius
	"NO": regexp.MustCompile("^NO[0-9]{2}[0-9]{4}[0-9]{6}[0-9]$"),                  //Norway
	"PL": regexp.MustCompile("^PL[0-9]{2}[0-9]{3}[0-9]{4}[0-9]{1}[0-9]{16}$"),      //Poland
	"PK": regexp.MustCompile("^PK[0-9]{2}[A-Z]{4}[A-Z0-9]{16}$"),                   //Pakistan
	"ME": regexp.MustCompile("^ME[0-9]{2}[0-9]{3}[0-9]{13}[0-9]{2}$"),              //Montenegro
	"NL": regexp.MustCompile("^NL[0-9]{2}[A-Z]{4}[0-9]{10}$"),                      //Netherland
	"PT": regexp.MustCompile("^PT[0-9]{2}[0-9]{4}[0-9]{4}[0-9]{11}[0-9]{2}$"),      //Portugal
	"QA": regexp.MustCompile("^QA[0-9]{2}[A-Z]{4}[A-Z0-9]{21}$"),                   //Qatar
	"RO": regexp.MustCompile("^RO[0-9]{2}[A-Z]{4}[A-Z0-9]{16}$"),                   //Romania
	"SM": regexp.MustCompile("^SM[0-9]{2}[A-Z]{1}[0-9]{5}[0-9]{5}[A-Z0-9]{12}$"),   //San Marino
	"SA": regexp.MustCompile("^SA[0-9]{2}[0-9]{2}[0-9]{18}$"),                      //Saudi Arabia
	"SK": regexp.MustCompile("^SK[0-9]{2}[0-9]{4}[0-9]{6}[0-9]{10}$"),              //Slovakia
	"SE": regexp.MustCompile("^SE[0-9]{2}[0-9]{3}[0-9]{16}[0-9]$"),                 //Sweden
	"TR": regexp.MustCompile("^TR[0-9]{2}[0-9]{5}[0-9]{1}[A-Z0-9]{16}$"),           //Turkey
	"UA": regexp.MustCompile("^UA[0-9]{2}[0-9]{6}[A-Z0-9]{19}$"),                   //Ukraine
}
