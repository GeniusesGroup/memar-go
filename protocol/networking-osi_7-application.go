/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

/*
**********************************************************************************
Application (OSI Layer 7: Application)
**********************************************************************************
*/

// OSI_Application usually use to save state and release thread(goroutine) in waiting state
type OSI_Application interface {
	Handler() Network_Application_Handler //
	Service() Service                     //
	Request() any                         // Codec
	Response() any                        // Codec
	Error() Error                         // just indicate peer error that receive by response of the request.

	ObjectLifeCycle
	OSI_Application_LowLevelAPIs
}

// OSI_Application_LowLevelAPIs is low level APIs, don't use them in the services layer, if you don't know how it can be effect the application.
type OSI_Application_LowLevelAPIs interface {
	// TODO::: below methods must call just once,
	// TODO::: But some protocol like http allow to change it after first set in a reusable socket like IP/TCP, Allow them??
	SetHandler(nah Network_Application_Handler)
	SetService(ser Service)
	SetError(err Error)
	SetRequest(req any)
	SetResponse(res any)
}

type Network_Application_ProtocolID = MediaTypeID

// Network_Application_Multiplexer
type Network_Application_Multiplexer interface {
	GetNetworkApplicationHandler(pID Network_Application_ProtocolID) Network_Application_Handler
	SetNetworkApplicationHandler(nah Network_Application_Handler)
	DeleteNetworkApplicationHandler(pID Network_Application_ProtocolID)
}

// Network_Application_Handler
type Network_Application_Handler interface {
	ProtocolID() Network_Application_ProtocolID
	// HandleIncomeRequest must check socket status
	HandleIncomeRequest(sk Socket) (err Error)

	ObjectLifeCycle
	Stringer // e.g. "http", ...

	// SendBidirectionalRequest()
	// SendUnidirectionalRequest()
	// Due to each application handler wants its signature, implement it as a pure function inside each package.
	// srpc.SendBidirectionalRequest(sk Socket, sr Service, req Codec) (res Codec, err Error)
	// http.SendBidirectionalRequest(sk Socket, sr Service, httpReq *Request) (httpRes *Response, err Error) {
}
