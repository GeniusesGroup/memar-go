/* For license and copyright information please see LEGAL file in repository */

package event

import (
	"sync"

	"../protocol"
)

// EventTarget must declare separately for each protocol.EventMainType types.
// otherwise need this struct that use 2KB for each instance.
// 		cl   [256]*[]customListener // 256 is max protocol.EventMainType
// Or use map but need some benchmarks to check performance.
type EventTarget struct {
	sync sync.Mutex
	lls  *[]customListener
}

type customListener struct {
	eventSubType  protocol.EventSubType
	eventListener protocol.EventListener
}

func (et *EventTarget) DispatchEvent(event protocol.LogEvent) {
	// TODO::: add atomic mechanism??
	var lls = *et.lls
	var eventSubType = event.SubType()
	for i := 0; i < len(lls); i++ {
		var cl = lls[i]
		if cl.eventSubType == protocol.EventSubType_Unset || cl.eventSubType == eventSubType {
			cl.eventListener.EventHandler(event)
		}
	}
}

func (et *EventTarget) AddEventListener(mainType protocol.EventMainType, subType protocol.EventSubType, callback protocol.EventListener, options protocol.AddEventListenerOptions) {
	et.sync.Lock()
	var lls = *et.lls
	var ln = len(lls)
	var newLLS = make([]customListener, ln+1)
	copy(newLLS, lls)
	newLLS[ln-1] = customListener{subType, callback}
	et.lls = &newLLS
	// TODO::: handle options here or caller layer must handle it?
	et.sync.Unlock()
}

func (et *EventTarget) RemoveEventListener(mainType protocol.EventMainType, subType protocol.EventSubType, callback protocol.EventListener, options protocol.EventListenerOptions) {
	et.sync.Lock()
	var lls = *et.lls
	var ln = len(lls)
	for i := 0; i < ln; i++ {
		var cl = lls[i]
		if cl.eventSubType == subType && cl.eventListener == callback {
			var newLLS = make([]customListener, ln-1)
			copy(newLLS, lls[:i])
			copy(newLLS[i:], lls[i+1:])
			et.lls = &newLLS
			break
		}
	}
	et.sync.Unlock()
}
