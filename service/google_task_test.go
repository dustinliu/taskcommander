package service

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	tasks "google.golang.org/api/tasks/v1"
)

func TestNewGoogleTaskTest(t *testing.T) {
	now := time.Now().Format(time.RFC3339)
	gt := &tasks.Task{}
	gt.Id = "id"
	gt.Title = "title"
	gt.Notes = "Notes" + note_seperator + `{"focus":true,"project":"project","tags":["tag1","tag2"],"category":1}`
	gt.Due = now
	gt.Completed = &now

	gtask := newGoogleTask(gt)

	assert.Equal(t, "id", gtask.GetId())
	assert.Equal(t, "title", gtask.GetTitle())
	assert.Equal(t, "Notes", gtask.GetNote())
	assert.Equal(t, true, gtask.GetFocus())
	assert.Equal(t, "project", gtask.GetProject())
	assert.Equal(t, []string{"tag1", "tag2"}, gtask.GetTags())
	assert.Equal(t, CategoryNext, gtask.GetCategory())
	assert.Equal(t, now, gtask.GetDue())
	assert.Equal(t, now, gtask.GetCompleted())
}

func TestSetter(t *testing.T) {
	gtask := newGoogleTask(&tasks.Task{Id: "id1"})
	now := time.Now().Format(time.RFC3339)

	gtask.SetTitle("title1")
	gtask.SetNote("Notes1")
	gtask.SetFocus(false)
	gtask.SetProject("project1")
	gtask.SetTag("tag3")
	gtask.SetTag("tag4")
	gtask.SetCategory(CategorySomeday)
	gtask.SetDue(now)
	gtask.SetCompleted(now)

	assert.Equal(t, "id1", gtask.GetId())
	assert.Equal(t, "title1", gtask.GetTitle())
	assert.Equal(t, "Notes1", gtask.GetNote())
	assert.Equal(t, false, gtask.GetFocus())
	assert.Equal(t, "project1", gtask.GetProject())
	assert.Equal(t, []string{"tag3", "tag4"}, gtask.GetTags())
	assert.Equal(t, CategorySomeday, gtask.GetCategory())
	assert.Equal(t, now, gtask.GetDue())
	assert.Equal(t, now, gtask.GetCompleted())
}
