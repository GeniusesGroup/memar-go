//go:build persiaos

/* For license and copyright information please see LEGAL file in repository */

package os

import (
	"../protocol"
	"./persiaos"
)

func init() {
	protocol.OS = &persiaos.OS
}
