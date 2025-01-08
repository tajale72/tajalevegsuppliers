package integration

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// CalculatorTestSuite defines the suite for testing the Add function.
type CalculatorTestSuite struct {
	suite.Suite
}

// SetupSuite runs before the suite starts.
func (suite *CalculatorTestSuite) SetupSuite() {
	suite.T().Log("SetupSuite: Initialize resources")
}

// TearDownSuite runs after the suite finishes.
func (suite *CalculatorTestSuite) TearDownSuite() {
	suite.T().Log("TearDownSuite: Cleanup resources")
}

// TestAddPositiveNumbers tests the addition of two positive numbers.
func (suite *CalculatorTestSuite) TestAddPositiveNumbers() {
	result := Add(1, 2)
	suite.Equal(3, result, "Add(1, 2) should return 3")
}

// TestAddNegativeNumbers tests the addition of two negative numbers.
func (suite *CalculatorTestSuite) TestAddNegativeNumbers() {
	result := Add(-1, -2)
	suite.Equal(-3, result, "Add(-1, -2) should return -3")
}

// TestAddMixedNumbers tests the addition of a positive and a negative number.
func (suite *CalculatorTestSuite) TestAddMixedNumbers() {
	result := Add(5, -3)
	suite.Equal(2, result, "Add(5, -3) should return 2")
}

// TestAddWithZero tests the addition of a number and zero.
func (suite *CalculatorTestSuite) TestAddWithZero() {
	result := Add(7, 0)
	suite.Equal(7, result, "Add(7, 0) should return 7")
}

// TestRunSuite is the entry point for running the suite.
func TestRunSuite(t *testing.T) {
	suite.Run(t, new(CalculatorTestSuite))
}
