/* For license and copyright information please see the LEGAL file in the code repository */

package achaemenid

import (
	"golang.org/x/crypto/sha3"

	"memar/convert"
	"memar/protocol"
)

// Manifest store server manifest data
// All string slice is multi language and in order by ManifestLanguages order
type Manifest struct {
	Society     string
	SocietyUUID [32]byte
	SocietyID   protocol.SocietyID
	domainName  string
	appUUID     protocol.UUID // Hash of domain act as Application ID too
	appID       uint16        // Hash of domain act as Application ID too
	email       string
	icon        []byte
	AppDetail   map[protocol.LanguageID]AppDetail

	RequestedPermission []protocol.ServiceID // e.g. InternetInBackground, Notification, ...
}

//memar:impl memar/protocol.ObjectLifeCycle
func (ma *Manifest) init() {
	var sha3 = sha3.Sum256(convert.UnsafeStringToByteSlice(ma.domainName))
	copy(ma.appUUID[:], sha3[:])
}

type AppDetail struct {
	Organization   string
	Name           string
	Description    string
	TermsOfService string
	License        string
	TAGS           []string // Use to categorized apps e.g. Music, GPS, ...
}

//memar:impl memar/protocol.ApplicationManifest
func (ma *Manifest) Icon() []byte             { return ma.icon }
func (ma *Manifest) DomainName() string       { return ma.domainName }
func (ma *Manifest) Email() string            { return ma.email }
func (ma *Manifest) AppUUID() protocol.UUID   { return ma.appUUID }
func (ma *Manifest) AppID() (appID uint16)    { return ma.appID }
func (ma *Manifest) ContentPreferences()      {}
func (ma *Manifest) PresentationPreferences() {}
