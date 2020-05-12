/* For license and copyright information please see LEGAL file in repository */

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

// codeGenerationLog hold logs until end of build proccess.
// Use BuildLog when you don't want to stop build process but need to store related log.
var codeGenerationLog []byte // code-generate.log

// buildLog will show log in console & append log to buffer to save them later.
// It is just to use for build proccess not running apps!!
func buildLog(a ...interface{}) {
	// print error to stderr
	fmt.Fprintf(os.Stderr, "%v\n", a)
	// Append log to CodeGenerateLog for saving to code-generate.log later.
	codeGenerationLog = append(codeGenerationLog, fmt.Sprintf("%v\n", a)...)
}

// saveLog use to make||flush to code-generate.log
func saveLog() {
	// Check if code-generate.log not exist create it || Flush old logs!
	err := ioutil.WriteFile(path.Join(ServiceRootLocation, "code-generate.log"), codeGenerationLog, 0755)
	if err != nil {
		panic(fmt.Sprintf("Unable to write code-generate.log: %v", err))
	}
}
