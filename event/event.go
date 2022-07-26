/* For license and copyright information please see LEGAL file in repository */

package event

import (
	"github.com/GeniusesGroup/libgo/protocol"
	"github.com/GeniusesGroup/libgo/time/unix"
)

// Event implement protocol.LogEvent
type Event struct {
	subType protocol.EventSubType
	domain  string
	nodeID  protocol.NodeID
	time    unix.Time
}

func (e *Event) Init(subType protocol.EventSubType, domain string, nodeID [16]byte, time unix.Time) {
	e.subType = subType
	e.domain = domain
	e.nodeID = nodeID
	e.time = time
}

func (e *Event) MainType() protocol.EventMainType { return protocol.EventMainType_Unset }
func (e *Event) SubType() protocol.EventSubType   { return e.subType }
func (e *Event) Domain() string                   { return e.domain }
func (e *Event) NodeID() protocol.NodeID          { return e.nodeID }
func (e *Event) Time() protocol.Time              { return &e.time }
func (e *Event) Cancelable() bool                 { return false }
func (e *Event) DefaultPrevented() bool           { return false }
func (e *Event) Bubbles() bool                    { return false }
func (e *Event) PreventDefault()                  {}

func (e *Event) SetSubType(subType protocol.EventSubType)   { e.subType = subType }
func (e *Event) SetDomain(domain string)             { e.domain = domain }
func (e *Event) SetNodeID(nodeID protocol.NodeID)    { e.nodeID = nodeID }
func (e *Event) SetTime(time unix.Time)              { e.time = time }

/*
	-- protocol.Syllab interface Encoder & Decoder --
*/
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
func (e *Event) LenAsSyllab() uint64      { return uint64(e.LenOfSyllabStack() + e.LenOfSyllabHeap()) }
func (e *Event) LenOfSyllabStack() uint32 { return 33 }
func (e *Event) LenOfSyllabHeap() (ln uint32) {
	return uint32(len(e.domain))
}
