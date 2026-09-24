/* For license and copyright information please see the LEGAL file in the code repository */

package mp

import (
	capsule_p "memar/computer/capsule/protocol"
)

var testString = "test"
var _ capsule_p.Accessor[Path] = (*Path)(&testString)
