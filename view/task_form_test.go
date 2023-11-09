package view

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TaskFormTestSuite struct {
	suite.Suite
}

func TestTaskFormTestSuite(t *testing.T) {
	suite.Run(t, new(TaskFormTestSuite))
}

func (suite *TaskFormTestSuite) TestGetKeyword() {
	t := newTagsInputField()

	t.SetText("abc")
	suite.Equal("abc", t.GetKeywords())

	t.SetText("abc def")
	suite.Equal("def", t.GetKeywords())

	t.SetText("abc   def")
	suite.Equal("def", t.GetKeywords())

	t.SetText("abc def ghi")
	suite.Equal("ghi", t.GetKeywords())
}
