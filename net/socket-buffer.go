/* For license and copyright information please see the LEGAL file in the code repository */

package net

import (
	buffer_p "memar/computer/buffer/protocol"
	error_p "memar/process/error/protocol"
)

type buf[BUF buffer_p.Buffer] struct {
	buf BUF
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (b *buf[BUF]) Init() (err error_p.Error) {
	// TODO:::
	return
}
func (b *buf[BUF]) Reinit() (err error_p.Error) {
	// TODO:::
	return
}
func (b *buf[BUF]) Deinit() (err error_p.Error) {
	// TODO:::
	return
}
