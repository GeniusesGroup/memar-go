/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// FileDirectory is the descriptor interface that must implement by any to be an file.
// File owner is one app so it must handle concurrent protection internally not by file it self.
type FileDirectory interface {
	Metadata() FileDirectoryMetadata
	ParentDirectory() (dir FileDirectory)

	Directories(offset, limit uint64) (dirs []FileDirectory)
	Directory(name string) (dir FileDirectory, err Error) // make if not exist before

	Files(offset, limit uint64) (files []File)
	File(name string) (file File, err Error) // make if not exist before
	FileByPath(uriPath string) (file File, err Error)

	FindFiles(partName string, num uint) (files []File)
	FindFile(partName string) (file File) // return first match file. It will prevent unneeded slice allocation.

	Rename(oldURIPath, newURIPath string) (err Error)
	Copy(uriPath, newURIPath string) (err Error)
	Move(uriPath, newURIPath string) (err Error)
	Delete(uriPath string) (err Error)            // make invisible by move to recycle bin
	PermanentlyDelete(uriPath string) (err Error) // make invisible just by remove from primary index
	Erase(uriPath string) (err Error)             // make invisible by remove from primary index & write zero data to all file locations
}

// FileDirectoryMetadata is the interface that must implement by any file and directory.
type FileDirectoryMetadata interface {
	DirNum() uint  // return number of directory save in this directory
	FileNum() uint // return number of file save in this directory

	FileMetadata
}

// File is the descriptor interface that must implement by any to be an file.
// File owner is one app so it must handle concurrent protection internally not by file it self.
type File interface {
	Metadata() FileMetadata
	Data() FileData
	ParentDirectory() (dir FileDirectory)

	Rename(newName string)
}

// FileMetadata is the interface that must implement by any file and directory.
type FileMetadata interface {
	URI() FileURI
	Size() uint64 // in Byte
	Created() Time
	Accessed() Time
	Modified() Time
}

// https://datatracker.ietf.org/doc/html/rfc8089
// https://en.wikipedia.org/wiki/File_URI_scheme
type FileURI interface {
	URI
	// URI() string    // always return full file uri as "{{file}}://{{domain}}/{{path/to/{{{{the file}}.{{html}}}}}}"
	// Scheme() string // always return "file"
	Domain() string // same as URI.Authority
	FileURI_Path
	Name() string // Full name with extension if exist
	Extension() string
	NameWithoutExtension() string

	// Helper functions
	IsDirectory() bool // Path end with "/" if it is a file directory
}

type FileURI_Path interface {
	// File location from root directory include file name if not point to a directory.
	Path() string
}

type FileData interface {
	// Depend on OS, file data can be cache on ram until Save() called.
	Save() (err Error)

	Prepend(data []byte)
	Append(data []byte)
	Replace(old, new []byte, n int)
	Codec
}
