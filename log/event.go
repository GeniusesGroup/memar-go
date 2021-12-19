/* For license and copyright information please see LEGAL file in repository */

package log

import (
	"io"
	"runtime/debug"

	// "../mediatype"
	"../protocol"
)

func NewEvent(level protocol.LogType, domian, message string) (event *Event) {
	return &Event{
		level:   level,
		domain:  domian,
		stack:   nil,
		message: message,
	}
}

func TraceEvent(level protocol.LogType, domian, message string) (event *Event) {
	return &Event{
		level:   level,
		domain:  domian,
		stack:   debug.Stack(),
		message: message,
	}
}

func InfoEvent(domian, message string) (event *Event) {
	return &Event{
		level:   protocol.Log_Information,
		domain:  domian,
		stack:   nil,
		message: message,
	}
}

func WarnEvent(domian, message string) (event *Event) {
	return &Event{
		level:   protocol.Log_Warning,
		domain:  domian,
		stack:   nil,
		message: message,
	}
}

// FatalEvent return new event with panic level and added stack trace.
func PanicEvent() (event *Event) {
	return &Event{
		level: protocol.Log_Panic,
		stack: debug.Stack(),
	}
}

// FatalEvent return new event with fatal level and added stack trace.
func FatalEvent(domian, message string) (event *Event) {
	return &Event{
		level:   protocol.Log_Fatal,
		domain:  domian,
		stack:   debug.Stack(),
		message: message,
	}
}

// Event implement protocol.LogEvent
type Event struct {
	level   protocol.LogType
	id      uint64
	domain  string
	stack   []byte
	message string
}

func (e *Event) Level() protocol.LogType { return e.level }
func (e *Event) ID() uint64              { return e.id }
func (e *Event) Domain() string          { return e.domain }
func (e *Event) Stack() []byte           { return e.stack }
func (e *Event) Message() string         { return e.message }

/*
********** protocol.Codec interface **********
 */
func (e *Event) MediaType() protocol.MediaType                      { return nil } // mediatype.LOG }
func (e *Event) CompressType() protocol.CompressType                { return nil }
func (e *Event) Decode(reader protocol.Reader) (err protocol.Error) { return }
func (e *Event) Encode(writer protocol.Writer) (err protocol.Error) {
	var _, goErr = e.WriteTo(writer)
	if goErr != nil {
		// err =
	}
	return
}
func (e *Event) Marshal() (data []byte)                     { return }
func (e *Event) MarshalTo(data []byte) []byte               { return data }
func (e *Event) Unmarshal(data []byte) (err protocol.Error) { return }
func (e *Event) Len() (ln int)                              { return }

/*
********** io package interfaces **********
 */
func (e *Event) ReadFrom(reader io.Reader) (n int64, err error) { return }
func (e *Event) WriteTo(w io.Writer) (totalWrite int64, err error) {
	// var writeLen int
	// writeLen, err = w.Write(r)
	// totalWrite = int64(writeLen)
	return
}

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
	return uint32(len(e.domain) + len(e.stack) + len(e.message))
}
