/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import "time"

// Manifest use to store server manifest data
// All string slice is multi language and in order by ManifestLanguages order
type Manifest struct {
	SocietyID           uint32
	AppID               [16]byte // Application ID
	DomainID            [16]byte // Usually hash of domain
	DomainName          string
	Email               string
	Icon                string
	AuthorizedAppDomain [][16]byte // Just accept request from these domains, neither guest nor users!
	SupportedLanguages  []uint32
	ManifestLanguages   []uint32
	Organization        []string
	Name                []string
	Description         []string
	TermsOfService      []string
	Licence             []string
	TAGS                []string // Use to categorized apps e.g. Music, GPS, ...
	RequestedPermission []uint32 // ServiceIDs from PersiaOS services e.g. InternetInBackground, Notification, ...
	TechnicalInfo       TechnicalInfo
}

// TechnicalInfo store some technical information but may different from really server condition!
type TechnicalInfo struct {
	// Shutdown settings
	ShutdownDelay time.Duration // the server will wait for at least this amount of time for active streams to finish!

	// Server Overal rate limit
	MaxOpenConnection     uint64 // The maximum number of concurrent connections the app may serve.
	ConnectionIdleTimeout time.Duration
	MaxStreamHeaderSize   uint64 // For stream protocols with variable header size like HTTP
	MaxStreamPayloadSize  uint64 // For stream protocols with variable payload size like sRPC, HTTP, ...

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

	// Minimum hardware specification for each instance of application.
	CPUCores uint8  // Number
	CPUSpeed uint64 // Hz
	RAM      uint64 // Byte
	GPU      uint64 // Hz
	Network  uint64 // Byte per second
	Storage  uint64 // Byte, HHD||SSD||... indicate by DataCentersClassForDataStore

	// Distribution
	DistributeOutOfSociety       bool   // Allow to run service-only instance of app out of original society belong to.
	DataCentersClass             uint8  // 0:FirstClass 256:Low-Quality default:5
	DataCentersClassForDataStore uint8  // 0:FirstClass 256:Low-Quality default:0
	ReplicationNumber            uint8  // deafult:3
	MaxNodeNumber                uint32 // default:3

	// DataStore
	TransactionTimeOut uint16 // in ms, default:500ms, Max 65.535s timeout
	NodeFailureTimeOut uint16 // in minute, default:60m, other corresponding node same range will replace failed node! not use in network failure, it is handy proccess!
}
