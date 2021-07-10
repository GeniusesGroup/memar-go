/* For license and copyright information please see LEGAL file in repository */

package object

import (
	etime "../earth-time"
	"../giti"
	"../syllab"
)

// MetaDataAccess store header structure of an object!
type MetaDataAccess []byte

func (md MetaDataAccess) CheckSyllab(payload []byte) (err giti.Error) {
	if len(md) < int(md.LenOfSyllabStack()) {
		err = syllab.ErrShortArrayDecode
	}
	return
}
func (md MetaDataAccess) LenOfSyllabStack() uint32 { return 88 }

func (md MetaDataAccess) ObjectID() (id [32]byte)         { copy(id[:], md[0:]); return }
func (md MetaDataAccess) ObjectStructureID() uint64       { return syllab.GetUInt64(md, 32) }
func (md MetaDataAccess) ObjectSize() uint64              { return syllab.GetUInt64(md, 40) }
func (md MetaDataAccess) ObjectWriteTime() giti.Time      { return etime.Time(syllab.GetInt64(md, 48)) }
func (md MetaDataAccess) ObjectOwnerAppID() (id [32]byte) { copy(id[:], md[56:]); return }
