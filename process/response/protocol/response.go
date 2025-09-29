/* For license and copyright information please see the LEGAL file in the code repository */

package operation_p

import (
	datatype_p "memar/datatype/protocol"
)

// Field_Response ...
type Field_Response/*[RES Response]*/ interface {
	Response() datatype_p.DataType
}

// Response indicate response type that MUST use by any other types.
type Response interface {
	datatype_p.DataType
}
