/* For license and copyright information please see the LEGAL file in the code repository */

package directory_p

import (
	container_p "memar/adt/container/protocol"
	error_p "memar/process/error/protocol"
	file_p "memar/storage/file/protocol"
)

// Metadata is the interface that must implement by any file and directory.
type Metadata interface {
	DirNumber() container_p.NumberOfElement  // return number of directory save in this directory
	FileNumber() container_p.NumberOfElement // return number of file save in this directory

	file_p.Metadata
}

type Field_Metadata interface {
	// File descriptor is just a capsule that need to read data from other sources,
	// So this field get call can return error.
	Metadata() (md Metadata, err error_p.Error)
}
