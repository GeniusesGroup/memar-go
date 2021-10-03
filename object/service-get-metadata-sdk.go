/* For license and copyright information please see LEGAL file in repository */

package object

import (
	"../protocol"
	"../service"
	"../srpc"
	"../syllab"
)

var GetMetadataService = getMetadataService{
	Service: service.New("urn:giti:object.protocol:service:get-metadata", "", protocol.ServiceStatePreAlpha, 1587282740).
		SetDetail(protocol.LanguageEnglish, "Get Metadata",
			`use to get an object by the objectID and structureID! It must send to proper node otherwise get not found error!`,
			[]string{}).
		SetAuthorization(protocol.CRUDRead, protocol.UserTypeApp).Expired(0, ""),
}

type getMetadataService struct {
	service.Service
}

func (ser *getMetadataService) DoSRPC(req GetMetadataRequest) (res protocol.ObjectMetadata, err protocol.Error) {
	var node protocol.ApplicationNode
	node, err = protocol.App.GetNodeByObjectID(req.objectID)
	if err != nil {
		return
	}

	if node.Status() == protocol.ApplicationStateLocalNode {
		return getMetadata(&req)
	}

	var srpcRes srpc.Response
	srpcRes, err = srpc.HandleOutcomeRequest(node.Conn(), ser, &req)
	if err != nil {
		return
	}
	res = Metadata(syllab.GetByteArray(srpcRes.Payload(), 0))
	return
}

/*
	Service request and response shape
*/

type getMetadataRequest interface {
	ObjectID() [32]byte
	ObjectStructureID() uint64
}

/*
	Service Request
*/

// GetMetadataRequest is request structure of Object()
type GetMetadataRequest struct {
	objectID          [32]byte
	objectStructureID uint64
}

// methods to implements getRequest interface
func (req *GetMetadataRequest) ObjectID() [32]byte        { return req.objectID }
func (req *GetMetadataRequest) ObjectStructureID() uint64 { return req.objectStructureID }

// methods to implements protocol.Syllab interface
func (req *GetMetadataRequest) CheckSyllab(payload []byte) (err protocol.Error) {
	if len(payload) < int(req.LenOfSyllabStack()) {
		err = syllab.ErrShortArrayDecode
	}
	return
}
func (req *GetMetadataRequest) FromSyllab(payload []byte, stackIndex uint32) {
	copy(req.objectID[:], payload[:])
}
func (req *GetMetadataRequest) ToSyllab(payload []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32) {
	copy(payload[4:], req.objectID[:])
	syllab.SetUInt64(payload, 36, req.objectStructureID)
	return heapIndex
}
func (req *GetMetadataRequest) LenAsSyllab() uint64          { return 40 }
func (req *GetMetadataRequest) LenOfSyllabStack() uint32     { return 40 }
func (req *GetMetadataRequest) LenOfSyllabHeap() (ln uint32) { return }
