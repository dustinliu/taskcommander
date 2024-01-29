package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCateogry(t *testing.T) {
	a := Category(0)
	assert.Equal(t, "Inbox", a.Name())
}
