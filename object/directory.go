/* For license and copyright information please see LEGAL file in repository */

package object

import (
	"../protocol"
)

// Directory implement protocol.ObjectDirectory to be distributed object storage
type Directory struct{}

// Get return the whole object as metadata and data
func (dir *Directory) Get(uuid [32]byte, structureID uint64) (obj protocol.Object, err protocol.Error) {
	// TODO::: First read from local OS (related lib) as cache

	var req = GetRequest{
		objectID:          uuid,
		objectStructureID: structureID,
	}
	obj, err = GetService.DoSRPC(req)
	if err != nil {
		return
	}

	// TODO::: Write to local OS as cache if not enough storage exist do GC(Garbage Collector)
	return
}

// Read return requested part of object data.
func (dir *Directory) Read(uuid [32]byte, osID uint64, offset, limit uint64) (data []byte, err protocol.Error) {
	// TODO::: First read from local storage as cache

	var req = ReadRequest{
		objectID:          uuid,
		objectStructureID: osID,
		offset:            offset,
		limit:             limit,
	}
	var res readResponse
	res, err = ReadService.DoSRPC(req)
	if err != nil {
		return
	}
	data = res.Data()
	return
}

func (dir *Directory) Save(data protocol.Codec) (metadata protocol.ObjectMetadata, err protocol.Error) {
	return SaveService.DoSRPC(data)
}

// Delete delete the object by object-UUID
func (dir *Directory) Delete(uuid [32]byte, structureID uint64) (err protocol.Error) {
	var req = DeleteRequest{
		requestType:       RequestTypeBroadcast,
		objectID:          uuid,
		objectStructureID: structureID,
	}
	err = DeleteService.DoSRPC(req)
	return
}

// Wipe make invisible by remove from primary index & write random data to object location
func (dir *Directory) Wipe(uuid [32]byte, structureID uint64) (err protocol.Error) {
	return
}
