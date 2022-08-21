/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Streams indicate how stream pools mechanism must behave.
type Streams interface {
	// OutcomeStream make the stream and returns it or return error if any problems occur
	OutcomeStream(service Service) (stream Stream, err Error)
	// IncomeStream make the stream and returns it or return error if any problems occur
	IncomeStream(id uint64) (stream Stream, err Error)
	// Stream returns Stream from pool if exists by given ID
	Stream(id uint64) (stream Stream, err Error)

	// OpenStream opens a new bidirectional QUIC stream.
	// There is no signaling to the peer about new streams:
	// The peer can only accept the stream after data has been sent on the stream.
	// If the error is non-nil, it satisfies the net.Error interface.
	// When reaching the peer's stream limit, err.Temporary() will be true.
	// If the connection was closed due to a timeout, Timeout() will be true.
	OpenStream() (Stream, error)
	// OpenUniStream opens a new outgoing unidirectional QUIC stream.
	// If the error is non-nil, it satisfies the net.Error interface.
	// When reaching the peer's stream limit, Temporary() will be true.
	// If the connection was closed due to a timeout, Timeout() will be true.
	OpenUniStream() (Stream, error)
}
