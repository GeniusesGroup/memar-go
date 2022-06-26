//go:build freebsd

/* For license and copyright information please see LEGAL file in repository */

package os

import (
	"../freebsd"
	"../protocol"
)

func init() {
	protocol.OS = &freebsd.OS
}
