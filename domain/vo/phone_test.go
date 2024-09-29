package vo_test

import (
	"testing"

	"github.com/DanielAgostinhoSilva/go-clean-architecture/domain/vo"
	"github.com/stretchr/testify/suite"
)

type PhoneSuite struct {
	suite.Suite
}

func (suite *PhoneSuite) TestValidPhone() {
	validPhones := []string{
		"(11) 91234-5678",
		"(11)91234-5678",
		"11912345678",
		"(21) 2233-4455",
		"2122334455",
	}

	for _, phone := range validPhones {
		_, err := vo.NewPhoneVo(phone)
		suite.Nil(err)
	}
}

func (suite *PhoneSuite) TestInvalidPhone() {
	invalidPhones := []string{
		"12345",
		"abcd-efgh",
		"(11) 1234-567",
		"212233-445",
	}

	for _, phone := range invalidPhones {
		_, err := vo.NewPhoneVo(phone)
		suite.NotNil(err)
	}
}

func TestPhoneSuite(t *testing.T) {
	suite.Run(t, new(PhoneSuite))
}
