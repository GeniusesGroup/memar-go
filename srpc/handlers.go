/* For license and copyright information please see LEGAL file in repository */

package srpc

import (
	"../giti"
)

const (
	// 8-BYTE as Service or Error ID
	MinLength = 8
)

// IncomeRequestHandler handle incoming sRPC request streams!
func IncomeRequestHandler(app giti.Application, stream giti.Stream) {
	var err giti.Error
	var connection = stream.Connection()
	var srpcReq = Request(stream.IncomeData())
	var srpcRes Response

	// Check request
	err = srpcReq.Check(MinLength)
	if err != nil {
		connection.ServiceCallFail()
		stream.SetError(err)
		srpcRes = MakeNewResponse(err.ID(), 0)
		stream.SetOutcomeData(srpcRes)
		return
	}

	// Find service
	var serviceID = srpcReq.ServiceID()
	var service = app.GetServiceByID(serviceID)
	if service == nil {
		err = ErrServiceNotFound
		connection.ServiceCallFail()
		stream.SetError(err)
		srpcRes = MakeNewResponse(err.ID(), 0)
		stream.SetOutcomeData(srpcRes)
		return
	}
	stream.SetService(service)

	// call request service
	err = service.SRPCHandler(app, stream)
	if err != nil {
		connection.ServiceCallFail()
		stream.SetError(err)
		srpcRes = MakeNewResponse(err.ID(), 0)
		stream.SetOutcomeData(srpcRes)
		return
	}

	connection.ServiceCalled()
}

// IncomeResponseHandler use to handle incoming sRPC response streams!
func IncomeResponseHandler(app giti.Application, stream giti.Stream) {
	var srpcRes = Response(stream.IncomeData())
	var errID = srpcRes.ErrorID()
	var err = app.GetErrorByID(errID)
	stream.SetError(err)
	stream.SetState(giti.ConnectionStateReady)
}

// OutcomeRequestHandler use to handle outcoming sRPC request stream!
// given stream can't be nil, otherwise panic will occur!
// It block caller until get response or error!!
func OutcomeRequestHandler(conn giti.NetworkApplicationConnection, req Request) (err giti.Error) {
	// TODO::: get new stream

	// err = stream.Send()

	return
}
