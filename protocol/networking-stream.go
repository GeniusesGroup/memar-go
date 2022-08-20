/* For license and copyright information please see LEGAL file in repository */

package protocol

// Stream is the interface that must implement by any struct to be Stream
type Streams interface {
	// OutcomeStream make the stream and returns it or return error if any problems occur
	OutcomeStream(service Service) (stream Stream, err Error)
	// IncomeStream make the stream and returns it or return error if any problems occur
	IncomeStream(id uint64) (stream Stream, err Error)
	// Stream returns Stream from pool if exists by given ID
	Stream(id uint64) (stream Stream, err Error)
}

// Stream is the interface that must implement by any struct to be a stream!
type Stream interface {
	ID() uint64
	Connection() Connection
	Protocol() NetworkApplicationHandler // usage is like TCP||UDP ports that indicate payload protocol is GitiURN ID
	Service() Service
	Error() Error                // just indicate peer error that receive by response of the request.
	Status() ConnectionState     // return last stream state
	State() chan ConnectionState // return state channel to listen to new stream state
	Weight() Weight              // sum of connection and service weight.

	// Authorize request by data in related stream and connection by any data like service, time, ...
	Authorize() (err Error)

	SendRequest(req Codec) (err Error)  // Listen to stream state to check request successfully send, response ready to serve, ...
	SendResponse(res Codec) (err Error) // Listen to stream state to check response successfully send, ...

	// Below methods are low level APIs, don't use them in services layer, if you don't know how it can be effect the application.
	Close() (err Error)                        // Just once, must deregister the stream from the connection and send close message to Socket in some types.
	Socket() Codec                             // Chunks manager like sRPC, QUIC, TCP, UDP, ...
	SetSocket(codec Codec)                     // Just once, use SendRequest||SendResponse methods
	SetProtocol(nah NetworkApplicationHandler) // Just once, (But some protocol like http allow to change it after first set in a reusable stream like IP/TCP, Allow them??)
	SetService(ser Service)                    // Just once, (But some protocol like http allow to change it after first set in a reusable stream like IP/TCP, Allow them??)
	SetError(err Error)                        // Just once
	SetState(state ConnectionState)            // change state of stream and send notification on stream StateChannel.

	// SetDeadline sets the read and write deadlines associated with the connection.
	// It is equivalent to calling both SetReadDeadline and SetWriteDeadline.
	SetDeadline(d Duration) Error
	// SetReadDeadline(d Duration) Error
	// SetWriteDeadline(d Duration) Error
}
