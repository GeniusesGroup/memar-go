/* For license and copyright information please see LEGAL file in repository */

package object

import (
	"../protocol"
	"../service"
	"../srpc"
	"../syllab"
)

var DeleteService = deleteService{
	Service: service.New("urn:giti:object.protocol:service:delete", "", protocol.ServiceStatePreAlpha, 1587282740).
		SetDetail(protocol.LanguageEnglish, "Delete Object",
			`Delete specific object by given ID in all cluster!
We don't suggest use this service, due to we strongly suggest think about data as immutable entity(stream and time)
It won't delete object history or indexes associate to it!`,
			[]string{}).
		SetAuthorization(protocol.CRUDDelete, protocol.UserTypeApp).Expired(0, ""),
}

type deleteService struct {
	service.Service
}

func (ser *deleteService) DoSRPC(req DeleteRequest) (err protocol.Error) {
	var node protocol.ApplicationNode
	node, err = protocol.App.GetNodeByObjectID(req.objectID)
	if err != nil {
		return
	}

	if node.Status() == protocol.ApplicationStateLocalNode {
		return delete(req)
	}

	_, err = srpc.HandleOutcomeRequest(node.Conn(), ser, &req)
	return
}

/*
	Service request and response shape
*/

type deleteRequest interface {
	RequestType() RequestType
	ObjectID() [32]byte
	ObjectStructureID() uint64
}

/*
	Service Request
*/

// DeleteRequest is request structure of Delete()
type DeleteRequest struct {
	requestType       RequestType
	objectID          [32]byte
	objectStructureID uint64
}

// methods to implements deleteRequest interface
func (req *DeleteRequest) RequestType() RequestType        { return req.requestType }
func (req *DeleteRequest) ObjectID() [32]byte              { return req.objectID }
func (req *DeleteRequest) ObjectStructureID() uint64       { return req.objectStructureID }
func (req *DeleteRequest) SetRequestType(rt RequestType)   { req.requestType = rt }
func (req *DeleteRequest) SetObjectID(id [32]byte)         { req.objectID = id }
func (req *DeleteRequest) SetObjectStructureID(sID uint64) { req.objectStructureID = sID }

// methods to implements protocol.Syllab interface
func (req *DeleteRequest) CheckSyllab(payload []byte) (err protocol.Error) {
	if len(payload) < int(req.LenOfSyllabStack()) {
		err = syllab.ErrShortArrayDecode
	}
	return
}
func (req *DeleteRequest) FromSyllab(payload []byte, stackIndex uint32) {
	req.requestType = RequestType(syllab.GetUInt8(payload, 0))
	copy(req.objectID[:], payload[1:])
	req.objectStructureID = syllab.GetUInt64(payload, 33)
	return
}
func (req *DeleteRequest) ToSyllab(payload []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32) {
	syllab.SetUInt8(payload, 0, uint8(req.requestType))
	copy(payload[1:], req.objectID[:])
	syllab.SetUInt64(payload, 33, req.objectStructureID)
	return heapIndex
}
func (req *DeleteRequest) LenAsSyllab() uint64          { return 41 }
func (req *DeleteRequest) LenOfSyllabStack() uint32     { return 41 }
func (req *DeleteRequest) LenOfSyllabHeap() (ln uint32) { return }

type deleteRequestSyllab []byte

// methods to implements deleteRequest interface
func (req deleteRequestSyllab) RequestType() RequestType         { return RequestType(syllab.GetUInt8(req, 0)) }
func (req deleteRequestSyllab) ObjectID() (objectID [32]byte)    { copy(objectID[:], req[1:]); return }
func (req deleteRequestSyllab) ObjectStructureID() (osID uint64) { return syllab.GetUInt64(req, 33) }
func (req deleteRequestSyllab) SetRequestType(rt RequestType)    { syllab.SetUInt8(req, 0, uint8(rt)) }

// methods to implements protocol.Syllab interface
func (req deleteRequestSyllab) CheckSyllab(payload []byte) (err protocol.Error) {
	if len(req) < int(req.LenOfSyllabStack()) {
		err = syllab.ErrShortArrayDecode
	}
	return
}
func (req deleteRequestSyllab) FromSyllab(payload []byte, stackIndex uint32) {
	// err = ErrSourceNotChangeable
}
func (req deleteRequestSyllab) ToSyllab(payload []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32) {
	copy(payload[stackIndex:], req)
	return heapIndex
}
func (req deleteRequestSyllab) LenAsSyllab() uint64          { return 41 }
func (req deleteRequestSyllab) LenOfSyllabStack() uint32     { return 41 }
func (req deleteRequestSyllab) LenOfSyllabHeap() (ln uint32) { return }
