/* For license and copyright information please see the LEGAL file in the code repository */

package datatypes_p

import (
	datatype_p "memar/computer/datatype/protocol"
	error_p "memar/process/error/protocol"
)

type DataTypes_Map interface {
	Register(dt datatype_p.DataType) (err error_p.Error)
	GetByID(id datatype_p.ID) (dt datatype_p.DataType, err error_p.Error)
}
