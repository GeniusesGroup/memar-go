//go:build openbsd

/* For license and copyright information please see LEGAL file in repository */

package os

import (
	"../openbsd"
	"../protocol"
)

func init() {
	protocol.OS = &openbsd.OS
}
