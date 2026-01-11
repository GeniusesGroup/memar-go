/* For license and copyright information please see the LEGAL file in the code repository */

package file_p

import (
	string_p "memar/codec/string/protocol"
)

// A filename extension, file name extension or file extension is
// a suffix to the name of a computer file (for example, .txt, .docx, .md).
// The extension indicates a characteristic of the file contents or its intended use.
// A filename extension is typically delimited from the rest of the filename with a period,
// but in some systems it is separated with spaces.
// https://en.wikipedia.org/wiki/Filename_extension
type Extension interface {
	string_p.String
}

type Field_Extension interface {
	// FileExtension return extension name without a period or space or ...
	FileExtension() Extension
}
