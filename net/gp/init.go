/* For license and copyright information please see the LEGAL file in the code repository */

package gp

import (
	"memar/process/log"
	error_p "memar/process/error/protocol"
)

// must assign before use the package
var conns Connections

//memar:impl memar/computer/capsule/protocol.LifeCycle
func Init(cs Connections) (err error_p.Error) {
	if conns != nil {
		// err =
		return
	}
	log.Info(&domain, "GP network begin listening...")
	conns = cs
	return
}
