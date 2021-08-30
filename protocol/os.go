/* For license and copyright information please see LEGAL file in repository */

package protocol

// Application is the interface that must implement by any OS!
type OS interface {
	Manifest() ApplicationManifest

	RegisterNetworkTransportMultiplexer(tMux NetworkTransportAppMultiplexer)

	ObjectDirectory() ObjectDirectory // Local object storage
	FileDirectory() FileDirectory     // Local file storage
}
