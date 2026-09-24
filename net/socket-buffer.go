/* For license and copyright information please see the LEGAL file in the code repository */

package net

import (
	buffer_p "memar/computer/buffer/protocol"
	error_p "memar/process/error/protocol"
)

type buf struct {
	sendBuffer    buffer_p.Buffer
	receiveBuffer buffer_p.Buffer
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (b *buf) Init() (err error_p.Error) {
	// TODO:::
	return
}
func (b *buf) Reinit() (err error_p.Error) {
	// TODO:::
	return
}
func (b *buf) Deinit() (err error_p.Error) {
	// TODO:::
	return
}
