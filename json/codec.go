/* For license and copyright information please see LEGAL file in repository */

package json

import (
	"bytes"
	"io"

	"../mediatype"
	"../protocol"
)

// Codec is a wrapper to use anywhere need protocol.Codec interface instead of protocol.JSON interface
type Codec struct {
	json    protocol.JSON
	payload []byte
	len     int
}

func NewCodec(json protocol.JSON) (codec Codec) {
	codec = Codec{
		json: json,
		len:  json.LenAsJSON(),
	}
	return
}

/*
********** protocol.Codec interface **********
 */

// https://www.iana.org/assignments/media-types/application/json
func (c *Codec) MediaType() protocol.MediaType       { return mediatype.JSON }
func (c *Codec) CompressType() protocol.CompressType { return nil }

func (c *Codec) Decode(reader protocol.Reader) (err protocol.Error) {
	var buf bytes.Buffer
	var _, goErr = buf.ReadFrom(reader)
	if goErr != nil {
		// err =
		return
	}
	c.payload = buf.Bytes()
	err = c.json.FromJSON(c.payload)
	return
}
func (c *Codec) Encode(writer protocol.Writer) (err protocol.Error) {
	var _, goErr = c.WriteTo(writer)
	if goErr != nil {
		// err =
	}
	return
}
func (c *Codec) Len() int { return c.len }

func (c *Codec) Unmarshal(data []byte) (err protocol.Error) {
	err = c.json.FromJSON(data)
	c.payload = data
	return
}

func (c *Codec) Marshal() (data []byte) {
	if c.payload == nil {
		c.payload = make([]byte, 0, c.len)
		c.payload = c.json.ToJSON(c.payload)
	}
	return c.payload
}

func (c *Codec) MarshalTo(data []byte) []byte {
	return c.json.ToJSON(data)
}

/*
********** io package interfaces **********
 */

func (c *Codec) ReadFrom(reader io.Reader) (n int64, err error) {
	var buf bytes.Buffer
	n, err = buf.ReadFrom(reader)
	if err != nil {
		return
	}
	c.payload = buf.Bytes()
	err = c.json.FromJSON(c.payload)
	return
}

func (c *Codec) WriteTo(writer io.Writer) (n int64, err error) {
	if c.payload == nil {
		c.Marshal()
	}
	var writeLength int
	writeLength, err = writer.Write(c.payload)
	n = int64(writeLength)
	return
}
