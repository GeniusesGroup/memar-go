/* For license and copyright information please see LEGAL file in repository */

package json

import "../errors"

// Declare Errors Details
var (
	ErrJSONNeededTypeNotExist      = errors.New("JSONNeededTypeNotExist", "")
	ErrJSONTypeIncludeIllegalChild = errors.New("JSONTypeIncludeIllegalChild", "Requested type may include function, interface, int, uint, ... type that can't encode||decode")
	ErrJSONArrayLenNotSupported    = errors.New("JSONArrayLenNotSupported", "")

	ErrJSONEncodedIncludeNotDeffiendKey = errors.New("JSONEncodedIncludeNotDeffiendKey", "Given encoded json string include a key that must not be in the encoded string")
	ErrJSONEncodedStringCorrupted       = errors.New("JSONEncodedStringCorrupted", "Given encoded json string corruputed and not encode in the way that can decode")
)
