/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Stream indicate a minimum networking stream functionality usually occur in layer 4.
// It must also implement chunks managing like sRPC, QUIC, TCP, UDP, ...
type Stream interface {
	Connection() Connection
	Handler() NetworkCommonHandler // usage is like TCP||UDP ports that indicate payload protocol like TLS, HTTPv1, HTTPv2, ...
	Service() Service              //
	Error() Error                  // just indicate peer error that receive by response of the request.

	// Authorize request by data in related stream and connection by any data like service, time, ...
	// Dev must extend this method in each service by it uses.
	Authorize() (err Error)

	Stream_ID
	Network_Status
	Timeout
	OperationImportance // base on the connection and the service priority and weight
	StreamOptions
	StreamLowLevelAPIs
}

// StreamLowLevelAPIs is low level APIs, don't use them in the services layer, if you don't know how it can be effect the application.
// It will use in chunks managing packages e.g. sRPC, QUIC, TCP, UDP, ... or Application layer protocols e.g. HTTP, ...
type StreamLowLevelAPIs interface {
	Send(data Codec) (err Error) // Listen to stream state to check request successfully send, response ready to serve, ...
	Close() (err Error)          // Just once, must deregister the stream from the connection and notify peer in some proper way.

	// use to save state and release thread(goroutine) in waiting state
	Request() any
	Response() any
	SetRequest(req any)
	SetResponse(res any)

	SetHandler(nch NetworkCommonHandler) // Just once, (But some protocol like http allow to change it after first set in a reusable stream like IP/TCP, Allow them??)
	SetService(ser Service)              // Just once, (But some protocol like http allow to change it after first set in a reusable stream like IP/TCP, Allow them??)
	SetError(err Error)                  // Just once
	// Put in related queue to process income stream in non-blocking mode, means It must not block the caller in any ways.
	// Stream must start with NetworkStatus_NeedMoreData if it doesn't need to call the service when the state changed for the first time
	ScheduleProcessingStream()

	Codec
}

type StreamOptions interface {
	// release any underling data reference until call time without need to release socket itself
	Discard(n int) (discarded int, err Error)
	SetLinger(d Duration) error
	SetKeepAlivePeriod(d Duration) error
	SetNoDelay(noDelay bool) error
}
