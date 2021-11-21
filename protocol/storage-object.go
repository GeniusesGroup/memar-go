/* For license and copyright information please see LEGAL file in repository */

package protocol

// ObjectDirectory is the interface that show how an object directory work.
// Object owner is one app so it must handle concurrent protection internally not by object it self!
// 32 byte uuid choose to defeat https://en.wikipedia.org/wiki/Birthday_problem due to we suggest use hash of data (sha3-256) as object key
type ObjectDirectory interface {
	Length(uuid [32]byte, structureID uint64) (ln int, err Error)
	Get(uuid [32]byte, structureID uint64) (object Codec, err Error)                      // make if not exist before
	Read(uuid [32]byte, structureID uint64, offset, limit uint64) (data Codec, err Error) // act like block storage API
	Save(data Codec) (uuid [32]byte, err Error)                                           //
	SaveRaw(object []byte) (err Error)                                                    // just use to store distributed object that store by other app node
	Update(uuid [32]byte, data Codec) (err Error)                                         // use to update or save by specific ID
	Write(uuid [32]byte, structureID uint64, offset uint64, data Codec) (err Error)       // Rewrite an object without get it first!
	Delete(uuid [32]byte, structureID uint64) (err Error)                                 // make invisible just by remove from primary index
	Wipe(uuid [32]byte, structureID uint64) (err Error)                                   // make invisible by remove from primary index & write random data to object location
}

// Objects or records or rows (RDBMS) all are same just difference in where schema must handle.
// Strongly think fundementally due to all data in computer must store as k/v in storages finally even files or rows.
type TypicalObject interface {
	WriteTime() Time
	Codec
}
