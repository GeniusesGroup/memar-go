/* For license and copyright information please see the LEGAL file in the code repository */

package directory_p

import (
	container_p "memar/adt/container/protocol"
	error_p "memar/process/error/protocol"
	file_p "memar/storage/file/protocol"
)

// A directory is a special-purpose database that contains typed information.
// A directory usually supports both read and search of the information it contains,
// and can support creation and modification of the information as well.
// 
// A directory is just a kind of file like other files that store some data.
// A directory file store hierarchical file information and may serve other services like AAA.
// Directory is the descriptor interface that must implement by any to be an file directory.
// Directory owner is one app so it must handle concurrent protection internally not by file it self.
// 
// It can implement in many ways e.g. linux inode(https://en.wikipedia.org/wiki/Inode), ...
// `text/directory` chose by `rfc2425` for directory MIME Content-Type.
// https://www.w3.org/2002/12/cal/rfc2425.html
// 
// https://en.wikipedia.org/wiki/Directory_(computing)
type Directory interface {
	Field_Metadata

	Directories(offset container_p.Offset, limit container_p.Limit) (dirs []Directory, err error_p.Error)
	Directory(name file_p.Name) (dir Directory, err error_p.Error) // make if not exist before

	Files(offset container_p.Offset, limit container_p.Limit) (files []file_p.File, err error_p.Error)
	File(name file_p.Name) (file file_p.File, err error_p.Error) // make if not exist before

	// Just delete file from directory not delete file itself from NVM device.
	Delete(name file_p.Name) (err error_p.Error)

	Method_Directory_Helpers
}

type Method_Directory_Helpers interface {
	FileByPath(uriPath uri_p.Path) (file file_p.File, err error_p.Error)

	FindFiles(partName file_p.Name, offset container_p.Offset, limit container_p.Limit) (files []file_p.File, err error_p.Error)
	FindFile(partName file_p.Name) (file file_p.File, err error_p.Error) // return first match file. It will prevent unneeded slice allocation.

	// = Call File() and write `Metadata`&`Data`
	Copy(uriPath, newURIPath uri_p.Path) (err error_p.Error)
	// = Call Copy() and Delete()
	Move(uriPath, newURIPath uri_p.Path) (err error_p.Error)

	// = Call File() and Delete()
	// It will return two types of errors, `File()` and `Delete()` errors
	DeleteByPath(uriPath uri_p.Path) (err error_p.Error)
	// or PermanentlyDelete make invisible by call both `Delete()` in directory and file to remove from directory index and nvm.
	Remove(name Name) (err error_p.Error) 
	RemoveByPath(uriPath uri_p.Path) (err error_p.Error) 
}
