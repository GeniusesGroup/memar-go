/* For license and copyright information please see the LEGAL file in the code repository */

package json

import (
	"libgo/buffer"
	"libgo/protocol"
)

func NewCodec(json protocol.JSON) (codec Codec) {
	codec.Init(json)
	return
}

// Codec is a wrapper to use anywhere need protocol.Codec interface instead of protocol.JSON interface
type Codec struct {
	json    protocol.JSON
	payload []byte
	len     int
}

//libgo:impl libgo/protocol.ObjectLifeCycle
func (c *Codec) Init(json protocol.JSON) {
	c.json = json
	c.len = json.LenAsJSON()
}
func (c *Codec) Reinit() {}
func (c *Codec) Deinit() {}

// https://www.iana.org/assignments/media-types/application/json
//
//libgo:impl libgo/protocol.Codec
func (c *Codec) MediaType() protocol.MediaType       { return &MediaType }
func (c *Codec) CompressType() protocol.CompressType { return nil }

//libgo:impl libgo/protocol.Decoder
func (c *Codec) Decode(source protocol.Codec) (n int, err protocol.Error) {
	c.payload, err = source.Marshal()
	if err != nil {
		return
	}
	_, err = c.json.FromJSON(c.payload)
	return
}

//libgo:impl libgo/protocol.Encoder
func (c *Codec) Encode(destination protocol.Codec) (n int, err protocol.Error) {
	n, err = destination.Decode(c)
	return
}
func (c *Codec) Len() int { return c.len }

//libgo:impl libgo/protocol.Unmarshaler
func (c *Codec) Unmarshal(data []byte) (n int, err protocol.Error) {
	_, err = c.json.FromJSON(data)
	c.payload = data
	return
}
func (c *Codec) UnmarshalFrom(data []byte) (remaining []byte, err protocol.Error) {
	// TODO:::
	return
}

func (c *Codec) Marshal() (data []byte, err protocol.Error) {
	if c.payload == nil {
		c.payload = make([]byte, 0, c.len)
		c.payload, err = c.json.ToJSON(c.payload)
	}
	return c.payload, nil
}

func (c *Codec) MarshalTo(data []byte) (added []byte, err protocol.Error) {
	return c.json.ToJSON(data)
}

//libgo:impl libgo/protocol.Buffer
func (c *Codec) ReadFrom(reader protocol.Reader) (n int, err protocol.Error) {
	var buf buffer.Flat
	n, err = buf.ReadFrom(reader)
	if err != nil {
		return
	}
	c.payload = buf.Bytes()
	_, err = c.json.FromJSON(c.payload)
	return
}

func (c *Codec) WriteTo(writer protocol.Writer) (n int, err protocol.Error) {
	if c.payload == nil {
		c.Marshal()
	}
	n, err = writer.Write(c.payload)
	return
}
