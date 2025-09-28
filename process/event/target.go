/* For license and copyright information please see the LEGAL file in the code repository */

package event

import (
	"sync"

	error_p "memar/error/protocol"
	event_p "memar/event/protocol"
)

// TODO::: Change to use atomic operation instead of sync.

// Target dispatch events to listeners on desire event_p.EventMainType types.
type Target[E event_p.Event] struct {
	sync sync.Mutex
	ls   []event_p.EventListener[E]
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (self *Target[E]) Init() (err error_p.Error) {
	self.ls = make([]event_p.EventListener[E], 0, CNF_InitialListenersLength)
	return
}

//memar:impl memar/event/protocol.Target
func (self *Target[E]) DispatchEvent(event E) (err error_p.Error) {
	self.sync.Lock()
	var ls = self.ls
	for i := 0; i < len(ls); i++ {
		ls[i].EventHandler(event)
	}
	self.sync.Unlock()
	return
}
func (self *Target[E]) AddEventListener(callback event_p.EventListener[E]) (err error_p.Error) {
	self.sync.Lock()
	var dls = self.ls
	dls = append(dls, callback)
	self.ls = dls
	self.sync.Unlock()
	return
}
func (self *Target[E]) RemoveEventListener(callback event_p.EventListener[E]) (err error_p.Error) {
	self.sync.Lock()
	var dls = self.ls
	var ln = len(dls)
	for i := 0; i < ln; i++ {
		if dls[i] == callback {
			copy(dls[i:], dls[i+1:])
			dls = dls[:ln-1]
			self.ls = dls
			break
		}
	}
	self.sync.Unlock()
	return
}
