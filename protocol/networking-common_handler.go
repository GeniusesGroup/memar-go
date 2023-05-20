/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// NetworkCommonHandler usually use to indicate layer 5, 6, 7 handlers
type NetworkCommonHandler interface {
	ProtocolID() MediaTypeID
	// HandleIncomeRequest must check stream status
	HandleIncomeRequest(st Stream) (err Error)

	ObjectLifeCycle
	Stringer // e.g. "tls", ...
}
