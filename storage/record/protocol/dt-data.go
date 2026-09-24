/* For license and copyright information please see the LEGAL file in the code repository */

package record_p

import (
	buffer_p "memar/computer/buffer/protocol"
)

type Field_Data interface {
	Data() Data
}

// TODO::: can provide buffer protocol here??
// If it is buffer we MUST change the model to something like `file-descriptor`
type Data buffer_p.Buffer
