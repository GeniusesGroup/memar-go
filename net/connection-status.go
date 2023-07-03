/* For license and copyright information please see the LEGAL file in the code repository */

package net

import (
	"libgo/protocol"
)

// STATUS is the same as the Status.
// Use this type when embed in other struct to solve field & method same name problem(Status struct and Status() method) to satisfy interfaces.
type STATUS = Status

type Status struct {
	status protocol.NetworkStatus
	state  chan protocol.NetworkStatus
}

//libgo:impl libgo/protocol.Network_Status
func (s *Status) Status() protocol.NetworkStatus     { return s.status }
func (s *Status) State() chan protocol.NetworkStatus { return s.state }
func (s *Status) SetStatus(ns protocol.NetworkStatus) {
	// TODO::: atomic and non blocking logic
	// atomic.StoreUInt64(&st.State, state)
	s.status = ns
	// notify stream listener that stream state has been changed!
	s.state <- ns
}
