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
