/* For license and copyright information please see the LEGAL file in the code repository */

package event

import (
	"memar/protocol"
	"memar/time/unix"
)

// Event implement protocol.LogEvent
type Event struct {
	domain   protocol.MediaType
	domainID protocol.MediaTypeID
	nodeID   protocol.NodeID
	time     unix.Time
}

//memar:impl memar/protocol.ObjectLifeCycle
func (e *Event) Init(domain protocol.MediaType, nodeID [16]byte, time unix.Time) {
	e.domain = domain
	e.domainID = domain.ID()
	e.nodeID = nodeID
	e.time = time
}

//memar:impl memar/protocol.Event
func (e *Event) Domain() protocol.MediaType     { return e.domain }
func (e *Event) DomainID() protocol.MediaTypeID { return e.domainID }
func (e *Event) NodeID() protocol.NodeID        { return e.nodeID }
func (e *Event) Time() protocol.Time            { return &e.time }
func (e *Event) Cancelable() bool               { return false }
func (e *Event) DefaultPrevented() bool         { return false }
func (e *Event) Bubbles() bool                  { return false }
func (e *Event) PreventDefault()                {}

//memar:impl memar/protocol.Syllab
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
