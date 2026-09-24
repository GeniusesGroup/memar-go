/* For license and copyright information please see the LEGAL file in the code repository */

package datatypes_p

import (
	container_p "memar/computer/adt/container/protocol"
	datatype_p "memar/computer/datatype/protocol"
	// error_p "memar/process/error/protocol"
)

type Field_DataTypes interface {
	DataTypes() DataTypes
}

// DataTypes is a container that store some datatype
type DataTypes interface {
	container_p.Container_READONLY[datatype_p.DataType]
}
