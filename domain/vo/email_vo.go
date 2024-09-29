package vo

import (
	"github.com/DanielAgostinhoSilva/go-clean-architecture/domain/errors"
	"regexp"
)

var (
	ErrInvalidEmail = errors.NewValidationError("invalid email address")
)

// Email struct
type Email struct {
	value string
}

func (vo Email) Value() string {
	return vo.value
}

// NewEmailVo constructs a new Email value object
func NewEmailVo(value string) (*Email, error) {
	if !isValidEmailFormat(value) {
		return nil, ErrInvalidEmail
	}

	return &Email{value: value}, nil
}

// isValidEmailFormat validates the format of an email address.
func isValidEmailFormat(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9]+(\.[a-z]{2,4})+$`)
	//re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	//re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9\-]+\.[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
