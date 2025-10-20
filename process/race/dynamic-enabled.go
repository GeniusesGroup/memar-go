//go:build race

/* For license and copyright information please see the LEGAL file in the code repository */

package race

import (
	error_p "memar/process/error/protocol"
	memory_p "memar/storage/memory/protocol"
)

// Public race detection API, present iff build with -race.
const DetectorEnabled = true

type Dynamic struct {
	// Race context used while executing related functions e.g. timers, ...
	raceCTX uintptr
}

// Init will initialize the capsule.
// f MUST be any function or method.
//
//memar:impl memar/computer/capsule/protocol.LifeCycle
func (self *Dynamic) Init(f any) (err error_p.Error) {
	// self.raceCTX = racegostart(abi.FuncPCABIInternal(f) + sys.PCQuantum)
	return
}
func (self *Dynamic) Reinit(f any) (err error_p.Error) {
	return
}
func (self *Dynamic) Deinit() (err error_p.Error) {
	// racectxend(self.raceCTX)
	self.raceCTX = 0
	return
}

func (self *Dynamic) Acquire(addr memory_p.Pointer) (err error_p.Error) { 
	// unsafe.Pointer
	return
}
func (self *Dynamic) Release(addr memory_p.Pointer) (err error_p.Error) {
	return
}
