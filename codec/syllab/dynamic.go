/* For license and copyright information please see the LEGAL file in the code repository */

package syllab

import (
	error_p "memar/process/error/protocol"
)

// Dynamic try to implement a logic over syllab to be like Flat-buffers, ...
type Dynamic struct {
	buf          []byte
	fieldsNumber int
	heapIndex    int
	// fields []dynamic
}

func (d *Dynamic) Init(initialSize int) {}

func (d *Dynamic) Unmarshal(buf []byte) (n int, err error_p.Error) {
	d.buf = buf
	var firstFiled = d.Filed(0)
	d.heapIndex = int(firstFiled.loc)
	d.fieldsNumber = int(firstFiled.loc) / 8
	return
}

func (d *Dynamic) Filed(num int) dynamic {
	var stackIndex = num * 8
	return dynamic{
		loc:    GetUInt32(d.buf, uint32(stackIndex)),
		length: GetUInt32(d.buf, uint32(stackIndex+4)),
	}
}
