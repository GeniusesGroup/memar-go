/* For license and copyright information please see LEGAL file in repository */

package tcp

import (
	"../protocol"
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
func (s *Socket) Marshal() (data []byte) {
	if !s.recv.buf.Full() {
		select {
		case <-s.readDeadline.C:
			break
		case <-s.recv.pushFlag:
			break
		}
	}
	return s.recv.buf.Marshal()
}
func (s *Socket) MarshalTo(data []byte) []byte {
	// TODO::: check buffer and return
	// TODO::: listen to state channel and return when data ready
	return s.recv.buf.Marshal(data)
}
func (s *Socket) Unmarshal(data []byte) (err protocol.Error) { return }
func (s *Socket) UnmarshalFrom(data []byte) (remaining []byte, err protocol.Error) {
	return
}
func (s *Socket) Len() (ln int) { return }
