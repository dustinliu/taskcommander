package view

import (
	"github.com/dustinliu/taskcommander/service"
	"github.com/gdamore/tcell/v2"
)

var ch chan tcell.Event = make(chan tcell.Event)

type EventCategoryChange struct {
	*tcell.EventTime
	Category service.Category
}

func NewEventCategoryChange(cat service.Category) EventCategoryChange {
	return EventCategoryChange{
		newEventTime(),
		cat,
	}
}

func SendEvent(event tcell.Event) {
	ch <- event
}

func PollEvent() tcell.Event {
	return <-ch
}

func newEventTime() *tcell.EventTime {
	e := tcell.EventTime{}
	e.SetEventNow()
	return &e
}
