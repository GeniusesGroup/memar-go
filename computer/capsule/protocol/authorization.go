/* For license and copyright information please see the LEGAL file in the code repository */

package capsule_p

import (
	datatypes_p "memar/computer/datatypes/protocol"
)

type Authorization interface {
	// LimitToUseBy or ImportableBy
	LimitToUseBy() datatypes_p.DataTypes
}
