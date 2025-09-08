/* For license and copyright information please see the LEGAL file in the code repository */

package operation_p

import (
	datatype_p "memar/datatype/protocol"
)

type Field_OperationID interface {
	// Usually easily return s.DataTypeID()
	// or it can return some old manual way numbering like HTTP:80, HTTPS:443, ...
	OperationID() ID
}

type ID = datatype_p.ID
