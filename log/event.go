/* For license and copyright information please see LEGAL file in repository */

package log

import (
	"runtime/debug"

	"github.com/GeniusesGroup/libgo/event"
	"github.com/GeniusesGroup/libgo/protocol"
	// "github.com/GeniusesGroup/libgo/syllab"
	"github.com/GeniusesGroup/libgo/time/unix"
)

// Event implement protocol.LogEvent
type Event struct {
	event.Event

	message string
	stack   []byte
}

func (e *Event) MainType() protocol.EventMainType { return protocol.EventMainType_Log }

func (e *Event) Level() protocol.LogType { return e.Event.SubType() }
func (e *Event) Message() string         { return e.message }
func (e *Event) Stack() []byte           { return e.stack }

func (e *Event) Init(level protocol.LogType, domain, message string, stack bool) {
	e.message = message
	if stack {
		e.stack = debug.Stack()
	}
	// e.Event.SetSubType(level)
	// e.Event.SetDomain(domain)
	// e.Event.SetNodeID([16]byte{})
	// e.Event.SetTime(unix.Now())
	e.Event.Init(level, domain, [16]byte{}, unix.Now())
}

/*
	-- protocol.Syllab interface --
*/
func (e *Event) CheckSyllab(payload []byte) (err protocol.Error) {
	if len(payload) < int(e.LenOfSyllabStack()) {
		// err = &syllab.ErrShortArrayDecode
	}
	return
}
func (e *Event) FromSyllab(payload []byte, stackIndex uint32) {
}
func (e *Event) ToSyllab(payload []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32) {
	return
}
func (e *Event) LenAsSyllab() uint64      { return uint64(e.LenOfSyllabStack() + e.LenOfSyllabHeap()) }
func (e *Event) LenOfSyllabStack() uint32 { return 16 + e.Event.LenOfSyllabStack() }
func (e *Event) LenOfSyllabHeap() (ln uint32) {
	return uint32(len(e.stack)+len(e.message)) + e.Event.LenOfSyllabHeap()
}
