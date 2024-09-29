package vo

import "github.com/DanielAgostinhoSilva/go-clean-architecture/domain/errors"

var (
	ErrInvalidName = errors.NewValidationError("invalid name")
)

type Name struct {
	value string
}

func (vo Name) Value() string {
	return vo.value
}

func NewNameVo(value string) (*Name, error) {
	if len(value) < 3 {
		return nil, ErrInvalidName
	}

	return &Name{value: value}, nil
}
