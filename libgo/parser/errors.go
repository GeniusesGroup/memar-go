/* For license and copyright information please see the LEGAL file in the code repository */

package parser

import "errors"

//Declare Errors Details
var (
	ErrEmptyPackageFolder   = errors.New("You can't have empty logic folder or logic package contain no compatible function")
	ErrBadServiceName       = errors.New("Each service file must start with ServiceID")
	ErrBadServiceParameters = errors.New("Service functions must have just one struct as parameters")
)
