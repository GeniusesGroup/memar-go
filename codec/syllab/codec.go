/* For license and copyright information please see the LEGAL file in the code repository */

package syllab

import (
	container_p "memar/computer/adt/container/protocol"
	buffer_p "memar/computer/buffer/protocol"
	syllab_p "memar/codec/syllab/protocol"
	mediatype_p "memar/identifier/mediatype/protocol"
	error_p "memar/process/error/protocol"
)

// Codec is a wrapper to use anywhere need codec_p.Codec interface instead of protocol.Syllab interface
type Codec struct {
	syllab syllab_p.Codec
}

func NewCodec(syllab syllab_p.Codec) (codec Codec) {
	codec.syllab = syllab
	return
}

//memar:impl memar/codec/protocol.Codec
func (c *Codec) MediaType() mediatype_p.MediaType { return MediaType.MediaType() }
func (c *Codec) Decode(source buffer_p.Buffer) (err error_p.Error) {
	err = c.syllab.CheckSyllab(source)
	if err != nil {
		return
	}
	c.syllab.FromSyllab(source, 0)
	return
}
func (c *Codec) Encode(destination buffer_p.Buffer) (err error_p.Error) {
	c.syllab.ToSyllab(destination, 0, 0)
	return
}

//memar:impl memar/codec/protocol.Field_Length
func (c *Codec) SerializationLength() container_p.NumberOfElement { return c.syllab.Syllab_Length() }
