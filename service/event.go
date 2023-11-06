package service

import (
	"github.com/gdamore/tcell/v2"
)

var Events chan tcell.Event = make(chan tcell.Event, 5)

type EventQuit struct {
	*tcell.EventTime
}

type EventAddTask struct {
	*tcell.EventTime
	Task Task
}

type EventCategoryChange struct {
	*tcell.EventTime
	Category Category
}

type EventTaskChange struct {
	*tcell.EventTime
	Task Task
}

func NewEventQuit() EventQuit {
	return EventQuit{
		newEventTime(),
	}
}

func NewEventAddTask(task Task) EventAddTask {
	return EventAddTask{
		newEventTime(),
		task,
	}
}

func NewEventCategoryChange(cat Category) EventCategoryChange {
	return EventCategoryChange{
		newEventTime(),
		cat,
	}
}

func NewEventTaskChange(task Task) EventTaskChange {
	return EventTaskChange{
		newEventTime(),
		task,
	}
}

func newEventTime() *tcell.EventTime {
	e := tcell.EventTime{}
	e.SetEventNow()
	return &e
}
