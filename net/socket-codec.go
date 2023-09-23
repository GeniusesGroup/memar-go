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

	if !sk.buf.Full() {
		err = sk.blockInSelect()
	}
	// TODO::: check and wrap above error?
	return sk.buf.Marshal()
}
func (sk *Socket) MarshalTo(data []byte) (added []byte, err protocol.Error) {
	err = sk.Check()
	if err != nil {
		return
	}

	if !sk.buf.Full() {
		err = sk.blockInSelect()
	}
	// TODO::: check and wrap above error?
	return sk.buf.MarshalTo(data)
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
func (sk *Socket) Len() (ln int) { return sk.buf.Len() }

// BlockInSelect waits for something to happen, which is one of the following conditions in the function body.
func (sk *Socket) blockInSelect() (err protocol.Error) {
	// TODO::: check auto scheduling or block??

loop:
	for {
		select {
		// TODO::: if buffer not full but before get push flag go to full state??
		// I think we must send custom package level flag here when process last segment change buffer state to full.
		case state := <-sk.State():
			switch state {
			case protocol.NetworkStatus_ReceivedCompletely:
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

func (sk *Socket) sendPayload(b []byte) (n int, err protocol.Error) {
	return
}
