package event

import (
	"time"

	"github.com/dustinliu/taskcommander/service"
)

type Event interface {
	When() time.Time
}

type EventCategoryChanged struct {
	t        time.Time
	Category service.Category
}

func NewEventCategoryChanged(c service.Category) EventCategoryChanged {
	return EventCategoryChanged{time.Now(), c}
}

func (e EventCategoryChanged) When() time.Time {
	return e.t
}

type EventError struct {
	t     time.Time
	Err   error
	Fatal bool
}

func NewEventError(err error, fatal bool) EventError {
	return EventError{time.Now(), err, fatal}
}

func (e EventError) When() time.Time {
	return e.t
}

type EventWorker struct {
	t time.Time
	f func()
}

func (e EventWorker) When() time.Time {
	return e.t
}
