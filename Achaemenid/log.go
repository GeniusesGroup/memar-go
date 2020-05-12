/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

// log hold logs until app running.
var log []byte

// Log will show log in standard console & append log to buffer to save them later.
func Log(a ...interface{}) {
	// print error to stderr
	fmt.Fprintf(os.Stderr, "%v\n", a)
	// Append log to CodeGenerateLog for saving to achaemenid.log later.
	log = append(log, fmt.Sprintf("%v\n", a)...)
}

// SaveLogToStorage use to make||flush to achaemenid.log if needed!
func SaveLogToStorage(saveLocation string) {
	// TODO::: Check if achaemenid.log not exist create it || Flush old logs!
	err := ioutil.WriteFile(path.Join(saveLocation, "achaemenid.log"), log, 0755)
	if err != nil {
		panic(fmt.Sprintf("Unable to write achaemenid.log: %v", err))
	}
}
