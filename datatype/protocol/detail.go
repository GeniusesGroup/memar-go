/* For license and copyright information please see the LEGAL file in the code repository */

package datatype_p

// Detail is some piece of information that write for humans to understand some thing.
//
// When defining a new detail, the information included must be carefully vetted.
// Likewise, when actually generating a problem (however it is serialized) the details given must also be scrutinized.
// Risks include leaking information that can be exploited to compromise the system,
// access to the system, or the privacy of users of the system.
type Detail interface {
	Quiddity

	// Domain return locale domain name that MediaType belongs to it.
	// More user friendly domain name to show to users on screens.
	Domain() string
	// Summary return locale general summary text that gives the main points in a concise form.
	// Usually it is one line text to shown in the '<app> help' output of service or notify errors to user.
	Summary() string
	// Overview return locale general text that gives the main ideas without explaining all the details.
	// Usually it is multi line text to shown in the '<app> help <this-command>' output or expand error details in GUI screen.
	Overview() string
	// UserNote return locale note that user do when face this MediaType
	// Description text that gives the main ideas with explaining all the details and purposes.
	UserNote() string
	// DevNote return locale technical advice for developers
	// Description text that gives the main ideas with explaining all the details and purposes.
	DevNote() string
	// TAGS return locale MediaType tags to sort MediaType in groups for any purpose e.g. in GUI to help org manager to give service delegate authorization to staffs.
	TAGS() []string
}

// Quiddity is the essence that makes something the kind of thing it is and makes it different from any other
type Quiddity interface {
	Name() string         // e.g. initialize
	Abbreviation() string // e.g. init
	// These are not suggested to the user in the shell completion,
	// but accepted if entered manually.
	Aliases() []string // e.g. []string{"initialise", "create"}
}
