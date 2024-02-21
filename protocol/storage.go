/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

type NumberOfVersion uint64
type VersionOffset uint64

const (
	StorageRecord_NoVersion         VersionOffset = 0 // also as first version in Get logic
	StorageRecord_LastLocalVersion  VersionOffset = 18446744073709551613
	StorageRecord_LastEdgeVersion   VersionOffset = 18446744073709551614
	StorageRecord_LastSourceVersion VersionOffset = 18446744073709551615
)

// Some other framework also introduce this protocol too:
// https://docs.abp.io/en/abp/latest/Entities#versioning-entities
type Storage_Versioned interface {
	VersionNumber() VersionOffset
}

// MaxVersion == StorageRecord_NoVersion means this record don't need versioning or just one version.
// MaxVersion > 0 indicate max version. e.g. 6 means just 6 version must store for the record.
// MaxVersion == StorageRecord_LastSourceVersion indicate no version limit but logically it has limit up to uint64.
type Storage_MaxVersioned interface {
	MaxVersion() VersionOffset
}

type Storage_SaveTime interface {
	SaveTime() Time // Request time or save Time of the request not the created record by this record.
}

// TTL(Time-To-Live) or Expiration Number of nanoseconds until record expires.
type Storage_TTL interface {
	TTL() Duration
}
