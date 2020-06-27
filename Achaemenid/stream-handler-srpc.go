/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	"../srpc"
	"../errors"
)

const (
	// ProtocolPortSRPC indicate standard port number register for sRPC protocol
	ProtocolPortSRPC uint16 = 4
)

// SRPCHandler use to standard services handlers in any layer!
type SRPCHandler func(*Server, *Stream)

// SrpcIncomeRequestHandler handle incoming sRPC request streams!
func SrpcIncomeRequestHandler(s *Server, st *Stream) {
	var err error
	err = srpc.CheckPacket(st.Payload, 4)
	if err != nil {
		st.Connection.FailedPacketsReceived++
		st.Connection.FailedServiceCall++
		// Attack??
		return
	}

	st.ServiceID = srpc.GetID(st.Payload)

	var service *Service
	service = s.Services.GetServiceHandlerByID(st.ServiceID)
	if service == nil {
		err = ErrSRPCServiceNotFound
		st.Connection.ServiceCallCount++
		st.Connection.FailedServiceCall++
		// handle err
		// Send response or just ignore packet, Attack??
		return
	}

	// Handle request stream
	service.SRPCHandler(s, st)

	// Handle response stream
	SrpcOutcomeResponseHandler(s, st.ReqRes)

	// Close active stream!
	st.Connection.CloseStream(st)
}

// SrpcIncomeResponseHandler use to handle incoming sRPC response streams!
func SrpcIncomeResponseHandler(s *Server, st *Stream) {
	// Get error code from
	st.Err = errors.GetErrByCode(srpc.GetID(st.Payload))

	// tell request stream that response stream ready to use!
	st.ReqRes.StateChannel <- StreamStateReady
}

// SrpcOutcomeRequestHandler use to handle outcoming sRPC request stream!
// given stream can't be nil, otherwise panic will occur!
// It block caller until get response or error!!
func SrpcOutcomeRequestHandler(s *Server, st *Stream) (err error) {
	// Set ServiceID
	srpc.SetID(st.Payload, st.ServiceID)

	// send stream to outcome pool by weight
	err = s.SendStream(st)

	// Listen to response stream and decode error ID and return it to caller
	var responseStatus streamState = <-st.StateChannel
	if responseStatus == StreamStateReady {
		
	} else {

	}

	return
}

// SrpcOutcomeResponseHandler use to handle outcoming sRPC response stream!
func SrpcOutcomeResponseHandler(s *Server, st *Stream) (err error) {
	// Convert error to errors.ExtendedError and write error code to stream payload.
	var ee, ok = st.Err.(*errors.ExtendedError)
	if ok {
		srpc.SetID(st.Payload, ee.Code)
	}

	return
}
