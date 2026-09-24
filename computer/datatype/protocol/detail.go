/* For license and copyright information please see the LEGAL file in the code repository */

package datatype_p

// Detail is some piece of information that write for humans to understand some thing.
//
// When defining a new detail, the information included must be carefully vetted.
// Likewise, when actually generating a problem (however it is serialized) the details given must also be scrutinized.
// Risks include leaking information that can be exploited to compromise the system,
// access to the system, or the privacy of users of the system.
//
// Summary/Overview vs UserActionNote/DevActionNote: Summary and Overview document the DataType itself, for
// someone inspecting or managing the type in general (e.g. a manager reviewing what a service/role does, or a
// CLI/GUI listing available types) — read independently of any specific occurrence. UserActionNote and
// DevActionNote, by contrast, are addressed to whoever is actually facing an occurrence of this DataType
// (typically an Error) at the moment it happens, and describe what action to take next: UserActionNote for the
// end user facing the situation, DevActionNote for the developer/integrator handling it.
type Detail interface {
	Quiddity

	// Domain return locale domain name that DataType belongs to it.
	// More user friendly domain name to show to users on screens.
	Domain() string
	// Summary return locale general summary text that gives the main points in a concise form.
	// Usually it is one line text to shown in the '<app> help' output of service or notify errors to user.
	Summary() string
	// Overview return locale general text that gives the main ideas without explaining all the details.
	// Usually it is multi line text to shown in the '<app> help <this-command>' output or expand error details in GUI screen.
	Overview() string
	// UserActionNote return locale note describing what action the end user should take when facing this DataType.
	// Description text that gives the main ideas with explaining all the details and purposes, addressed directly to the end user.
	UserActionNote() string
	// DevActionNote return locale technical note describing what action the developer/integrator should take when facing this DataType.
	// Description text that gives the main ideas with explaining all the details and purposes, addressed to a developer.
	DevActionNote() string
	// TAGS return locale DataType tags to sort DataType in groups for any purpose e.g. in GUI to help org manager to give service delegate authorization to staffs.
	TAGS() []string
}

// Quiddity is the essence that makes something the kind of thing it is and makes it different from any other.
//
// Abbreviation and Aliases are human-lookup conveniences only (e.g. helping a manager find a type under a name
// they happen to remember, even if imperfectly). Neither is guaranteed unique and neither may be used for
// dispatch, equality, or any mechanism requiring a guaranteed-unique identifier — that role belongs to DataTypeID.
type Quiddity interface {
	Name() string         // e.g. initialize
	Abbreviation() string // e.g. init
	// These are not suggested to the user in the shell completion,
	// but accepted if entered manually.
	Aliases() []string // e.g. []string{"initialise", "create"}
}
