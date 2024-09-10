/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

import (
	datatype_p "memar/datatype/protocol"
	error_p "memar/error/protocol"
)

type DataTypes interface {
	Register(dt datatype_p.DataType) (err error_p.Error)
	GetByID(id datatype_p.DataTypeID) (dt datatype_p.DataType, err error_p.Error)
}
