/* For license and copyright information please see LEGAL file in repository */

package protocol

// ObjectDirectory is the interface that show how an object directory work.
type ObjectDirectory interface {
	Metadata(uuid [32]byte, structureID uint64) (metadata ObjectMetadata, err Error)
	Get(uuid [32]byte, structureID uint64) (object Object, err Error)                      // make if not exist before
	Read(uuid [32]byte, structureID uint64, offset, limit uint64) (data []byte, err Error) // act like block storage API
	Save(data Codec) (metadata ObjectMetadata, err Error)                                  //
	SaveRaw(object []byte) (err Error)                                                     // just use to store distributed object that store by other app node
	Write(uuid [32]byte, structureID uint64, offset uint64, data []byte) (err Error)       // Rewrite an object without get it first!
	Delete(uuid [32]byte, structureID uint64) (err Error)                                  // make invisible just by remove from primary index
	Wipe(uuid [32]byte, structureID uint64) (err Error)                                    // make invisible by remove from primary index & write random data to object location
}

// Object is the descriptor interface that must implement by any to be an object!
// Object owner is one app so it must handle concurrent protection internally not by object it self!
type Object interface {
	Metadata() ObjectMetadata
	Data() Codec
	Codec
}

// ObjectMetadata is the interface that must implement by any object!
type ObjectMetadata interface {
	ID() [32]byte
	WriteTime() Time        // codec as int64
	MediaTypeID() uint64    // GitiURN.ID()
	CompressTypeID() uint64 // GitiURN.ID()
	DataLength() uint64     // just data part

	// Due to we think objects as immutable data we don't support below metadata as standard! implement by any object such as files if you require them!
	// CreationTime() Time
	// LastAccessTime() Time
	// LastWriteTime() Time
}
