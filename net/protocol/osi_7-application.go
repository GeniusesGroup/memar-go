/* For license and copyright information please see the LEGAL file in the code repository */

package net_p

import (
	request_p "memar/audit/request/protocol"
	capsule_p "memar/computer/capsule/protocol"
	error_p "memar/error/protocol"
	operation_p "memar/operation/protocol"
	service_p "memar/operation/service/protocol"
)

/*
**********************************************************************************
Application (OSI Layer 7: Application)
**********************************************************************************
*/

// OSI_Application usually use to save state and release thread(goroutine) in waiting state
type OSI_Application interface {
	capsule_p.LifeCycle

	request_p.Field_RequestID
	service_p.Field_Service
	
	operation_p.Field_Request
	operation_p.Field_Response

	// just indicate peer error that receive by response of the request.
	error_p.Field_Error

	OSI_Application_LowLevelAPIs
}

// OSI_Application_LowLevelAPIs is low level APIs, don't use them in the services layer, if you don't know how it can be effect the application.
type OSI_Application_LowLevelAPIs interface {
	// Below Set methods must call just once,
	// But some protocol like http allow to change it after first set in a reusable socket like IP/TCP.
	SetService(ser service_p.Service)
	SetRequest(req operation_p.Request)
	SetResponse(res operation_p.Response)
	SetError(err error_p.Error)

	operation_p.Importance // base on the connection and the service priority and weight
	OSI_Application_Handler
	// Stringer_To[String] // e.g. "http", ...
}

type OSI_Application_Handler interface {
	// Put in related queue to process income socket in non-blocking mode, means It must not block the caller in any ways.
	// Socket must start with NetworkStatus_NeedMoreData if it doesn't need to call the service when the state changed for the first time
	ScheduleProcessingSocket()

	// HandleIncomeRequest must check socket status
	HandleIncomeRequest(sk Socket) (err error_p.Error)

	// SendBidirectionalRequest()
	// SendUnidirectionalRequest()
	// Due to each application handler wants its signature, implement it as a pure function inside each package.
	// srpc.SendBidirectionalRequest(sk Socket, sr Service, req Codec) (res Codec, err error_p.Error)
	// http.SendBidirectionalRequest(sk Socket, sr Service, httpReq *Request) (httpRes *Response, err error_p.Error) {
}
