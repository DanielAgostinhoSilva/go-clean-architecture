package vo

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// CpfTestSuite defines a test suite for the Cpf VO
type CpfTestSuite struct {
	suite.Suite
}

// SetupSuite runs once when the suite starts
func (suite *CpfTestSuite) SetupSuite() {
	// Initialization if needed
}

// TearDownSuite runs once when the suite finishes
func (suite *CpfTestSuite) TearDownSuite() {
	// Cleanup if needed
}

// TestValidCpf tests valid CPF inputs
func (suite *CpfTestSuite) TestValidCpf() {
	validCpfWithFormat := "123.456.789-09"
	validCpfWithoutFormat := "12345678909"

	cpfVo1, err1 := NewCpfVo(validCpfWithFormat)
	cpfVo2, err2 := NewCpfVo(validCpfWithoutFormat)

	assert.NoError(suite.T(), err1)
	assert.NotNil(suite.T(), cpfVo1)
	assert.Equal(suite.T(), validCpfWithFormat, cpfVo1.Value()) // Assuming we keep the format

	assert.NoError(suite.T(), err2)
	assert.NotNil(suite.T(), cpfVo2)
	assert.Equal(suite.T(), validCpfWithoutFormat, cpfVo2.Value())
}

// TestInvalidCpf tests invalid CPF inputs
func (suite *CpfTestSuite) TestInvalidCpf() {
	invalidCpfCases := []string{
		"123",            // too short
		"123.456.789-0X", // invalid character
		"000.000.000-00", // invalid CPF number
		"12345678900",    // incorrect check digits
	}

	for _, invalidCpf := range invalidCpfCases {
		cpfVo, err := NewCpfVo(invalidCpf)
		assert.Error(suite.T(), err)
		assert.Nil(suite.T(), cpfVo)
		assert.Equal(suite.T(), ErrInvalidCpf, err)
	}
}

// TestCpfTestSuite executes the test suite
func TestCpfTestSuite(t *testing.T) {
	suite.Run(t, new(CpfTestSuite))
}
