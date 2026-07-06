/* For license and copyright information please see the LEGAL file in the code repository */

package http

import (
	"memar/protocol"
)

// body is represent HTTP body.
// Due to many performance impact, MediaType() method of body not return any true data.
// use header ContentType() method instead. This can be change if ...
// https://datatracker.ietf.org/doc/html/rfc2616#section-4.3
type body struct {
	protocol.Codec
}

//memar:impl memar/protocol.ObjectLifeCycle
func (b *body) Init() (err protocol.Error)   { return }
func (b *body) Reinit() (err protocol.Error) { b.Codec = nil; return }
func (b *body) Deinit() (err protocol.Error) { return }

func (b *body) Body() protocol.Codec         { return b }
func (b *body) SetBody(codec protocol.Codec) { b.Codec = codec }

// Below methods exists to prevent panic errors cause of nil interface access.
//
//memar:impl memar/protocol.Codec
func (b *body) Len() int {
	if b.Codec != nil {
		return b.Codec.Len()
	}
	return 0
}
func (b *body) MediaType() protocol.MediaType {
	if b.Codec != nil {
		return b.Codec.MediaType()
	}
	return nil
}
func (b *body) CompressType() protocol.CompressType {
	if b.Codec != nil {
		return b.Codec.CompressType()
	}
	return nil
}
func (b *body) Decode(source protocol.Codec) (n int, err protocol.Error) {
	if b.Codec != nil {
		n, err = b.Codec.Decode(source)
	}
	return
}
func (b *body) Encode(destination protocol.Codec) (n int, err protocol.Error) {
	if b.Codec != nil {
		n, err = b.Codec.Encode(destination)
	}
	return
}
func (b *body) Marshal() (data []byte, err protocol.Error) {
	if b.Codec != nil {
		data, err = b.Codec.Marshal()
	}
	return
}
func (b *body) MarshalTo(data []byte) (added []byte, err protocol.Error) {
	if b.Codec != nil {
		return b.Codec.MarshalTo(data)
	}
	return data, nil
}
func (b *body) Unmarshal(data []byte) (n int, err protocol.Error) {
	if b.Codec != nil {
		n, err = b.Codec.Unmarshal(data)
	}
	return
}
func (b *body) UnmarshalFrom(data []byte) (remaining []byte, err protocol.Error) {
	if b.Codec != nil {
		return b.Codec.UnmarshalFrom(data)
	}
	return
}
