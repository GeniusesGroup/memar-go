/* For license and copyright information please see LEGAL file in repository */

package giti

// Object is the descriptor interface that must implement by any to be an object!
// Object owner is one app so it must handle concurrent protection internally not by object it self!
type Object interface {
	ObjectMetaData
	Buffer
}

// ObjectMetaData is the interface that must implement by any object!
type ObjectMetaData interface {
	ObjectID() [32]byte
	ObjectStructureID() uint64
	ObjectSize() uint64
	ObjectWriteTime() Time
	ObjectOwnerAppID() [32]byte

	// Due to we think objects as immutable data we don't support below metadata as standard! implement by any object such as files if you require them!
	// CreationTime() Time
	// LastAccessTime() Time
	// LastWriteTime() Time
}
