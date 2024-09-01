/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

type DataTypes interface {
	Register(dt DataType) (err Error)
	GetByID(id DataTypeID) (dt DataType, err Error)
}
