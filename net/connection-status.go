/* For license and copyright information please see the LEGAL file in the code repository */

package net

import (
	"sync/atomic"

	"memar/protocol"
)

// STATUS is the same as the Status.
// Use this type when embed in other struct to solve field & method same name problem(Status struct and Status() method) to satisfy interfaces.
type STATUS = Status

type Status struct {
	status atomic.Uint32
	state  chan protocol.NetworkStatus
}

//memar:impl memar/protocol.Network_Status
func (s *Status) Status() protocol.NetworkStatus     { return protocol.NetworkStatus(s.status.Load()) }
func (s *Status) State() chan protocol.NetworkStatus { return s.state }
func (s *Status) SetStatus(ns protocol.NetworkStatus) {
	s.status.Store(uint32(ns))

	// TODO::: non blocking logic??
	// notify stream listener that stream state has been changed!
	s.state <- ns
}
