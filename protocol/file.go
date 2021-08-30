/* For license and copyright information please see LEGAL file in repository */

package protocol

// FileDirectory is the descriptor interface that must implement by any to be an file!
// File owner is one app so it must handle concurrent protection internally not by file it self!
type FileDirectory interface {
	MetaData() FileMetaData

	Directories(offset, limit uint64) (names []string)
	Directory(name string) (dir FileDirectory, err Error) // make if not exist before

	Files(offset, limit uint64) (names []string)
	File(name string) (file File, err Error)

	FindFiles(partName string, num uint64) (uriPaths []string)
	FindFile(partName string) (uriPath string)

	Rename(oldURIPath, newURIPath string) (err Error)
	Delete(uriPath string) (err Error) // make invisible just by remove from primary index
	Wipe(uriPath string) (err Error)   // make invisible by remove from primary index & write random data to all file locations
}

// File is the descriptor interface that must implement by any to be an file!
// File owner is one app so it must handle concurrent protection internally not by file it self!
type File interface {
	MetaData() FileMetaData
	Codec
}

// FileMetaData is the interface that must implement by any file and directory!
type FileMetaData interface {
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
	// URI() string    // always return full file uri as "{{file}}://{{sabz.city}}/{{path/to/{{{{the file}}.{{html}}}}}}"
	// Scheme() string // always return "file"
	Domain() string
	Path() string // File location from root directory include file name
	Name() string
	Extension() string
}
