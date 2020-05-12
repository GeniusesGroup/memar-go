/* For license and copyright information please see LEGAL file in repository */

package syllab

import "errors"

// Declare Errors Details
var (
	ErrNeededTypeNotExist      = errors.New("")
	ErrTypeIncludeIllegalChild = errors.New("Requested type may include function, interface, int, uint, ... type that can't encode||decode")
	ErrArrayLenNotSupported    = errors.New("")
)
