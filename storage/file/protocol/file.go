/* For license and copyright information please see the LEGAL file in the code repository */

package file_p

import (
	error_p "memar/process/error/protocol"
)

// File is the descriptor interface that must implement by any to be an file.
// File owner is one app so it must handle concurrent protection internally not by file it self.
// 
// https://en.wikipedia.org/wiki/Computer_file
type File interface {
	Field_Metadata
	Field_Data

	Method_File
}

type Field_File interface {
	File() File
}

type Method_File interface {
	// Depend on OS, file data can be cache on ram until `Save()` or `Flush()` called.
	Save() (err error_p.Error)
	
	// Just delete file from NVM device not from any directory that point to it.
	Delete() (err error_p.Error)
	// Make invisible by write zero data to prevent any recovery process.
	Erase(uriPath string) (err error_p.Error) 

	Method_Rename
}
