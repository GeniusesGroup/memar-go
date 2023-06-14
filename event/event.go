/* For license and copyright information please see the LEGAL file in the code repository */

package event

import (
	"libgo/protocol"
	"libgo/time/unix"
)

// Event implement protocol.LogEvent
type Event struct {
	mediaType protocol.MediaType
	domain    protocol.MediaTypeID
	nodeID    protocol.NodeID
	time      unix.Time
}

//libgo:impl libgo/protocol.ObjectLifeCycle
func (e *Event) Init(mediaType protocol.MediaType, nodeID [16]byte, time unix.Time) {
	e.mediaType = mediaType
	e.domain = mediaType.ID()
	e.nodeID = nodeID
	e.time = time
}

//libgo:impl libgo/protocol.Event
func (e *Event) MediaType() protocol.MediaType { return e.mediaType }
func (e *Event) Domain() protocol.MediaTypeID  { return e.domain }
func (e *Event) NodeID() protocol.NodeID       { return e.nodeID }
func (e *Event) Time() protocol.Time           { return &e.time }
func (e *Event) Cancelable() bool              { return false }
func (e *Event) DefaultPrevented() bool        { return false }
func (e *Event) Bubbles() bool                 { return false }
func (e *Event) PreventDefault()               {}

//libgo:impl libgo/protocol.Syllab
func (e *Event) CheckSyllab(payload []byte) (err protocol.Error) {
	if len(payload) < int(e.LenOfSyllabStack()) {
		// err = syllab.ErrShortArrayDecode
	}
	return
}
func (e *Event) FromSyllab(payload []byte, stackIndex uint32) {
}
func (e *Event) ToSyllab(payload []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32) {
	return
}
func (e *Event) LenAsSyllab() uint64          { return uint64(e.LenOfSyllabStack() + e.LenOfSyllabHeap()) }
func (e *Event) LenOfSyllabStack() uint32     { return 36 }
func (e *Event) LenOfSyllabHeap() (ln uint32) { return }
