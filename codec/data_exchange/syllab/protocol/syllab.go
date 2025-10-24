/* For license and copyright information please see the LEGAL file in the code repository */

package syllab_p

import (
	buffer_p "memar/buffer/protocol"
	error_p "memar/process/error/protocol"
)

// Syllab is the interface that must implement by any `Capsule` to be a Syllab object transmittable over networks.
// Standards in https://github.com/GeniusesGroup/memar/blob/main/Syllab.md
type Syllab interface {
	Field_Lengths

	// **Be Aware** that below methods has runtime impact and MUST use `Codec` methods instead for codec purposes.
	Syllab_Stack() (stack buffer_p.Buffer, err error_p.Error)
	Syllab_Heap() (heap buffer_p.Buffer, err error_p.Error)
}
