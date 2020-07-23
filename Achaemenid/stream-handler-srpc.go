/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	"../errors"
	"../srpc"
)

// Indicate standard listen and send port number register for sRPC protocol.
const (
	ProtocolPortSRPCReceive uint16 = 4
	ProtocolPortSRPCSend    uint16 = 5
)

// SRPCHandler use to standard services handlers in any layer!
type SRPCHandler func(*Server, *Stream)

// SrpcIncomeRequestHandler handle incoming sRPC request streams!
func SrpcIncomeRequestHandler(s *Server, st *Stream) {
	st.ReqRes.Err = srpc.CheckPacket(st.Payload, 4)
	if st.ReqRes.Err != nil {
		st.Connection.FailedServiceCall++
		// TODO::: Attack?? tel router to block
		st.ReqRes.Payload = make([]byte, 4)
		// Handle response stream
		SrpcOutcomeResponseHandler(s, st.ReqRes)
		return
	}

	st.ServiceID = srpc.GetID(st.Payload)

	var service *Service = s.Services.GetServiceHandlerByID(st.ServiceID)
	if service == nil {
		st.Connection.FailedServiceCall++
		// TODO::: Attack?? tel router to block
		st.ReqRes.Payload = make([]byte, 4)
		st.ReqRes.Err = ErrSRPCServiceNotFound
		// Handle response stream
		SrpcOutcomeResponseHandler(s, st.ReqRes)
		return
	}

	// Handle request stream
	service.SRPCHandler(s, st)
	if st.ReqRes.Err != nil {
		st.Connection.FailedServiceCall++
		// TODO::: Attack?? tel router to block
		st.ReqRes.Payload = make([]byte, 4)
		// Handle response stream
		SrpcOutcomeResponseHandler(s, st.ReqRes)
		return
	}

	// Handle response stream
	SrpcOutcomeResponseHandler(s, st.ReqRes)
}

// SrpcIncomeResponseHandler use to handle incoming sRPC response streams!
func SrpcIncomeResponseHandler(s *Server, st *Stream) {
	// Get error code from
	st.Err = errors.GetErrByCode(srpc.GetID(st.Payload))

	st.SetState(StreamStateReady)
}

// SrpcOutcomeRequestHandler use to handle outcoming sRPC request stream!
// given stream can't be nil, otherwise panic will occur!
// It block caller until get response or error!!
func SrpcOutcomeRequestHandler(s *Server, st *Stream) (err error) {
	srpc.SetID(st.Payload, st.ServiceID)

	err = st.SendReq()
	// TODO::: handle send error almost due to no network available or connection closed!

	return
}

// SrpcOutcomeResponseHandler use to handle outcoming sRPC response stream!
func SrpcOutcomeResponseHandler(s *Server, st *Stream) (err error) {
	// Close request stream!
	st.Connection.CloseStream(st.ReqRes)

	// write error code to stream payload if exist.
	srpc.SetID(st.Payload, errors.GetCode(st.Err))

	// send stream to outcome pool by weight
	err = st.Send()
	// TODO::: handle send error almost due to no network available or connection closed!

	return
}
