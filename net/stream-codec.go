/* For license and copyright information please see the LEGAL file in the code repository */

package net

import (
	"memar/protocol"
)

//memar:impl memar/protocol.Codec
func (sk *Socket) MediaType() protocol.MediaType       { return nil }
func (sk *Socket) CompressType() protocol.CompressType { return nil }
func (sk *Socket) Decode(source protocol.Codec) (n int, err protocol.Error) {
	return source.Encode(sk)
}
func (sk *Socket) Encode(destination protocol.Codec) (n int, err protocol.Error) {
	return destination.Decode(sk)
}
func (sk *Socket) Marshal() (data []byte, err protocol.Error) {
	err = sk.Check()
	if err != nil {
		return
	}

	if !sk.buf.buf.Full() {
		err = sk.blockInSelect()
	}
	// TODO::: check and wrap above error?
	return sk.buf.buf.Marshal()
}
func (sk *Socket) MarshalTo(data []byte) (added []byte, err protocol.Error) {
	err = sk.Check()
	if err != nil {
		return
	}

	if !sk.buf.buf.Full() {
		err = sk.blockInSelect()
	}
	// TODO::: check and wrap above error?
	return sk.buf.buf.MarshalTo(data)
}
func (sk *Socket) Unmarshal(data []byte) (n int, err protocol.Error) {
	for len(data) > 0 {
		err = sk.Check()
		if err != nil {
			return
		}

		var sendNumber int
		sendNumber, err = sk.sendPayload(data)
		if err != nil {
			return
		}
		n += sendNumber
		data = data[sendNumber:]
	}
	return
}
func (sk *Socket) UnmarshalFrom(data []byte) (remaining []byte, err protocol.Error) {
	return
}
func (sk *Socket) Len() (ln int) { return sk.buf.buf.Len() }
