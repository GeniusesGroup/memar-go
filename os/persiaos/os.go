/* For license and copyright information please see LEGAL file in repository */

package persiaos

import (
	"../../gp"
	"../../mediatype"
	"../../protocol"
)

// OS implements protocol.OperatingSystem
var OS os

type os struct {
	Manifest     Manifest
	State        protocol.ApplicationState
	StateChannel chan protocol.ApplicationState

	// localFileDirectory        FileDirectory   // Local file storage
	// localObjectDirectory      ObjectDirectory // Local object storage
	// localCacheObjectDirectory ObjectDirectory // Local object storage

	gp gp.OSMultiplexer
	ipv4 ip.OSMultiplexerIPv4

	mediatype.MediaTypes
}

func (os *os) AppManifest() protocol.ApplicationManifest { return &os.Manifest }

// func (os *os) ObjectDirectory() protocol.ObjectDirectory      { return &os.localObjectDirectory }
// func (os *os) CacheObjectDirectory() protocol.ObjectDirectory { return &os.localCacheObjectDirectory }
// func (os *os) FileDirectory() protocol.FileDirectory          { return &os.localFileDirectory }

// Init method use to auto initialize App object with default data.
func init() {
	protocol.App.LogInfo("Application Run on PersiaOS")

	OS.gp.Init()
}

// Shutdown use to graceful stop os
func (os *os) Shutdown() {
	// os.localFileDirectory.Save()
}
