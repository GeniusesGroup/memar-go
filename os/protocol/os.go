/* For license and copyright information please see the LEGAL file in the code repository */

package os_p

// OperatingSystem is the interface that must implement by any OS object
type OperatingSystem interface {
	Screens() []GUIScreen

	OperatingSystem_User
	OperatingSystem_Storage
	net_p.PacketTarget
}
