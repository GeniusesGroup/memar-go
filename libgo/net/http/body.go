/* For license and copyright information please see the LEGAL file in the code repository */

package http

import (
	container_p "memar/computer/adt/container/protocol"
	buffer_p "memar/computer/buffer/protocol"
	http_p "memar/libgo/net/http/protocol"
	error_p "memar/process/error/protocol"
)

// body is represent HTTP body.
// Due to many performance impact, MediaType() method of body not return any true data.
// use header ContentType() method instead. This can be change if ...
// https://datatracker.ietf.org/doc/html/rfc2616#section-4.3
type body struct {
	body http_p.Body
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (b *body) Init() (err error_p.Error)   { return }
func (b *body) Reinit() (err error_p.Error) { b.body = nil; return }
func (b *body) Deinit() (err error_p.Error) { return }

func (b *body) Body() http_p.Body        { return b.body }
func (b *body) SetBody(body http_p.Body) { b.body = body }

// Below methods exists to prevent panic errors cause of nil interface access.
//
//memar:impl memar/codec/protocol.Codec
func (b *body) SerializationLength() container_p.NumberOfElement {
	if b.body != nil {
		return b.body.SerializationLength()
	}
	return 0
}
func (b *body) Decode(source buffer_p.Buffer) (err error_p.Error) {
	if b.body != nil {
		err = b.body.Decode(source)
	}
	return
}
func (b *body) Encode(destination buffer_p.Buffer) (err error_p.Error) {
	if b.body != nil {
		err = b.body.Encode(destination)
	}
	return
}
