/* For license and copyright information please see the LEGAL file in the code repository */

package net

import (
	net_p "memar/net/protocol"
	"sync/atomic"
)

// STATUS is the same as the Status.
// Use this type when embed in other struct to solve field & method same name problem(Status struct and Status() method) to satisfy interfaces.
type STATUS = Status

type Status struct {
	status atomic.Uint64
	state  chan net_p.SocketStatus
}

//memar:impl memar/protocol.Network_Status
func (s *Status) Status() net_p.SocketStatus     { return net_p.SocketStatus(s.status.Load()) }
func (s *Status) State() chan net_p.SocketStatus { return s.state }
func (s *Status) SetStatus(ns net_p.SocketStatus) {
	s.status.Store(uint64(ns))

	// TODO::: non blocking logic??
	// notify stream listener that stream state has been changed!
	s.state <- ns
}
