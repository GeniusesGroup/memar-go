/* For license and copyright information please see the LEGAL file in the code repository */

package capsule_p

import (
	string_p "memar/codec/string/protocol"
	datatype_p "memar/computer/datatype/protocol"
)

type Field_Field interface {
	Field() Field
}

type Field_Fields interface {
	Fields() []Field
}

type Field interface {
	datatype_p.DataType

	string_p.Stringer
}
