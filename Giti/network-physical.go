/* For license and copyright information please see LEGAL file in repository */

package giti

/*
**********************************************************************************
Physical - (OSI Layer 1: Physical)
**********************************************************************************
*/

// PhysicalConnection or Hardware2HardwareConnection
type PhysicalConnection interface {
	Send(frame []byte) (err Error)
	// SendAsync(frame []byte) (err Error)
}
