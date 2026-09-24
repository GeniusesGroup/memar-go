/* For license and copyright information please see the LEGAL file in the code repository */

package error

import (
	"memar/computer/datatype"
	"memar/identifier/mediatype"
	error_p "memar/process/error/protocol"
)

// Error is a reusable base struct that a concrete Error capsule MAY embed.
// to obtain default implementations of ImplError and the ADT methods, reducing per-capsule boilerplate.
// This is purely a convenience — embedding it is not required to satisfy error_p.Error, and a concrete
// capsule remains free to write ImplError() itself instead, with no difference in what error_p.Error accepts either way.
// KEEP IN MIND that concrete errors SHOULD typically be developed via code generation.
type Error struct {
	datatype.DataType
	mediatype.MT
}

func (err *Error) ImplError() {}

func (err *Error) IsNil() bool { return err == nil }

// TODO::: IsNull is currently a hard-coded default (always true, regardless of state). Flagged during
// review: as written, ANY non-nil Error embedding this base struct still reports IsNull() == true, which is
// likely not the intended semantics once a concrete capsule has real content. Is this a deliberate
// "unset-by-default, concrete capsules override" placeholder, or does it need real logic? This is directly
// gated on the separate, not-yet-resolved ADT IsNil/IsNull/IsEmpty semantics discussion (deferred to its own
// dedicated chat) — resolve there first, then fix this accordingly.
func (err *Error) IsNull() bool { return true }

func (err *Error) IsEmpty() bool { return err.IsNil() && err.IsNull() }

// IsEqual compares two Error values. Equality is DataTypeID-based (see error_p.Error's doc comment).
// Suggest to add more logic to Error.Equal() logic such as chain situation.
//
// TODO::: MediaType() is also compared here. If MediaType is a fixed, type-level property (i.e. every
// instance sharing the same DataTypeID always reports the same MediaType), this check is redundant with the
// DataTypeID comparison and can be dropped. If MediaType can genuinely vary per-instance for the same
// DataTypeID, this reopens the "DataTypeID alone determines identity" principle documented on error_p.Error
// and needs to be resolved explicitly, not left as an incidental extra check. Needs a decision before this
// is considered final.
func IsEqual(base, with error_p.Error) bool {
	if base == nil && with == nil {
		// TODO::: return true?? Likely depends on the separate ADT IsNil/IsNull/IsEmpty semantics
		// discussion (deferred to its own dedicated chat) — resolve there, then reflect the answer here.
		return true
	}
	if (base == nil && with != nil) || base != nil && with == nil {
		return false
	}
	if base.MediaType() == with.MediaType() && base.DataTypeID() == with.DataTypeID() {
		return true
	}
	return false
}
