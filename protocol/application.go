/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Application immutable runtime settings
const (
	// AppLanguage store global language to use by any locale text selector.
	AppLanguage = LanguageEnglish

	// AppMode_Dev use to indicate that app can do some more logic e.g.
	// - Save more logs
	// - Add more services like net/http/pprof for better debugging
	// - Add more pages that just need only for developers
	AppMode_Dev = true
)

// Application is the interface that must implement by any Application.
// It introduce just local computing mechanism not network, storage, distributed, gui, ...
type Application interface {
	Engine() ApplicationEngine
	Manifest() ApplicationManifest

	SoftwareStatus() SoftwareStatus

	Application_Status

	EventTarget

	OS_Signal_Listener
	Network_PacketListener

	ObjectLifeCycle
}
