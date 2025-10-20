//go:build !race

/* For license and copyright information please see the LEGAL file in the code repository */

package race

import (
	error_p "memar/process/error/protocol"
)

const DetectorEnabled = false

type Dynamic struct {}

// Init will initialize the capsule.
//
//memar:impl memar/computer/capsule/protocol.LifeCycle
func (self *Dynamic) Init(f any) (err error_p.Error)   { return }
func (self *Dynamic) Reinit(f any) (err error_p.Error) { return }
func (self *Dynamic) Deinit() (err error_p.Error)      { return }

func (self *Dynamic) Acquire(addr memory_p.Pointer) (err error_p.Error) { return }
func (self *Dynamic) Release(addr memory_p.Pointer) (err error_p.Error) { return }
