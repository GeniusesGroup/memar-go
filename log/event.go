/* For license and copyright information please see the LEGAL file in the code repository */

package log

import (
	"runtime/debug"

	"memar/event"
	"memar/protocol"

	// "memar/syllab"
	"memar/time/unix"
)

// Event implement protocol.LogEvent
type Event struct {
	event.Event

	level   protocol.LogLevel
	message string
	stack   []byte
}

//memar:impl memar/protocol.LogEvent
func (e *Event) Level() protocol.LogLevel { return e.level }
func (e *Event) Message() string          { return e.message }
func (e *Event) Stack() []byte            { return e.stack }

//memar:impl memar/protocol.ObjectLifeCycle
func (e *Event) Init(domain protocol.MediaType, level protocol.LogLevel, message string, stack bool) {
	e.level = level
	e.message = message
	if stack {
		e.stack = debug.Stack()
	}
	// TODO::: set NodeID???
	e.Event.Init(domain, [16]byte{}, unix.Now())
}

//memar:impl memar/protocol.Codec
func (e *Event) MediaType() protocol.MediaType       { return e.Event.Domain() }
func (e *Event) CompressType() protocol.CompressType { return nil }

//memar:impl memar/protocol.Decoder
func (e *Event) Decode(source protocol.Codec) (n int, err protocol.Error) {
	return
}

//memar:impl memar/protocol.Encoder
func (e *Event) Encode(destination protocol.Codec) (n int, err protocol.Error) {
	return
}

//memar:impl memar/protocol.Unmarshaler
func (e *Event) Unmarshal(data []byte) (n int, err protocol.Error) {
	return
}
func (e *Event) UnmarshalFrom(data []byte) (remaining []byte, err protocol.Error) {
	return
}

//memar:impl memar/protocol.Marshaler
func (e *Event) Marshal() (data []byte, err protocol.Error) {
	return
}
func (e *Event) MarshalTo(data []byte) (added []byte, err protocol.Error) {
	return
}

//memar:impl memar/protocol.Len
func (e *Event) Len() int { return 0 }

//memar:impl memar/protocol.Syllab
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
