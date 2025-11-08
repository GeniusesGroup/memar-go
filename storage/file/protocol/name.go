/* For license and copyright information please see the LEGAL file in the code repository */

package file_p

import (
	uri_p "memar/net/uri/protocol"
	string_p "memar/codec/string/protocol"
)

type Name interface {
	string_p.String
}

type Field_Name interface {
	FileName() Name // Full name with extension if exist
	FileNameWithoutExtension() Name
}

type Method_Rename interface {
	// Usually it MUST call by upper function that call change the directory database information.
	Rename(newName Name) (err error_p.Error)
}
