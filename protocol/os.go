/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// OS is default global protocol.OperatingSystem
// You must assign to it by any object implement protocol.OperatingSystem on your main.go file
// Strongly suggest assign easily by `import _ "libgo/os"`
var OS OperatingSystem

// OperatingSystem is the interface that must implement by any OS object
type OperatingSystem interface {
	AppManifest() ApplicationManifest

	RegisterNetworkNetwork_Multiplexer(nMux NetworkNetwork_Multiplexer)
	RegisterNetworkTransportMultiplexer(tMux NetworkTransport_Multiplexer)

	Screens() []GUIScreen

	OperatingSystem_User
	OperatingSystem_Storage
	MediaTypes
	CompressTypes
}
