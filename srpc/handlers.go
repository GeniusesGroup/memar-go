/* For license and copyright information please see LEGAL file in repository */

package srpc

import (
	"net"

	"../protocol"
)

const (
	// 8-BYTE as Service or Error ID
	MinLength uint64 = 8
)

type SRPCSyllabHandler struct{}

// HandleIncomeRequest handle incoming sRPC request streams that carry on Syllab codec!
func (ssh *SRPCSyllabHandler) HandleIncomeRequest(stream protocol.Stream) (err protocol.Error) {
	var srpcReq = Request(stream.IncomeData().Marshal())
	var srpcRes Response

	// Check request
	err = srpcReq.Check(MinLength)
	if err != nil {
		srpcRes = NewResponse(0)
		srpcRes.SetErrorID(err.URN().ID())
		stream.SetOutcomeData(srpcRes)
		return
	}

	// Find service
	var serviceID = srpcReq.ServiceID()
	var service = protocol.App.GetServiceByID(serviceID)
	if service == nil {
		err = ErrServiceNotFound
		srpcRes = NewResponse(0)
		srpcRes.SetErrorID(err.URN().ID())
		stream.SetOutcomeData(srpcRes)
		return
	}
	stream.SetService(service)

	// call request service
	var res protocol.Syllab
	res, err = service.ServeSRPC(stream, srpcReq)
	if err != nil {
		srpcRes = NewResponse(0)
		srpcRes.SetErrorID(err.URN().ID())
		stream.SetOutcomeData(srpcRes)
		return
	}

	srpcRes = NewResponse(res.LenAsSyllab())
	res.ToSyllab(srpcRes.Payload(), 0, res.LenOfSyllabStack())
	stream.SetOutcomeData(srpcRes)
	return
}

func (hh *SRPCSyllabHandler) HandleStreamConnection(stream protocol.Stream, conn net.Conn) {
	// TODO:::
}

// HandleOutcomeRequest use to handle outcoming sRPC request.
// It block caller until get response or error.
func HandleOutcomeRequest(conn protocol.NetworkTransportConnection, service protocol.Service, payload protocol.Syllab) (srpcRes Response, err protocol.Error) {
	var stream protocol.Stream
	stream, err = conn.OutcomeStream()
	if err != nil {
		return
	}

	var srpcReq = NewRequest(payload.LenAsSyllab())
	srpcReq.SetServiceID(service.URN().ID())
	payload.ToSyllab(srpcReq.Payload(), 0, payload.LenOfSyllabStack())
	stream.SetOutcomeData(srpcReq)

	err = conn.Send(stream)
	if err != nil {
		return
	}

	srpcRes = Response(stream.IncomeData().Marshal())
	err = srpcRes.Error()
	return
}
