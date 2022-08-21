/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	"github.com/GeniusesGroup/libgo/protocol"
)

/*
********** protocol.Codec interface **********
 */

func (s *Socket) MediaType() protocol.MediaType       { return nil }
func (s *Socket) CompressType() protocol.CompressType { return nil }

func (s *Socket) Decode(reader protocol.Reader) (err protocol.Error) { return }
func (s *Socket) Encode(writer protocol.Writer) (err protocol.Error) {
	var _, goErr = s.WriteTo(writer)
	if goErr != nil {
		// err =
	}
	return
}
func (s *Socket) Marshal() (data []byte, err protocol.Error) {
	err = s.checkSocket()
	if err != nil {
		return
	}

	if !s.recv.buf.Full() {
		err = s.blockInSelect()
	}
	return s.recv.buf.Marshal()
}
func (s *Socket) MarshalTo(data []byte) (added []byte, err protocol.Error) {
	err = s.checkSocket()
	if err != nil {
		return
	}

	if !s.recv.buf.Full() {
		err = s.blockInSelect()
	}
	return s.recv.buf.MarshalTo(data)
}
func (s *Socket) Unmarshal(data []byte) (n int, err protocol.Error) {
	err = s.checkSocket()
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
func (s *Socket) UnmarshalFrom(data []byte) (remaining []byte, err protocol.Error) {
	return
}
func (s *Socket) Len() (ln int) { return s.recv.buf.Len() }
