/* For license and copyright information please see LEGAL file in repository */

package object

import (
	"../protocol"
	"../service"
	"../srpc"
	"../syllab"
)

var GetService = getService{
	Service: service.New("urn:giti:object.protocol:service:get", "", protocol.ServiceStatePreAlpha, 1587282740).
		SetDetail(protocol.LanguageEnglish, "Get",
			`use to get an object by the objectID and structureID! In multi node application, Request must send to proper node otherwise get not found error!`,
			[]string{}).
		SetAuthorization(protocol.CRUDRead, protocol.UserTypeApp).Expired(0, ""),
}

type getService struct {
	service.Service
}

func (ser *getService) DoSRPC(req GetRequest) (res protocol.Object, err protocol.Error) {
	var node protocol.ApplicationNode
	node, err = protocol.App.GetNodeByObjectID(req.objectID)
	if err != nil {
		return
	}

	if node.Status() == protocol.ApplicationStateLocalNode {
		return get(&req)
	}

	var srpcRes srpc.Response
	srpcRes, err = srpc.HandleOutcomeRequest(node.Conn(), ser, &req)
	if err != nil {
		return
	}
	res = Object(syllab.GetByteArray(srpcRes.Payload(), 0))
	return
}

/*
	Service request and response shape
*/

type getRequest interface {
	ObjectID() [32]byte
	ObjectStructureID() uint64
}

/*
	Service Request
*/

// GetRequest is request structure of Get()
type GetRequest struct {
	objectID          [32]byte
	objectStructureID uint64
}

// methods to implements getRequest interface
func (req *GetRequest) ObjectID() [32]byte               { return req.objectID }
func (req *GetRequest) ObjectStructureID() uint64        { return req.objectStructureID }
func (req *GetRequest) SetObjectID(oID [32]byte)         { req.objectID = oID }
func (req *GetRequest) SetObjectStructureID(osID uint64) { req.objectStructureID = osID }

// methods to implements protocol.Syllab interface
func (req *GetRequest) CheckSyllab(payload []byte) (err protocol.Error) {
	if len(payload) < int(req.LenOfSyllabStack()) {
		err = syllab.ErrShortArrayDecode
	}
	return
}
func (req *GetRequest) FromSyllab(payload []byte, stackIndex uint32) {
	copy(req.objectID[:], payload[:])
	req.objectStructureID = syllab.GetUInt64(payload, 32)
}
func (req *GetRequest) ToSyllab(payload []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32) {
	copy(payload[4:], req.objectID[:])
	syllab.SetUInt64(payload, 36, req.objectStructureID)
	return heapIndex
}
func (req *GetRequest) LenAsSyllab() uint64          { return 40 }
func (req *GetRequest) LenOfSyllabStack() uint32     { return 40 }
func (req *GetRequest) LenOfSyllabHeap() (ln uint32) { return }

type getRequestSyllab []byte

// methods to implements getRequest interface
func (req getRequestSyllab) ObjectID() (objectID [32]byte)    { copy(objectID[:], req[0:]); return }
func (req getRequestSyllab) ObjectStructureID() (osID uint64) { return syllab.GetUInt64(req, 32) }

// methods to implements protocol.Syllab interface
func (req getRequestSyllab) CheckSyllab(payload []byte) (err protocol.Error) {
	if len(req) < int(req.LenOfSyllabStack()) {
		err = syllab.ErrShortArrayDecode
	}
	return
}
func (req getRequestSyllab) FromSyllab(payload []byte, stackIndex uint32) {
	// err = ErrSourceNotChangeable
}
func (req getRequestSyllab) ToSyllab(payload []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32) {
	copy(payload[stackIndex:], req)
	return heapIndex
}
func (req getRequestSyllab) LenAsSyllab() uint64          { return 40 }
func (req getRequestSyllab) LenOfSyllabStack() uint32     { return 40 }
func (req getRequestSyllab) LenOfSyllabHeap() (ln uint32) { return }
