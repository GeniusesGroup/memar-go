/* For license and copyright information please see LEGAL file in repository */

package object

import (
	"../protocol"
	"../service"
	"../srpc"
	"../syllab"
)

var WriteService = writeService{
	Service: service.New("urn:giti:object.protocol:service:write", "", protocol.ServiceStatePreAlpha, 1587282740).
		SetDetail(protocol.LanguageEnglish, "Write",
			`write some part of a object! Don't use this service until you force to use!
Recalculate checksum do in database server that is not so efficient!`,
			[]string{}).
		SetAuthorization(protocol.CRUDUpdate, protocol.UserTypeApp).Expired(0, ""),
}

type writeService struct {
	service.Service
}

func (ser *writeService) DoSRPC(req *WriteReq) (err protocol.Error) {
	var node protocol.ApplicationNode
	node, err = protocol.App.GetNodeByObjectID(req.objectID)
	if err != nil {
		return
	}

	if node.Status() == protocol.ApplicationStateLocalNode {
		return write(req)
	}

	_, err = srpc.HandleOutcomeRequest(node.Conn(), ser, req)
	return
}

/*
	Service request and response shape
*/

type writeRequest interface {
	RequestType() RequestType
	ObjectID() [32]byte
	ObjectStructureID() uint64
	Offset() uint64
	Data() []byte
}

/*
	Service Request
*/

// WriteReq is request structure of Write()
type WriteReq struct {
	requestType       RequestType
	objectID          [32]byte
	objectStructureID uint64
	offset            uint64 // start location of write data
	data              []byte
}

// methods to implements writeRequest interface
func (req *WriteReq) RequestType() RequestType      { return req.requestType }
func (req *WriteReq) ObjectID() [32]byte            { return req.objectID }
func (req *WriteReq) ObjectStructureID() uint64     { return req.objectStructureID }
func (req *WriteReq) Offset() uint64                { return req.offset }
func (req *WriteReq) Data() (data []byte)           { return req.data }
func (req *WriteReq) SetRequestType(rt RequestType) { req.requestType = rt }

// methods to implements protocol.Syllab interface
func (req *WriteReq) CheckSyllab(payload []byte) (err protocol.Error) {
	if len(payload) < int(req.LenOfSyllabStack()) {
		err = syllab.ErrShortArrayDecode
	}
	return
}
func (req *WriteReq) FromSyllab(payload []byte, stackIndex uint32) {
	req.requestType = RequestType(syllab.GetUInt8(payload, 0))
	copy(req.objectID[:], payload[1:])
	req.objectStructureID = syllab.GetUInt64(payload, 33)
	req.offset = syllab.GetUInt64(payload, 41)
	req.data = syllab.GetByteArray(payload, 49)
}
func (req *WriteReq) ToSyllab(payload []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32) {
	syllab.SetUInt8(payload, 4, uint8(req.requestType))
	copy(payload[5:], req.objectID[:])
	syllab.SetUInt64(payload, 33, req.objectStructureID)
	syllab.SetUInt64(payload, 41, req.offset)
	freeHeapIndex = syllab.SetByteArray(payload, req.data, 49, heapIndex)
	return
}
func (req *WriteReq) LenAsSyllab() uint64 {
	return uint64(req.LenOfSyllabStack() + req.LenOfSyllabHeap())
}
func (req *WriteReq) LenOfSyllabStack() uint32     { return 57 }
func (req *WriteReq) LenOfSyllabHeap() (ln uint32) { return uint32(len(req.data)) }

type writeRequestSyllab []byte

// methods to implements writeRequest interface
func (req writeRequestSyllab) RequestType() RequestType      { return RequestType(syllab.GetUInt8(req, 0)) }
func (req writeRequestSyllab) ObjectID() (oID [32]byte)      { copy(oID[:], req[1:]); return }
func (req writeRequestSyllab) ObjectStructureID() uint64     { return syllab.GetUInt64(req, 33) }
func (req writeRequestSyllab) Offset() uint64                { return syllab.GetUInt64(req, 41) }
func (req writeRequestSyllab) Data() (data []byte)           { return syllab.GetByteArray(req, 1) }
func (req writeRequestSyllab) SetRequestType(rt RequestType) { syllab.SetUInt8(req, 0, uint8(rt)) }

// methods to implements protocol.Syllab interface
func (req writeRequestSyllab) CheckSyllab(payload []byte) (err protocol.Error) {
	if len(req) < int(req.LenOfSyllabStack()) {
		err = syllab.ErrShortArrayDecode
	}
	return
}
func (req writeRequestSyllab) FromSyllab(payload []byte, stackIndex uint32) {
	// err = ErrSourceNotChangeable
}
func (req writeRequestSyllab) ToSyllab(payload []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32) {
	copy(payload[stackIndex:], req)
	return heapIndex
}
func (req writeRequestSyllab) LenAsSyllab() uint64 {
	return uint64(req.LenOfSyllabStack() + req.LenOfSyllabHeap())
}
func (req writeRequestSyllab) LenOfSyllabStack() uint32     { return 57 }
func (req writeRequestSyllab) LenOfSyllabHeap() (ln uint32) { return uint32(len(req) - 57) }
