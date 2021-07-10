/* For license and copyright information please see LEGAL file in repository */

package object

import (
	etime "../earth-time"
	"../giti"
	"../syllab"
)

const MetaDataSize uint32 = 88

// MetaData store header structure of an object!
type MetaData struct {
	ID          [32]byte
	StructureID uint64
	Size        uint64
	WriteTime   etime.Time
	OwnerAppID  [32]byte
}

func (md *MetaData) ObjectID() [32]byte         { return md.ID }
func (md *MetaData) ObjectStructureID() uint64  { return md.StructureID }
func (md *MetaData) ObjectSize() uint64         { return md.Size }
func (md *MetaData) ObjectWriteTime() giti.Time { return md.WriteTime }
func (md *MetaData) ObjectOwnerAppID() [32]byte { return md.OwnerAppID }

func (md *MetaData) CheckSyllab(payload []byte) (err giti.Error) {
	if len(payload) < int(md.LenOfSyllabStack()) {
		err = syllab.ErrShortArrayDecode
	}
	return
}
func (md *MetaData) FromSyllab(buf []byte, stackIndex uint32) {
	copy(md.ID[:], buf[0:])
	md.StructureID = syllab.GetUInt64(buf, 32)
	md.Size = syllab.GetUInt64(buf, 40)
	md.WriteTime = etime.Time(syllab.GetInt64(buf, 48))
	copy(md.OwnerAppID[:], buf[56:])
}
func (md *MetaData) ToSyllab(payload []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32) {
	copy(payload[0:], md.ID[:])
	syllab.SetUInt64(payload, 32, md.StructureID)
	syllab.SetUInt64(payload, 40, md.Size)
	syllab.SetInt64(payload, 48, int64(md.WriteTime))
	copy(payload[56:], md.OwnerAppID[:])
	return heapIndex
}

func (md *MetaData) LenOfSyllabStack() uint32     { return MetaDataSize }
func (md *MetaData) LenOfSyllabHeap() (ln uint32) { return }
func (md *MetaData) LenAsSyllab() uint64          { return uint64(md.LenOfSyllabStack() + md.LenOfSyllabHeap()) }
