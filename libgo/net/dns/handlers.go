/* For license and copyright information please see the LEGAL file in the code repository */

package dns

import (
	"libgo/protocol"
)

// Handler use to standard services handlers in any layer!
type Handler func(protocol.Socket)

// IncomeRequestHandler handle incoming DNS request streams!
func IncomeRequestHandler(sk protocol.Socket) {

	// Handle response stream
	OutcomeResponseHandler(sk)
}

// IncomeResponseHandler use to handle incoming DNS response streams!
func IncomeResponseHandler(sk protocol.Socket) {

}

// OutcomeRequestHandler use to handle outcoming DNS request stream!
func OutcomeRequestHandler(sk protocol.Socket) (err protocol.Error) {
	return
}

// OutcomeResponseHandler use to handle outcoming DNS response stream!
func OutcomeResponseHandler(sk protocol.Socket) (err protocol.Error) {
	return
}
