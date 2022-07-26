/* For license and copyright information please see LEGAL file in repository */

package log

import (
	"runtime"

	"github.com/GeniusesGroup/libgo/protocol"
)

func CheckLevelEnabled(level protocol.LogType) bool {
	return protocol.LogMode&level != 0 
}

func CallerInfo(calldepth int) (file string, line int) {
	var ok bool
	_, file, line, ok = runtime.Caller(calldepth)
	if !ok {
		file = "?"
		line = 0
	}
	return
}

// FatalError use as log.FatalError(function()) and not check return error from function.
// It will just panic error not exit app and return to OS, Because all goroutine without any notify will terminate and can't recover in any way.
// So we just panic it and wait to some logic recover it or let app close in main function.
func FatalError(err protocol.Error) {
	if err != nil {
		// os.Exit(125)
		panic(err)
	}
}
