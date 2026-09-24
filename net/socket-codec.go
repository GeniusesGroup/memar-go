/* For license and copyright information please see the LEGAL file in the code repository */

package net

import (
	container_p "memar/computer/adt/container/protocol"
	buffer_p "memar/computer/buffer/protocol"
	mediatype_p "memar/identifier/mediatype/protocol"
	net_p "memar/net/protocol"
	error_p "memar/process/error/protocol"
)

//memar:impl memar/codec/protocol.Codec
func (sk *Socket) MediaType() mediatype_p.MediaType { return "" }

func (sk *Socket) Decode(source buffer_p.Buffer) (err error_p.Error) {
	return source.Encode(sk.Buffer())
}
func (sk *Socket) Encode(destination buffer_p.Buffer) (err error_p.Error) {
	return destination.Decode(sk)
}
func (sk *Socket) Marshal() (data []byte, err error_p.Error) {
	err = sk.Check()
	if err != nil {
		return
	}

	if !sk.buf.Full() {
		err = sk.blockInSelect()
	}
	// TODO::: check and wrap above error?
	return sk.buf.Marshal()
}
func (sk *Socket) Unmarshal(source []byte) (n container_p.NumberOfElement, err error_p.Error) {
	for len(source) > 0 {
		err = sk.Check()
		if err != nil {
			return
		}

		var sendNumber container_p.NumberOfElement
		sendNumber, err = sk.sendPayload(source)
		if err != nil {
			return
		}
		n += sendNumber
		source = source[sendNumber:]
	}
	return
}

//memar:impl memar/codec/protocol.Field_Length
func (sk *Socket) SerializationLength() (ln container_p.NumberOfElement) {
	return sk.buf.SerializationLength()
}

// BlockInSelect waits for something to happen, which is one of the following conditions in the function body.
func (sk *Socket) blockInSelect() (err error_p.Error) {
	// TODO::: check auto scheduling or block??

loop:
	for {
		select {
		// TODO::: if buffer not full but before get push flag go to full state??
		// I think we must send custom package level flag here when process last segment change buffer state to full.
		case state := <-sk.State():
			switch state {
			case net_p.Status_ReceivedCompletely:
				sk.socketTimer.Stop()
				break loop
			default:
				// TODO::: attack??
				goto loop
			}
		}
	}
	return
}

func (sk *Socket) sendPayload(b []byte) (n int, err error_p.Error) {
	return
}
