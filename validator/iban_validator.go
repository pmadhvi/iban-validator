// https://www.morfoedro.it/doc.php?n=219&lang=en#:~:text=The%20IBAN%20must%20have%20a,digits%20from%200%20to%209.
package validator

import (
	"math/big"
	"strings"
)

func IbanValidation(iban string) error {
	//check if iban is empty string
	if iban == "" {
		return ErrInvalidIBANEmpty
	}
	iban = strings.ToUpper(iban)

	// check minimum length(5) of iban
	if len(iban) < 5 {
		return ErrInvalidIBANMinimumLength
	}

	// check iban pattern for the specific country
	if err := verifyCountryIbanPattern(iban); err != nil {
		return err
	}

	// verify the checksum(should be mod97)
	if err := verifyIbanChecksum(iban); err != nil {
		return err
	}
	return nil
}

func verifyCountryIbanPattern(iban string) error {
	countryCode := iban[:2]
	// verify the valid country code (alpha2 country code)
	// and also if the country and its iban pattern is present in countryIbanFormats
	countryIbanFormat, ok := countryIbanFormats[countryCode]
	if !ok {
		return ErrWrongInvalidIBANCountryCode
	}
	// check if iban matches the pattern defined for the specific country
	if !countryIbanFormat.MatchString(iban) {
		return ErrInvalidIBANCountryFormat
	}
	return nil
}

func verifyIbanChecksum(iban string) error {
	//swap the fist 4 chars with remaining iban chars
	swappedIban := iban[4:] + iban[:4]

	var convertedIbanNumber string
	//convert each iban chars into numbers
	for _, c := range swappedIban {
		number := convertedNumber[c]
		convertedIbanNumber += number
	}

	// calculate the mod97
	ibanNum, _ := new(big.Int).SetString(convertedIbanNumber, 10)
	mod97 := new(big.Int).SetInt64(97)
	rem := new(big.Int).Mod(ibanNum, mod97)
	if rem.Int64() != 1 {
		return ErrInvalidIBANChecksum
	}
	return nil
}
