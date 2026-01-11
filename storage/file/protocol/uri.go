/* For license and copyright information please see the LEGAL file in the code repository */

package file_p

import (
	uri_p "memar/net/uri/protocol"
)

// https://datatracker.ietf.org/doc/html/rfc8089
// https://en.wikipedia.org/wiki/File_URI_scheme
type URI interface {
	// always return full file uri as "{{file}}://{{authority}}/{{path/to/{{{{the file}}.{{html}}}}}}"
	uri_p.URI
}

type Field_URI interface {
	// URI Return full file URI
	URI() URI
}

type Field_URI_Parsed interface {
	// always return "file"
	uri_p.Field_Scheme
	// TODO::: Is it ok to have `port` in authority?
	uri_p.Field_Authority
	// File location from root directory include file name if not point to a directory.
	uri_p.Field_Path

	Field_Name
	Field_Extension
}
