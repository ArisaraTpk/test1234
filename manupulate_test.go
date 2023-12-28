package test

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type ManipulateTestSuite struct {
	suite.Suite
	service func(text []string) []string
}

func (suite *ManipulateTestSuite) SetupSuite() {
	suite.service = Manipulate
}

func (suite *ManipulateTestSuite) TestZeroCharShuffle() {
	result := suite.service([]string{""})
	suite.Equal([]string{""}, result)
}

func (suite *ManipulateTestSuite) TestOneCharShuffle() {
	result := suite.service([]string{"a"})
	suite.Equal([]string{"a"}, result)
}

func (suite *ManipulateTestSuite) TestTwoCharShuffle() {
	result := suite.service([]string{"a", "b"})
	suite.Equal([]string{"ab", "ba"}, result)
}

func (suite *ManipulateTestSuite) TestTreeCharShuffle() {
	result := suite.service([]string{"a", "b", "c"})
	suite.Equal([]string{"abc", "acb", "bac", "bca", "cab", "cba"}, result)
}

func (suite *ManipulateTestSuite) TestFourCharShuffle() {
	result := suite.service([]string{"a", "a", "b", "b"})
	suite.Equal([]string{"aabb", "abab", "abba", "baab", "baba", "bbaa"}, result)
}

func TestManipulateTestSuite(t *testing.T) {
	suite.Run(t, new(ManipulateTestSuite))
}
