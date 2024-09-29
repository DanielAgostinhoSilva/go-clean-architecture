package vo

import (
	"github.com/DanielAgostinhoSilva/go-clean-architecture/domain/errors"
	"regexp"
)

var (
	ErrInvalidPhone = errors.NewValidationError("invalid phone number")
)

// Phone struct
type Phone struct {
	value string
}

func (vo Phone) Value() string {
	return vo.value
}

// NewPhoneVo constructs a new Phone value object
func NewPhoneVo(value string) (*Phone, error) {
	if !isValidPhoneFormat(value) {
		return nil, ErrInvalidPhone
	}

	return &Phone{value: value}, nil
}

// isValidPhoneFormat validates the format of a phone number.
func isValidPhoneFormat(phone string) bool {
	// Regex for validating Brazilian phone numbers (landline and mobile)
	re := regexp.MustCompile(`^\(?[1-9]{2}\)? ?9?[0-9]{4}-?[0-9]{4}$`)
	return re.MatchString(phone)
}
