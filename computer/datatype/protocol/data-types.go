/* For license and copyright information please see the LEGAL file in the code repository */

package datatype_p

import (
	container_p "memar/adt/container/protocol"
	error_p "memar/process/error/protocol"
)

type Field_DataTypes interface {
	DataTypes() DataTypes
}

// DataTypes is a container that store some datatype
type DataTypes interface {
	container_p.Container_READONLY[DataType]
}
