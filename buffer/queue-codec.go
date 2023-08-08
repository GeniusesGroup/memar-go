/* For license and copyright information please see the LEGAL file in the code repository */

package buffer

import (
	"memar/protocol"
)

//memar:impl memar/protocol.Codec
func (q *Queue) MediaType() protocol.MediaType                            { return nil }
func (q *Queue) CompressType() protocol.CompressType                      { return nil }
func (q *Queue) Decode(source protocol.Codec) (n int, err protocol.Error) { return }
func (q *Queue) Encode(destination protocol.Codec) (n int, err protocol.Error) {
	return
}
func (q *Queue) Marshal() (data []byte, err protocol.Error) {
	return
}
func (q *Queue) MarshalTo(data []byte) (added []byte, err protocol.Error) {
	return
}
func (q *Queue) Unmarshal(data []byte) (n int, err protocol.Error) {
	return
}
func (q *Queue) UnmarshalFrom(data []byte) (remaining []byte, err protocol.Error) {
	return
}
func (q *Queue) Len() (ln int) { return q.totalLen }
