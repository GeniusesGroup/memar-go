/* For license and copyright information please see LEGAL file in repository */

package protocol

// Stream is the interface that must implement by any struct to be Stream!
type Streams interface {
	OutcomeStream(service Service) (stream Stream, err Error)
	IncomeStream(id uint32) (stream Stream, err Error)
	Stream(id uint32) (stream Stream, err Error)
}

// Stream is the interface that must implement by any struct to be a stream!
type Stream interface {
	ID() uint32
	Connection() Connection
	Protocol() NetworkApplicationHandler // usage is like TCP||UDP ports that indicate payload protocol is GitiURN ID
	Service() (ser Service, err Error)
	Error() Error
	Status() ConnectionState     // return last stream state
	State() chan ConnectionState // return state channel to listen to new stream state
	Weight() ConnectionWeight

	SetProtocolID(id NetworkApplicationProtocolID) // Just once
	SetService(ser Service)                        // Just once
	SetError(err Error)                            // Just once
	SetState(state ConnectionState)                //

	SendRequest() (err Error)  // Caller block until response ready to serve
	SendResponse() (err Error) // Caller block until response ready to serve
	Close() // Just once

	// Authorize request by data in related stream and connection by any data like service, time, ...
	Authorize() (err Error)

	IncomeData() Codec
	OutcomeData() Codec
	SetIncomeData(codec Codec)
	SetOutcomeData(codec Codec)
}
