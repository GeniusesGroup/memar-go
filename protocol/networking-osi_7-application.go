/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

/*
**********************************************************************************
Application (OSI Layer 7: Application)
**********************************************************************************
*/

type NetworkApplication_ProtocolID = MediaTypeID

// NetworkApplication_Multiplexer
type NetworkApplication_Multiplexer interface {
	GetNetworkApplicationHandler(protocolID NetworkApplication_ProtocolID) NetworkApplication_Handler
	SetNetworkApplicationHandler(nah NetworkApplication_Handler)
	DeleteNetworkApplicationHandler(protocolID NetworkApplication_ProtocolID)
}

// NetworkApplication_Handler
type NetworkApplication_Handler interface {
	NetworkCommonHandler

	// SendBidirectionalRequest()
	// SendUnidirectionalRequest()
	// Due to each application handler wants its signature, implement it as a pure function inside each package.
	// srpc.SendBidirectionalRequest(conn Connection, service Service, req Codec) (res Codec, err Error)
	// http.SendBidirectionalRequest(conn Connection, service Service, httpReq *Request) (httpRes *Response, err Error) {
}
