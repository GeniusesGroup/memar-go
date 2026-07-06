/* For license and copyright information please see the LEGAL file in the code repository */

package json

import (
	container_p "memar/computer/adt/container/protocol"
	buffer_p "memar/computer/buffer/protocol"
	json_p "memar/codec/data_exchange/json/protocol"
	error_p "memar/process/error/protocol"
)

func NewCodec(json json_p.Codec) (codec Codec) {
	codec.Init(json)
	return
}

// Codec is a wrapper to use anywhere need codec_p.Codec interface instead of json_p.JSON interface
type Codec struct {
	json json_p.Codec
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (c *Codec) Init(json json_p.Codec) (err error_p.Error) {
	c.json = json
	return
}
func (c *Codec) Reinit(json json_p.Codec) (err error_p.Error) {
	c.json = json
	return
}
func (c *Codec) Deinit() (err error_p.Error) {
	return
}

// https://www.iana.org/assignments/media-types/application/json
//
//memar:impl memar/codec/protocol.Codec
// func (c *Codec) MediaType() mediatype_p.MediaType    { return &DT }
// func (c *Codec) CompressType() compress_p.CompressType { return nil }

//memar:impl memar/protocol.Decoder
func (c *Codec) Decode(source buffer_p.Buffer) (err error_p.Error) {
	err = c.json.FromJSON(source)
	return
}

//memar:impl memar/protocol.Encoder
func (c *Codec) Encode(destination buffer_p.Buffer) (err error_p.Error) {
	err = c.json.ToJSON(destination)
	return
}

//memar:impl memar/codec/protocol.Field_Length
func (c *Codec) SerializationLength() container_p.NumberOfElement { return c.json.JSON_Length() }
