/* For license and copyright information please see the LEGAL file in the code repository */

package error_p

import (
	adt_p "memar/computer/adt/protocol"
	datatype_p "memar/computer/datatype/protocol"
	mediatype_p "memar/identifier/mediatype/protocol"
)

type Field_Error interface {
	Error() Error
}

// Error is base behaviors that any Error capsule must implement.
// Error MUST NOT mix with `Log Event`.
// Error has static data for any client, to tell about system fault situation when processing client request.
// Log event carry on static and dynamic data for developers, to indicate system fault situation, and help them troubleshoot potential bugs.
//
// Equality between two Error values is determined solely by comparing their DataTypeID() (via datatype_p.Field_ID,
// itself part of datatype_p.DataType). A dedicated Equivalence method is intentionally NOT part of this
// abstraction: Error carries no per-instance dynamic field meaningful to compare beyond "which error occurred",
// so a separate equivalence check would only duplicate the DataTypeID comparison under a second name. See
// the "error" package's IsEqual for a generic, interface-level comparison helper.
type Error interface {
	ImplementsError

	datatype_p.DataType
	mediatype_p.Field_MediaType

	// Error embeds adt_p.Nil only, as a Go-specific realization — not because Error's canonical (memar-level)
	// definition only needs Nil. Khayyam's Error requires the full adt_p.ADT family (IsNil/IsNull/IsEmpty); Go's
	// narrowing to Nil is legitimate specifically because Go variables are true pointers with a real nil value,
	// while Khayyam variables are always references (no pointer/nil concept the same way), so Go's memar-go
	// realization can rely on a plain pointer-nil check where Khayyam cannot. This divergence is analogous to
	// ImplementsError's domain-specific naming (RFC 495413): a per-backend realization choice, not a change to
	// the abstraction's canonical composition. See RFC 495442 for the full reasoning, and the dedicated ADT
	// session (not yet held) for what IsNull/IsEmpty should actually mean for a value like Error.
	adt_p.Nil

	// Below methods comment in favor of log_p.Event_Message interface.
	// Error() string
	// string_p.Stringer
}

// ImplementsError is Error's own domain-specific realization of the general, tooling-facing
// abstraction_p.Implements pattern (see RFC 495413). A distinct name (ImplError, rather than the fully
// generic Implements()) is used so a single capsule can unambiguously declare intent for several different
// abstractions at once, each under its own distinctly-named method — a shared, generic Implements() alone
// cannot indicate which abstraction is being claimed once more than one is in play.
//
// This is a disambiguation convenience only. It provides no stronger safety or "sealing" guarantee than the
// generic abstraction_p.Implements would, in either Go or Khayyam: deliberately embedding a shared marker
// struct and deliberately writing this same empty method independently are equally trivial for anyone
// intent on misrepresenting a type. An earlier draft (RFC 495414, now Rejected) explored a Go-specific
// unexported/embedded-marker version claiming a stronger, compiler-enforced guarantee; that claim did not
// hold up and was withdrawn, which is why this method is a plain exported method with no required embedding.
type ImplementsError interface {
	ImplError()
}
