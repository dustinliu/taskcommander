package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCateogry(t *testing.T) {
	assert.Equal(t, "Inbox", CategoryInbox.Name())
	assert.Equal(t, "Next", CategoryNext.Name())
	assert.Equal(t, "Someday", CategorySomeday.Name())
	assert.Equal(t, "Focus", CategoryFocus.Name())

	a := Category(4) // must hard code the number
	assert.Equal(t, false, a.IsValid())
	assert.Equal(t, false, a.IsValid())
}

func TestStatus(t *testing.T) {
	assert.Equal(t, "Todo", StatusTodo.Name())
	assert.Equal(t, "Done", StatusDone.Name())

	a := Status(2) // must hard code the number
	assert.Equal(t, false, a.IsValid())
	assert.Equal(t, false, a.IsValid())
}
