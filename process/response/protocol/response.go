/* For license and copyright information please see the LEGAL file in the code repository */

package response_p

import (
	codec_p "memar/codec/protocol"
	datatype_p "memar/computer/datatype/protocol"
)

// Field_Response ...
type Field_Response/*[RES Response]*/ interface {
	Response() Response
}

// Response indicate response type that MUST use by any other types.
type Response interface {
	datatype_p.DataType   
	codec_p.Codec
}
