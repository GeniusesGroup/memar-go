/* For license and copyright information please see LEGAL file in repository */

package object

import (
	"../protocol"
	"../service"
	"../srpc"
	"../syllab"
)

var ReadService = readService{
	Service: service.New("urn:giti:object.protocol:service:read", "", protocol.ServiceStatePreAlpha, 1587282740).
		SetDetail(protocol.LanguageEnglish, "Read",
			`use to read some part of an object! It must send to proper node otherwise get not found error!
Mostly use to get metadata first to know about object size before get it to split to some shorter part!`,
			[]string{}).
		SetAuthorization(protocol.CRUDRead, protocol.UserTypeApp).Expired(0, ""),
}

type readService struct {
	service.Service
}

func (ser *readService) DoSRPC(req ReadRequest) (res readResponse, err protocol.Error) {
	var node protocol.ApplicationNode
	node, err = protocol.App.GetNodeByObjectID(req.objectID)
	if err != nil {
		return
	}

	if node.Status() == protocol.ApplicationStateLocalNode {
		return read(&req)
	}

	var srpcRes srpc.Response
	srpcRes, err = srpc.HandleOutcomeRequest(node.Conn(), ser, &req)
	if err != nil {
		return
	}
	var srpcResponsePayload = srpcRes.Payload()
	var resAsSyllab = readResponseSyllab(srpcResponsePayload)
	err = resAsSyllab.CheckSyllab(srpcResponsePayload)
	if err != nil {
		return
	}
	res = resAsSyllab
	return
}

/*
	Service request and response shape
*/

type readRequest interface {
	ObjectID() [32]byte
	ObjectStructureID() uint64
	Offset() uint64
	Limit() uint64
}

type readResponse interface {
	Data() []byte
}

/*
	Service Request
*/

// ReadRequest is request structure of Read()
type ReadRequest struct {
	objectID          [32]byte
	objectStructureID uint64
	offset            uint64
	limit             uint64
}

// methods to implements readRequest interface
func (req *ReadRequest) ObjectID() [32]byte               { return req.objectID }
func (req *ReadRequest) ObjectStructureID() uint64        { return req.objectStructureID }
func (req *ReadRequest) Offset() uint64                   { return req.offset }
func (req *ReadRequest) Limit() uint64                    { return req.limit }
func (req *ReadRequest) SetObjectID(oID [32]byte)         { req.objectID = oID }
func (req *ReadRequest) SetObjectStructureID(osID uint64) { req.objectStructureID = osID }
func (req *ReadRequest) SetOffset(offset uint64)          { req.offset = offset }
func (req *ReadRequest) SetLimit(limit uint64)            { req.limit = limit }

// methods to implements protocol.Syllab interface
func (req *ReadRequest) CheckSyllab(payload []byte) (err protocol.Error) {
	if len(payload) < int(req.LenOfSyllabStack()) {
		err = syllab.ErrShortArrayDecode
	}
	return
}
func (req *ReadRequest) FromSyllab(payload []byte, stackIndex uint32) {
	copy(req.objectID[:], payload[:])
	req.objectStructureID = syllab.GetUInt64(payload, 32)
	req.offset = syllab.GetUInt64(payload, 40)
	req.limit = syllab.GetUInt64(payload, 48)
}
func (req *ReadRequest) ToSyllab(payload []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32) {
	copy(payload[0:], req.objectID[:])
	syllab.SetUInt64(payload, 32, req.objectStructureID)
	syllab.SetUInt64(payload, 40, req.offset)
	syllab.SetUInt64(payload, 48, req.limit)
	return heapIndex
}
func (req *ReadRequest) LenAsSyllab() uint64          { return 56 }
func (req *ReadRequest) LenOfSyllabStack() uint32     { return 56 }
func (req *ReadRequest) LenOfSyllabHeap() (ln uint32) { return }

type readRequestSyllab []byte

// methods to implements readRequest interface
func (req readRequestSyllab) ObjectID() (objectID [32]byte)    { copy(objectID[:], req[0:]); return }
func (req readRequestSyllab) ObjectStructureID() (osID uint64) { return syllab.GetUInt64(req, 32) }
func (req readRequestSyllab) Offset() uint64                   { return syllab.GetUInt64(req, 40) }
func (req readRequestSyllab) Limit() uint64                    { return syllab.GetUInt64(req, 48) }

// methods to implements protocol.Syllab interface
func (req readRequestSyllab) CheckSyllab(payload []byte) (err protocol.Error) {
	if len(req) < int(req.LenOfSyllabStack()) {
		err = syllab.ErrShortArrayDecode
	}
	return
}
func (req readRequestSyllab) FromSyllab(payload []byte, stackIndex uint32) {
	// err = ErrSourceNotChangeable
}
func (req readRequestSyllab) ToSyllab(payload []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32) {
	copy(payload[stackIndex:], req)
	return heapIndex
}
func (req readRequestSyllab) LenAsSyllab() uint64          { return 56 }
func (req readRequestSyllab) LenOfSyllabStack() uint32     { return 56 }
func (req readRequestSyllab) LenOfSyllabHeap() (ln uint32) { return }

/*
	Service Response
*/

// ReadResponse is response structure of Read Serice
type ReadResponse struct {
	data []byte
}

// methods to implements readResponse interface
func (req ReadResponse) Data() []byte { return req.data }

// methods to implements protocol.Syllab interface
func (req ReadResponse) CheckSyllab(payload []byte) (err protocol.Error) {
	if len(payload) < int(req.LenOfSyllabStack()) {
		err = syllab.ErrShortArrayDecode
	}
	return
}
func (res ReadResponse) FromSyllab(payload []byte, stackIndex uint32) {
	res.data = syllab.GetByteArray(payload, stackIndex)
}
func (res ReadResponse) ToSyllab(payload []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32) {
	freeHeapIndex = syllab.SetByteArray(payload, res.data, stackIndex, heapIndex)
	return
}
func (res ReadResponse) LenAsSyllab() uint64 {
	return uint64(res.LenOfSyllabStack() + res.LenOfSyllabHeap())
}
func (res ReadResponse) LenOfSyllabStack() uint32     { return 8 }
func (res ReadResponse) LenOfSyllabHeap() (ln uint32) { return uint32(len(res.data)) }

type readResponseSyllab []byte

// methods to implements readResponse interface
func (res readResponseSyllab) Data() []byte { return syllab.GetByteArray(res, 0) }

// methods to implements protocol.Syllab interface
func (res readResponseSyllab) CheckSyllab(payload []byte) (err protocol.Error) {
	if len(res) < int(res.LenOfSyllabStack()) {
		err = syllab.ErrShortArrayDecode
	}
	return
}
func (res readResponseSyllab) FromSyllab(payload []byte, stackIndex uint32) {
	// err = ErrSourceNotChangeable
}
func (res readResponseSyllab) ToSyllab(payload []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32) {
	freeHeapIndex = syllab.SetByteArray(payload, res.Data(), stackIndex, heapIndex)
	return heapIndex
}
func (res readResponseSyllab) LenAsSyllab() uint64 {
	return uint64(res.LenOfSyllabStack() + res.LenOfSyllabHeap())
}
func (res readResponseSyllab) LenOfSyllabStack() uint32     { return 8 }
func (res readResponseSyllab) LenOfSyllabHeap() (ln uint32) { return uint32(len(res) - 8) }
