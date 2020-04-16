/* For license and copyright information please see LEGAL file in repository */

package generator

import "errors"

// Declare Errors Details
var (
	ErrEmptyPackageFolder    = errors.New("You can't have empty logic folder or logic package contain no compatible function")
	ErrBadServiceServiceName = errors.New("Each service file must start with ServiceID")
	ErrBadServiceParameters  = errors.New("Service functions must have just one struct as parameters")
)
