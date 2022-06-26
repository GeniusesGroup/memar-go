//go:build windows

/* For license and copyright information please see LEGAL file in repository */

package os

import (
	"../protocol"
	"./windows"
)

func init() {
	protocol.OS = &windows.OS
}
