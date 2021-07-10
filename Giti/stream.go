/* For license and copyright information please see LEGAL file in repository */

package giti

// Stream is the interface that must implement by any struct to be Stream!
type Stream interface {
	ID() uint32

	Connection() NetworkTransportConnection
	SetConnection(conn NetworkTransportConnection) // Just once and register stream in connection streams

	Service() Service
	SetService(Service) // Just once

	State() ConnectionState
	SetState(ConnectionState)

	Error() (err Error)
	SetError(err Error) // Just once

	IncomeData() Buffer
	OutcomeData() Buffer

	Send() (err Error)
	SendAsync() (err Error)
}
