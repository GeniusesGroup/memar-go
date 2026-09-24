/* For license and copyright information please see the LEGAL file in the code repository */

package queue

import (
	container_p "memar/computer/adt/container/protocol"
	buffer_p "memar/computer/buffer/protocol"
	error_p "memar/process/error/protocol"
)

//memar:impl memar/codec/protocol.Codec
func (q *Queue) Decode(source buffer_p.Buffer) (err error_p.Error) { return }
func (q *Queue) Encode(destination buffer_p.Buffer) (err error_p.Error) {
	return
}
func (q *Queue) Marshal() (data []byte, err error_p.Error) {
	return
}
func (q *Queue) Unmarshal(source []byte) (n container_p.NumberOfElement, err error_p.Error) {
	return
}
func (q *Queue) SerializationLength() (ln container_p.NumberOfElement) { return q.totalLen }
