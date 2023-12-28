package test

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type CountSmilyFaceTestSuite struct {
	suite.Suite
	service func(text []string) int
}

func (suite *CountSmilyFaceTestSuite) SetupSuite() {
	suite.service = CountSmilyFace
}

func (suite *CountSmilyFaceTestSuite) TestTwoCountSmilyFace() {
	result := suite.service([]string{":)", ";(", ";}", ":-D"})
	suite.Equal(2, result)
}

func (suite *CountSmilyFaceTestSuite) TestTreeCountSmilyFace() {
	result := suite.service([]string{";D", ":-(", ":-)", ";~)"})
	suite.Equal(3, result)
}

func (suite *CountSmilyFaceTestSuite) TestOneCountSmilyFace() {
	result := suite.service([]string{";]", ":[", ";*", ":$", ";-D"})
	suite.Equal(1, result)
}

func (suite *CountSmilyFaceTestSuite) TestNotFoundSmil() {
	result := suite.service([]string{";]", ":[", ";*", ":$", ";-&"})
	suite.Equal(0, result)
}

func TestCountSmilyFaceTestSuite(t *testing.T) {
	suite.Run(t, new(CountSmilyFaceTestSuite))
}
