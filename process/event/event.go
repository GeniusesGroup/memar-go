/* For license and copyright information please see the LEGAL file in the code repository */

package event

import (
	datatype_p "memar/datatype/protocol"
	event_p "memar/event/protocol"
	time_p "memar/time/protocol"
	"memar/time/unix"
)

// Event implement memar/event/protocol.Event
//
//memar:impl memar/event/protocol.Event
type Event struct {
	dt   datatype_p.DataType
	time unix.Time
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (self *Event) Init(dt datatype_p.DataType, time unix.Time) {
	self.dt = dt
	self.time = time
}

//memar:impl memar/datatype/protocol.Field_DataType
func (self *Event) DataType() datatype_p.DataType { return self.dt }

//memar:impl memar/time/protocol.Field_Time
func (self *Event) Time() time_p.Time { return &self.time }

//memar:impl memar/math/logic/protocol.Equivalence
func (self *Event) Equivalence(with event_p.Event) bool { return self.DataType() == with.DataType() }
