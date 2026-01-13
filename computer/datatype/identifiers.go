/* For license and copyright information please see the LEGAL file in the code repository */

package datatype

import (
	datatype_p "memar/computer/datatype/protocol"
)

type Identifiers struct{}

//memar:impl memar/datatype/protocol.Identifiers
func (self *Identifiers) DataTypeID() datatype_p.ID { return 0 }
func (self *Identifiers) DataTypeID_Base64() string { return "" }
