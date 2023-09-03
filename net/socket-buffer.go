/* For license and copyright information please see the LEGAL file in the code repository */

package net

import (
	"memar/buffer"
	"memar/protocol"
)

type buf struct {
	buf buffer.Queue
}

//memar:impl memar/protocol.ObjectLifeCycle
func (b *buf) Init() (err protocol.Error) {
	// TODO:::
	return
}
func (b *buf) Reinit() (err protocol.Error) {
	// TODO:::
	return
}
func (b *buf) Deinit() (err protocol.Error) {
	// TODO:::
	return
}
