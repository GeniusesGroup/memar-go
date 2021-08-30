/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	"crypto/sha512"
	"time"

	etime "../earth-time"
	"../protocol"
)

// Manifest store server manifest data
// All string slice is multi language and in order by ManifestLanguages order
type Manifest struct {
	Society     string
	SocietyUUID [32]byte
	SocietyID   protocol.SocietyID
	DomainName  string
	AppID       [32]byte // Hash of domain act as Application ID too
	Email       string
	Icon        string
	AppDeatil   map[protocol.LanguageID]AppDetail

	RequestedPermission []uint32 // ServiceIDs from PersiaOS services e.g. InternetInBackground, Notification, ...
	TechnicalInfo       TechnicalInfo
	NetworkInfo         NetworkInfo
	DeployInfo          DeployInfo
}

func (ma *Manifest) init() {
	ma.AppID = sha512.Sum512_256(convert.UnsafeStringToByteSlice(ma.DomainName))
}

type AppDetail struct {
	Organization   string
	Name           string
	Description    string
	TermsOfService string
	Licence        string
	TAGS           []string // Use to categorized apps e.g. Music, GPS, ...
}

// TechnicalInfo store some technical information but may different from really server condition!
type TechnicalInfo struct {
	// Shutdown settings
	ShutdownDelay time.Duration // the server will wait for at least this amount of time for active streams to finish!

	// Minimum hardware specification for each instance of application.
	CPUCores uint8  // Number
	CPUSpeed uint64 // Hz
	RAM      uint64 // Byte
	GPU      uint64 // Hz
	Network  uint64 // Byte per second
	Storage  uint64 // Byte, HHD||SSD||... indicate by DataCentersClassForDataStore
}

// NetworkInfo store some network information.
type NetworkInfo struct {
	// Application Overal rate limit
	MaxOpenConnection     uint64         // The maximum number of concurrent connections the app may serve.
	ConnectionIdleTimeout etime.Duration // In seconds
	// MaxStreamHeaderSize   uint64 // For stream protocols with variable header size like HTTP
	// MaxStreamPayloadSize  uint64 // For stream protocols with variable payload size like sRPC, HTTP, ...

	// Guest rete limit - Connection.OwnerType==0
	GuestMaxConnections            uint64 // 0 means App not accept guest connection.
	GuestMaxUserConnectionsPerAddr uint64
	GuestMaxConcurrentStreams      uint32
	GuestMaxStreamConnectionDaily  uint32 // Max open stream per day for a guest connection. overflow will drop on creation!
	GuestMaxServiceCallDaily       uint64 // 0 means no limit and good for PayAsGo strategy!
	GuestMaxBytesSendDaily         uint64
	GuestMaxBytesReceiveDaily      uint64
	GuestMaxPacketsSendDaily       uint64
	GuestMaxPacketsReceiveDaily    uint64

	// Registered rate limit - Connection.OwnerType==1
	RegisteredMaxConnections            uint64
	RegisteredMaxUserConnectionsPerAddr uint64
	RegisteredMaxConcurrentStreams      uint32
	RegisteredMaxStreamConnectionDaily  uint32 // Max open stream per day for a Registered user connection. overflow will drop on creation!
	RegisteredMaxServiceCallDaily       uint64 // 0 means no limit and good for PayAsGo strategy!
	RegisteredMaxBytesSendDaily         uint64
	RegisteredMaxBytesReceiveDaily      uint64
	RegisteredMaxPacketsSendDaily       uint64
	RegisteredMaxPacketsReceiveDaily    uint64

	// If you want to know Connection.OwnerType>1 rate limit strategy, You must read server codes!!
}

// DeployInfo store some application deployment information.
type DeployInfo struct {
	// Distribution
	DistributeOutOfSociety bool          // Allow to run service-only instance of app out of original society belong to.
	DataCentersClass       uint8         // 0:FirstClass 256:Low-Quality default:5
	MaxNodeNumber          uint32        // default:3
	NodeFailureTimeOut     time.Duration // Max suggestion is 6 hour, other service only node replace failed node! not use in network failure, it is handy proccess!
}
