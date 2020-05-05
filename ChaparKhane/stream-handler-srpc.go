/* For license and copyright information please see LEGAL file in repository */

package chaparkhane

import (
	"../srpc"
)

// SrpcIncomeRequestHandler use to handle incoming sRPC request streams!
func SrpcIncomeRequestHandler(s *Server, st *Stream) {
	var err error
	err = srpc.CheckPacket(st.Payload, 4)
	if err != nil {
		st.Connection.FailedPacketsReceived++
		// Attack??
		return
	}

	st.ServiceID = srpc.GetID(st.Payload)

	var ok bool
	var streamHandler StreamHandler
	streamHandler, ok = s.Services.GetServiceHandlerByID(st.ServiceID)
	if !ok {
		err = ErrSRPCServiceNotFound
		st.Connection.FailedServiceCall++
		// handle err
		// Send response or just ignore packet, Attack??
		return
	}
	streamHandler(s, st)

	// Close active stream!
	st.Connection.CloseStream(st)
}

// SrpcIncomeResponseHandler use to handle incoming sRPC response streams!
func SrpcIncomeResponseHandler(s *Server, st *Stream) {
	// Get error code from
	st.ServiceID = srpc.GetID(st.Payload)
	// TODO::: convert ErrorID to error!!

	// tell request stream that response stream ready to use!
	st.StatusChannel <- StreamStateReady
}

// SrpcOutcomeRequestHandler use to handle outcoming sRPC request stream!
func (st *Stream) SrpcOutcomeRequestHandler(s *Server) (err error) {
	// Set ServiceID
	srpc.SetID(st.Payload, st.ServiceID)

	// send stream to outcome pool by weight
	err = s.SendStream(st)

	return
}

// SrpcOutcomeResponseHandler use to handle outcoming sRPC response stream!
func (st *Stream) SrpcOutcomeResponseHandler(s *Server) (err error) {
	// TODO::: Convert error to ErrorID and write it to stream payload
	// srpc.SetID(st.Payload, ErrorID)

	return
}
