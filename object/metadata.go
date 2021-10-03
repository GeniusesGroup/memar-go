/* For license and copyright information please see LEGAL file in repository */

package object

import (
	etime "../earth-time"
	"../protocol"
	"../syllab"
)

const MetadataLength uint32 = 64

// Metadata is the header structure of an object!
type Metadata []byte

// methods to implements protocol.ObjectMetadata interface
func (md Metadata) ID() (id [32]byte)        { copy(id[:], md[0:]); return }
func (md Metadata) WriteTime() protocol.Time { return etime.Time(syllab.GetInt64(md, 32)) }
func (md Metadata) MediaTypeID() uint64      { return syllab.GetUInt64(md, 40) }
func (md Metadata) CompressTypeID() uint64   { return syllab.GetUInt64(md, 48) }
func (md Metadata) DataLength() uint64       { return syllab.GetUInt64(md, 56) }

func (md Metadata) setID(id [32]byte)                       { copy(md[0:], id[:]) }
func (md Metadata) setWriteTime(writeTime etime.Time)       { syllab.SetInt64(md, 32, int64(writeTime)) }
func (md Metadata) setMediaTypeID(mediaTypeID uint64)       { syllab.SetUInt64(md, 40, mediaTypeID) }
func (md Metadata) setCompressTypeID(compressTypeID uint64) { syllab.SetUInt64(md, 48, compressTypeID) }
func (md Metadata) setDataLength(dataLength int)            { syllab.SetUInt64(md, 56, uint64(dataLength)) }

// methods to implements protocol.Syllab interface
func (md Metadata) CheckSyllab(payload []byte) (err protocol.Error) {
	if len(payload) < int(md.LenOfSyllabStack()) {
		err = syllab.ErrShortArrayDecode
	}
	return
}
func (md Metadata) FromSyllab(payload []byte, stackIndex uint32) {
	// err = ErrSourceNotChangeable
}
func (md Metadata) ToSyllab(payload []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32) {
	freeHeapIndex = syllab.SetByteArray(payload, md, stackIndex, heapIndex)
	return
}
func (md Metadata) LenAsSyllab() uint64          { return uint64(md.LenOfSyllabStack() + md.LenOfSyllabHeap()) }
func (md Metadata) LenOfSyllabStack() uint32     { return 8 }
func (md Metadata) LenOfSyllabHeap() (ln uint32) { ln = uint32(len(md)); return }

/*
********** metadata structure **********
 */

type metadata struct {
	// Describe the object
	id        [32]byte
	writeTime etime.Time

	// Describe the object data
	mediaTypeID    uint64
	compressTypeID uint64
	dataLength     uint64
}

// methods to implements protocol.ObjectMetadata interface
func (md *metadata) ID() [32]byte             { return md.id }
func (md *metadata) WriteTime() protocol.Time { return md.writeTime }
func (md *metadata) MediaTypeID() uint64      { return md.mediaTypeID }
func (md *metadata) CompressTypeID() uint64   { return md.compressTypeID }
func (md *metadata) DataLength() uint64       { return md.dataLength }

// methods to implements protocol.Syllab interface
func (md *metadata) CheckSyllab(payload []byte) (err protocol.Error) {
	if len(payload) < int(md.LenOfSyllabStack()) {
		err = syllab.ErrShortArrayDecode
	}
	return
}
func (md *metadata) FromSyllab(buf []byte, stackIndex uint32) {
	copy(md.id[:], buf[0:])
	md.writeTime = etime.Time(syllab.GetInt64(buf, 32))

	md.mediaTypeID = syllab.GetUInt64(buf, 40)
	md.compressTypeID = syllab.GetUInt64(buf, 48)
	md.dataLength = syllab.GetUInt64(buf, 56)
}
func (md *metadata) ToSyllab(payload []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32) {
	copy(payload[0:], md.id[:])
	syllab.SetInt64(payload, 32, int64(md.writeTime))

	syllab.SetUInt64(payload, 40, md.mediaTypeID)
	syllab.SetUInt64(payload, 48, md.compressTypeID)
	syllab.SetUInt64(payload, 56, md.dataLength)
	return heapIndex
}
func (md *metadata) LenOfSyllabStack() uint32     { return MetadataLength }
func (md *metadata) LenOfSyllabHeap() (ln uint32) { return }
func (md *metadata) LenAsSyllab() uint64          { return uint64(MetadataLength) }
