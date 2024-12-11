package sorcerydualbackend

import "context"

// event type
const (
	EVENT_ENTER    = iota
	EVENT_DAMAGE   = iota
	EVENT_DIE      = iota
	EVENT_ADD_LIFE = iota
)

// timing
const (
	EVENT_BEFORE = iota
	EVENT_AFTER  = iota
)

type Game struct {
}

type Trigger interface {
	Execute(context.Context, *Game, *Executor) (context.Context, *Game, error)
}

type BasicEvent interface {
	Execute(context.Context, *Game, *Executor) (context.Context, *Game, error)
	GetEventType() int
}

type Executor struct {
	basicEventQueue    []BasicEvent
	registeredTriggers map[int]map[int][]Trigger
}

func (e *Executor) SubmitBasicEvent(basicEvent BasicEvent) error {
	e.basicEventQueue = append(e.basicEventQueue, basicEvent)
	return nil
}

func (e *Executor) RegisterTrigger(trigger Trigger, event int, timing int) error {
	if _, ok := e.registeredTriggers[event]; !ok {
		e.registeredTriggers[event] = make(map[int][]Trigger)
	}
	if _, ok := e.registeredTriggers[event][timing]; !ok {
		e.registeredTriggers[event][timing] = make([]Trigger, 0)
	}
	e.registeredTriggers[event][timing] = append(e.registeredTriggers[event][timing], trigger)
	return nil
}

func (e *Executor) ExecuteAll() {
	for _, basicEvent := range e.basicEventQueue {
		for _, trigger := range e.registeredTriggers[basicEvent.GetEventType()][EVENT_BEFORE] {
			trigger.Execute(context.Background(), nil, e)
		}
		basicEvent.Execute(context.Background(), nil, e)
		for _, trigger := range e.registeredTriggers[basicEvent.GetEventType()][EVENT_AFTER] {
			trigger.Execute(context.Background(), nil, e)
		}
	}
}
