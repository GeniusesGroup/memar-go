/* For license and copyright information please see LEGAL file in repository */

package object

import (
	etime "../earth-time"
	"../protocol"
)

func (ser *getMetadataService) ServeSRPC(st protocol.Stream, srpcReq protocol.SRPCRequest) (res protocol.Syllab, err protocol.Error) {
	var srpcRequestPayload = srpcReq.Payload()
	var reqAsSyllab = getRequestSyllab(srpcRequestPayload)
	err = reqAsSyllab.CheckSyllab(srpcRequestPayload)
	if err != nil {
		return
	}

	var objMetadata protocol.ObjectMetadata
	objMetadata, err = getMetadata(reqAsSyllab)
	if err != nil {
		return
	}
	switch md := objMetadata.(type) {
	case Metadata:
		res = md
	default:
		var tempMetadata = metadata{
			id:             objMetadata.ID(),
			writeTime:      etime.Time(objMetadata.WriteTime().Unix()),
			mediaTypeID:    objMetadata.MediaTypeID(),
			compressTypeID: objMetadata.CompressTypeID(),
			dataLength:     objMetadata.DataLength(),
		}
		res = &tempMetadata
	}
	return
}

func getMetadata(req getMetadataRequest) (metadata protocol.ObjectMetadata, err protocol.Error) {
	metadata, err = protocol.OS.ObjectDirectory().Metadata(req.ObjectID(), req.ObjectStructureID())
	return
}
