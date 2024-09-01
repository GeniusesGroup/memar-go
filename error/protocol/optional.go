/* For license and copyright information please see the LEGAL file in the code repository */

package error_p

// Belows are some other error behaviors that Errors capsule CAN implement,
// but any error package can introduce new behaviors for any purpose.

type Internal interface {
	// Who cause the error?
	// Internal	: means calling process logic has runtime bugs like HTTP server error status codes ( 500 – 599 ).
	// Caller	: Opposite of internal that indicate caller give some data that cause the error like HTTP client error status codes ( 400 – 499 )
	Internal() bool
}

type Temporary interface {
	// opposite is permanent situation
	Temporary() bool
}

type Timeout interface {
	Timeout() bool
}
