package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGoogleTaskService(t *testing.T) {
	service, err := NewGoogleTaskService()
	if err != nil {
		t.Fatal(err)
	}
	assert.Nil(t, service.Init())
	lists, err := service.getTaskLists()
	assert.Nil(t, err)

	for _, list := range lists {
		t.Log("xxxxxxxxxxxxxxxxxxxxxxxxxxxx [" + list.Id + "]")
		t.Log("xxxxxxxxxxxxxxxxxxxxxxxxxxxx [" + list.Title + "]")
		t.Log("xxxxxxxxxxxxxxxxxxxxxxxxxxxx [" + list.Kind + "]")
	}
}
