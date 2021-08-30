/* For license and copyright information please see LEGAL file in repository */

package protocol

// Stream is the interface that must implement by any struct to be Stream!
type Streams interface {
	GetStreamByID(ID uint32) Stream
	RegisterStream(Stream)
	CloseStream(Stream)
}

// Stream is the interface that must implement by any struct to be a stream!
type Stream interface {
	ID() uint32
	Connection() NetworkTransportConnection
	Service() Service
	State() ConnectionState
	Error() (err Error)

	SetConnection(conn NetworkTransportConnection) // Just once and register stream in connection streams
	SetService(Service) // Just once
	SetState(ConnectionState)
	SetError(err Error) // Just once

	// Authorize request by data in related stream and connection by any data like service, time, ...
	Authorize() (err Error)

	IncomeData() Codec
	OutcomeData() Codec
	SetIncomeData(codec Codec)
	SetOutcomeData(codec Codec)
}
