/* For license and copyright information please see the LEGAL file in the code repository */

package module_p

import (
	datatype_p "memar/computer/datatype/protocol"
)

// https://doc.rust-lang.org/cargo/reference/manifest.html
type Module interface {
	datatype_p.DataType
	datatype_p.Field_DataTypes
}
