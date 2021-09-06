package validator

import (
	"testing"
)

func TestIbanValidation(t *testing.T) {
	tests := []struct {
		name    string
		iban    string
		wantErr bool
		Err     error
	}{
		{
			name:    "IBAN of short length",
			iban:    "AJ08",
			wantErr: true,
			Err:     ErrInvalidIBANMinimumLength,
		},
		{
			name:    "IBAN with invalid country code",
			iban:    "ZZ070331234567890123456",
			wantErr: true,
			Err:     ErrWrongInvalidIBANCountryCode,
		},
		{
			name:    "IBAN with invalid checksum",
			iban:    "AE07BARC20201530093459",
			wantErr: true,
			Err:     ErrInvalidIBANChecksum,
		},
		{
			name:    "IBAN with wrong countrycode",
			iban:    "AJ07BARC20201530093459",
			wantErr: true,
			Err:     ErrWrongInvalidIBANCountryCode,
		},
		{
			name:    "IBAN with invalid iban format for 'AE'",
			iban:    "AE070331234567890ABC456",
			wantErr: true,
			Err:     ErrInvalidIBANCountryFormat,
		},
		{
			name: "Valid IBAN for country code 'AE'",
			iban: "AE070331234567890123456",
		},
		{
			name: "Valid IBAN for country code 'SE'",
			iban: "SE4550000000058398257466",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := IbanValidation(tt.iban); (err != nil) != tt.wantErr {
				t.Errorf("IbanValidation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_verifyIbanChecksum(t *testing.T) {
	tests := []struct {
		name    string
		iban    string
		wantErr bool
		Err     error
	}{
		{
			name:    "IBAN with invalid checksum for UK",
			iban:    "GB94BARC20201530093459",
			wantErr: true,
		},
		{
			name: "Valid IBAN checksum for Switzerland",
			iban: "CH9300762011623852957",
		},
		{
			name: "Valid IBAN checksum for Turkey",
			iban: "TR330006100519786457841326",
		},
		{
			name: "Valid IBAN checksum for Ireland",
			iban: "IE29AIBK93115212345678",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := verifyIbanChecksum(tt.iban); (err != nil) != tt.wantErr {
				t.Errorf("verifyIbanChecksum() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_verifyCountryIbanPattern(t *testing.T) {
	tests := []struct {
		name    string
		iban    string
		wantErr bool
	}{
		{
			name:    "IBAN with invalid iban format for Greeland(GL)",
			iban:    "GL8964711000206",
			wantErr: true,
		},
		{
			name:    "IBAN with wrong country code for UAE(AE)",
			iban:    "AJ070331234567890ABC456",
			wantErr: true,
		},
		{
			name: "Valid IBAN for country code Denmark(DK) '",
			iban: "DK5000400440116243",
		},
		{
			name:    "Invalid IBAN for country code Estonia(EE) '",
			iban:    "EE3822221020145685",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := verifyCountryIbanPattern(tt.iban); (err != nil) != tt.wantErr {
				t.Errorf("verifyCountryIbanPattern() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
