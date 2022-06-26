//go:build netbsd

/* For license and copyright information please see LEGAL file in repository */

package os

import (
	"../protocol"
	"./netbsd"
)

func init() {
	protocol.OS = &netbsd.OS
}
