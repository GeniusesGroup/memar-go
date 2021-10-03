/* For license and copyright information please see LEGAL file in repository */

package object

import (
	"../protocol"
	"../srpc"
	"../syllab"
)

func (ser *wipeService) ServeSRPC(st protocol.Stream, srpcReq protocol.SRPCRequest) (res protocol.Syllab, err protocol.Error) {
	var srpcRequestPayload = srpcReq.Payload()
	var reqAsSyllab = wipeRequestSyllab(srpcRequestPayload)
	err = reqAsSyllab.CheckSyllab(srpcRequestPayload)
	if err != nil {
		return
	}

	if reqAsSyllab.RequestType() == RequestTypeBroadcast {
		// tell other node that this node handle request and don't send this request to other nodes!
		reqAsSyllab.SetRequestType(RequestTypeStandalone)

		var replicatedLocalNodes = protocol.App.ReplicatedLocalNode()
		// send request to other related nodes
		for i := 0; i < len(replicatedLocalNodes); i++ {
			var conn = replicatedLocalNodes[i].Conn()

			// Make new request-response streams
			var stream protocol.Stream
			stream, err = conn.OutcomeStream(ser)
			if err != nil {
				// TODO::: Can we easily return error if some nodes did their job and not have enough resource to send request to other nodes??
				return
			}
			stream.SetOutcomeData(srpcReq)

			err = conn.Send(stream)
			if err != nil {
				return
			}

			var srpcRes = srpc.Response(stream.IncomeData().Marshal())
			err = srpcRes.Error()
			if err != nil {
				// TODO::: Can we easily return error if some nodes do their job and just one node connection lost??
				return
			}
		}
	}

	// Do for local node
	err = protocol.OS.ObjectDirectory().Wipe(reqAsSyllab.ObjectID(), reqAsSyllab.ObjectStructureID())
	return
}

func wipe(req WipeRequest) (err protocol.Error) {
	if req.RequestType() == RequestTypeBroadcast {
		// tell other node that this node handle request and don't send this request to other nodes!
		req.SetRequestType(RequestTypeStandalone)
		var reqAsSyllab = syllab.NewCodec(&req)

		var replicatedLocalNodes = protocol.App.ReplicatedLocalNode()
		// send request to other related nodes
		for i := 0; i < len(replicatedLocalNodes); i++ {
			var conn = replicatedLocalNodes[i].Conn()

			// Make new request-response streams
			var stream protocol.Stream
			stream, err = conn.OutcomeStream(&WipeService)
			if err != nil {
				// TODO::: Can we easily return error if some nodes did their job and not have enough resource to send request to other nodes??
				return
			}
			stream.SetOutcomeData(&reqAsSyllab)

			err = conn.Send(stream)
			if err != nil {
				return
			}

			var srpcRes = srpc.Response(stream.IncomeData().Marshal())
			err = srpcRes.Error()
			if err != nil {
				// TODO::: Can we easily return error if some nodes do their job and just one node connection lost??
				return
			}
		}
	}

	// Do for local node
	err = protocol.OS.ObjectDirectory().Wipe(req.ObjectID(), req.ObjectStructureID())
	return
}
