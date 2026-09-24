/* For license and copyright information please see the LEGAL file in the code repository */

package error_p

// Belows are some other error behaviors that Errors capsule CAN implement,
// but any error package can introduce new behaviors for any purpose.

// Chian errors can use in situation that error occurred but can't handle in that layer and
// Upper layer also can't handle it directly and decide to chain it!
// Use chain errors to return multiple error to caller and
// Usually caller log them and return another error to caller that almost always SDK client e.g. GUI, Apps, ...
type Chain interface {
	PastErrorChain() (last Error)
	// CauseError() (cause Error)
}
