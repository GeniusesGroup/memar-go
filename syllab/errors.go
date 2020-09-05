/* For license and copyright information please see LEGAL file in repository */

package syllab

import "../errors"

// Declare Errors Details
var (
	ErrNeededTypeNotExist      = errors.New("NeededTypeNotExist", "Custom type exist in type that generator can't access it to know basic type")
	ErrTypeIncludeIllegalChild = errors.New("TypeIncludeIllegalChild", "Requested type may include function, interface, int, uint, ... type that can't encode||decode")
	ErrArrayLenNotSupported    = errors.New("ArrayLenNotSupported", "Length of array larger than 32 bit space that syllab can encode||decode")

	ErrSyllabDecodingFailedSmallSlice   = errors.New("SyllabDecodingFailedSmallSlice", "Given slice smaller than expected to decode data from it")
	ErrSyllabDecodingFailedHeapOverFlow = errors.New("SyllabDecodingFailedHeapOverFlow", "Encoded syllab want to access to out of slice.")
)
