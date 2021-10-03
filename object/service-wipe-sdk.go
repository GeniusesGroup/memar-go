/* For license and copyright information please see LEGAL file in repository */

package object

import (
	"../protocol"
	"../service"
	"../srpc"
	"../syllab"
)

var WipeService = wipeService{
	Service: service.New("urn:giti:object.protocol:service:wipe", "", protocol.ServiceStatePreAlpha, 1587282740).
		SetDetail(protocol.LanguageEnglish, "Wipe Object",
			`Wipe specific object by given ID in all cluster!
We don't suggest use this service, due to we strongly suggest think about data as immutable entity(stream and time)
It won't wipe object history or indexes associate to it!`,
			[]string{}).
		SetAuthorization(protocol.CRUDDelete, protocol.UserTypeApp).Expired(0, ""),
}

type wipeService struct {
	service.Service
}

func (ser *wipeService) DoSRPC(req WipeRequest) (err protocol.Error) {
	var node protocol.ApplicationNode
	node, err = protocol.App.GetNodeByObjectID(req.objectID)
	if err != nil {
		return
	}

	if node.Status() == protocol.ApplicationStateLocalNode {
		return wipe(req)
	}

	_, err = srpc.HandleOutcomeRequest(node.Conn(), ser, &req)
	return
}

/*
	Service request and response shape
*/

type wipeRequest interface {
	RequestType() RequestType
	ObjectID() [32]byte
	ObjectStructureID() uint64
}

/*
	Service Request
*/

// WipeRequest is request structure of Wipe()
type WipeRequest struct {
	requestType       RequestType
	objectID          [32]byte
	objectStructureID uint64
}

// methods to implements wipeRequest interface
func (req *WipeRequest) RequestType() RequestType        { return req.requestType }
func (req *WipeRequest) ObjectID() [32]byte              { return req.objectID }
func (req *WipeRequest) ObjectStructureID() uint64       { return req.objectStructureID }
func (req *WipeRequest) SetRequestType(rt RequestType)   { req.requestType = rt }
func (req *WipeRequest) SetObjectID(id [32]byte)         { req.objectID = id }
func (req *WipeRequest) SetObjectStructureID(sID uint64) { req.objectStructureID = sID }

// methods to implements protocol.Syllab interface
func (req *WipeRequest) CheckSyllab(payload []byte) (err protocol.Error) {
	if len(payload) < int(req.LenOfSyllabStack()) {
		err = syllab.ErrShortArrayDecode
	}
	return
}
func (req *WipeRequest) FromSyllab(payload []byte, stackIndex uint32) {
	req.requestType = RequestType(syllab.GetUInt8(payload, 0))
	copy(req.objectID[:], payload[1:])
	req.objectStructureID = syllab.GetUInt64(payload, 33)
	return
}
func (req *WipeRequest) ToSyllab(payload []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32) {
	syllab.SetUInt8(payload, 0, uint8(req.requestType))
	copy(payload[1:], req.objectID[:])
	syllab.SetUInt64(payload, 33, req.objectStructureID)
	return heapIndex
}
func (req *WipeRequest) LenAsSyllab() uint64          { return 41 }
func (req *WipeRequest) LenOfSyllabStack() uint32     { return 41 }
func (req *WipeRequest) LenOfSyllabHeap() (ln uint32) { return }

type wipeRequestSyllab []byte

// methods to implements wipeRequest interface
func (req wipeRequestSyllab) RequestType() RequestType         { return RequestType(syllab.GetUInt8(req, 0)) }
func (req wipeRequestSyllab) ObjectID() (objectID [32]byte)    { copy(objectID[:], req[1:]); return }
func (req wipeRequestSyllab) ObjectStructureID() (osID uint64) { return syllab.GetUInt64(req, 33) }
func (req wipeRequestSyllab) SetRequestType(rt RequestType)    { syllab.SetUInt8(req, 0, uint8(rt)) }

// methods to implements protocol.Syllab interface
func (req wipeRequestSyllab) CheckSyllab(payload []byte) (err protocol.Error) {
	if len(req) < int(req.LenOfSyllabStack()) {
		err = syllab.ErrShortArrayDecode
	}
	return
}
func (req wipeRequestSyllab) FromSyllab(payload []byte, stackIndex uint32) {
	// err = ErrSourceNotChangeable
}
func (req wipeRequestSyllab) ToSyllab(payload []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32) {
	copy(payload[stackIndex:], req)
	return heapIndex
}
func (req wipeRequestSyllab) LenAsSyllab() uint64          { return 41 }
func (req wipeRequestSyllab) LenOfSyllabStack() uint32     { return 41 }
func (req wipeRequestSyllab) LenOfSyllabHeap() (ln uint32) { return }
