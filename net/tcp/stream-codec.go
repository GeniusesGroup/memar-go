/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	"memar/protocol"
)

//memar:impl memar/protocol.Codec
func (s *Stream) MediaType() protocol.MediaType       { return nil }
func (s *Stream) CompressType() protocol.CompressType { return nil }
func (s *Stream) Decode(source protocol.Codec) (n int, err protocol.Error) {
	return source.Encode(s)
}
func (s *Stream) Encode(destination protocol.Codec) (n int, err protocol.Error) {
	return destination.Decode(s)
}
func (s *Stream) Marshal() (data []byte, err protocol.Error) {
	err = s.checkStream()
	if err != nil {
		return
	}

	if !s.recv.buf.Full() {
		err = s.blockInSelect()
	}
	// TODO::: check and wrap above error?
	return s.recv.buf.Marshal()
}
func (s *Stream) MarshalTo(data []byte) (added []byte, err protocol.Error) {
	err = s.checkStream()
	if err != nil {
		return
	}

	if !s.recv.buf.Full() {
		err = s.blockInSelect()
	}
	// TODO::: check and wrap above error?
	return s.recv.buf.MarshalTo(data)
}
func (s *Stream) Unmarshal(data []byte) (n int, err protocol.Error) {
	err = s.checkStream()
	if err != nil {
		return
	}

	for len(data) > 0 {
		select {
		case <-s.writeTimer.Signal():
			// err =
			return
		default:
			var sendNumber int
			sendNumber, err = s.sendPayload(data)
			if err != nil {
				return
			}
			n += sendNumber
			data = data[sendNumber:]
		}
	}
	return
}
func (s *Stream) UnmarshalFrom(data []byte) (remaining []byte, err protocol.Error) {
	return
}
func (s *Stream) Len() (ln int) { return s.recv.buf.Len() }
