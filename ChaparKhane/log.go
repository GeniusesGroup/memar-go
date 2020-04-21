/* For license and copyright information please see LEGAL file in repository */

package chaparkhane

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
	// Append log to CodeGenerateLog for saving to ChaparKhane.log later.
	log = append(log, fmt.Sprintf("%v\n", a)...)
}

// SaveLogToStorage use to make||flush to ChaparKhane.log if needed!
func SaveLogToStorage(saveLocation string) {
	// TODO::: Check if ChaparKhane.log not exist create it || Flush old logs!
	err := ioutil.WriteFile(path.Join(saveLocation, "ChaparKhane.log"), log, 0755)
	if err != nil {
		panic(fmt.Sprintf("Unable to write ChaparKhane.log: %v", err))
	}
}
