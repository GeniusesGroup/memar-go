//go:build darwin

/* For license and copyright information please see LEGAL file in repository */

package os

import (
	"../protocol"
	"./darwin"
)

func init() {
	protocol.OS = &darwin.OS
}
