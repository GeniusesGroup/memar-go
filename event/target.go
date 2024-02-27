/* For license and copyright information please see the LEGAL file in the code repository */

package event

import (
	"sync"

	"memar/protocol"
)

// TODO::: Can implement anything instead use of sync??

// EventTarget dispatch events to listeners on desire protocol.EventMainType types.
type EventTarget[E protocol.Event] struct {
	sync sync.Mutex
	ls   []listener[E]
}

type listener[E protocol.Event] struct {
	eventListener protocol.EventListener[E]
	options       protocol.AddEventListenerOptions
}

//memar:impl memar/protocol.ObjectLifeCycle
func (et *EventTarget[E]) Init() (err protocol.Error) {
	et.ls = make([]listener[E], 0, CNF_InitialListenersLength)
	return
}

//memar:impl memar/protocol.EventTarget
func (et *EventTarget[E]) DispatchEvent(event E) (err protocol.Error) {
	et.sync.Lock()
	var ls = et.ls
	for i := 0; i < len(ls); i++ {
		// TODO::: handle options here or caller layer must handle it?
		ls[i].eventListener.EventHandler(event)
	}
	et.sync.Unlock()
	return
}
func (et *EventTarget[E]) AddEventListener(callback protocol.EventListener[E], options protocol.AddEventListenerOptions) (err protocol.Error) {
	et.sync.Lock()
	var dls = et.ls
	dls = append(dls, listener[E]{callback, options})
	et.ls = dls
	et.sync.Unlock()
	return
}
func (et *EventTarget[E]) RemoveEventListener(callback protocol.EventListener[E], options protocol.EventListenerOptions) (err protocol.Error) {
	et.sync.Lock()
	var dls = et.ls
	var ln = len(dls)
	for i := 0; i < ln; i++ {
		// TODO::: handle options here or caller layer must handle it?
		if dls[i].eventListener == callback {
			copy(dls[i:], dls[i+1:])
			dls = dls[:ln-1]
			et.ls = dls
			break
		}
	}
	et.sync.Unlock()
	return
}
