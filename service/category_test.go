package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
)

type CategoryTestSuite struct {
	suite.Suite
}

func TestCategoryTestSuite(t *testing.T) {
	suite.Run(t, new(CategoryTestSuite))
}

func (suite *CategoryTestSuite) TestCategoryFromName() {
	testCases := []struct {
		name     string
		expected Category
		err      error
	}{
		{
			name:     "Inbox",
			expected: 0,
		},
		{
			name:     "Next",
			expected: Next,
		},
		{
			name:     "Someday",
			expected: Someday,
		},
		{
			name:     "Focus",
			expected: Focus,
		},
		{
			name:     "dfdsf",
			expected: -1,
		},
	}

	for _, tc := range testCases {
		actual := CategoryFromName(tc.name)
		suite.Equal(tc.expected, actual)
	}
}

func (suite *CategoryTestSuite) TestMarshalJSON() {
	cats := []struct {
		cat      Category
		expected []byte
		err      error
	}{
		{
			2,
			[]byte("Next"),
			nil,
		},
		{
			5,
			nil,
			errors.New("invalid category (5)"),
		},
	}

	for _, tc := range cats {
		actual, err := tc.cat.MarshalJSON()
		suite.Equal(tc.expected, actual)
		suite.Equal(tc.err, err)
	}
}
