package test

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type FindOddNumberTestSuite struct {
	suite.Suite
	service func(text []int) int
}

func (suite *FindOddNumberTestSuite) SetupSuite() {
	suite.service = FindOddNumber
}

func (suite *FindOddNumberTestSuite) TestOneFindOddNumber_1() {
	result := suite.service([]int{7})
	suite.Equal(7, result)
}

func (suite *FindOddNumberTestSuite) TestOneFindOddNumber_2() {
	result := suite.service([]int{0})
	suite.Equal(0, result)
}

func (suite *FindOddNumberTestSuite) TestTreeFindOddNumber() {
	result := suite.service([]int{1, 1, 2})
	suite.Equal(2, result)
}

func (suite *FindOddNumberTestSuite) TestManyFindOddNumber_1() {
	result := suite.service([]int{0, 1, 0, 1, 0})
	suite.Equal(0, result)
}

func (suite *FindOddNumberTestSuite) TestManyFindOddNumber_2() {
	result := suite.service([]int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1})
	suite.Equal(4, result)
}

func (suite *FindOddNumberTestSuite) TestNotFoundOddNumber() {
	result := suite.service([]int{1, 2, 2, 3, 3, 3, 4, 4, 3, 3, 3, 2, 2, 1})
	suite.Equal(-1, result)
}

func TestFindOddNumberTestSuite(t *testing.T) {
	suite.Run(t, new(FindOddNumberTestSuite))
}
