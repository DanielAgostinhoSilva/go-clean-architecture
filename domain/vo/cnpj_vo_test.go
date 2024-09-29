package vo_test

import (
	"testing"

	"github.com/DanielAgostinhoSilva/go-clean-architecture/domain/vo"
	"github.com/stretchr/testify/suite"
)

type CnpjSuite struct {
	suite.Suite
}

func (suite *CnpjSuite) TestValidCnpj() {
	validCnpjs := []string{
		"12.345.678/0001-95",
		"12345678000195",
	}

	for _, cnpj := range validCnpjs {
		_, err := vo.NewCnpjVo(cnpj)
		suite.Nil(err)
	}
}

func (suite *CnpjSuite) TestInvalidCnpj() {
	invalidCnpjs := []string{
		"00.000.000/0000-00",
		"11111111000111",
		"12.345.678/0001-00",
	}

	for _, cnpj := range invalidCnpjs {
		_, err := vo.NewCnpjVo(cnpj)
		suite.NotNil(err)
	}
}

func TestCnpjSuite(t *testing.T) {
	suite.Run(t, new(CnpjSuite))
}
