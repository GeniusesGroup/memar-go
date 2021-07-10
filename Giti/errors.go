/* For license and copyright information please see LEGAL file in repository */

package giti

// Errors is the interface that must implement by any Application!
type Errors interface {
	GetErrorByID(id uint64) (err Error)
}
