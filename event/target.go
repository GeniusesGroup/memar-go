/* For license and copyright information please see the LEGAL file in the code repository */

package event

import (
	"sync"

	"memar/protocol"
)

// TODO::: Can implement a parallel hash table here instead use of map and sync?

// EventTarget dispatch events to listeners on desire protocol.EventMainType types.
type EventTarget struct {
	sync sync.Mutex
	ls   map[protocol.MediaTypeID][]listener
}

type listener struct {
	eventListener protocol.EventListener
	options       protocol.AddEventListenerOptions
}

//memar:impl memar/protocol.ObjectLifeCycle
func (et *EventTarget) Init() (err protocol.Error) {
	et.ls = make(map[protocol.ID][]listener)
	return
}

//memar:impl memar/protocol.EventTarget
func (et *EventTarget) DispatchEvent(event protocol.Event) (err protocol.Error) {
	et.sync.Lock()
	var eventDomain = event.DomainID()
	var ls = et.ls[eventDomain]
	for i := 0; i < len(ls); i++ {
		// TODO::: handle options here or caller layer must handle it?
		ls[i].eventListener.EventHandler(event)
	}
	et.sync.Unlock()
	return
}
func (et *EventTarget) AddEventListener(domain protocol.MediaTypeID, callback protocol.EventListener, options protocol.AddEventListenerOptions) (err protocol.Error) {
	et.sync.Lock()
	var dls = et.ls[domain]
	dls = append(dls, listener{callback, options})
	et.ls[domain] = dls
	et.sync.Unlock()
	return
}
func (et *EventTarget) RemoveEventListener(domain protocol.MediaTypeID, callback protocol.EventListener, options protocol.EventListenerOptions) (err protocol.Error) {
	et.sync.Lock()
	var dls = et.ls[domain]
	var ln = len(dls)
	for i := 0; i < ln; i++ {
		// TODO::: handle options here or caller layer must handle it?
		if dls[i].eventListener == callback {
			copy(dls[i:], dls[i+1:])
			dls = dls[:ln-1]
			et.ls[domain] = dls
			break
		}
	}
	et.sync.Unlock()
	return
}
