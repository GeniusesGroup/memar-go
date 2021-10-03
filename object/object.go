/* For license and copyright information please see LEGAL file in repository */

package object

import (
	"io"

	"../compress"
	"../mediatype"
	"../protocol"
	"../syllab"
)

// Object has needed methods to implements protocol.Object interface!
type Object []byte

func (ob Object) Metadata() protocol.ObjectMetadata { return Metadata(ob) }
func (ob Object) Data() protocol.Codec              { return ob[MetadataLength:] }

func NewObject(data protocol.Codec) Object { return make([]byte, 0, int(MetadataLength)+data.Len()) }

/*
********** protocol.Codec interface **********
 */

func (ob Object) MediaType() protocol.MediaType {
	return mediatype.MediaTypeByID(ob.Metadata().MediaTypeID())
}
func (ob Object) CompressType() protocol.CompressType {
	return compress.CompressTypeByID(ob.Metadata().CompressTypeID())
}

func (ob Object) Decode(reader io.Reader) (err protocol.Error) { err = ErrSourceNotChangeable; return }
func (ob Object) Encode(writer io.Writer) (err error)          { _, err = ob.WriteTo(writer); return }
func (ob Object) Unmarshal(data []byte) (err protocol.Error)   { err = ErrSourceNotChangeable; return }
func (ob Object) Marshal() (data []byte)                       { return ob }
func (ob Object) MarshalTo(data []byte) []byte                 { return append(data, ob...) }
func (ob Object) Len() int                                     { return len(ob) }

/*
********** io package interfaces **********
 */

func (ob Object) ReadFrom(reader io.Reader) (n int64, err error) {
	err = ErrSourceNotChangeable
	return
}
func (ob Object) WriteTo(writer io.Writer) (n int64, err error) {
	var writeLength int
	writeLength, err = writer.Write(ob)
	n = int64(writeLength)
	return
}

/*
********** protocol.Syllab interface **********
 */

func (ob Object) CheckSyllab(payload []byte) (err protocol.Error) {
	if len(payload) < int(ob.LenOfSyllabStack()) {
		err = syllab.ErrShortArrayDecode
	}
	return
}
func (ob Object) FromSyllab(payload []byte, stackIndex uint32) {
	// err = ErrSourceNotChangeable
}
func (ob Object) ToSyllab(payload []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32) {
	freeHeapIndex = syllab.SetByteArray(payload, ob, stackIndex, heapIndex)
	return
}
func (ob Object) LenAsSyllab() uint64          { return uint64(ob.LenOfSyllabStack() + ob.LenOfSyllabHeap()) }
func (ob Object) LenOfSyllabStack() uint32     { return 8 }
func (ob Object) LenOfSyllabHeap() (ln uint32) { ln = uint32(len(ob)); return }
