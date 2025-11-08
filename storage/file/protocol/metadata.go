/* For license and copyright information please see the LEGAL file in the code repository */

package file_p

import (
	container_p "memar/adt/container/protocol"
	string_p "memar/codec/string/protocol"
	time_p "memar/time/protocol"
)

// Metadata is the interface that must implement by any file and directory.
type Metadata interface {
	URI
	Size() container_p.NumberOfElement // in Byte??
	Created() time_p.Time
	Accessed() time_p.Time
	Modified() time_p.Time
}

type Field_Metadata interface {
	// File descriptor is just a capsule that need to read data from other sources,
	// So this field get call can return error.
	Metadata() (md Metadata, err error_p.Error)
}
