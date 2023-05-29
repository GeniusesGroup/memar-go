/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Quiddity is the essence that makes something the kind of thing it is and makes it different from any other
type Quiddity interface {
	Name() string         // e.g. initialize
	Abbreviation() string // e.g. init
	// These are not suggested to the user in the shell completion,
	// but accepted if entered manually.
	Aliases() []string // e.g. []string{"initialise", "create"}
}
