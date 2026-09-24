/* For license and copyright information please see the LEGAL file in the code repository */

package buffer

import (
	"io"

	container_p "memar/computer/adt/container/protocol"
	buffer_p "memar/computer/buffer/protocol"

	// compress_errs "memar/codec/compress/errors"
	compress_p "memar/codec/compress/protocol"
	mediatype_p "memar/identifier/mediatype/protocol"
	error_p "memar/process/error/protocol"
)

type ComDecom struct {
	mt  mediatype_p.MediaType
	ct  compress_p.CompressType
	Buf []byte
	// reader  protocol.Reader
	readLen int
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (r *ComDecom) Init(mt mediatype_p.MediaType, ct compress_p.CompressType) (err error_p.Error) {
	r.mt = mt
	r.ct = ct
	return
}
func (r *ComDecom) Reinit(mt mediatype_p.MediaType, ct compress_p.CompressType) (err error_p.Error) {
	r.mt = mt
	r.ct = ct
	return
}
func (r *ComDecom) Deinit() (err error_p.Error) {
	return
}

//memar:impl memar/codec/protocol.Codec
func (r *ComDecom) SerializationLength() container_p.NumberOfElement {
	return container_p.NumberOfElement(len(r.Buf))
}
func (r *ComDecom) Decode(source buffer_p.Buffer) (err error_p.Error) {
	// TODO:::
	return
}
func (r *ComDecom) Encode(destination buffer_p.Buffer) (err error_p.Error) {
	// TODO:::
	return
}
func (r *ComDecom) Marshal(destination []byte) (n container_p.NumberOfElement, err error_p.Error) {
	// TODO::: check enough cap on destination or return error.

	_ = append(destination, r.Buf...)
	n = container_p.NumberOfElement(len(r.Buf))
	return
}
func (r *ComDecom) Unmarshal(source []byte) (n container_p.NumberOfElement, err error_p.Error) {
	// err = &compress_errs.SourceNotChangeable
	return
}

/*
********** buffer_p.Buffer interface **********
 */

//memar:impl memar/buffer_p.Buffer
func (r *ComDecom) ReadFrom(reader io.Reader) (n int64, err error) {
	// err = &compress_errs.SourceNotChangeable
	return
}
func (r *ComDecom) WriteTo(w io.Writer) (totalWrite int64, err error) {
	var writeLen int
	writeLen, err = w.Write(r.Buf)
	totalWrite = int64(writeLen)
	return
}
