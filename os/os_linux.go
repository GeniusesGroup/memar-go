//go:build linux

/* For license and copyright information please see LEGAL file in repository */

package os

import (
	"../protocol"
	"./linux"
)

func init() {
	protocol.OS = &linux.OS
}
