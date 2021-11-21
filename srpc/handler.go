/* For license and copyright information please see LEGAL file in repository */

package srpc

import (
	"../protocol"
)

// Read more about this protocol : https://github.com/GeniusesGroup/RFCs/blob/master/sRPC.md
type SRPCHandler struct{}

// HandleIncomeRequest handle incoming sRPC request streams that carry on Syllab codec!
func (srpc *SRPCHandler) HandleIncomeRequest(stream protocol.Stream) (err protocol.Error) {
	var service protocol.Service
	service, err = stream.Service()
	if err != nil {
		stream.SetError(err)
		return
	}

	// call request service
	err = service.ServeSRPC(stream)
	if err != nil {
		stream.SetError(err)
	}

	stream.SendResponse()
	stream.Close()
	return
}

// HandleOutcomeRequest use to handle outcoming sRPC request.
// It block caller until get response or error.
func HandleOutcomeRequest(conn protocol.Connection, service protocol.Service, payload protocol.Codec) (stream protocol.Stream, err protocol.Error) {
	stream, err = conn.OutcomeStream(service)
	if err != nil {
		return
	}

	// stream.SetOutcomeData(syllab.NewCodec(payload))
	stream.SetOutcomeData(payload)

	err = stream.SendRequest()
	if err != nil {
		return
	}

	err = stream.Error()
	return
}
