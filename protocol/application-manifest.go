/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// ApplicationManifest introduce all data about an applications
type ApplicationManifest interface {
	Icon() []byte
	DomainName() string
	Email() string

	AppUUID() UUID
	AppID() (appID uint16) // local OS application ID

	ContentPreferences()
	PresentationPreferences()
}
