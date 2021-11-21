/* For license and copyright information please see LEGAL file in repository */

package protocol

// OS is default global protocol.OperatingSystem
// You must assign to it by any object implement protocol.OperatingSystem on your main.go file
// Strongly suggest assign easily by `import _ "./libgo/os"`
var OS OperatingSystem

// OperatingSystem is the interface that must implement by any OS object
type OperatingSystem interface {
	AppManifest() ApplicationManifest

	RegisterNetworkTransportMultiplexer(tMux NetworkTransportAppMultiplexer)

	ObjectDirectory() ObjectDirectory      // Local object storage
	CacheObjectDirectory() ObjectDirectory // Local object storage
	FileDirectory() FileDirectory          // Local file storage

	MediaTypes
	CompressTypes
}
