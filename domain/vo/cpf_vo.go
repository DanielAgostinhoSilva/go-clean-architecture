package vo

import (
	"github.com/DanielAgostinhoSilva/go-clean-architecture/domain/errors"
	"regexp"
	"strings"
)

var (
	ErrInvalidCpf = errors.NewValidationError("invalid CPF")
)

// Cpf struct
type Cpf struct {
	value string
}

func (vo Cpf) Value() string {
	return vo.value
}

// NewCpfVo constructs a new Cpf value object
func NewCpfVo(value string) (*Cpf, error) {
	if !isValidCpfFormat(value) || !isValidCpfDigits(value) {
		return nil, ErrInvalidCpf
	}

	return &Cpf{value: value}, nil
}

// isValidCpfFormat validates the format of CPF with and without punctuation.
func isValidCpfFormat(cpf string) bool {
	re := regexp.MustCompile(`^\d{3}\.\d{3}\.\d{3}\-\d{2}$`)
	reNoFormat := regexp.MustCompile(`^\d{11}$`)
	return re.MatchString(cpf) || reNoFormat.MatchString(cpf)
}

// isValidCpfDigits validates the CPF digits based on Brazilian CPF rules.
func isValidCpfDigits(cpf string) bool {
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")

	if len(cpf) != 11 {
		return false
	}

	invalids := []string{
		"00000000000",
		"11111111111",
		"22222222222",
		"33333333333",
		"44444444444",
		"55555555555",
		"66666666666",
		"77777777777",
		"88888888888",
		"99999999999",
	}
	for _, invalid := range invalids {
		if cpf == invalid {
			return false
		}
	}

	sum := 0
	for i := 0; i < 9; i++ {
		digit := int(cpf[i] - '0')
		sum += digit * (10 - i)
	}

	firstCheckDigit := (sum * 10 % 11) % 10

	if firstCheckDigit != int(cpf[9]-'0') {
		return false
	}

	sum = 0
	for i := 0; i < 10; i++ {
		digit := int(cpf[i] - '0')
		sum += digit * (11 - i)
	}

	secondCheckDigit := (sum * 10 % 11) % 10

	return secondCheckDigit == int(cpf[10]-'0')
}
