/* For license and copyright information please see the LEGAL file in the code repository */

package service_p

import (
	datatype_p "memar/datatype/protocol"
)

type Field_ServiceID interface {
	// Usually easily return s.DataTypeID()
	// or it can return some old manual way numbering like HTTP:80, HTTPS:443, ...
	ServiceID() ID
}

type ID = datatype_p.ID
