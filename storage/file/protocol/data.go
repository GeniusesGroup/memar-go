/* For license and copyright information please see the LEGAL file in the code repository */

package file_p

import (
	buffer_p "memar/buffer/protocol"
)

type Data interface {
	buffer_p.Buffer
}

type Field_Data interface {
	// File descriptor is just a capsule that need to read data from other sources,
	// So this field get call can return error.
	Data() (d Data, err error_p.Error)
}
