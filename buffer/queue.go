/* For license and copyright information please see the LEGAL file in the code repository */

package buffer

import "memar/protocol"

// Concurrent safe
type Queue struct {
	// queue    []ringBuf
	head     *ringBuf
	tail     *ringBuf
	totalLen int
	totalCap int
}

//memar:impl memar/protocol.ObjectLifeCycle
func (q *Queue) Init(initCap int) (err protocol.Error) {
	q.totalCap = initCap
	// TODO:::
	return
}
func (q *Queue) Reinit() (err protocol.Error) {
	// TODO:::
	return
}
func (q *Queue) Deinit() (err protocol.Error) {
	// TODO:::
	return
}

func (q *Queue) Full() bool { return false }

func (q *Queue) Read(b []byte) (n int, err protocol.Error) {
	return
}
func (q *Queue) Write(b []byte) (n int, err protocol.Error) {
	var lenB = len(b)
	if q.head.ln == 0 {
		q.head.data = make([]byte, lenB)
		copy(q.head.data, b)
	} else {
		var rb = ringBuf{
			data: make([]byte, lenB),
			ln:   lenB,
		}
		copy(rb.data, b)
		q.head.next = &rb
		q.tail = &rb
	}
	return
}
func (q *Queue) WriteIn(b []byte, at uint32) (n int, err protocol.Error) {
	// TODO:::
	return
}

// ----------/-
type ringBuf struct {
	data   []byte
	offset int
	ln     int
	next   *ringBuf
}
