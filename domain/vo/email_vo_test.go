package vo_test

import (
	"testing"

	"github.com/DanielAgostinhoSilva/go-clean-architecture/domain/vo"
	"github.com/stretchr/testify/suite"
)

type EmailSuite struct {
	suite.Suite
}

func (suite *EmailSuite) TestValidEmail() {
	validEmails := []string{
		"test@example.com",
		"user.name+tag+sorting@example.com",
		"x@example.com",
		"user.name@domain.co.uk",
		"user_name@domain.com",
		"user+name@example.co.in",
		"user-name@domain.org",
	}

	for _, email := range validEmails {
		_, err := vo.NewEmailVo(email)
		suite.Nil(err)
	}
}

func (suite *EmailSuite) TestInvalidEmail() {
	invalidEmails := []string{
		"plainaddress",
		"@missingusername.com",
		"user@.com",
		"user@domain",
		"@domain.com",
		"user@domain..com",
		"user@@domain.com",
		"username.com",
		"user@domain.c",
	}

	for _, email := range invalidEmails {
		_, err := vo.NewEmailVo(email)
		suite.NotNil(err)
	}
}

func TestEmailSuite(t *testing.T) {
	suite.Run(t, new(EmailSuite))
}
