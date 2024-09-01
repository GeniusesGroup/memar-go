/* For license and copyright information please see the LEGAL file in the code repository */

package storage_p

import (
	buffer_p "memar/buffer/protocol"
	uri_p "memar/net/uri/protocol"
	"memar/protocol"
	string_p "memar/string/protocol"
	time_p "memar/time/protocol"
)

// File is the descriptor interface that must implement by any to be an file.
// File owner is one app so it must handle concurrent protection internally not by file it self.
type File interface {
	Metadata() FileMetadata
	Data() buffer_p.Buffer

	File_Methods
}

// FileMetadata is the interface that must implement by any file and directory.
type FileMetadata interface {
	ParentDirectory() Directory
	FileURI[string_p.String]
	Size() adt_p.NumberOfElement // in Byte
	Created() time_p.Time
	Accessed() time_p.Time
	Modified() time_p.Time
}

// https://datatracker.ietf.org/doc/html/rfc8089
// https://en.wikipedia.org/wiki/File_URI_scheme
type FileURI[STR string_p.String] interface {
	// URI[STR] // always return full file uri as "{{file}}://{{domain}}/{{path/to/{{{{the file}}.{{html}}}}}}"
	uri_p.Scheme[STR] // always return "file"
	Domain() string   // same as URI.Authority
	// File location from root directory include file name if not point to a directory.
	uri_p.Path[STR]

	FileName[STR]
	FileExtension

	// Helper functions
	IsDirectory() bool // Path end with "/" if it is a file directory
}

type FileName[STR string_p.String] interface {
	FileName() STR // Full name with extension if exist
	FileNameWithoutExtension() STR
}

// A filename extension, file name extension or file extension is
// a suffix to the name of a computer file (for example, .txt, .docx, .md).
// The extension indicates a characteristic of the file contents or its intended use.
// A filename extension is typically delimited from the rest of the filename with a period,
// but in some systems it is separated with spaces.
// https://en.wikipedia.org/wiki/Filename_extension
type FileExtension interface {
	// FileExtension return extension name without a period or space or ...
	FileExtension() string
}

type File_Methods interface {
	// Depend on OS, file data can be cache on ram until Save() called.
	Save() (err protocol.Error)

	Rename(newName string)
}
