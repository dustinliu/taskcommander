package view

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TaskFormTestSuite struct {
	suite.Suite
}

func TestTaskFormTestSuite(t *testing.T) {
	t.Skip("skipping test in short mode.")
	suite.Run(t, new(TaskFormTestSuite))
}

func (suite *TaskFormTestSuite) TestGetKeyword() {
	// t := newTagsInputField()

	// suite.Equal("abc", t.GetKeywords("abc"))
	// suite.Equal("def", t.GetKeywords("abc def"))
	// suite.Equal("def", t.GetKeywords("abc   def"))
	// suite.Equal("ghi", t.GetKeywords("abc   def ghi"))
	// suite.Equal("", t.GetKeywords("abc   def ghi "))
}

func (suite *TaskFormTestSuite) TestComplete() {
	// t := newTagsInputField()
	// t.compList = []string{"test1", "test2", "tag1", "tag2"}

	// suite.Equal([]string{"test1", "test2", "tag1", "tag2"}, t.complete("t"))
	// suite.Equal([]string{"test1", "test2"}, t.complete("te"))
	// suite.Equal([]string{"tag1", "tag2"}, t.complete("ta"))
	// suite.Equal([]string{}, t.complete("test1 "))
	// suite.Equal([]string{}, t.complete("test1 tag2 "))
}

func (suite *TaskFormTestSuite) TestCompleted() {
	// t := newTagsInputField()
	// t.compList = []string{"test1", "test2", "tag1", "tag2"}

	// t.SetText("te")
	// suite.False(t.completed("test1", 0, tview.AutocompletedNavigate))
	// suite.Equal("test1", t.GetText())
}
