/* For license and copyright information please see the LEGAL file in the code repository */

package gp

import (
	"libgo/log"
	"libgo/protocol"
)

// must assign before use the package
var conns Connections

//libgo:impl libgo/protocol.ObjectLifeCycle
func Init(cs Connections) (err protocol.Error) {
	if conns != nil {
		// err =
		return
	}
	log.Info(&Package_MediaType, "GP network begin listening...")
	conns = cs
	return
}
