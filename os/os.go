/* For license and copyright information please see LEGAL file in repository */

package os

import (
	"runtime"

	"../protocol"
	dos "./default"
)

func init() {
	switch runtime.GOOS {
	case "persiaos":
		// TODO:::
	case "windows":
		// TODO:::
	case "linux":
		// TODO:::
	case "bsd":
		// TODO:::
	default:
		protocol.OS = &dos.OS
	}
}
