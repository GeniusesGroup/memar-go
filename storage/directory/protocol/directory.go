/* For license and copyright information please see the LEGAL file in the code repository */

package file_p

import (
	container_p "memar/adt/container/protocol"
	error_p "memar/error/protocol"
)

// Directory is the descriptor interface that must implement by any to be an file directory.
// Directory owner is one app so it must handle concurrent protection internally not by file it self.
type Directory interface {
	Metadata() (md DirectoryMetadata, err error_p.Error)
	ParentDirectory() (dir Directory, err error_p.Error)

	Directories(offset container_p.ElementIndex, limit container_p.NumberOfElement) (dirs []Directory, err error_p.Error)
	Directory(name string) (dir Directory, err error_p.Error) // make if not exist before

	Files(offset container_p.ElementIndex, limit container_p.NumberOfElement) (files []File, err error_p.Error)
	File(name string) (file File, err error_p.Error) // make if not exist before
	FileByPath(uriPath string) (file File, err error_p.Error)

	FindFiles(partName string, num uint) (files []File, err error_p.Error)
	FindFile(partName string) (file File, err error_p.Error) // return first match file. It will prevent unneeded slice allocation.

	Rename(oldURIPath, newURIPath string) (err error_p.Error)
	Copy(uriPath, newURIPath string) (err error_p.Error)
	Move(uriPath, newURIPath string) (err error_p.Error)
	Delete(uriPath string) (err error_p.Error) // make invisible by move to recycle bin
	Remove(uriPath string) (err error_p.Error) // or PermanentlyDelete make invisible just by remove from primary index
	Erase(uriPath string) (err error_p.Error)  // make invisible by remove from primary index & write zero data to all file locations
}

// DirectoryMetadata is the interface that must implement by any file and directory.
type DirectoryMetadata interface {
	DirNumber() container_p.NumberOfElement  // return number of directory save in this directory
	FileNumber() container_p.NumberOfElement // return number of file save in this directory

	Metadata
}
