/* For license and copyright information please see the LEGAL file in the code repository */

package net_p

import (
	buffer_p "memar/computer/buffer/protocol"
)

type Field_Socket_Buffers interface {
	SendBuffer() buffer_p.Buffer
	ReceiveBuffer() buffer_p.Buffer
}
