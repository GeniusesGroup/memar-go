/* For license and copyright information please see LEGAL file in repository */

package object

import (
	"../protocol"
)

func (ser *readService) ServeSRPC(st protocol.Stream, srpcReq protocol.SRPCRequest) (res protocol.Syllab, err protocol.Error) {
	var srpcRequestPayload = srpcReq.Payload()
	var reqAsSyllab = readRequestSyllab(srpcRequestPayload)
	err = reqAsSyllab.CheckSyllab(srpcRequestPayload)
	if err != nil {
		return
	}

	res, err = read(reqAsSyllab)
	return
}

func read(req readRequest) (res ReadResponse, err protocol.Error) {
	res.data, err = protocol.OS.ObjectDirectory().Read(req.ObjectID(), req.ObjectStructureID(), req.Offset(), req.Limit())
	return
}
