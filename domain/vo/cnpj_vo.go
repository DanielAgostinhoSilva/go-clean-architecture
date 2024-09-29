package vo

import (
	"github.com/DanielAgostinhoSilva/go-clean-architecture/domain/errors"
	"regexp"
	"strings"
)

var (
	ErrInvalidCnpj = errors.NewValidationError("invalid CNPJ")
)

// Cnpj struct
type Cnpj struct {
	value string
}

func (vo Cnpj) Value() string {
	return vo.value
}

// NewCnpjVo constructs a new Cnpj value object
func NewCnpjVo(value string) (*Cnpj, error) {
	if !isValidCnpjFormat(value) || !isValidCnpjDigits(value) {
		return nil, ErrInvalidCnpj
	}

	return &Cnpj{value: value}, nil
}

// isValidCnpjFormat validates the format of CNPJ with and without punctuation.
func isValidCnpjFormat(cnpj string) bool {
	re := regexp.MustCompile(`^\d{2}\.\d{3}\.\d{3}/\d{4}\-\d{2}$`)
	reNoFormat := regexp.MustCompile(`^\d{14}$`)
	return re.MatchString(cnpj) || reNoFormat.MatchString(cnpj)
}

// isValidCnpjDigits validates the CNPJ digits based on Brazilian CNPJ rules.
func isValidCnpjDigits(cnpj string) bool {
	cnpj = strings.ReplaceAll(cnpj, ".", "")
	cnpj = strings.ReplaceAll(cnpj, "-", "")
	cnpj = strings.ReplaceAll(cnpj, "/", "")

	if len(cnpj) != 14 {
		return false
	}

	invalids := []string{
		"00000000000000",
		"11111111111111",
		"22222222222222",
		"33333333333333",
		"44444444444444",
		"55555555555555",
		"66666666666666",
		"77777777777777",
		"88888888888888",
		"99999999999999",
	}
	for _, invalid := range invalids {
		if cnpj == invalid {
			return false
		}
	}

	multipliers1 := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	multipliers2 := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	sum := 0
	for i := 0; i < 12; i++ {
		digit := int(cnpj[i] - '0')
		sum += digit * multipliers1[i]
	}

	firstCheckDigit := (sum % 11)
	if firstCheckDigit < 2 {
		firstCheckDigit = 0
	} else {
		firstCheckDigit = 11 - firstCheckDigit
	}

	if firstCheckDigit != int(cnpj[12]-'0') {
		return false
	}

	sum = 0
	for i := 0; i < 13; i++ {
		digit := int(cnpj[i] - '0')
		sum += digit * multipliers2[i]
	}

	secondCheckDigit := (sum % 11)
	if secondCheckDigit < 2 {
		secondCheckDigit = 0
	} else {
		secondCheckDigit = 11 - secondCheckDigit
	}

	return secondCheckDigit == int(cnpj[13]-'0')
}
