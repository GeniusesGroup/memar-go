/* For license and copyright information please see the LEGAL file in the code repository */

package file_p

import (
	container_p "memar/adt/container/protocol"
	error_p "memar/process/error/protocol"
	audit_time_p "memar/process/time/protocol"
)

// Metadata is the interface that must implement by any file and directory.
type Metadata interface {
	Field_URI
	Field_URI_Parsed

	// Capacity use in sparse file usage.
	// In computer science, a sparse file is a type of computer file that attempts to use file system space more efficiently when the file itself is partially empty.
	// https://en.wikipedia.org/wiki/Sparse_file
	container_p.Field_Capacity // in Byte??
	container_p.Field_OccupiedLength
	container_p.Field_AvailableLength

	audit_time_p.Field_Creation
	audit_time_p.Field_Modification
	audit_time_p.Field_Access
}

type Field_Metadata interface {
	// File descriptor is just a capsule that need to read data from other sources,
	// So this field get call can return error.
	Metadata() (md Metadata, err error_p.Error)
}
