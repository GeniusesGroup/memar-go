/* For license and copyright information please see LEGAL file in repository */

package protocol

// ObjectDirectory is the interface that must implement by any Application!
type ObjectDirectory interface {
	RegisterObjectStructure(StructureID uint64)
	Object(uuid [32]byte, structureID uint64) (object Object, err Error) // make if not exist before
	SaveObject(object Object) (err Error)                                // Also can rewrite an object without get it first!

	Delete(uuid [32]byte, structureID uint64) (err Error) // make invisible just by remove from primary index
	Wipe(uuid [32]byte, structureID uint64) (err Error)   // make invisible by remove from primary index & write random data to object location
}

// Object is the descriptor interface that must implement by any to be an object!
// Object owner is one app so it must handle concurrent protection internally not by object it self!
type Object interface {
	MetaData() ObjectMetaData
	Codec
}

// ObjectMetaData is the interface that must implement by any object!
type ObjectMetaData interface {
	ID() [32]byte
	StructureID() uint64
	Size() uint64
	WriteTime() Time // codec size equal to int64
	OwnerAppID() [32]byte

	// Due to we think objects as immutable data we don't support below metadata as standard! implement by any object such as files if you require them!
	// CreationTime() Time
	// LastAccessTime() Time
	// LastWriteTime() Time
}
