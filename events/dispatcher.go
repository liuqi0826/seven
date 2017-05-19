package events

import (
	"fmt"
)

type IEventDispatcher interface {
	DispatchEvent(event *Event)
	AddEventListener(eventType string, listener func(*Event))
	RemoveEventListener(eventType string, listener func(*Event))
	HasEventListener(eventType string) bool
}

type EventDispatcher struct {
	list map[string]*dispatcher
}

func (this *EventDispatcher) EventDispatcher() {
	this.list = make(map[string]*dispatcher)
}
func (this *EventDispatcher) DispatchEvent(event *Event) {
	if this.list == nil {
		this.list = make(map[string]*dispatcher)
	}
	if d, ok := this.list[event.Type]; ok {
		d.dispatch(event)
	}
}
func (this *EventDispatcher) AddEventListener(eventType string, listener func(*Event)) {
	if this.list == nil {
		this.list = make(map[string]*dispatcher)
	}
	if d, ok := this.list[eventType]; ok {
		d.addListener(listener)
	} else {
		dph := new(dispatcher)
		dph.dispatcher(eventType)
		dph.addListener(listener)
		this.list[eventType] = dph
	}
}
func (this *EventDispatcher) RemoveEventListener(eventType string, listener func(*Event)) {
	if this.list == nil {
		this.list = make(map[string]*dispatcher)
		return
	}
	if d, ok := this.list[eventType]; ok {
		d.removeListener(listener)
	}
}
func (this *EventDispatcher) HasEventListener(eventType string) bool {
	if _, ok := this.list[eventType]; ok {
		return true
	}
	return false
}

//++++++++++++++++++++ dispatch ++++++++++++++++++++

type dispatcher struct {
	Type     string
	listener []func(*Event)
}

func (this *dispatcher) dispatcher(eventType string) {
	this.Type = eventType
}
func (this *dispatcher) dispatch(event *Event) {
	for _, lis := range this.listener {
		go lis(event)
	}
}
func (this *dispatcher) addListener(listener func(*Event)) {
	if this.hasListener(listener) < 0 {
		this.listener = append(this.listener, listener)
	}
}
func (this *dispatcher) removeListener(listener func(*Event)) {
	if this.listener == nil {
		return
	}
	idx := this.hasListener(listener)
	if idx > -1 {
		this.listener = append(this.listener[:idx], this.listener[idx+1:]...)
	}
}
func (this *dispatcher) hasListener(listener func(*Event)) int {
	if this.listener == nil {
		return -1
	}
	for idx, l := range this.listener {
		la := fmt.Sprintf("%d", l)
		lb := fmt.Sprintf("%d", listener)
		if la == lb {
			return idx
		}
	}
	return -1
}
func (this *dispatcher) dispose() {
	this.listener = nil
}
