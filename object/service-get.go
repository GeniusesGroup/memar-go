/* For license and copyright information please see LEGAL file in repository */

package object

import (
	"../protocol"
)

func (ser *getService) ServeSRPC(st protocol.Stream, srpcReq protocol.SRPCRequest) (res protocol.Syllab, err protocol.Error) {
	var srpcRequestPayload = srpcReq.Payload()
	var reqAsSyllab = getRequestSyllab(srpcRequestPayload)
	err = reqAsSyllab.CheckSyllab(srpcRequestPayload)
	if err != nil {
		return
	}

	var obj protocol.Object
	obj, err = get(reqAsSyllab)
	if err != nil {
		return
	}
	res = Object(obj.Marshal())
	return
}

func get(req getRequest) (obj protocol.Object, err protocol.Error) {
	obj, err = protocol.OS.ObjectDirectory().Get(req.ObjectID(), req.ObjectStructureID())
	return
}
