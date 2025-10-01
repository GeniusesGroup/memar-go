/* For license and copyright information please see the LEGAL file in the code repository */

package request_p

import (
	codec_p "memar/codec/protocol"
	datatype_p "memar/computer/datatype/protocol"
)

// Field_Request ...
type Field_Request/*[REQ Request]*/ interface {
	Request() Request
}

// Request indicate request type that MUST use by any other types.
type Request interface {
	datatype_p.DataType
	codec_p.Codec
}
