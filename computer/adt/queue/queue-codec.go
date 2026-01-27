/* For license and copyright information please see the LEGAL file in the code repository */

package queue

import (
	adt_p "memar/adt/protocol"
	buffer_p "memar/buffer/protocol"
	"memar/protocol"
)

//memar:impl memar/codec/protocol.Codec
func (q *Queue) Decode(source buffer_p.Buffer) (err protocol.Error) { return }
func (q *Queue) Encode(destination buffer_p.Buffer) (err protocol.Error) {
	return
}
func (q *Queue) Marshal() (data []byte, err protocol.Error) {
	return
}
func (q *Queue) Unmarshal(source []byte) (n adt_p.NumberOfElement, err protocol.Error) {
	return
}
func (q *Queue) SerializationLength() (ln adt_p.NumberOfElement) { return q.totalLen }
