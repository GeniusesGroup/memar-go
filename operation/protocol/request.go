/* For license and copyright information please see the LEGAL file in the code repository */

package operation_p

import (
	datatype_p "memar/datatype/protocol"
)

// Field_Request ...
type Field_Request/*[REQ Request]*/ interface {
	Request() datatype_p.DataType
}

// Request indicate request type that MUST use by any other types.
type Request interface {
	datatype_p.DataType
}
