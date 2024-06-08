/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

/*
**********************************************************************************
Application (OSI Layer 7: Application)
**********************************************************************************
*/

// OSI_Application usually use to save state and release thread(goroutine) in waiting state
type OSI_Application interface {
	Service() Service //
	Request() any     // Codec
	Response() any    // Codec
	Error() Error     // just indicate peer error that receive by response of the request.

	ObjectLifeCycle
	OSI_Application_LowLevelAPIs
}

// OSI_Application_LowLevelAPIs is low level APIs, don't use them in the services layer, if you don't know how it can be effect the application.
type OSI_Application_LowLevelAPIs interface {
	// Below Set methods must call just once,
	// But some protocol like http allow to change it after first set in a reusable socket like IP/TCP.
	SetService(ser Service)
	SetRequest(req any)
	SetResponse(res any)
	SetError(err Error)

	OperationImportance // base on the connection and the service priority and weight
	OSI_Application_Handler
	// Stringer_To[String] // e.g. "http", ...
}

type OSI_Application_Handler interface {
	// Put in related queue to process income socket in non-blocking mode, means It must not block the caller in any ways.
	// Socket must start with NetworkStatus_NeedMoreData if it doesn't need to call the service when the state changed for the first time
	ScheduleProcessingSocket()

	// HandleIncomeRequest must check socket status
	HandleIncomeRequest(sk Socket) (err Error)

	// SendBidirectionalRequest()
	// SendUnidirectionalRequest()
	// Due to each application handler wants its signature, implement it as a pure function inside each package.
	// srpc.SendBidirectionalRequest(sk Socket, sr Service, req Codec) (res Codec, err Error)
	// http.SendBidirectionalRequest(sk Socket, sr Service, httpReq *Request) (httpRes *Response, err Error) {
}
