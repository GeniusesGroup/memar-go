/* For license and copyright information please see LEGAL file in repository */

package chaparkhane

import (
	"../srpc"
)

// srpcHandler use to handle sRPC protcol!
func srpcHandler(s *Server, st *Stream) {
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
		st.Connection.FailedPacketsReceived++
		// handle err
		// Send response or just ignore packet, Attack??
		return
	}
	streamHandler(s, st)
}
