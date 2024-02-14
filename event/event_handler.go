package event

import (
	"sync"

	"github.com/dustinliu/taskcommander/core"
)

var (
	eventChannel = make(chan Event, 10)
	listeners    = make([]func(Event) Event, 0)
	Wg           = new(sync.WaitGroup)
)

func init() {
	go worker()
}

func SendEvent(event Event) {
	core.GetLogger().Debugf("queued %T: %+v", event, event)
	eventChannel <- event
}

func QueueFunc(f func()) {
	SendEvent(EventWorker{f: f})
}

func ListenEvent(f func(Event) Event) {
	listeners = append(listeners, f)
}

func worker() {
	for {
		event := <-eventChannel
		core.GetLogger().Debugf("received %T: %+v", event, event)
		switch event := event.(type) {
		case EventWorker:
			executeWork(event.f)
		default:
			dispatch(event)
		}
	}
}

func executeWork(f func()) {
	Wg.Add(1)
	defer Wg.Done()
	f()
}

func dispatch(event Event) {
	Wg.Add(1)
	defer Wg.Done()
	for _, listener := range listeners {
		core.GetLogger().Debugf("dispatching %T: %+v", event, event)
		e := listener(event)
		if e == nil {
			return
		}
	}
}
