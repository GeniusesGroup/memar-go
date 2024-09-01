/* For license and copyright information please see the LEGAL file in the code repository */

package storage_p

import (
	"memar/protocol"
)

// Directory is the descriptor interface that must implement by any to be an file directory.
// Directory owner is one app so it must handle concurrent protection internally not by file it self.
type Directory interface {
	Metadata() DirectoryMetadata
	ParentDirectory() Directory

	Directories(offset, limit uint64) (dirs []Directory)
	Directory(name string) (dir Directory, err protocol.Error) // make if not exist before

	Files(offset, limit uint64) (files []File)
	File(name string) (file File, err protocol.Error) // make if not exist before
	FileByPath(uriPath string) (file File, err protocol.Error)

	FindFiles(partName string, num uint) (files []File)
	FindFile(partName string) (file File) // return first match file. It will prevent unneeded slice allocation.

	Rename(oldURIPath, newURIPath string) (err protocol.Error)
	Copy(uriPath, newURIPath string) (err protocol.Error)
	Move(uriPath, newURIPath string) (err protocol.Error)
	Delete(uriPath string) (err protocol.Error) // make invisible by move to recycle bin
	Remove(uriPath string) (err protocol.Error) // or PermanentlyDelete make invisible just by remove from primary index
	Erase(uriPath string) (err protocol.Error)  // make invisible by remove from primary index & write zero data to all file locations
}

// DirectoryMetadata is the interface that must implement by any file and directory.
type DirectoryMetadata interface {
	DirNum() adt_p.NumberOfElement  // return number of directory save in this directory
	FileNum() adt_p.NumberOfElement // return number of file save in this directory

	FileMetadata
}
