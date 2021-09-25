/* For license and copyright information please see LEGAL file in repository */

package dos

import (
	goos "os"
	"runtime"

	"../../file"
	"../../log"
	"../../protocol"
)

// OS implements protocol.OperatingSystem
var OS os

type os struct {
	AppID        [32]byte // Hash of domain act as Application ID too
	State        protocol.ApplicationState
	StateChannel chan protocol.ApplicationState

	appFileURI           file.URI // Just fill in old type OS like linux, windows, ...
	localFileDirectory   FileDirectory
	localObjectDirectory protocol.ObjectDirectory

	netTransMux
}

// Init method use to auto initialize App object with default data.
func init() {
	log.Info("Application Run on ", runtime.GOOS, " OS")

	// Indicate repoLocation
	var ex, err = goos.Executable()
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Application binary file location is", ex)
	OS.appFileURI.Init(ex)

	goos.Chdir(ex)
	OS.localFileDirectory.init("/") // TODO::: add more data like domain

	// Register data-structures
	// persiaos.StorageRegisterStructure(d.ID)
}

// Shutdown use to graceful stop os
func (os *os) Shutdown() {
	// os.localFileDirectory.Save()
}
