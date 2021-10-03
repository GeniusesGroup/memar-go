/* For license and copyright information please see LEGAL file in repository */

package object

import (
	"crypto/sha512"

	etime "../earth-time"
	"../protocol"
	"../service"
	"../srpc"
	"../syllab"
)

type saveService struct {
	service.Service
}

var SaveService = saveService{
	Service: service.New("urn:giti:object.protocol:service:save", "", protocol.ServiceStatePreAlpha, 1587282740).
		SetDetail(protocol.LanguageEnglish, "Save",
			`Write the object to storage engine. To replace old one you need to delete old`,
			[]string{}).
		SetAuthorization(protocol.CRUDCreate, protocol.UserTypeApp).Expired(0, ""),
}

func (ser *saveService) DoSRPC(data protocol.Codec) (metadata protocol.ObjectMetadata, err protocol.Error) {
	var object = make([]byte, MetadataLength, int(MetadataLength)+data.Len())
	var objectMetadata = Metadata(object)
	metadata = objectMetadata
	objectMetadata.setWriteTime(etime.Now())
	objectMetadata.setMediaTypeID(data.MediaType().URN().ID())
	objectMetadata.setCompressTypeID(data.CompressType().URN().ID())
	objectMetadata.setDataLength(data.Len())

	object = data.MarshalTo(object[MetadataLength:])
	var objectID = sha512.Sum512_256(object[32:])
	objectMetadata.setID(objectID)

	var node protocol.ApplicationNode
	node, err = protocol.App.GetNodeByObjectID(objectID)
	if err != nil {
		return
	}

	var req = SaveRequest{
		requestType: RequestTypeBroadcast,
		object:      object,
	}
	if node.Status() == protocol.ApplicationStateLocalNode {
		err = save(req)
		return
	}

	_, err = srpc.HandleOutcomeRequest(node.Conn(), ser, &req)
	return
}

/*
	Service request and response shape
*/

type saveRequest interface {
	RequestType() RequestType
	Object() []byte
}

/*
	Service Request
*/

// SaveRequest is request structure of Save()
type SaveRequest struct {
	requestType RequestType
	object      []byte
}

// methods to implements saveRequest interface
func (req *SaveRequest) RequestType() RequestType      { return req.requestType }
func (req *SaveRequest) Object() (objectID []byte)     { return req.object }
func (req *SaveRequest) SetRequestType(rt RequestType) { req.requestType = rt }

// methods to implements protocol.Syllab interface
func (req *SaveRequest) CheckSyllab(payload []byte) (err protocol.Error) {
	if len(payload) < int(req.LenOfSyllabStack()) {
		err = syllab.ErrShortArrayDecode
	}
	return
}
func (req *SaveRequest) FromSyllab(payload []byte, stackIndex uint32) {
	req.requestType = RequestType(syllab.GetUInt8(payload, 0))
	req.object = syllab.GetByteArray(payload, 1)
}
func (req *SaveRequest) ToSyllab(payload []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32) {
	syllab.SetUInt8(payload, 4, uint8(req.requestType))
	freeHeapIndex = syllab.SetByteArray(payload, req.object, stackIndex, heapIndex)
	return
}
func (req *SaveRequest) LenAsSyllab() uint64 {
	return uint64(req.LenOfSyllabStack() + req.LenOfSyllabHeap())
}
func (req *SaveRequest) LenOfSyllabStack() uint32     { return 9 }
func (req *SaveRequest) LenOfSyllabHeap() (ln uint32) { return uint32(len(req.object)) }

type saveRequestSyllab []byte

// methods to implements saveRequest interface
func (req saveRequestSyllab) RequestType() RequestType      { return RequestType(syllab.GetUInt8(req, 0)) }
func (req saveRequestSyllab) Object() (objectID []byte)     { return syllab.GetByteArray(req, 1) }
func (req saveRequestSyllab) SetRequestType(rt RequestType) { syllab.SetUInt8(req, 0, uint8(rt)) }

// methods to implements protocol.Syllab interface
func (req saveRequestSyllab) CheckSyllab(payload []byte) (err protocol.Error) {
	if len(req) < int(req.LenOfSyllabStack()) {
		err = syllab.ErrShortArrayDecode
	}
	return
}
func (req saveRequestSyllab) FromSyllab(payload []byte, stackIndex uint32) {
	// err = ErrSourceNotChangeable
}
func (req saveRequestSyllab) ToSyllab(payload []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32) {
	copy(payload[stackIndex:], req)
	return heapIndex
}
func (req saveRequestSyllab) LenAsSyllab() uint64          { return 9 }
func (req saveRequestSyllab) LenOfSyllabStack() uint32     { return 9 }
func (req saveRequestSyllab) LenOfSyllabHeap() (ln uint32) { return uint32(len(req) - 9) }
